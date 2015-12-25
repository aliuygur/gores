package gores

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const (
	// CONNECT HTTP method
	CONNECT = "CONNECT"
	// DELETE HTTP method
	DELETE = "DELETE"
	// GET HTTP method
	GET = "GET"
	// HEAD HTTP method
	HEAD = "HEAD"
	// OPTIONS HTTP method
	OPTIONS = "OPTIONS"
	// PATCH HTTP method
	PATCH = "PATCH"
	// POST HTTP method
	POST = "POST"
	// PUT HTTP method
	PUT = "PUT"
	// TRACE HTTP method
	TRACE = "TRACE"

	//-------------
	// Media types
	//-------------

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

	//---------
	// Charset
	//---------

	CharsetUTF8 = "charset=utf-8"

	//---------
	// Headers
	//---------

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
	//-----------
	// Protocols
	//-----------

	WebSocket = "websocket"
)

// HTML sends an HTTP response with status code.
func HTML(w http.ResponseWriter, code int, html string) (err error) {
	w.Header().Set(ContentType, TextHTMLCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(html))
	return
}

// String sends a string response with status code.
func String(w http.ResponseWriter, code int, s string) (err error) {
	w.Header().Set(ContentType, TextPlainCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(s))
	return
}

// JSON sends a JSON response with status code.
func JSON(w http.ResponseWriter, code int, i interface{}) (err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	w.Header().Set(ContentType, ApplicationJSONCharsetUTF8)
	w.WriteHeader(code)
	w.Write(b)
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
	w.Header().Set(ContentType, ApplicationXMLCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(xml.Header))
	w.Write(b)
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

// File sends a response with the content of the file. If `attachment` is set
// to true, the client is prompted to save the file with provided `name`,
// name can be empty, in that case name of the file is used.
// func File(w http.ResponseWriter, path, name string, attachment bool) (err error) {
// 	dir, file := filepath.Split(path)
// 	if attachment {
// 		w.Header().Set(ContentDisposition, "attachment; filename="+name)
// 	}
//
//   fs := http.Dir(dir)
//   	f, err := fs.Open(file)
//   	if err != nil {
//       return
//   	}
//   	defer f.Close()
//
//   	fi, _ := f.Stat()
//
//   	http.ServeContent(c.response, c.request, fi.Name(), fi.ModTime(), f)
//
// 	return
// }

// NoContent sends a response with no body and a status code.
func NoContent(w http.ResponseWriter, code int) error {
	w.WriteHeader(code)
	return nil
}

// Error invokes the registered HTTP error handler. Generally used by middleware.
func Error(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
}
