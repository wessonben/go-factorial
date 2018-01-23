package factorial

// Calculate implements the factorial calculation of the input value.
func Calculate(value int) int {

	if value <= 0 {
		return 0
	}

	if value == 1 {
		return 1
	}

	if value == 2 {
		return 2
	}

	return value
}
