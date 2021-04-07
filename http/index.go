package http

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gitlab.xtc.home/xtc/redisearchd/internal/json"

	"github.com/RediSearch/redisearch-go/redisearch"
	"gitlab.xtc.home/xtc/redisearchd/conn"
	self "gitlab.xtc.home/xtc/redisearchd/pkg/redisearch"
)

type IndexRouter struct {
	*fiber.Group
}

func NewIndexRouter(r fiber.Router) *IndexRouter {
	g, ok := r.(*fiber.Group)
	if ok {
		return &IndexRouter{g}
	}
	return nil
}

func (r *IndexRouter) Route() {
	r.Get("", List)
	r.Get("/:index", Info)
	r.Post("/:index", CreateIndex)
	r.Delete("/:index", DropIndex)
}

// @Summary List all indexes
// @Description List all indexes
// @Produce application/json
// @Tags index
// @Router /indexes [GET]
// @Success 200 {array} string
func List(c *fiber.Ctx) error {
	cli := conn.DummyClient()
	indexes, err := self.ListIndexes(c.Context(), cli)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(indexes)
}

// @Summary Get Index Info
// @Description Get Index Info
// @Produce application/json
// @Tags index
// @Router /indexes/{index} [GET]
// @Param index path string true "index name"
// @Success 200 {object} redisearch.IndexInfo
func Info(c *fiber.Ctx) error {
	index := c.Params("index")
	cli := conn.Client(index)
	info, err := self.Info(c.Context(), cli)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(info)
}

// @Summary Delete an index
// @Description Delete an index
// @Produce application/json
// @Tags index
// @Router /indexes/{index} [DELETE]
// @Param index path string true "index name"
// @Param deldocs query bool false "delete document"
// @Success 204 {string} string ""
func DropIndex(c *fiber.Ctx) error {
	deldocs := false
	index := c.Params("index")
	if len(c.Query("deldocs")) > 0 && strings.ToLower(c.Query("deldocs")) == "true" {
		deldocs = true
	}
	cli := conn.Client(index)
	err := self.DropIndex(c.Context(), cli, deldocs)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(http.StatusNoContent)
}

type CreateIndexReq struct {
	Schema          *redisearch.Schema
	IndexDefinition *redisearch.IndexDefinition
}

// @Summary Create an index
// @Description Create an index. `schema.fields.type`: `0`-`Text`;`1`-`Numeric`;`2`-`Geo`;`3`-`Tag`
// @Produce application/json
// @Tags index
// @Router /indexes/{index} [POST]
// @Param index path string true "index name"
// @Success 200 {string} string ""
func CreateIndex(c *fiber.Ctx) error {
	var req CreateIndexReq
	body := c.Request().Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	index := c.Params("index")
	cli := conn.Client(index)
	if err := self.CreateIndex(c.Context(), cli, req.Schema, req.IndexDefinition); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(http.StatusCreated)
}
