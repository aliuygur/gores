package gores_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alioygur/gores"
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

	gores.HTML(resp, responseCode, responseString)

	if resp.Body.String() != responseString || resp.Header().Get(gores.ContentType) != gores.TextHTMLCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	responseString := "Hello World"
	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	gores.String(resp, responseCode, responseString)

	if resp.Body.String() != responseString || resp.Header().Get(gores.ContentType) != gores.TextPlainCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestJSON(t *testing.T) {

	responseString := `{"Name":"Jhon","Email":"jhon@example.com","Age":20}`
	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	gores.JSON(resp, responseCode, user)

	if resp.Body.String() != responseString || resp.Header().Get(gores.ContentType) != gores.ApplicationJSONCharsetUTF8 || resp.Code != responseCode {
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

	gores.JSONIndent(resp, responseCode, user, "**", "%%")

	if resp.Body.String() != responseString || resp.Header().Get(gores.ContentType) != gores.ApplicationJSONCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestJSONP(t *testing.T) {

	responseString := `parseResponse({"Name":"Jhon","Email":"jhon@example.com","Age":20});`

	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	gores.JSONP(resp, responseCode, "parseResponse", user)

	if resp.Body.String() != responseString || resp.Header().Get(gores.ContentType) != gores.ApplicationJavaScriptCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestXML(t *testing.T) {

	responseString := `<?xml version="1.0" encoding="UTF-8"?>
<User><Name>Jhon</Name><Email>jhon@example.com</Email><Age>20</Age></User>`

	responseCode := http.StatusOK

	resp := httptest.NewRecorder()

	gores.XML(resp, responseCode, user)

	if resp.Body.String() != responseString || resp.Header().Get(gores.ContentType) != gores.ApplicationXMLCharsetUTF8 || resp.Code != responseCode {
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

	gores.XMLIndent(resp, responseCode, user, "**", "%%")

	if resp.Body.String() != responseString || resp.Header().Get(gores.ContentType) != gores.ApplicationXMLCharsetUTF8 || resp.Code != responseCode {
		t.Fail()
	}
}

func TestNoContent(t *testing.T) {

	responseString := ``

	responseCode := http.StatusNoContent

	resp := httptest.NewRecorder()

	gores.NoContent(resp)

	if resp.Body.String() != responseString || resp.Code != responseCode {
		t.Fail()
	}
}

func TestError(t *testing.T) {

	responseString := `error`

	responseCode := http.StatusBadRequest

	resp := httptest.NewRecorder()

	gores.Error(resp, responseCode, responseString)

	if resp.Body.String() != responseString+"\n" || resp.Code != responseCode {
		t.Fail()
	}
}
