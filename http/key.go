package http

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"gitlab.xtc.home/xtc/redisearchd/conn/redis"
	"gitlab.xtc.home/xtc/redisearchd/pkg/http"
)

type KeyRouter struct {
	*fiber.Group
}

func NewKeyRouter(r fiber.Router) *KeyRouter {
	g, ok := r.(*fiber.Group)
	if ok {
		return &KeyRouter{g}
	}
	return nil
}

func (r *KeyRouter) Route() {
	r.Get("/:key", GetKeyAll)
	r.Put("/:key", UpdateKeyAll)
	r.Delete("/:key", DeleteKey)
}

// GetKeyAll
// @Summary Get key
// @Description Get key
// @Produce application/json
// @Tags key
// @Router /keys/{key} [GET]
// @Param key path string true "index name"
// @Success 200
func GetKeyAll(c *fiber.Ctx) error {
	key := c.Params("key")
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
// @Tags key
// @Router /keys/{key} [PUT]
// @Param key path string true "index name"
// @Success 200
func UpdateKeyAll(c *fiber.Ctx) error {
	key := c.Params("key")
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
// @Tags key
// @Router /keys/{key} [DELETE]
// @Param key path string true "index name"
// @Success 200
func DeleteKey(c *fiber.Ctx) error {
	key := c.Params("key")
	client := redis.Client()
	_, err := client.Del(context.TODO(), key).Result()
	if err != nil {
		return http.Error(c, err)
	}
	return http.Success(c, fiber.Map{})
}
