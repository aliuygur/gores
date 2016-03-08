// Package gores http response utility library for GO
package gores

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"path/filepath"
)

// HTTP Methods
const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

// Media Types
const (
	ApplicationJSON                  = "application/json"
	ApplicationJSONCharsetUTF8       = ApplicationJSON + "; " + CharsetUTF8
	ApplicationJavaScript            = "application/javascript"
	ApplicationJavaScriptCharsetUTF8 = ApplicationJavaScript + "; " + CharsetUTF8
	ApplicationXML                   = "application/xml"
	ApplicationXMLCharsetUTF8        = ApplicationXML + "; " + CharsetUTF8
	ApplicationForm                  = "application/x-www-form-urlencoded"
	ApplicationProtobuf              = "application/protobuf"
	ApplicationMsgpack               = "application/msgpack"
	TextHTML                         = "text/html"
	TextHTMLCharsetUTF8              = TextHTML + "; " + CharsetUTF8
	TextPlain                        = "text/plain"
	TextPlainCharsetUTF8             = TextPlain + "; " + CharsetUTF8
	MultipartForm                    = "multipart/form-data"
)

// Headers
const (
	AcceptEncoding     = "Accept-Encoding"
	Authorization      = "Authorization"
	ContentDisposition = "Content-Disposition"
	ContentEncoding    = "Content-Encoding"
	ContentLength      = "Content-Length"
	ContentType        = "Content-Type"
	Location           = "Location"
	Upgrade            = "Upgrade"
	Vary               = "Vary"
	WWWAuthenticate    = "WWW-Authenticate"
	XForwardedFor      = "X-Forwarded-For"
	XRealIP            = "X-Real-IP"
)

const (
	// CharsetUTF8 utf8 character set
	CharsetUTF8 = "charset=utf-8"

	// WebSocket web socket protocol
	WebSocket = "websocket"
)

// HTML sends an HTTP response with status code.
func HTML(w http.ResponseWriter, code int, html string) {
	w.Header().Set(ContentType, TextHTMLCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(html))
}

// String sends a string response with status code.
func String(w http.ResponseWriter, code int, s string) {
	w.Header().Set(ContentType, TextPlainCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(s))
}

// JSON sends a JSON response with status code.
func JSON(w http.ResponseWriter, code int, i interface{}) (err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	_json(w, code, b)
	return
}

// JSONIndent sends a JSON response with status code, but it applies prefix and indent to format the output.
func JSONIndent(w http.ResponseWriter, code int, i interface{}, prefix string, indent string) (err error) {
	b, err := json.MarshalIndent(i, prefix, indent)
	if err != nil {
		return err
	}
	_json(w, code, b)
	return
}

func _json(w http.ResponseWriter, code int, b []byte) {
	w.Header().Set(ContentType, ApplicationJSONCharsetUTF8)
	w.WriteHeader(code)
	w.Write(b)
}

// JSONP sends a JSONP response with status code. It uses `callback` to construct
// the JSONP payload.
func JSONP(w http.ResponseWriter, code int, callback string, i interface{}) (err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	w.Header().Set(ContentType, ApplicationJavaScriptCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(callback + "("))
	w.Write(b)
	w.Write([]byte(");"))
	return
}

// XML sends an XML response with status code.
func XML(w http.ResponseWriter, code int, i interface{}) (err error) {
	b, err := xml.Marshal(i)
	if err != nil {
		return err
	}
	_xml(w, code, b)
	return
}

// XMLIndent sends an XML response with status code, but it applies prefix and indent to format the output.
func XMLIndent(w http.ResponseWriter, code int, i interface{}, prefix string, indent string) (err error) {
	b, err := xml.MarshalIndent(i, prefix, indent)
	if err != nil {
		return err
	}
	_xml(w, code, b)
	return
}

func _xml(w http.ResponseWriter, code int, b []byte) {
	w.Header().Set(ContentType, ApplicationXMLCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(xml.Header))
	w.Write(b)
}

// File sends a response with the content of the file
func File(w http.ResponseWriter, r *http.Request, path string) (err error) {
	err = file(w, r, path, "", false)
	return
}

// Download the client is prompted to save the file with provided `name`,
// name can be empty, in that case name of the file is used.
func Download(w http.ResponseWriter, r *http.Request, path string, name string) (err error) {
	err = file(w, r, path, name, true)
	return
}

func file(w http.ResponseWriter, r *http.Request, path, name string, attachment bool) (err error) {
	dir, file := filepath.Split(path)
	if attachment {
		w.Header().Set(ContentDisposition, "attachment; filename="+name)
	}

	fs := http.Dir(dir)
	f, err := fs.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	fi, _ := f.Stat()

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)

	return
}

// NoContent sends a response with no body and a status code.
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// Error sends a error response with a status code
func Error(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}
