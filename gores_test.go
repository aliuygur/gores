package gores

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	Name  string
	Email string
	Age   int
}

var user = User{
	Name:  "Jhon",
	Email: "jhon@example.com",
	Age:   20,
}

func TestHTML(t *testing.T) {
	responseString := "<h1>Hello World</h1>"
	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	HTML(resp, responseCode, responseString)

	if resp.Body.String() != responseString || resp.Header().Get(ContentType) != TextHTMLCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	responseString := "Hello World"
	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	String(resp, responseCode, responseString)

	if resp.Body.String() != responseString || resp.Header().Get(ContentType) != TextPlainCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestJSON(t *testing.T) {

	responseString := `{"Name":"Jhon","Email":"jhon@example.com","Age":20}`
	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	JSON(resp, responseCode, user)

	if resp.Body.String() != responseString || resp.Header().Get(ContentType) != ApplicationJSONCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestJSONIndent(t *testing.T) {

	responseString := `{
**%%"Name": "Jhon",
**%%"Email": "jhon@example.com",
**%%"Age": 20
**}`

	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	JSONIndent(resp, responseCode, user, "**", "%%")

	if resp.Body.String() != responseString || resp.Header().Get(ContentType) != ApplicationJSONCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestJSONP(t *testing.T) {

	responseString := `parseResponse({"Name":"Jhon","Email":"jhon@example.com","Age":20});`

	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	JSONP(resp, responseCode, "parseResponse", user)

	if resp.Body.String() != responseString || resp.Header().Get(ContentType) != ApplicationJavaScriptCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestXML(t *testing.T) {

	responseString := `<?xml version="1.0" encoding="UTF-8"?>
<User><Name>Jhon</Name><Email>jhon@example.com</Email><Age>20</Age></User>`

	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	XML(resp, responseCode, user)

	if resp.Body.String() != responseString || resp.Header().Get(ContentType) != ApplicationXMLCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestXMLIndent(t *testing.T) {

	responseString := `<?xml version="1.0" encoding="UTF-8"?>
**<User>
**%%<Name>Jhon</Name>
**%%<Email>jhon@example.com</Email>
**%%<Age>20</Age>
**</User>`

	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	XMLIndent(resp, responseCode, user, "**", "%%")

	if resp.Body.String() != responseString || resp.Header().Get(ContentType) != ApplicationXMLCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestNoContent(t *testing.T) {

	responseString := ``

	responseCode := http.StatusNoContent

	resp := httptest.NewRecorder()

	NoContent(resp)

	if resp.Body.String() != responseString || resp.Code != responseCode {
		t.Fail()
	}
}

func TestError(t *testing.T) {

	responseString := `error`

	responseCode := http.StatusBadRequest

	resp := httptest.NewRecorder()

	Error(resp, responseCode, responseString)

	if resp.Body.String() != responseString+"\n" || resp.Code != responseCode {
		t.Fail()
	}
}
