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
// @Param num query int false "maximum number of documents returned. default is `10`;max is `1_000_000`. when num is `0`, just return the count" default(10) minimum(0) maximum(1000000)
// @Param offset query int false "number of documents to skip，default is `0`" default(0) minimum(0)
func SearchByGet(c *gin.Context) {
	index := c.Param("index")
	cli := conn.Client(index)

	raw := c.Query("raw")

	var offset, num int
	var err error

	poffset := c.Query("offset")
	if len(poffset) > 0 {
		offset, err = strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
	}

	pnum := c.Query("num")
	if len(pnum) > 0 {
		num, err = strconv.Atoi(pnum)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		if num > PAGE_NUM_LIMIT_MAX {
			num = PAGE_NUM_LIMIT_MAX
		}
	} else {
		num = 10
	}

	query := &redisearch.Query{
		Raw: raw,
		Paging: redisearch.Paging{
			Offset: offset,
			Num:    num,
		},
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
// @Router /search/{index} [GET]
// @Param index path string true "index name"
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
