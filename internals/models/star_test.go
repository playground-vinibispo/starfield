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
	star := models.NewStar()
	firstZ := star.Z
	star.Update(1)
	if star.Z == firstZ {
		t.Error("Expected star Z to be different from Pz")
	}
}
