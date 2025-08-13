// 代码生成时间: 2025-08-14 01:52:49
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
)

// ApiResponse is a structure to define a standard API response format.
type ApiResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
    Error   string      `json:"error"`
    Time    string      `json:"time"`
}

// ErrorResponse is a helper method for creating API error responses.
func ErrorResponse(code int, message, error string) ApiResponse {
    return ApiResponse{
        Code:    code,
        Message: message,
        Error:   error,
        Time:    time.Now().Format(time.RFC1123),
    }
}

// SuccessResponse is a helper method for creating API success responses.
func SuccessResponse(code int, message string, data interface{}) ApiResponse {
    return ApiResponse{
        Code:    code,
        Message: message,
        Data:    data,
        Time:    time.Now().Format(time.RFC1123),
    }
}

func main() {
    app := iris.New()
    app.RegisterView(iris.JSON{})

    // Define a route that handles requests to the root path and responds with a formatted API response.
    app.Get("/", func(ctx iris.Context) {
        // Prepare a success response with sample data.
        response := SuccessResponse(200, "Success", "Hello, world!")

        // Write the response to the client.
        ctx.JSON(response)
    })

    // Define a route that handles requests to the error path and responds with a formatted API error.
    app.Get("/error", func(ctx iris.Context) {
        // Prepare an error response.
        response := ErrorResponse(500, "Internal Server Error", "Something went wrong.")

        // Write the response to the client.
        ctx.JSON(response)
    })

    // Start the Iris web server.
    log.Printf("Server is running at :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
