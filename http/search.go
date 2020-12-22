package http

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"gitlab.xtc.home/xtc/redisearchd/conn"
	self "gitlab.xtc.home/xtc/redisearchd/pkg/redisearch"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// 分页最大数量限制
const PAGE_NUM_LIMIT_MAX = 1_000_000

type SearchRouter struct {
	*gin.RouterGroup
}

func NewSearchRouter(r *gin.RouterGroup) *SearchRouter {
	return &SearchRouter{r}
}

func (r *SearchRouter) Route() {
	r.GET("/:index", SearchByGet)
	r.POST("/:index", SearchByPost)
}

// @Summary Search in an index with GET
// @Description Searches the index with a textual query, returning either documents or just count(when num=0 and offset=0).
// @Produce application/json
// @Tags search
// @Router /search/{index} [GET]
// @Param index path string true "index name"
// @Param raw query string true " the text query to search"
// @Param limit query int false "maximum number of documents returned. default is `10`;max is `1_000_000`. when num is `0`, just return the count" default(10) minimum(0) maximum(1000000)
// @Param offset query int false "number of documents to skip，default is `0`" default(0) minimum(0)
// @Param in_keys query []string false "If set, we limit the result to a given set of keys specified in the list. " collectionFormat(csv)
// @Param in_fields query []string false " If set, filter the results to ones appearing only in specific fields of the document, like title or URL. num is the number of specified field arguments " collectionFormat(csv)
// @Param return_fields query []string false "Use this keyword to limit which fields from the document are returned.e.g: `return_fields=id,name,age` " collectionFormat(csv)
// @Param sort_by query string false "If specified, the results are ordered by the value of this field. This applies to both text and numeric fields. e.g: `sort_by=name|asc`"
// @Param language query string false "If set, we use a stemmer for the supplied language during search for query expansion"
// @Success 200 {array} redisearch.Document
func SearchByGet(c *gin.Context) {
	index := c.Param("index")
	cli := conn.Client(index)

	raw := c.Query("raw")

	var limit, offset int
	var err error

	plimit := c.Query("limit")
	if len(plimit) > 0 {
		limit, err = strconv.Atoi(plimit)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
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
			c.String(http.StatusBadRequest, err.Error())
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

	docs, total, err := self.Search(c.Request.Context(), cli, query)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"docs":  docs,
		"total": total,
	})
}

// @Summary Search in an index with POST
// @Description Searches the index with a textual query, returning either documents or just count(when num=0 and offset=0).
// @Produce application/json
// @Tags search
// @Router /search/{index} [POST]
// @Param index path string true "index name"
// @Success 200 {array} redisearch.Document
func SearchByPost(c *gin.Context) {
	index := c.Param("index")
	cli := conn.Client(index)
	var query = new(redisearch.Query)
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err := jsoniter.Unmarshal(body, query); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	docs, total, err := self.Search(c.Request.Context(), cli, query)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"docs":  docs,
		"total": total,
	})
}
