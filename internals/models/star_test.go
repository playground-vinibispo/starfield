package models_test

import (
	"starfield/internals/models"
	"testing"
)

func TestNewStar(t *testing.T) {
	star := models.NewStar()
	if star == nil {
		t.Error("Expected star to be created")
	}
}

func TestStarUpdate(t *testing.T) {
	star := models.Star{X: 0, Y: 0, Z: 100, Pz: 100}
	firstZ := star.Z
	star.Update(1)
	if star.Z == firstZ {
		t.Error("Expected star Z to be different from Pz")
	}
}
