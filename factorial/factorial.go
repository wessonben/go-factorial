package factorial

// Calculate implements the factorial calculation of the input value.
func Calculate(value int) int {

	if value <= 0 {
		return 0
	}

	if value <= 2 {
		return value
	}

	factorial := value

	for i := value - 1; i > 1; i-- {
		factorial *= i
	}

	return factorial
}
