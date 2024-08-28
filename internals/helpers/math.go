package helpers

func MapValue(value, start1, stop1, start2, stop2 float32) float32 {
	newVal := (value-start1)/(stop1-start1)*(stop2-start2) + start2
	if start2 < stop2 {
		return Constrain(newVal, start2, stop2)
	} else {
		return Constrain(newVal, stop2, start2)
	}
}

func Constrain(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
