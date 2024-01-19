package utils

import (
	"errors"
	"log"
	"strconv"
)

func HandleError(err error) {
	if err != nil {
		//TODO panic and recover, do not stop the server
		log.Fatal(err)
	}
}

func ParseStringToUIntGT0(number string) (uint, error) {
	value, err := strconv.ParseUint(number, 10, 32)
	if err == nil && value < 1 {
		err = errors.New("less than 1")
	}
	return uint(value), err
}
