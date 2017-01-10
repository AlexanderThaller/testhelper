package testhelper

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/juju/errgo"
	"github.com/sergi/go-diff/diffmatchpatch"

	log "github.com/Sirupsen/logrus"
)

func TestOutput(t *testing.T, got, expected interface{}) {
	message := "Did not get the expected output"
	Test(t, message, got, expected)
}

func Test(t *testing.T, message string, got, expected interface{}) {
	TestErr(t, message, errors.New("no error"), got, expected)
}

func TestErrOutput(t *testing.T, err error, got, expected interface{}) {
	message := "Did not get the expected output"
	TestErr(t, message, err, got, expected)
}

func TestErr(t *testing.T, message string, err error, got, expected interface{}) {
	if reflect.DeepEqual(got, expected) {
		return
	}

	log.Error("MESSAGE: ", message)
	if err != nil {
		log.Error("ERROR: ", errgo.Details(err))
	}
	log.Info("GOT:\n", got)
	fmt.Println("")

	log.Info("EXPECTED:\n", expected)

	differ := diffmatchpatch.New()
	diff := differ.DiffMain(fmt.Sprintf("%+v", expected),
		fmt.Sprintf("%+v", got), true)

	fmt.Println("")

	log.Info("DIFF:")
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
	fmt.Println("")
	fmt.Println("")

	t.Fail()
}
