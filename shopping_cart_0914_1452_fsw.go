// 代码生成时间: 2025-09-14 14:52:06
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/kataras/iris/v12"
)

// ShoppingCart represents a shopping cart entity
type ShoppingCart struct {
    ID        string       `json:"id"`
    Items     []CartItem   `json:"items"`
    TotalCost float64      `json:"totalCost"`
}

// CartItem represents an item in the shopping cart
type CartItem struct {
    ProductID string  `json:"productID"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

// NewShoppingCart creates a new instance of ShoppingCart
func NewShoppingCart(id string) *ShoppingCart {
    return &ShoppingCart{ID: id}
}

// AddItem adds an item to the shopping cart
func (cart *ShoppingCart) AddItem(productID string, quantity int, price float64) error {
    // Check if the item already exists in the cart
    for _, item := range cart.Items {
        if item.ProductID == productID {
            return fmt.Errorf("product with ID %s already exists in the cart", productID)
        }
    }
    
    // Add new item to the cart
    cart.Items = append(cart.Items, CartItem{ProductID: productID, Quantity: quantity, Price: price})
    return nil
}

// CalculateTotalCost calculates the total cost of items in the cart
func (cart *ShoppingCart) CalculateTotalCost() {
    cart.TotalCost = 0
    for _, item := range cart.Items {
        cart.TotalCost += float64(item.Quantity) * item.Price
    }
}

// Handler for adding items to the cart
func addItemToCart(ctx iris.Context) {
    id := ctx.URLParam("id")
    productID := ctx.URLParam("productID")
    quantity := ctx.URLParamDefault("quantity", "1")
    priceStr := ctx.URLParam("price")
    price, err := strconv.ParseFloat(priceStr, 64)
    if err != nil {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.JSON(iris.Map{"error": "Invalid price format"})
        return
    }
    
    // Assume cart is retrieved from a store or database
    // For simplicity, we'll use a mock cart
    cart := NewShoppingCart(id)
    
    if err := cart.AddItem(productID, parseInt(quantity, 10, 0), price); err != nil {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.JSON(iris.Map{"error": err.Error()})
        return
    }
    
    cart.CalculateTotalCost()
    
    ctx.JSON(cart)
}

func main() {
    app := iris.New()
    app.Logger().SetLevel("github.com/kataras/iris/v12/logger".LevelInfo)

    // Define route for adding items to the cart
    app.Post("/cart/{id}/add", addItemToCart)

    // Start the Iris server
    log.Fatal(app.Listen(":8080"))
}

// parseInt is a helper function to parse a string to an integer with a base
func parseInt(str string, base int, bitSize int) int {
    result, err := strconv.Atoi(str)
    if err != nil {
        return 0
    }
    return result
}
