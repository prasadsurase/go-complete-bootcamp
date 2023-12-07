package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	type digit [5]string

	zero := digit{
		"000",
		"0 0",
		"0 0",
		"0 0",
		"000",
	}
	one := digit{
		"  0",
		"  0",
		"  0",
		"  0",
		"  0",
	}

	two := digit{
		"000",
		"  0",
		"000",
		"0  ",
		"000",
	}

	three := digit{
		"000",
		"  0",
		"000",
		"  0",
		"000",
	}

	four := digit{
		"0 0",
		"0 0",
		"000",
		"  0",
		"  0",
	}

	five := digit{
		"000",
		"0  ",
		"000",
		"  0",
		"000",
	}

	six := digit{
		"000",
		"0  ",
		"000",
		"0 0",
		"000",
	}
	seven := digit{
		"000",
		"  0",
		"  0",
		"  0",
		"  0",
	}
	eight := digit{
		"000",
		"0 0",
		"000",
		"0 0",
		"000",
	}

	nine := digit{
		"000",
		"0 0",
		"000",
		"  0",
		"000",
	}

	separator := digit{
		"   ",
		" 0 ",
		"   ",
		" 0 ",
		"   ",
	}

	digits := [...]digit{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		time.Sleep(1 * time.Second)

		// Get current time
		ct := time.Now()
		hour, min, sec := ct.Hour(), ct.Minute(), ct.Second()
		fmt.Printf("Hour: %d, Minute: %d, Second: %d\n\n", hour, min, sec)

		// Assemble the clock
		clock := [...]digit{
			digits[hour/10], digits[hour%10],
			separator,
			digits[min/10], digits[min%10],
			separator,
			digits[sec/10], digits[sec%10],
		}

		// Print the clock
		for i := 0; i < 5; i++ {
			for _, digit := range clock {
				fmt.Printf("%v  ", digit[i])
			}
			fmt.Println("")
		}

		time.Sleep(1 * time.Second)
		c.Run()
	}
}
