// 代码生成时间: 2025-09-19 22:23:46
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/adaptors/httprouter"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// DatabaseMigrationTool represents the database migration tool
type DatabaseMigrationTool struct {
    db *gorm.DB
}

// NewDatabaseMigrationTool initializes a new instance of DatabaseMigrationTool
func NewDatabaseMigrationTool() *DatabaseMigrationTool {
    // SQLite DB
    db, err := gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    return &DatabaseMigrationTool{db: db}
}

// MigrateUp applies all migrations up to the latest version
func (tool *DatabaseMigrationTool) MigrateUp() error {
    // Implement your migration logic here
    // For example, create tables if they don't exist
    err := tool.db.AutoMigrate(&YourModel{})
    if err != nil {
        return err
    }
    fmt.Println("Migrations applied successfully")
    return nil
}

// MigrateDown applies all migrations down to the first version
func (tool *DatabaseMigrationTool) MigrateDown() error {
    // Implement your migration logic here
    // For example, drop tables
    // Note: GORM does not support down migrations out of the box, this is a custom implementation
    err := tool.db.Migrator().DropTableIfExists(&YourModel{})
    if err != nil {
        return err
    }
    fmt.Println("Migrations rolled back successfully")
    return nil
}

func main() {
    // Initialize Iris
    app := iris.New()
    app.Adapt(httprouter.New())

    // Initialize DatabaseMigrationTool
    migrationTool := NewDatabaseMigrationTool()

    // Routes
    app.Post("/migrate/up", func(ctx iris.Context) {
        err := migrationTool.MigrateUp()
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(iris.Map{
            "message": "Migrations applied successfully",
        })
    })

    app.Post("/migrate/down", func(ctx iris.Context) {
        err := migrationTool.MigrateUp()
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(iris.Map{
            "message": "Migrations rolled back successfully",
        })
    })

    // Start the server
    log.Fatal(app.Listen(":8080"))
}

// YourModel represents the model that will be used for migrations
// This should be replaced with your actual model
type YourModel struct {
    ID uint `gorm:"primaryKey"`
    Name string
}

// TableName overrides the default table name
func (YourModel) TableName() string {
    return "your_table_name"
}
