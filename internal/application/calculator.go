package application

import (
	"math"

	"goalculator/internal/models"
	"goalculator/internal/utils"

	"gorm.io/gorm"
)

type Calculator struct {
	dB *gorm.DB
}

func NewCalculator(dB *gorm.DB) *Calculator {
	return &Calculator{dB}
}

func (c *Calculator) Add(a float64, b float64) (float64, error) {
	return a + b, nil
}

func (c *Calculator) Subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func (c *Calculator) Multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, utils.NewErr("Division by zero", 400)
	}
	return a / b, nil
}

func (c *Calculator) Power(a, b float64) (float64, error) {
	return math.Pow(a, b), nil
}

func (c *Calculator) PerformOperation(a, b float64, operation string) (float64, error) {
	switch operation {
	case "add":
		result, err := c.Add(a, b)
		if err != nil {
			return 0, err
		}
		c.dB.Create(&models.Operation{Operation: operation, A: a, B: b, Result: result, Message: "Operation performed successfully"})
		return result, nil
	case "subtract":
		result, err := c.Subtract(a, b)
		if err != nil {
			return 0, err
		}
		c.dB.Create(&models.Operation{Operation: operation, A: a, B: b, Result: result, Message: "Operation performed successfully"})
		return result, nil
	case "multiply":
		result, err := c.Multiply(a, b)
		if err != nil {
			return 0, err
		}
		c.dB.Create(&models.Operation{Operation: operation, A: a, B: b, Result: result, Message: "Operation performed successfully"})
		return result, nil
	case "divide":
		result, err := c.Divide(a, b)
		if err != nil {
			return 0, err
		}
		c.dB.Create(&models.Operation{Operation: operation, A: a, B: b, Result: result, Message: "Operation performed successfully"})
		return result, nil
	case "power":
		result, err := c.Power(a, b)
		if err != nil {
			return 0, err
		}
		c.dB.Create(&models.Operation{Operation: operation, A: a, B: b, Result: result, Message: "Operation performed successfully"})
		return result, nil
	default:
		c.dB.Create(&models.Operation{Operation: operation, A: a, B: b, Result: 0, Message: "Invalid operation"})
		return 0, utils.NewErr("Invalid operation", 400)
	}
}
