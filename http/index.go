package http

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"gitlab.xtc.home/xtc/redisearchd/conn"
	self "gitlab.xtc.home/xtc/redisearchd/pkg/redisearch"
)

type IndexRouter struct {
	*gin.RouterGroup
}

func NewIndexRouter(r *gin.RouterGroup) *IndexRouter {
	return &IndexRouter{r}
}

func (r *IndexRouter) Route() {
	r.GET("", List)
	r.GET("/:index", Info)
	r.POST("/:index", CreateIndex)
	r.DELETE("/:index", DropIndex)
}

// @Summary List all indexes
// @Description List all indexes
// @Produce application/json
// @Tags index
// @Router /indexes [GET]
// @Success 200 {array} string
func List(c *gin.Context) {
	cli := conn.DummyClient()
	indexes, _ := self.ListIndexes(c.Request.Context(), cli)
	c.JSON(http.StatusOK, indexes)
}

// @Summary Get Index Info
// @Description Get Index Info
// @Produce application/json
// @Tags index
// @Router /indexes/{index} [GET]
// @Param index path string true "index name"
// @Success 200 {object} redisearch.IndexInfo
func Info(c *gin.Context) {
	index := c.Param("index")
	cli := conn.Client(index)
	info, err := self.Info(c.Request.Context(), cli)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, info)
}

// @Summary Delete an index
// @Description Delete an index
// @Produce application/json
// @Tags index
// @Router /indexes/{index} [DELETE]
// @Param index path string true "index name"
// @Param deldocs query bool false "delete document"
// @Success 200 {string} string ""
func DropIndex(c *gin.Context) {
	deldocs := false
	index := c.Param("index")
	if len(c.Query("deldocs")) > 0 && strings.ToLower(c.Query("deldocs")) == "true" {
		deldocs = true
	}
	cli := conn.Client(index)
	err := self.DropIndex(c.Request.Context(), cli, deldocs)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "")
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
func CreateIndex(c *gin.Context) {
	var req CreateIndexReq
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err := jsoniter.Unmarshal(body, &req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	index := c.Param("index")
	cli := conn.Client(index)
	if err := self.CreateIndex(c.Request.Context(), cli, req.Schema, req.IndexDefinition); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}
