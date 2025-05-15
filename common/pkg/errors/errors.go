package errors

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 错误响应结构
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

// 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录错误日志
				log.Printf("panic: %v", err)

				// 返回统一错误响应
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "服务器内部错误",
					Detail:  "请稍后重试",
				})

				// 终止请求
				c.Abort()
			}
		}()

		// 继续处理请求
		c.Next()

		// 处理未处理的错误
		if len(c.Errors) > 0 {
			firstError := c.Errors[0]

			// 根据错误类型返回不同的HTTP状态码
			var statusCode int
			var message string

			switch firstError.Err.(type) {
			case *ValidationError:
				statusCode = http.StatusBadRequest
				message = "参数验证失败"
			case *UnauthorizedError:
				statusCode = http.StatusUnauthorized
				message = "未授权"
			case *ForbiddenError:
				statusCode = http.StatusForbidden
				message = "禁止访问"
			case *NotFoundError:
				statusCode = http.StatusNotFound
				message = "资源不存在"
			default:
				statusCode = http.StatusInternalServerError
				message = "服务器内部错误"
			}

			c.JSON(statusCode, ErrorResponse{
				Code:    statusCode,
				Message: message,
				Detail:  firstError.Error(),
			})
		}
	}
}

// 自定义错误类型
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

// 其他错误类型...
