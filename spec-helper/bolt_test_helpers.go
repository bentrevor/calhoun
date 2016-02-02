package spec_helper

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// taken from https://github.com/boltdb/bolt/blob/master/bolt_test.go

// assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\n\033[31m%s:%d: failure!!!\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	} else {
		fmt.Printf("\033[32m.\033[39m")
	}
}

// equals fails the test if exp is not equal to act.
func AssertEquals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\n\033[31m%s:%d:\n\n\texpected: %#v\n\n\tto eq: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	} else {
		fmt.Printf("\033[32m.\033[39m")
	}
}

func AssertNotEquals(tb testing.TB, exp, act interface{}) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\n\033[31m%s:%d:\n\n\texpected: %#v\n\n\tnot to eq: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	} else {
		fmt.Printf("\033[32m.\033[39m")
	}
}

func Describe(desc string) {
	underline := "  "

	for _ = range desc {
		underline += "-"
	}

	fmt.Print(fmt.Sprintf("\n\n\n| %s |\n%s", desc, underline))
}

func It(desc string) {
	fmt.Print(fmt.Sprintf("\n  %s: ", desc))
}

func Specify(desc string) {
	It(desc)
}
