package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(valueFloat float64, fileName string) {
	floatText := fmt.Sprint(valueFloat)
	os.WriteFile(fileName, []byte(floatText), 0644)

}

func GetFloatFromFile(accountBalanaceFileName string) (float64, error) {
	data, err := os.ReadFile(accountBalanaceFileName)

	if err != nil {
		return 1000, errors.New("unable to find the file")

	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("issue when parsing balance value to float")

	}
	return balance, nil
}
