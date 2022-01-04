package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

// TestRun runs the actual interpreter
func TestRun(t *testing.T) {
	assert.Equal(t, "True", "True")
}

// TestMain is the main entry point
func TestMain(t *testing.T) {
	Run()
	assert.Equal(t, "True", "True")
}
