package main

import "fmt"

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█░█",
		"█░█",
		"█░█",
		"███",
	}

	one := placeholder{
		"██",
		"░█",
		"░█",
		"░█",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	// This array's type is "like": [10][5]string
	//
	// However:
	// + "placeholder" is not equal to [5]string in type-wise.
	// + Because: "placeholder" is a defined type, which is different
	//   from [5]string type.
	// + [5]string is an unnamed type.
	// + placeholder is a named type.
	// + The underlying type of [5]string and placeholder is the same:
	//     [5]string

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	fmt.Println(digits[0])

}

// Digit character       : █

// Separator character   : ░