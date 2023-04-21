package calendly

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestGetUser(t *testing.T) {
	Convey("Given a client and a UUID", t, func() {
		os.Setenv("CALENDLY_MOCK_SERVER", "https://stoplight.io/mocks/calendly/api-docs/395")
		client := NewClient("test")

		Convey("When the GetUsers function is called", func() {
			results, err := client.Users.Get("f59e9bf9-67a5-4df4-a287-5aab29214269")

			Convey("Then the user details are returned", func() {
				So(err, ShouldBeNil)
				So(results.Resource.Name, ShouldEqual, "John Doe")
			})
		})
	})
}
