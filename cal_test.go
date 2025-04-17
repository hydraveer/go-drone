package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	// Define a struct for test cases
	tests := []struct {
		name          string
		input         string
		expectedValue float64
		expectError   bool
	}{
		// Addition tests
		{"Simple Addition", "5 + 3", 8.0, false},
		{"Addition With Decimals", "2.5 + 3.5", 6.0, false},
		{"Addition With Negative", "-5 + 3", -2.0, false},
		
		// Subtraction tests
		{"Simple Subtraction", "10 - 4", 6.0, false},
		{"Subtraction With Decimals", "5.5 - 2.5", 3.0, false},
		{"Subtraction Going Negative", "3 - 8", -5.0, false},
		
		// Multiplication tests
		{"Simple Multiplication", "4 * 5", 20.0, false},
		{"Multiplication With Decimals", "2.5 * 4", 10.0, false},
		{"Multiplication With Zero", "5 * 0", 0.0, false},
		
		// Division tests
		{"Simple Division", "20 / 4", 5.0, false},
		{"Division With Decimals", "5 / 2", 2.5, false},
		{"Division By Zero", "5 / 0", 0.0, true},
		
		// Modulo tests
		{"Simple Modulo", "7 % 3", 1.0, false},
		{"Modulo With Zero", "0 % 5", 0.0, false},
		{"Modulo By Zero", "5 % 0", 0.0, true},
		
		// Error cases
		{"No Operator", "123", 0.0, true},
		{"Invalid First Number", "abc + 5", 0.0, true},
		{"Invalid Second Number", "5 + abc", 0.0, true},
		{"Too Many Operators", "5 + 5 + 5", 0.0, true},
		{"Multiple Different Operators", "5 + 5 * 5", 0.0, true},
	}

	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculate(test.input)
			
			// Check if we expect an error
			if test.expectError {
				if err == nil {
					t.Errorf("Expected an error for input '%s', but got result %g", test.input, result)
				}
			} else {
				// We expect success
				if err != nil {
					t.Errorf("Unexpected error for input '%s': %v", test.input, err)
				}
				
				if result != test.expectedValue {
					t.Errorf("For input '%s', expected %g but got %g", test.input, test.expectedValue, result)
				}
			}
		})
	}
}

// Additional test for edge cases
func TestCalculateEdgeCases(t *testing.T) {
	// Test large numbers
	t.Run("Large Numbers", func(t *testing.T) {
		result, err := calculate("999999999 + 1")
		if err != nil {
			t.Errorf("Unexpected error with large numbers: %v", err)
		}
		if result != 1000000000 {
			t.Errorf("Expected 1000000000, got %g", result)
		}
	})
	
	// Test extra whitespace
	t.Run("Extra Whitespace", func(t *testing.T) {
		result, err := calculate("  5    +    5  ")
		if err != nil {
			t.Errorf("Unexpected error with extra whitespace: %v", err)
		}
		if result != 10 {
			t.Errorf("Expected 10, got %g", result)
		}
	})
	
	// Test empty input
	t.Run("Empty Input", func(t *testing.T) {
		_, err := calculate("")
		if err == nil {
			t.Error("Expected error for empty input, but got none")
		}
	})
}