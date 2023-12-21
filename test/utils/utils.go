package utils_test

import (
	"fmt"
	"testing"
)

func Success(t *testing.T, message string, args ...any) {
	t.Logf("\t%s\t%s", Checkmark, fmt.Sprintf(message, args...))
}

func Failure(t *testing.T, message string, args ...any) {
	t.Fatalf("\t%s\t%s", Cross, fmt.Sprintf(message, args...))
}
