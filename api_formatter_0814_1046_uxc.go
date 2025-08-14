// 代码生成时间: 2025-08-14 10:46:13
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
)

// ResponseFormatter is a struct that holds the data needed for formatting the API responses.
type ResponseFormatter struct {
    // Fields can be added here to include additional response data
    Data interface{} `json:"data"`
    Message string `json:"message"`
    Success bool `json:"success"`
    Timestamp time.Time `json:"timestamp"`
}

// NewResponseFormatter creates a new instance of ResponseFormatter.
func NewResponseFormatter(data interface{}, message string, success bool) *ResponseFormatter {
    return &ResponseFormatter{
        Data:      data,
        Message:   message,
        Success:   success,
        Timestamp: time.Now(),
    }
}

// FormatResponse is a middleware that formats the response in a standardized way.
func FormatResponse(ctx iris.Context) {
    ctx.Next()
    result := ctx.Values().GetString("result")
    if result == "" {
        ctx.JSON(http.StatusInternalServerError, NewResponseFormatter(nil, "No data to format", false))
        return
    }
    ctx.JSON(http.StatusOK, NewResponseFormatter(result, "Data retrieved successfully", true))
}

func main() {
    app := iris.New()
    app.Use(FormatResponse)

    // Define a route that demonstrates the use of the middleware
    app.Get("/api/data", func(ctx iris.Context) {
        ctx.Values().Set("result", "This is a sample data response")
    })

    // Start the server
    log.Fatal(app.Listen(":8080"))
}
