// 代码生成时间: 2025-08-10 20:57:17
package main

import (
    "fmt"
    "math"
    
    "github.com/kataras/iris/v12"
)

// MathService defines the structure for the math operations
type MathService struct{}

// Add handles the addition operation
func (s *MathService) Add(ctx iris.Context) {
	var params struct {
		A, B float64 `json:"a,string"`
	}
	
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error reading parameters: %s
", err))
		return
	}
	
	result := params.A + params.B
	
	ctx.JSON(iris.StatusOK, iris.Map{
		"result": result,
	})
}

// Subtract handles the subtraction operation
func (s *MathService) Subtract(ctx iris.Context) {
	var params struct {
		A, B float64 `json:"a,string"`
	}
	
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error reading parameters: %s
", err))
		return
	}
	
	result := params.A - params.B
	
	ctx.JSON(iris.StatusOK, iris.Map{
		"result": result,
	})
}

// Multiply handles the multiplication operation
func (s *MathService) Multiply(ctx iris.Context) {
	var params struct {
		A, B float64 `json:"a,string"`
	}
	
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error reading parameters: %s
", err))
		return
	}
	
	result := params.A * params.B
	
	ctx.JSON(iris.StatusOK, iris.Map{
		"result": result,
	})
}

// Divide handles the division operation
func (s *MathService) Divide(ctx iris.Context) {
	var params struct {
		A, B float64 `json:"a,string"`
	}
	
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error reading parameters: %s
", err))
		return
	}
	
	if params.B == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Error: cannot divide by zero
")
		return
	}
	
	result := params.A / params.B
	
	ctx.JSON(iris.StatusOK, iris.Map{
		"result": result,
	})
}

func newApp() *iris.Application {
	app := iris.New()

	mathService := &MathService{}

	app.Post("/add", mathService.Add)
	app.Post("/subtract", mathService.Subtract)
	app.Post("/multiply", mathService.Multiply)
	app.Post("/divide", mathService.Divide)

	return app
}

func main() {
	app := newApp()
	
	// Start the IRIS server
	app.Run(iris.Addr(":8080"))
}