// 代码生成时间: 2025-09-03 13:02:25
package main

import (
    "html"
    "log"
    "net/http"
    "github.com/kataras/iris/v12"
)

// XssMiddleware is a middleware function that prevents XSS attacks
// by escaping HTML special characters in the input.
func XssMiddleware(ctx iris.Context) {
    ctx.Next()
    // After processing the request, escape the output
    ctx.Stop()
    escapedHtml := html.EscapeString(ctx.ResponseWriter().Header().Get("X-Original-Response"))
    ctx.ResponseWriter().Header().Del("X-Original-Response")
    ctx.ResponseWriter().Header().Set("Content-Type", "text/html; charset=utf-8")
    ctx.WriteString(escapedHtml)
}

// Main function to start the Iris web server with XSS protection middleware
func main() {
    app := iris.New()

    // Register the middleware globally
    app.Use(XssMiddleware)

    // Sample route to demonstrate the middleware
    app.Get("/xss", func(ctx iris.Context) {
        // Simulate user input that could be vulnerable to XSS
        userInput := "<script>alert('XSS')</script>"
        // Respond with the original user input, which will be escaped by the middleware
        ctx.Header("X-Original-Response", userInput)
        ctx.StatusCode(iris.StatusOK)
    })

    // Start the Iris web server
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Error starting the server: %s
", err)
    }
}
