// 代码生成时间: 2025-08-05 01:29:43
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/kataras/iris/v12" // IRIS framework
)

// DatabaseConfig is a struct to hold database configuration details.
type DatabaseConfig struct {
    Host     string
    Port     int
    Username string
    Password string
    Database string
}

// DatabasePool is a struct to manage the database connection pool.
type DatabasePool struct {
    *sql.DB
    config *DatabaseConfig
}

// NewDatabasePool creates a new database connection pool.
func NewDatabasePool(config *DatabaseConfig) (*DatabasePool, error) {
    // Construct the connection string.
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.Username, config.Password, config.Host, config.Port, config.Database)

    // Open a new connection to the database.
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, err
    }

    // Set the maximum number of connections in the idle connection pool.
    db.SetMaxIdleConns(10)

    // Set the maximum number of open connections to the database.
    db.SetMaxOpenConns(100)

    // Set the connection maximum life time.
    db.SetConnMaxLifetime(5 * time.Minute)

    // Ping the database to verify the connection.
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &DatabasePool{DB: db, config: config}, nil
}

// Close closes the database connection pool.
func (p *DatabasePool) Close() error {
    return p.DB.Close()
}

func main() {
    // Define the database configuration.
    config := &DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        Username: "user",
        Password: "password",
        Database: "iris",
    }

    // Create a new database pool.
    dbPool, err := NewDatabasePool(config)
    if err != nil {
        log.Fatalf("Failed to create database pool: %s", err)
    }
    defer dbPool.Close()

    // Create an IRIS application.
    app := iris.New()

    // Define a route to test the database connection.
    app.Get("/test-db", func(ctx iris.Context) {
        // Use the database pool to perform a query.
        rows, err := dbPool.Query("SELECT 1")
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("Failed to query the database")
            return
        }
        defer rows.Close()

        // Check if there are any rows returned.
        if rows.Next() {
            ctx.WriteString("Database connection is working")
        } else {
            ctx.WriteString("No rows returned")
        }
    })

    // Start the IRIS server.
    log.Fatal(app.Listen(":8080"))
}
