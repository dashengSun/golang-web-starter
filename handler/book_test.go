package handler

import (
	"testing"
	"golang-web-starter/utility"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"golang-web-starter/models"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
	"encoding/xml"
	"strings"
)

func TestShowIndexPage(t *testing.T) {
	r := utility.GetRouter(true)
	r.GET("/", ShowIndexPage)

	Convey("Do not set the request header", t, func() {
		req, _ := http.NewRequest("GET", "/", nil)
		utility.TestHTTPResponse(r, req, func (w *httptest.ResponseRecorder) {
			p, _ := ioutil.ReadAll(w.Body)
			So(w.Code, ShouldEqual, http.StatusOK)
			So(string(p), ShouldContainSubstring, "<title>Hi</title>")
		})
	})

	Convey("Set the request header with application/json", t, func() {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Add("Accept", "application/json")

		utility.TestHTTPResponse(r, req, func(w *httptest.ResponseRecorder) {
			p, _ := ioutil.ReadAll(w.Body)
			var books []models.Book
			err := json.Unmarshal(p, &books)
			So(w.Code, ShouldEqual, http.StatusOK)
			So(err, ShouldBeNil)
		})
	})

	Convey("Set the request header with application/xml", t, func() {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Add("Accept", "application/xml")

		utility.TestHTTPResponse(r, req, func(w *httptest.ResponseRecorder) {
			p, _ := ioutil.ReadAll(w.Body)
			var books []models.Book
			err := xml.Unmarshal(p, &books)
			So(w.Code, ShouldEqual, http.StatusOK)
			So(err, ShouldBeNil)
		})
	})
}

func TestGetBook(t *testing.T) {
	r := utility.GetRouter(true)
	r.GET("/book/:book_id", GetBook)

	Convey("The book id is not integer", t, func() {
		req, _ := http.NewRequest("GET", "/book/a", nil)

		utility.TestHTTPResponse(r, req, func(w *httptest.ResponseRecorder) {
			So(w.Code, ShouldEqual, http.StatusBadRequest)
		})
	})

	Convey("The book exists", t, func() {
		req, _ := http.NewRequest("GET", "/book/10", nil)

		utility.TestHTTPResponse(r, req, func(w *httptest.ResponseRecorder) {
			p, _ := ioutil.ReadAll(w.Body)
			So(w.Code, ShouldEqual, http.StatusOK)
			So(string(p), ShouldContainSubstring, "<title>Book</title>")
		})
	})
}

func TestSaveBook(t *testing.T) {
	r := utility.GetRouter(true)
	r.POST("/book", SaveBook)

	Convey("The params can not convert to model book", t, func() {
		req, _ := http.NewRequest("POST", "/book", nil)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		utility.TestHTTPResponse(r, req, func(w *httptest.ResponseRecorder) {
			So(w.Code, ShouldEqual, http.StatusBadRequest)
		})
	})

	Convey("The params can convert to model book", t, func() {
		req, _ := http.NewRequest("POST", "/book", strings.NewReader("title=Hello world&author=will"))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		utility.TestHTTPResponse(r, req, func(w *httptest.ResponseRecorder) {
			p, _ := ioutil.ReadAll(w.Body)
			So(w.Code, ShouldEqual, http.StatusOK)
			So(string(p), ShouldContainSubstring, "保存成功")
		})
	})
}
