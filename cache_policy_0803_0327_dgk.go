// 代码生成时间: 2025-08-03 03:27:32
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/cache"
)

// CacheService is a struct that will hold the cache configuration.
type CacheService struct {
    // Cache is the cache instance provided by Iris.
    Cache *cache.Cache
}

// NewCacheService initializes a new cache service with a 5-minute expiry time for each item.
func NewCacheService() *CacheService {
    return &CacheService{
        Cache: cache.New(cache.Config{
            DefaultExpiration: 5 * time.Minute, // 5 minutes
            // You can add more cache settings here as needed.
        })},
}

// Get will attempt to retrieve a value from the cache with the given key.
// If the item is not found or expired, it returns an error.
func (cs *CacheService) Get(key string, out interface{}) error {
    if err := cs.Cache.Get(key, out); err != nil {
        if err == cache.ErrEntryNotFound {
            return fmt.Errorf("item with key '%s' not found", key)
        }
        return err
    }
    return nil
}

// Set will store the given key-value pair in the cache with the default expiration time.
func (cs *CacheService) Set(key string, value interface{}) error {
    if err := cs.Cache.Set(key, value, cache.DefaultExpiration); err != nil {
        return err
    }
    return nil
}

// main is the entry point of the application.
func main() {
    app := iris.New()
    // Create a new cache service.
    cacheService := NewCacheService()

    // Define a GET route to demonstrate caching.
    app.Get("/cache", func(ctx iris.Context) {
        key := "myCachedValue"
        var value string

        // Try to get the cached value.
        if err := cacheService.Get(key, &value); err != nil {
            // If there's an error or the item is not found, set a new value.
            value = "Hello, Cache!"
            if err := cacheService.Set(key, value); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                return
            }
        }

        // Return the cached (or newly set) value.
        ctx.JSON(iris.Map{
            "cachedValue": value,
        })
    })

    // Start the server.
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Failed to start the server: %s
", err)
    }
}
