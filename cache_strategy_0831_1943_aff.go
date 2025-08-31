// 代码生成时间: 2025-08-31 19:43:01
package main

import (
    "fmt"
    "time"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/cache"
)

// Cache configuration
var cacheConfig = cache.Config{
    // Specify a custom cache implementation, or use one of the built-in ones.
    // Here, we use the memory cache which is a simple in-memory cache that is not shared between different instances of the application.
    Adapter: cache.NewMemory(
        cache.WithTTL(10 * time.Minute), // Time to live
    ),
}

func main() {
    app := iris.New()
    app.Use(func(ctx iris.Context) {
        // This middleware will check if the cache is hit and if so, it will return the cached response.
        if resp, err := cacheConfig.Get(ctx.Request().URL.Path); err == nil {
            ctx.StopWithJSON(iris.StatusOK, iris.Map{
                "message": "Cache hit", "response": resp,
            })
            return
        }
        ctx.Next() // Call the next middleware/handler in the chain.
    })

    // Define a route that will be cached.
    app.Get("/cache", func(ctx iris.Context) {
        // Simulate a database operation or a heavy computation.
        time.Sleep(2 * time.Second)
        ctx.Writef("Hello from server!")
        // Store the response in the cache.
        cacheConfig.Set(ctx.Request().URL.Path, ctx.ResponseWriter(), nil)
    })

    // Start the server.
    app.Listen(":8080")
}
