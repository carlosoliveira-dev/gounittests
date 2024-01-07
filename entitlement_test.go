package main

import (
	"testing"
)

func TestCounter_Take(t *testing.T) {
	tt := []struct {
		name          string
		startingValue int64
		expectedValue int64
		valueToRemove int64
		hasError      bool
	}{
		{
			name:          "one item can be taken successfully",
			startingValue: 10,
			expectedValue: 9,
			valueToRemove: 1,
			hasError:      false,
		},
		{
			name:          "all items can be taken successfully",
			startingValue: 10,
			expectedValue: 10,
			valueToRemove: 0,
			hasError:      false,
		},
		{
			name:          "cannot take more items than what the counter has",
			startingValue: 10,
			expectedValue: 0,
			valueToRemove: 16,
			hasError:      true,
		},
		{
			name:          "cannot take more items than what the counter has",
			startingValue: 10,
			expectedValue: 0,
			valueToRemove: 16,
			hasError:      true,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			c := NewCounter(v.startingValue)

			err := c.TakeN(v.valueToRemove)

			if v.hasError {
				if err == nil {
					t.Fatal("expected an error but got none")
				}

				return
			}

			if c.Value() != v.expectedValue {
				t.Fatalf("Expected %d got %d", v.expectedValue, c.Value())
			}
		})
	}
}
