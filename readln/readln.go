package readln

import (
  "fmt"
  "os"
  "bufio"
	"strconv"
	"strings"
)

func Readln() []float64 {
	// read multiple input arguments
	reader := bufio.NewReader(os.Stdin)
	var nums []float64

	for {

		fmt.Print("Enter number: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			// error handling
			break
		}
		text = strings.TrimSpace(text)

		num, err := strconv.ParseFloat(text, 64)
		if err != nil {
			// error handling
			break
		}

		nums = append(nums, num)
		}

		return nums
}
