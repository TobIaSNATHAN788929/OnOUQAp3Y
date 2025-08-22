// 代码生成时间: 2025-08-22 14:54:35
package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"

    "github.com/kataras/iris/v12"
)

// MathTool struct encapsulates the necessary data for the math operations
type MathTool struct {
    // no fields needed for now
}

// Calculate performs the calculation based on the provided operation and operands
func (m *MathTool) Calculate(operation string, operands ...float64) (float64, error) {
    switch operation {
    case "add":
        return add(operands...), nil
    case "subtract":
        return subtract(operands...), nil
    case "multiply":
        return multiply(operands...), nil
    case "divide":
        return divide(operands...), nil
    case "sqrt":
        if len(operands) != 1 {
            return 0, fmt.Errorf("sqrt requires exactly one operand")
        }
        return math.Sqrt(operands[0]), nil
    default:
        return 0, fmt.Errorf("unsupported operation: %s", operation)
    }
}

// add performs the addition of all the operands
func add(operands ...float64) float64 {
    sum := 0.0
    for _, operand := range operands {
        sum += operand
    }
    return sum
}

// subtract performs the subtraction of all the operands
func subtract(operands ...float64) float64 {
    result := operands[0]
    for _, operand := range operands[1:] {
        result -= operand
    }
    return result
}

// multiply performs the multiplication of all the operands
func multiply(operands ...float64) float64 {
    product := 1.0
    for _, operand := range operands {
        product *= operand
    }
    return product
}

// divide performs the division of all the operands
func divide(operands ...float64) float64 {
    result := operands[0]
    for _, operand := range operands[1:] {
        if operand == 0 {
            return 0
        }
        result /= operand
    }
    return result
}

func main() {
    app := iris.New()
    mathTool := MathTool{}

    app.Get("/calculate", func(ctx iris.Context) {
        operation := ctx.URLParam("operation")
        operandsStr := ctx.URLParams().Get("operands")

        if operation == "" || operandsStr == "" {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "Missing operation or operands",
            })
            return
        }

        // Split the operands string by comma and convert each to float64
        operands := strings.Split(operandsStr, ",")
        var operandsFloat []interface{}
        for _, operandStr := range operands {
            operand, err := strconv.ParseFloat(operandStr, 64)
            if err != nil {
                ctx.StatusCode(iris.StatusBadRequest)
                ctx.JSON(iris.Map{
                    "error": "Invalid operand value",
                })
                return
            }
            operandsFloat = append(operandsFloat, operand)
        }

        // Perform the calculation
        result, err := mathTool.Calculate(operation, operandsFloat...)
        if err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        // Return the result
        ctx.JSON(iris.Map{
            "result": result,
        })
    })

    // Handle server error
    app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "error": "An internal server error occurred",
        })
    })

    // Start the iris server
    app.Listen(":8080")
}
