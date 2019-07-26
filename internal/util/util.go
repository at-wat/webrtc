package util

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"testing"
	"time"
)

// RandSeq generates a random alpha numeric sequence of the requested length
func RandSeq(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

// FlattenErrs flattens multiple errors into one
func FlattenErrs(errs []error) error {
	var errstrings []string

	for _, err := range errs {
		if err != nil {
			errstrings = append(errstrings, err.Error())
		}
	}

	if len(errstrings) == 0 {
		return nil
	}

	return fmt.Errorf(strings.Join(errstrings, "\n"))
}

func NewLeakTester(t *testing.T, name string, p interface{}) func() {
	finalized := false
	runtime.SetFinalizer(p, func(interface{}) {
		finalized = true
	})
	return func() {
		for i := 0; i < 5; i++ {
			runtime.GC()
			if finalized {
				return
			}
			time.Sleep(time.Second)
		}
		t.Errorf("%s: object is not finalized", name)
	}
}
