package main

import (
	"fmt"
)

// calculateBMI computes the BMI based on weight (kg) and height (m).
func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

// getBMICategory returns the health category based on BMI.
func getBMICategory(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi >= 18.5 && bmi < 24.9:
		return "Normal weight"
	case bmi >= 25.0 && bmi < 29.9:
		return "Overweight"
	default:
		return "Obesity"
	}
}

func main() {
	var weight, height float64

	fmt.Println("BMI Calculator")
	fmt.Println("=================")
	fmt.Print("Enter your weight in kilograms: ")
	_, errWeight := fmt.Scan(&weight)

	fmt.Print("Enter your height in meters: ")
	_, errHeight := fmt.Scan(&height)

	// Validate input
	if errWeight != nil || errHeight != nil || weight <= 0 || height <= 0 {
		fmt.Println("Error: Please enter valid positive numbers for weight and height.")
		return
	}

	// Calculate BMI
	bmi := calculateBMI(weight, height)
	category := getBMICategory(bmi)

	// Display results
	fmt.Printf("\nYour BMI is: %.2f\n", bmi)
	fmt.Printf("BMI Category: %s\n", category)
}
