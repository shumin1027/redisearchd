package http

import (
	"context"
	"gitlab.xtc.home/xtc/redisearchd/pkg/http"
	"gitlab.xtc.home/xtc/redisearchd/pkg/search"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gitlab.xtc.home/xtc/redisearchd/pkg/json"
	"gitlab.xtc.home/xtc/redisearchd/pkg/redis"
	self "gitlab.xtc.home/xtc/redisearchd/pkg/redisearch"
)

type DocRouter struct {
	*fiber.Group
}

func NewDocRouter(r fiber.Router) *DocRouter {
	g, ok := r.(*fiber.Group)
	if ok {
		return &DocRouter{g}
	}
	return nil
}

func (r *DocRouter) Route() {
	r.Get("/:id", GetDocById)
	r.Post("", CreateDocs)
	r.Delete("", DeleteDocs)
	r.Delete("/:id", DeleteDocById)

	r.Put("/:id", UpdateDocById)
}

// @Summary Get Doc By Id
// @Description Get Doc By Id
// @Produce application/json
// @Tags doc
// @Router /docs/{id} [GET]
// @Param id path string true "doc id"
// @Success 200 {object} http.Response
func GetDocById(c *fiber.Ctx) error {
	id := c.Params("id")
	fields := c.Query("fields")
	doc, err := self.GetDocById(c.Context(), search.NewConnPool(), id, strings.Split(fields, ",")...)
	if err != nil {
		return http.Error(c, err)
	}
	return http.Success(c, doc)
}

// @Summary Create Docs
// @Description Create Docs
// @Produce application/json
// @Tags doc
// @Router /docs [POST]
// @Success 200 {string} string ""
func CreateDocs(c *fiber.Ctx) error {
	var docs self.DocumentList
	body := c.Request().Body()

	if err := json.Unmarshal(body, &docs); err != nil {
		return http.Error(c, err)
	}
	conn := search.NewConnPool()
	err := self.AddDocs(c.Context(), conn, docs...)
	if err != nil {
		return http.Error(c, err)
	}
	return c.SendStatus(http.StatusCreated)
}

// @Summary Delete One Doc By Id
// @Description Delete One Doc By Id
// @Produce application/json
// @Tags doc
// @Router /docs/{id} [DELETE]
// @Success 204 {string} string ""
func DeleteDocById(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := redis.Del(c.Context(), id)
	if err != nil {
		return http.Error(c, err)
	}
	return c.SendStatus(http.StatusNoContent)
}

// @Summary Batch Delete Docs By Ids
// @Description Batch Delete Docs By Ids
// @Produce application/json
// @Tags doc
// @Router /docs [DELETE]
// @Success 204 {string} string ""
func DeleteDocs(c *fiber.Ctx) error {
	var ids []string
	body := c.Request().Body()

	if err := json.Unmarshal(body, &ids); err != nil {
		return http.Error(c, err)
	}
	_, err := redis.Del(c.Context(), ids...)

	if err != nil {
		return http.Error(c, err)
	}

	return c.SendStatus(http.StatusNoContent)
}

//UpdateDocById
//@Summary Update key,Use "HSET"
//@Description Update key,Use "HSET"
//@Produce application/json
//@Tags doc
//@Router /docs/{id} [PUT]
//@Param id path string true "document id"
//@Success 200
func UpdateDocById(c *fiber.Ctx) error {
	key := c.Params("id")
	var values map[string]interface{}
	data := c.Body()
	err := json.Unmarshal(data, &values)
	if err != nil {
		return http.Error(c, err)
	}
	_, err = redis.HSet(context.TODO(), key, values)
	if err != nil {
		return http.Error(c, err)
	}
	return http.Success(c, fiber.Map{})
}
