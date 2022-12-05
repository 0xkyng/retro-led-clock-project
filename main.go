package main

import (
	"fmt"
	"time"
	"github.com/inancgumus/screen"
)

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

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {
		screen.MoveTopLeft()
		now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	clock := [...]placeholder{
		digits[hour/10], digits[hour%10],
		colon,
		digits[min/10], digits[min%10],
		colon,
		digits[sec/10], digits[sec%10],
	}

	for line := range clock[0] {
		for index, digit := range clock {
			next := clock[index][line]
			if digit == colon && sec%2 == 0 {
				next = "   "
			}
			fmt.Print(next,"  ")
		}
		fmt.Println()
	}
	fmt.Println()
	time.Sleep(time.Second)

	}
}

// Digit character       : █

// Separator character   : ░
