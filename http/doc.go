package http

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"gitlab.xtc.home/xtc/redisearchd/conn"
	self "gitlab.xtc.home/xtc/redisearchd/pkg/redisearch"
)

type DocRouter struct {
	*gin.RouterGroup
}

func NewDocRouter(group *gin.RouterGroup) *DocRouter {
	return &DocRouter{group}
}

func (r *DocRouter) Route() {
	r.GET("/:id", GetDocById)
	r.POST("", CreateDocs)
	r.DELETE("", DeleteDocs)
	r.DELETE("/:id", DeleteDocById)
}

// @Summary Get Doc By Id
// @Description Get Doc By Id
// @Produce application/json
// @Tags doc
// @Router /docs/{id} [GET]
// @Param id path string true "doc id"
func GetDocById(c *gin.Context) {
	id := c.Param("id")
	fields := c.Query("fields")
	doc, err := self.GetDocById(c.Request.Context(), conn.ConnPool(), id, strings.Split(fields, ",")...)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, doc)
}

// @Summary Create Docs
// @Description Create Docs
// @Produce application/json
// @Tags doc
// @Router /docs [POST]
func CreateDocs(c *gin.Context) {
	var docs self.DocumentList
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err := jsoniter.Unmarshal(body, &docs); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	conn := conn.ConnPool()
	err = self.AddDocs(c.Request.Context(), conn, docs...)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}

// @Summary Delete One Doc By Id
// @Description Delete One Doc By Id
// @Produce application/json
// @Tags doc
// @Router /docs/{id} [DELETE]
func DeleteDocById(c *gin.Context) {
	id := c.Param("id")

	conn := conn.ConnPool()
	err := self.DeleteDocs(c.Request.Context(), conn, id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}

// @Summary Batch Delete Docs By Ids
// @Description Batch Delete Docs By Ids
// @Produce application/json
// @Tags doc
// @Router /docs [DELETE]
func DeleteDocs(c *gin.Context) {
	var ids []string
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err := jsoniter.Unmarshal(body, &ids); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	conn := conn.ConnPool()
	err = self.DeleteDocs(c.Request.Context(), conn, ids...)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}
