package controllers

import (
	"testing"

	"github.com/go-resty/resty"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetReport(t *testing.T) {
	Convey("test if you get response for xml", t, func() {
		r, err := resty.R().Get("http://127.0.0.1:8080/report/xml")

		So(err, ShouldBeNil)
		So(r.Status(), ShouldEqual, "200 OK")
		So(r.Header().Get("Content-Type"), ShouldEqual, "application/xml")
	})

	Convey("test with new context for pdf", t, func() {
		r, err := resty.R().Get("http://127.0.0.1:8080/report/pdf")

		So(err, ShouldBeNil)
		So(r.Status(), ShouldEqual, "200 OK")
		So(r.Header().Get("Content-Type"), ShouldEqual, "application/pdf")
	})
}