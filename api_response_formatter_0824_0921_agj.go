// 代码生成时间: 2025-08-24 09:21:37
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// ApiResponseFormatter provides a standardized way to format API responses.
type ApiResponseFormatter struct{}

// SuccessResponse formats a response with a success status.
func (formatter ApiResponseFormatter) SuccessResponse(data interface{}) map[string]interface{} {
    return map[string]interface{}{
        "status":  "success",
        "payload": data,
    }
}

// ErrorResponse formats a response with an error status.
func (formatter ApiResponseFormatter) ErrorResponse(errorCode int, errorMessage string) map[string]interface{} {
    return map[string]interface{}{
        "status":    "error",
        "error":     errorCode,
        "message":   errorMessage,
    }
}

// NewApiResponseFormatter creates a new instance of ApiResponseFormatter.
func NewApiResponseFormatter() ApiResponseFormatter {
    return ApiResponseFormatter{}
}

func main() {
    // Initialize the Iris application.
    app := iris.Default()

    // Create a new instance of ApiResponseFormatter.
    formatter := NewApiResponseFormatter()

    // Register a route for handling API requests.
    app.Get("/api/response", func(ctx iris.Context) {
        // Simulate a successful API response.
        data := map[string]string{"key": "value"}
        response := formatter.SuccessResponse(data)

        // Write the formatted response to the client.
        ctx.JSON(iris.StatusOK, response)
    })

    // Handle a simulated error scenario.
    app.Get("/api/error", func(ctx iris.Context) {
        // Simulate an error response.
        response := formatter.ErrorResponse(400, "Bad Request")

        // Write the formatted error response to the client.
        ctx.JSON(iris.StatusBadRequest, response)
    })

    // Start the Iris server.
    app.Listen(":8080")
}