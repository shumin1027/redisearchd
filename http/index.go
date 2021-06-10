package http

import (
	"gitlab.xtc.home/xtc/redisearchd/pkg/http"
	"gitlab.xtc.home/xtc/redisearchd/pkg/search"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gitlab.xtc.home/xtc/redisearchd/pkg/json"

	"github.com/RediSearch/redisearch-go/redisearch"
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

	r.Get("/:index/search", SearchIndexByGet)
	r.Post("/:index/search", SearchIndexByPost)
}

// @Summary List all indexes
// @Description List all indexes
// @Produce application/json
// @Tags index
// @Router /indexes [GET]
// @Success 200 {array} string
func List(c *fiber.Ctx) error {
	cli := search.NewClient("_redisearch_")
	indexes, err := self.ListIndexes(c.Context(), cli)
	if err != nil {
		return http.Fail(c, err.Error())
	}
	return http.Success(c, indexes)
}

// @Summary Get Index Info
// @Description Get Index Info
// @Produce application/json
// @Tags index
// @Router /indexes/{index} [GET]
// @Param index path string true "index name"
// @Success 200 {object} http.Response
func Info(c *fiber.Ctx) error {
	index := c.Params("index")
	client := search.NewClient(index)
	info, err := client.Info()
	if err != nil {
		return http.Fail(c, err.Error())
	}
	return http.Success(c, info)
}

// @Summary Delete an index
// @Description Delete an index
// @Produce application/json
// @Tags index
// @Router /indexes/{index} [DELETE]
// @Param index path string true "index name"
// @Param deldocs query bool false "delete document"
// @Success 204
func DropIndex(c *fiber.Ctx) error {
	deldocs := false
	index := c.Params("index")
	if len(c.Query("deldocs")) > 0 && strings.ToLower(c.Query("deldocs")) == "true" {
		deldocs = true
	}
	cli := search.NewClient(index)
	err := self.DropIndex(c.Context(), cli, deldocs)
	if err != nil {
		return http.Fail(c, err.Error())
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
// @Success 201
func CreateIndex(c *fiber.Ctx) error {
	var req CreateIndexReq
	body := c.Request().Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return http.Fail(c, err.Error())
	}
	index := c.Params("index")
	cli := search.NewClient(index)
	if err := self.CreateIndex(c.Context(), cli, req.Schema, req.IndexDefinition); err != nil {
		return http.Fail(c, err.Error())
	}
	return c.SendStatus(http.StatusCreated)
}

// @Summary Search in an index with GET
// @Description Searches the index with a textual query, returning either documents or just count(when num=0 and offset=0).
// @Produce application/json
// @Tags search
// @Router /indexes/{index}/search [GET]
// @Param index path string true "index name"
// @Param raw query string true " the text query to search"
// @Param limit query int false "maximum number of documents returned. default is `10`;max is `1_000_000`. when num is `0`, just return the count" default(10) minimum(0) maximum(1000000)
// @Param offset query int false "number of documents to skip，default is `0`" default(0) minimum(0)
// @Param in_keys query []string false "If set, we limit the result to a given set of keys specified in the list. " collectionFormat(csv)
// @Param in_fields query []string false " If set, filter the results to ones appearing only in specific fields of the document, like title or URL. num is the number of specified field arguments " collectionFormat(csv)
// @Param return_fields query []string false "Use this keyword to limit which fields from the document are returned.e.g: `return_fields=id,name,age` " collectionFormat(csv)
// @Param sort_by query string false "If specified, the results are ordered by the value of this field. This applies to both text and numeric fields. e.g: `sort_by=name|asc`"
// @Param language query string false "If set, we use a stemmer for the supplied language during search for query expansion"
// @Success 200 {array} http.Response
func SearchIndexByGet(c *fiber.Ctx) error {
	index := c.Params("index")
	cli := search.NewClient(index)

	raw := c.Query("raw")

	var limit, offset int
	var err error

	plimit := c.Query("limit")
	if len(plimit) > 0 {
		limit, err = strconv.Atoi(plimit)
		if err != nil {
			return http.Fail(c, err.Error())
		}
		if limit > PAGE_NUM_LIMIT_MAX {
			limit = PAGE_NUM_LIMIT_MAX
		}
	} else {
		limit = 10
	}

	poffset := c.Query("offset")
	if len(poffset) > 0 {
		offset, err = strconv.Atoi(c.Query("offset"))
		if err != nil {
			return http.Fail(c, err.Error())
		}
	}

	query := &redisearch.Query{
		Raw: raw,
		Paging: redisearch.Paging{
			Offset: offset,
			Num:    limit,
		},
	}

	if len(c.Query("in_keys")) > 0 {
		in_keys := strings.Split(c.Query("in_keys"), ",")
		query.InKeys = in_keys
	}

	if len(c.Query("in_fields")) > 0 {
		in_fields := strings.Split(c.Query("in_fields"), ",")
		query.InKeys = in_fields
	}

	if len(c.Query("return_fields")) > 0 {
		return_fields := strings.Split(c.Query("return_fields"), ",")
		query.ReturnFields = return_fields
	}

	if len(c.Query("sort_by")) > 0 {
		s := strings.Split(c.Query("sort_by"), "|")

		field := s[0]
		sort := strings.ToLower(s[1])

		var ascending bool

		if sort == "asc" {
			ascending = true
		} else if sort == "desc" {
			ascending = false
		}

		query.SortBy = &redisearch.SortingKey{
			Field:     field,
			Ascending: ascending,
		}

	}

	if len(c.Query("language")) > 0 {
		query.Language = c.Query("language")
	}

	docs, total, err := self.Search(c.Context(), cli, query)
	if err != nil {
		return http.Fail(c, err.Error())
	}
	return http.Success(c, fiber.Map{
		"docs":  docs,
		"total": total,
	})
}

// @Summary Search in an index with POST
// @Description Searches the index with a textual query, returning either documents or just count(when num=0 and offset=0).
// @Produce application/json
// @Tags index
// @Router /indexes/{index}/search [POST]
// @Param index path string true "index name"
// @Success 200 {array} http.Response
func SearchIndexByPost(c *fiber.Ctx) error {
	index := c.Params("index")
	cli := search.NewClient(index)
	var query = new(redisearch.Query)
	body := c.Request().Body()

	if err := json.Unmarshal(body, query); err != nil {
		return c.SendString(err.Error())
	}
	docs, total, err := self.Search(c.Context(), cli, query)
	if err != nil {
		return http.Fail(c, err.Error())
	}
	return http.Success(c, fiber.Map{
		"docs":  docs,
		"total": total,
	})
}
