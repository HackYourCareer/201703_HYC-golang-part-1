package myPackage

import (
	"testing"
	myname "strings"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMe(t *testing.T) {
	Convey("'ala ma kota' should", t, func() {
		testString := "ala ma kota"

		Convey("contain 'kota", func() {
			zawiera := myname.Contains(testString, "kota")
			So(zawiera, ShouldBeTrue)
		})
	})
}