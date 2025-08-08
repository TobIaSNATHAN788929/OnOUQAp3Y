// 代码生成时间: 2025-08-09 00:52:23
package main

import (
    "fmt"
    "log"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/cors"
)

// InventoryItem represents an item in the inventory system.
type InventoryItem struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Quantity  int    `json:"quantity"`
    Price     float64 `json:"price"`
}

// InventoryService is a service that manages the inventory items.
type InventoryService struct {
    items []InventoryItem
}

// NewInventoryService creates a new InventoryService with a slice of inventory items.
func NewInventoryService() *InventoryService {
    return &InventoryService{
        items: []InventoryItem{
            {ID: "1", Name: "Widget", Quantity: 10, Price: 2.99},
            {ID: "2", Name: "Gadget", Quantity: 5, Price: 5.99},
        },
    }
}

// GetInventoryItems returns a slice of all inventory items.
func (s *InventoryService) GetInventoryItems() []InventoryItem {
    return s.items
}

// GetInventoryItemByID finds and returns an inventory item by its ID.
func (s *InventoryService) GetInventoryItemByID(id string) (*InventoryItem, error) {
    for i := range s.items {
        if s.items[i].ID == id {
            return &s.items[i], nil
        }
    }
    return nil, fmt.Errorf("inventory item with id %s not found", id)
}

// AddInventoryItem adds a new inventory item.
func (s *InventoryService) AddInventoryItem(item InventoryItem) error {
    for _, existingItem := range s.items {
        if existingItem.ID == item.ID {
            return fmt.Errorf("inventory item with id %s already exists", item.ID)
        }
    }
    s.items = append(s.items, item)
    return nil
}

// UpdateInventoryItem updates an existing inventory item.
func (s *InventoryService) UpdateInventoryItem(id string, newDetails InventoryItem) (*InventoryItem, error) {
    for i := range s.items {
        if s.items[i].ID == id {
            s.items[i] = newDetails
            return &s.items[i], nil
        }
    }
    return nil, fmt.Errorf("inventory item with id %s not found", id)
}

// DeleteInventoryItem deletes an inventory item by its ID.
func (s *InventoryService) DeleteInventoryItem(id string) error {
    for i := range s.items {
        if s.items[i].ID == id {
            s.items = append(s.items[:i], s.items[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("inventory item with id %s not found", id)
}

func main() {
    app := iris.New()
    app.Use(cors.Default())

    inventoryService := NewInventoryService()

    // Inventory List
    app.Get("/inventory", func(ctx iris.Context) {
        items := inventoryService.GetInventoryItems()
        ctx.JSON(iris.StatusOK, items)
    })

    // Get Item by ID
    app.Get("/inventory/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        item, err := inventoryService.GetInventoryItemByID(id)
        if err != nil {
            ctx.StatusCode(iris.StatusNotFound)
            ctx.JSON(iris.StatusInternalServerError, iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(iris.StatusOK, item)
    })

    // Add Item
    app.Post("/inventory", func(ctx iris.Context) {
        var item InventoryItem
        if err := ctx.ReadJSON(&item); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.StatusInternalServerError, iris.Map{
                "error": "invalid item data",
            })
            return
        }
        if err := inventoryService.AddItem(item); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.StatusInternalServerError, iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(iris.StatusOK, iris.Map{
            "message": "item added successfully",
        })
    })

    // Update Item
    app.Put("/inventory/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        var newDetails InventoryItem
        if err := ctx.ReadJSON(&newDetails); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.StatusInternalServerError, iris.Map{
                "error": "invalid item data",
            })
            return
        }
        item, err := inventoryService.UpdateInventoryItem(id, newDetails)
        if err != nil {
            ctx.StatusCode(iris.StatusNotFound)
            ctx.JSON(iris.StatusInternalServerError, iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(iris.StatusOK, item)
    })

    // Delete Item
    app.Delete("/inventory/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        if err := inventoryService.DeleteInventoryItem(id); err != nil {
            ctx.StatusCode(iris.StatusNotFound)
            ctx.JSON(iris.StatusInternalServerError, iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(iris.StatusOK, iris.Map{
            "message": "item deleted successfully",
        })
    })

    // Start the server
    log.Printf("Server is running on http://localhost:8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatal(err)
    }
}
