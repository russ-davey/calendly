package calendly

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestGetEventType(t *testing.T) {
	Convey("Given a client and a UUID", t, func() {
		os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
		client := NewClient("test")

		Convey("When the Get function is called", func() {
			results, err := client.EventTypes.Get("8ead31de-0033-457a-8646-124e61742999")

			Convey("Then the event type details are returned", func() {
				So(err, ShouldBeNil)
				So(results.Resource.Name, ShouldEqual, "15 Minute Meeting")
			})
		})
	})
}
