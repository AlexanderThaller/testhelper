package testhelper

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestOutput(t *testing.T, l logger.Logger, got, expected interface{}) {
	message := "Did not get the expected output"
	Test(t, l, message, got, expected)
}

func Test(t *testing.T, l logger.Logger, message string, got, expected interface{}) {
	TestErr(t, l, message, errors.New("no error"), got, expected)
}

func TestErrOutput(t *testing.T, l logger.Logger, err error, got, expected interface{}) {
	message := "Did not get the expected output"
	TestErr(t, l, message, err, got, expected)
}

func TestErr(t *testing.T, l logger.Logger, message string, err error, got, expected interface{}) {
	if reflect.DeepEqual(got, expected) {
		return
	}

	l.Error("MESSAGE : ", message)
	if err != nil {
		l.Notice("ERROR: ", errgo.Details(err))
	}
	l.Notice("GOT:\n", got)
	l.Notice("EXPECTED:\n", expected)

	if reflect.TypeOf(got).Name() == "string" {
		differ := diffmatchpatch.New()
		diff := differ.DiffMain(expected.(string), got.(string), true)

		l.Notice("DIFF:")
		for _, line := range diff {
			switch line.Type {
			case diffmatchpatch.DiffDelete:
				fmt.Print("\033[32m" + line.Text + "\033[0m")
			case diffmatchpatch.DiffInsert:
				fmt.Print("\033[31m" + line.Text + "\033[0m")
			default:
				fmt.Print(line.Text)
			}
		}
	}

	t.Fail()
}
