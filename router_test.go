package baa

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var r = newRouter()
var c = newContext(nil, nil, nil)
var f = func(c *Context) {}

func TestRouteAdd1(t *testing.T) {
	Convey("测试路由添加", t, func() {
		r.add("GET", "/", []HandlerFunc{f})
		r.add("GET", "/abc", []HandlerFunc{f})
		r.add("GET", "/bcd", []HandlerFunc{f})
		r.add("GET", "/abcd", []HandlerFunc{f})
		r.add("GET", "/abd", []HandlerFunc{f})
		r.add("GET", "/abcdef", []HandlerFunc{f})
		r.add("GET", "/bcdefg", []HandlerFunc{f})
		r.print("", r.routeMap["GET"])
	})
}

func TestRouteAdd2(t *testing.T) {
	Convey("测试参数路由添加", t, func() {
		r.add("GET", "/", []HandlerFunc{f})
		r.add("GET", "/p/:id/id", []HandlerFunc{f})
		r.add("GET", "/p", []HandlerFunc{f})
		r.add("GET", "/p/:id", []HandlerFunc{f})
		r.add("GET", "/a/:ibb", []HandlerFunc{f})
		r.add("GET", "/a/:id/id", []HandlerFunc{f})
		r.add("GET", "/a/:ibb/name", []HandlerFunc{f})
		r.add("GET", "/a/:project/file/:name", []HandlerFunc{f})
		r.add("GET", "/a/", []HandlerFunc{f})
		r.add("GET", "/a/*/xxx", []HandlerFunc{f})
		r.print("", r.routeMap["GET"])
	})
}

func TestRouteAdd3(t *testing.T) {
	Convey("测试组路由添加", t, func() {
		app := New()
		app.SetRouter(r)
		app.Group("/user", func() {
			app.Get("/info", f)
			app.Get("/info2", f)
		})
		app.Group("/user", func() {
			app.Get("/pass", f)
			app.Get("/pass2", f)
		})
		r.print("", r.routeMap["GET"])
	})
}

func TestRoutematch1(t *testing.T) {
	Convey("测试参数路由获取", t, func() {
		ru := r.match("GET", "/", c)
		So(ru, ShouldNotBeNil)

		ru = r.match("GET", "/a/123/id", c)
		So(ru, ShouldNotBeNil)

		ru = r.match("GET", "/a/yst/file/a.jpg", c)
		So(ru, ShouldNotBeNil)

		ru = r.match("GET", "/user/info", c)
		So(ru, ShouldNotBeNil)

		ru = r.match("GET", "/user/pass", c)
		So(ru, ShouldNotBeNil)

		ru = r.match("GET", "/user/xxx", c)
		So(ru, ShouldBeNil)

		ru = r.match("GET", "/xxxx", c)
		So(ru, ShouldBeNil)
	})
}
