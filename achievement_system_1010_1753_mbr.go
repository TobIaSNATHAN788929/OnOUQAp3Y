// 代码生成时间: 2025-10-10 17:53:57
// achievement_system.go

package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12" // Iris框架
)

// Achievement represents a single achievement.
type Achievement struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Desc     string `json:"desc"`
    Points   int    `json:"points"`
    EarnedBy int    `json:"earned_by"`
}

// achievementStorage is a mock storage for achievements.
var achievementStorage = []Achievement{
    {
        ID:       "ach1",
        Title:    "First Login",
        Desc:     "Log in for the first time.",
        Points:   100,
        EarnedBy: 0,
    },
    {
        ID:       "ach2",
        Title:    "Daily Login",
        Desc:     "Log in everyday for a week.",
        Points:   200,
        EarnedBy: 0,
    },
    // ... more achievements can be added here
}

// Handler functions

// GetAchievements handles GET requests to retrieve all achievements.
func GetAchievements(ctx iris.Context) {
    achievements := make([]Achievement, len(achievementStorage))
    copy(achievements, achievementStorage)
    ctx.JSON(http.StatusOK, achievements)
}

// GetAchievement handles GET requests to retrieve a specific achievement by ID.
func GetAchievement(ctx iris.Context) {
    id := ctx.Param("id")
    for _, a := range achievementStorage {
        if a.ID == id {
            ctx.JSON(http.StatusOK, a)
            return
        }
    }
    ctx.StatusCode(http.StatusNotFound)
    ctx.Text("Achievement not found.")
}

// EarnAchievement handles POST requests to mark an achievement as earned by a user.
func EarnAchievement(ctx iris.Context) {
    id := ctx.Param("id")
    for i, a := range achievementStorage {
        if a.ID == id {
            achievementStorage[i].EarnedBy++
            ctx.JSON(http.StatusOK, achievementStorage[i])
            return
        }
    }
    ctx.StatusCode(http.StatusNotFound)
    ctx.Text("Achievement not found.")
}

// main function to setup the Iris server and routes.
func main() {
    app := iris.New()
    
    // Define routes for achievement management.
    app.Get("/achievements", GetAchievements)
    app.Get("/achievements/{id}", GetAchievement)
    app.Post("/achievements/{id}/earn", EarnAchievement)
    
    // Start the Iris server.
    log.Printf("Server is running on :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatal("Error starting server:", err)
    }
}