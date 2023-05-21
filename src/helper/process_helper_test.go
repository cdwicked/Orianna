package helper

import (
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestProcess(t *testing.T) {
	Convey("Process test", t, func() {
		Convey("test error", func() {
			var err error
			Process("test Do", func() {
				err = Do()
			})
			So(err, ShouldBeError)
		})
		Convey("test with param", func() {
			var err error
			Process("test DoParam", func() {
				err = DoParam("hello")
			})
			So(err, ShouldBeNil)
		})
	})
}

func Do() error {
	return errors.New("error")
}

func DoParam(hello string) error {
	log.Print(hello)
	return nil
}
