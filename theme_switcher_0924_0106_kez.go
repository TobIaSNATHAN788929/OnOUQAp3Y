// 代码生成时间: 2025-09-24 01:06:34
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// Theme represents the theme settings of the application
type Theme struct {
    Name string
}

// ThemeSwitcher handles theme switching functionality
type ThemeSwitcher struct {
    CurrentTheme *Theme
}

// NewThemeSwitcher initializes a new ThemeSwitcher with a default theme
func NewThemeSwitcher() *ThemeSwitcher {
    return &ThemeSwitcher{
        CurrentTheme: &Theme{Name: "light"},
    }
}

// SwitchTheme changes the current theme
func (ts *ThemeSwitcher) SwitchTheme(themeName string) error {
    if themeName == "" {
        return fmt.Errorf("theme name cannot be empty")
    }
    // Assume we have a predefined list of themes
    validThemes := []string{"light", "dark", "colorful"}
    if contains(validThemes, themeName) {
        ts.CurrentTheme.Name = themeName
        return nil
    }
    return fmt.Errorf("theme '%s' is not supported", themeName)
}

// contains checks if a slice contains a certain string
func contains(s []string, str string) bool {
    for _, v := range s {
        if v == str {
            return true
        }
    }
    return false
}

func main() {
    app := iris.New()
    
    // Initialize theme switcher with default theme
    themeSwitcher := NewThemeSwitcher()
    
    // Handle GET request to switch the theme
    app.Get("/switch_theme/{themeName}", func(ctx iris.Context) {
        themeName := ctx.Param("themeName")
        
        // Switch theme and handle errors
        if err := themeSwitcher.SwitchTheme(themeName); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        
        // On successful theme switch, inform the user
        ctx.JSON(iris.Map{
            "message": "Theme switched to " + themeName,
        })
    })
    
    // Handle GET request to get the current theme
    app.Get("/current_theme", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "current_theme": themeSwitcher.CurrentTheme.Name,
        })
    })
    
    // Start the Iris server
    app.Listen(":8080")
}