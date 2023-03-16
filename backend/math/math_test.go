package math_test

import (
	"github.com/batuhannoz/todo/backend/math"
	"testing"
)

func TestAdd(t *testing.T) {
	result := math.Add(3, 2)
	if result != 5 {
		t.Errorf("Expected: %d Recevied: %d", 5, result)
	}
}
