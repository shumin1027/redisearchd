package http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gitlab.xtc.home/xtc/redisearchd/conn/redis"
	"gitlab.xtc.home/xtc/redisearchd/pkg/http"
	"gitlab.xtc.home/xtc/redisearchd/pkg/json"
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
	r.Get("/:id", GetDocByID)
	r.Post("", CreateDocs)
	r.Patch("/:id", UpdateDocFields)
	r.Delete("", DeleteDocs)
	r.Delete("/:id", DeleteDocByID)

	r.Put("", UpdateDocs)
}

type UpdateDocFieldsPayload struct {
	Fields []string `json:"fields"`
	Values []string `json:"values"`
}

// GetDocByID
// @Summary Get Doc By Id
// @Description Get Doc By Id
// @Produce application/json
// @Tags doc
// @Router /docs/{id} [GET]
// @Param id path string true "doc id"
// @Success 200 {object} http.Response
func GetDocByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fields := c.Query("fields")
	doc, err := self.GetDocById(c.Context(), redis.Pool(), id, strings.Split(fields, ",")...)
	if err != nil {
		return http.Error(c, err)
	}

	return http.Success(c, doc)
}

// CreateDocs
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
	pool := redis.Pool()
	err := self.AddDocs(c.Context(), pool, docs...)
	if err != nil {
		return http.Error(c, err)
	}
	return c.SendStatus(http.StatusCreated)
}

// DeleteDocByID
// @Summary Delete One Doc By Id
// @Description Delete One Doc By Id
// @Produce application/json
// @Tags doc
// @Router /docs/{id} [DELETE]
// @Success 204 {string} string ""
func DeleteDocByID(c *fiber.Ctx) error {
	id := c.Params("id")
	connPool := redis.Pool()
	err := self.DeleteDocs(c.Context(), connPool, id)
	if err != nil {
		return http.Error(c, err)
	}
	return c.SendStatus(http.StatusNoContent)
}

// DeleteDocs
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
	connPool := redis.Pool()
	err := self.DeleteDocs(c.Context(), connPool, ids...)

	if err != nil {
		return http.Error(c, err)
	}

	return c.SendStatus(http.StatusNoContent)
}

//UpdateDocs
//@Summary Update docs,Use "HSET"
//@Description Update docs,Use "HSET"
//@Produce application/json
//@Tags doc
//@Router /docs [PUT]
//@Success 200
func UpdateDocs(c *fiber.Ctx) error {
	var docs self.DocumentList
	body := c.Request().Body()
	if err := json.Unmarshal(body, &docs); err != nil {
		return http.Error(c, err)
	}
	pool := redis.Pool()
	err := self.UpdateDocs(c.Context(), pool, docs...)
	if err != nil {
		return http.Error(c, err)
	}
	return http.Success(c, fiber.Map{})
}

//UpdateDocFields
//@Summary Update doc field,Use "HMSET"
//@Description Update doc field,Use "HMSET"
//@Produce application/json
//@Tags doc
//@Router /docs/{od} [PATCH]
//@Success 200
func UpdateDocFields(c *fiber.Ctx) error {
	var updatePayload UpdateDocFieldsPayload
	body := c.Request().Body()
	if err := json.Unmarshal(body, &updatePayload); err != nil {
		return http.Error(c, err)
	}
	id := c.Params("id")
	connPool := redis.Pool()
	err := self.UpdateDocFields(c.Context(), connPool, id, updatePayload.Fields, updatePayload.Values)
	if err != nil {
		return http.Error(c, err)
	}
	return http.Success(c, fiber.Map{})
}
