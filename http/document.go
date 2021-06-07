package http

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"gitlab.xtc.home/xtc/redisearchd/conn/redis"
	"gitlab.xtc.home/xtc/redisearchd/pkg/http"
)

type DocumentsRouter struct {
	*fiber.Group
}

func NewDocumentsRouter(r fiber.Router) *DocumentsRouter {
	g, ok := r.(*fiber.Group)
	if ok {
		return &DocumentsRouter{g}
	}
	return nil
}

func (r *DocumentsRouter) Route() {
	r.Get("/:document_id", GetKeyAll)
	r.Put("/:document_id", UpdateKeyAll)
	r.Delete("/:document_id", DeleteKey)
}

// GetKeyAll
// @Summary Get key
// @Description Get key
// @Produce application/json
// @Tags document
// @Router /documents/{document_id} [GET]
// @Param document_id path string true "index name"
// @Success 200
func GetKeyAll(c *fiber.Ctx) error {
	key := c.Params("document_id")
	client := redis.Client()
	result, err := client.HGetAll(context.TODO(), key).Result()
	if err != nil {
		return http.Error(c, err)
	}
	return http.Success(c, result)
}

// UpdateKeyAll
// @Summary Update key,Use "HSET"
// @Description Update key,Use "HSET"
// @Produce application/json
// @Tags document
// @Router /documents/{document_id} [PUT]
// @Param document_id path string true "index name"
// @Success 200
func UpdateKeyAll(c *fiber.Ctx) error {
	key := c.Params("document_id")
	var values map[string]interface{}
	data := c.Body()
	err := json.Unmarshal(data, &values)
	if err != nil {
		return http.Error(c, err)
	}
	client := redis.Client()
	_, err = client.HSet(context.TODO(), key, values).Result()
	if err != nil {
		return http.Error(c, err)
	}
	return http.Success(c, fiber.Map{})
}

// DeleteKey
// @Summary Delete key,Use "Del"
// @Description Delete key,Use "Del"
// @Produce application/json
// @Tags document
// @Router /documents/{document_id} [DELETE]
// @Param document_id path string true "index name"
// @Success 200
func DeleteKey(c *fiber.Ctx) error {
	key := c.Params("document_id")
	client := redis.Client()
	_, err := client.Del(context.TODO(), key).Result()
	if err != nil {
		return http.Error(c, err)
	}
	return http.Success(c, fiber.Map{})
}
