// 代码生成时间: 2025-09-08 17:09:22
package main

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12" // 引入IRIS框架
)

// ErrorResponse 结构体用于返回错误响应
type ErrorResponse struct {
    Error string `json:"error"`
}

// helloHandler 处理GET请求的处理器
func helloHandler(ctx iris.Context) {
    // 输出HTTP响应
    ctx.WriteString("Hello, World!")
}

// notFoundHandler 处理404页面的处理器
func notFoundHandler(ctx iris.Context) {
    // 发送404状态码
    ctx.StatusCode(iris.StatusNotFound)
    // 返回错误信息
    ctx.JSON(ErrorResponse{Error: "The requested resource was not found."})
}

// main 函数初始化并启动IRIS服务
func main() {
    // 创建IRIS应用
    app := iris.New()

    // 使用全局错误处理
    app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

    // 定义GET路由
    app.Get("/", helloHandler)

    // 配置端口和启动服务
    app.Listen(":8080")
}
