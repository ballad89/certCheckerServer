package libhelper

import (
	"testing"
)

// func TestNewTime(t *testing.T) {
// 	ti, err := NewTime(12, 11, 2009)

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(ti)
// }

func TestGetEnvOrDefault(t *testing.T) {
	v := GetEnvOrDefault("testing", "gimmicks")

	if v != "gimmicks" {
		t.Error("not using default value")
	}
}
