package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// 统一报文格式
type Response struct {
	Success   bool        `json:"success"`   // 是否调用成功：true表示调用成功，false表示调用失败
	Inventory interface{} `json:"inventory"` // 调用成功时，返回的数据清单
	Error     HttpError   `json:"error"`     // 调用失败时，返回的出错信息
}

// 2xx 处理成功
func Success(c *fiber.Ctx, inventory interface{}) error {
	rep := Response{
		Success:   true,
		Inventory: inventory,
	}
	return c.Status(http.StatusOK).JSON(rep)
}

// 4xx 处理失败
func Fail(c *fiber.Ctx, msg string, code ...int) error {
	err := ErrBadRequest
	err.Message = msg
	if len(code) > 0 {
		err.Code = code[0]
	}
	rep := Response{
		Success: false,
		Error:   err,
	}
	return c.Status(http.StatusBadRequest).JSON(rep)
}

// 5xx 处理出错
func Error(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusInternalServerError).JSON(err.Error())
}
