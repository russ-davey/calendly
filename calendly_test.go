package calendly

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestGetEventType(t *testing.T) {
	cy := Calendly{}

	os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
	client := NewClient("test")

	Convey("Given the Calendly API event_types endpoint needs to be accessed", t, func() {
		Convey("When the getting the event type details", func() {
			results, err := cy.GetEventType(client, "8ead31de-0033-457a-8646-124e61742999")
			Convey("Then the details are successfully returned", func() {
				So(err, ShouldBeNil)
				So(results.Resource.Name, ShouldEqual, "15 Minute Meeting")
			})
		})
	})
}

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
