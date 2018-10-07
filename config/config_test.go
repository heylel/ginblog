package config_test

import (
	"github.com/gq-tang/ginblog/test"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestString(t *testing.T) {
	convey.Convey("test string", t, func() {
		c, err := test.LoadConfig()
		convey.So(err, convey.ShouldBeNil)
		val, err := c.String("globaltitle")
		convey.So(err, convey.ShouldBeNil)
		convey.So(val, convey.ShouldEqual, "阿汤哥的博客")
	})
}
