package gorm

import (
	"testing"
	."github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestGetUser(t *testing.T) {

	Convey("a simple test", t, func() {
		user :=GetUserById(4)
		fmt.Printf("user:%v\n",user)
		So(user.ID, ShouldEqual, 4)

	})
}