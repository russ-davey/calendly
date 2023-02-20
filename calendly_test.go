package calendly

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMarshellAPIError(t *testing.T) {
	Convey("Given an error has been returned by the Calendly API", t, func() {
		Convey("When the error is unmarshalled", func() {
			jsonStr := "{\n  \"title\": \"Internal Server Error\",\n  \"message\": \"The server encountered an unexpected condition that prevented it from fulfilling the request.\",\n  \"details\": [\n    {\n      \"parameter\": \"string\",\n      \"message\": \"string\"\n    }\n  ]\n}"
			results := UnmarshallAPIError(errors.New(jsonStr))
			Convey("Then the error message is successfully unmarshalled", func() {
				So(results.Title, ShouldEqual, "Internal Server Error")
				So(results.Message, ShouldEqual, "The server encountered an unexpected condition that prevented it from fulfilling the request.")
			})
		})
		Convey("When a different error is unmarshalled", func() {
			results := UnmarshallAPIError(errors.New("unknown error"))
			Convey("Then the error message is successfully unmarshalled but the values are empty", func() {
				So(results.Title, ShouldEqual, "")
			})
		})
	})
}
