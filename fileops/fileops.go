package fileops

import (
	"errors"
	"fmt"
	"os"
)

func WriteFloatToFile(floatValue float64, file string) {
	os.WriteFile(file, []byte(fmt.Sprint(floatValue)), 0644)
}

func ReadFloatFromFile(file string) (floatValue float64, err error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return 0, errors.New("error reading from file")
	}

	fmt.Sscan(string(data), &floatValue)

	return floatValue, err
}
