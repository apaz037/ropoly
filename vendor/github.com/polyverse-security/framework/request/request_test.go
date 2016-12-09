package request

import (
	. "gopkg.in/check.v1"
	"net/url"
	"testing"
)

func TestBackend(t *testing.T) { TestingT(t) }

type BackendSuite struct {
}

var _ = Suite(&BackendSuite{})

func (b *BackendSuite) TestRequestTextUnmarshall(c *C) {
	requestText := `{"Method":"GET","URL":{"Scheme":"","Opaque":"","User":null,"Host":"","Path":"/","RawQuery":"","Fragment":""},"Proto":"HTTP/1.1","ProtoMajor":1,"ProtoMinor":1,"Header":{"Accept":["text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"],"Accept-Encoding":["gzip, deflate"],"Accept-Language":["en-us"],"Cache-Control":["max-age=0"],"Connection":["keep-alive"],"Referer":["http://polyverse.dockerhost:8080/"],"User-Agent":["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/600.4.10 (KHTML, like Gecko) Version/8.0.4 Safari/600.4.10"]},"Body":{"Closer":{"Reader":null}},"ContentLength":0,"TransferEncoding":null,"Close":false,"Host":"polyverse.dockerhost:8080","Form":null,"PostForm":null,"MultipartForm":null,"Trailer":null,"RemoteAddr":"192.168.59.3:49310","RequestURI":"/","TLS":null}`
	request := FromJson(requestText)

	c.Assert(request.Method, Equals, "GET")
	c.Assert(request.URL.Path, Equals, "/")
	c.Assert(request.Host, Equals, "polyverse.dockerhost:8080")
}

func (b *BackendSuite) TestRequestBytesUnmarshall(c *C) {
	requestBytes := []byte{123, 34, 77, 101, 116, 104, 111, 100, 34, 58, 34, 71, 69, 84, 34, 44, 34, 85, 82, 76, 34, 58, 123, 34, 83, 99, 104, 101, 109, 101, 34, 58, 34, 34, 44, 34, 79, 112, 97, 113, 117, 101, 34, 58, 34, 34, 44, 34, 85, 115, 101, 114, 34, 58, 110, 117, 108, 108, 44, 34, 72, 111, 115, 116, 34, 58, 34, 34, 44, 34, 80, 97, 116, 104, 34, 58, 34, 47, 34, 44, 34, 82, 97, 119, 81, 117, 101, 114, 121, 34, 58, 34, 34, 44, 34, 70, 114, 97, 103, 109, 101, 110, 116, 34, 58, 34, 34, 125, 44, 34, 80, 114, 111, 116, 111, 34, 58, 34, 72, 84, 84, 80, 47, 49, 46, 49, 34, 44, 34, 80, 114, 111, 116, 111, 77, 97, 106, 111, 114, 34, 58, 49, 44, 34, 80, 114, 111, 116, 111, 77, 105, 110, 111, 114, 34, 58, 49, 44, 34, 72, 101, 97, 100, 101, 114, 34, 58, 123, 34, 65, 99, 99, 101, 112, 116, 34, 58, 91, 34, 116, 101, 120, 116, 47, 104, 116, 109, 108, 44, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 120, 104, 116, 109, 108, 43, 120, 109, 108, 44, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 120, 109, 108, 59, 113, 61, 48, 46, 57, 44, 42, 47, 42, 59, 113, 61, 48, 46, 56, 34, 93, 44, 34, 65, 99, 99, 101, 112, 116, 45, 69, 110, 99, 111, 100, 105, 110, 103, 34, 58, 91, 34, 103, 122, 105, 112, 44, 32, 100, 101, 102, 108, 97, 116, 101, 34, 93, 44, 34, 65, 99, 99, 101, 112, 116, 45, 76, 97, 110, 103, 117, 97, 103, 101, 34, 58, 91, 34, 101, 110, 45, 117, 115, 34, 93, 44, 34, 67, 111, 110, 110, 101, 99, 116, 105, 111, 110, 34, 58, 91, 34, 107, 101, 101, 112, 45, 97, 108, 105, 118, 101, 34, 93, 44, 34, 85, 115, 101, 114, 45, 65, 103, 101, 110, 116, 34, 58, 91, 34, 77, 111, 122, 105, 108, 108, 97, 47, 53, 46, 48, 32, 40, 77, 97, 99, 105, 110, 116, 111, 115, 104, 59, 32, 73, 110, 116, 101, 108, 32, 77, 97, 99, 32, 79, 83, 32, 88, 32, 49, 48, 95, 49, 48, 95, 50, 41, 32, 65, 112, 112, 108, 101, 87, 101, 98, 75, 105, 116, 47, 54, 48, 48, 46, 52, 46, 49, 48, 32, 40, 75, 72, 84, 77, 76, 44, 32, 108, 105, 107, 101, 32, 71, 101, 99, 107, 111, 41, 32, 86, 101, 114, 115, 105, 111, 110, 47, 56, 46, 48, 46, 52, 32, 83, 97, 102, 97, 114, 105, 47, 54, 48, 48, 46, 52, 46, 49, 48, 34, 93, 125, 44, 34, 66, 111, 100, 121, 34, 58, 123, 34, 67, 108, 111, 115, 101, 114, 34, 58, 123, 34, 82, 101, 97, 100, 101, 114, 34, 58, 110, 117, 108, 108, 125, 125, 44, 34, 67, 111, 110, 116, 101, 110, 116, 76, 101, 110, 103, 116, 104, 34, 58, 48, 44, 34, 84, 114, 97, 110, 115, 102, 101, 114, 69, 110, 99, 111, 100, 105, 110, 103, 34, 58, 110, 117, 108, 108, 44, 34, 67, 108, 111, 115, 101, 34, 58, 102, 97, 108, 115, 101, 44, 34, 72, 111, 115, 116, 34, 58, 34, 100, 111, 99, 107, 101, 114, 104, 111, 115, 116, 58, 56, 48, 56, 48, 34, 44, 34, 70, 111, 114, 109, 34, 58, 110, 117, 108, 108, 44, 34, 80, 111, 115, 116, 70, 111, 114, 109, 34, 58, 110, 117, 108, 108, 44, 34, 77, 117, 108, 116, 105, 112, 97, 114, 116, 70, 111, 114, 109, 34, 58, 110, 117, 108, 108, 44, 34, 84, 114, 97, 105, 108, 101, 114, 34, 58, 110, 117, 108, 108, 44, 34, 82, 101, 109, 111, 116, 101, 65, 100, 100, 114, 34, 58, 34, 49, 57, 50, 46, 49, 54, 56, 46, 53, 57, 46, 51, 58, 54, 53, 52, 52, 54, 34, 44, 34, 82, 101, 113, 117, 101, 115, 116, 85, 82, 73, 34, 58, 34, 47, 34, 44, 34, 84, 76, 83, 34, 58, 110, 117, 108, 108, 125}
	request := FromJson(string(requestBytes))

	c.Assert(request.Method, Equals, "GET")
	c.Assert(request.URL.Path, Equals, "/")
	c.Assert(request.Host, Equals, "dockerhost:8080")
}

func (b *BackendSuite) TestMarshallUnmarshall(c *C) {
	var r SerializableHttpRequest
	r.Host = "Foobar"
	r.URL = &url.URL{Path: "Hello/World"}

	json := r.ToJson()
	r2 := FromJson(json)
	c.Assert(r2.Host, Equals, "Foobar")
	c.Assert(r2.URL.Path, Equals, "Hello/World")
}
