package utils

import (
	"errors"
	"log"
	"strconv"
)

// Stop the program with log.Fatal(...) if error isn't nil
func HandleError(err error) {
	if err != nil {
		//TODO panic and recover, do not stop the server. Maybe
		log.Fatal(err)
	}
}

// Convert string to uint
func ParseStringToUIntGT0(number string) (uint, error) {
	value, err := strconv.ParseUint(number, 10, 32)
	if err == nil && value < 1 {
		err = errors.New("less than 1")
	}
	return uint(value), err
}
