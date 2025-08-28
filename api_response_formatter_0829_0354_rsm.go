// 代码生成时间: 2025-08-29 03:54:54
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// ApiResponse 定义API响应的格式
type ApiResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
    Message string     `json:"message"`
}

// NewApiResponse 创建一个ApiResponse实例，用于响应
func NewApiResponse(data interface{}, message string) ApiResponse {
    return ApiResponse{
        Success: true,
        Data:    data,
        Message: message,
    }
}

// ErrorResponse 创建一个表示错误的ApiResponse实例
func ErrorResponse(message string) ApiResponse {
    return ApiResponse{
        Success: false,
        Data:    nil,
        Message: message,
    }
}

func main() {
    app := iris.New()

    // 定义GET /api/formatter路由，用于演示格式化响应
    app.Get("/api/formatter", func(ctx iris.Context) {
        // 模拟一些数据
        data := map[string]interface{}{
            "name": "John",
            "age": 30,
        }

        // 使用ApiResponse格式化响应
        response := NewApiResponse(data, "Data retrieved successfully")

        // 发送JSON响应
        ctx.JSON(response)
    })

    // 定义GET /api/error路由，用于演示错误响应
    app.Get("/api/error", func(ctx iris.Context) {
        // 发送错误响应
        ctx.JSON(ErrorResponse("An error occurred"))
    })

    // 启动服务器
    fmt.Println("Server is running at http://localhost:8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Printf("Failed to start server: %s
", err)
    }
}
