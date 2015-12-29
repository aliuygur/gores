# gores [![Build Status](https://travis-ci.org/alioygur/gores.svg?branch=master)](https://travis-ci.org/alioygur/gores)

http response utility library for GO


## installation

`go get github.com/alioygur/gores`


## usage

```go
package main

import (
	"log"
	"net/http"

	"github.com/alioygur/gores"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func StringResponse(w http.ResponseWriter, r *http.Request) {
	gores.String(w, http.StatusOK, "Hello World")
}

func HTMLResponse(w http.ResponseWriter, r *http.Request) {
	gores.HTML(w, http.StatusOK, "<h1>Hello World</h1>")
}

func JSONResponse(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Ali", Email: "alioygur@gmail.com", Age: 28}
	gores.JSON(w, http.StatusOK, user)
}

func main() {
	http.HandleFunc("/", StringResponse)
	http.HandleFunc("/html", HTMLResponse)
	http.HandleFunc("/json", JSONResponse)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

for more documentation [godoc](https://godoc.org/github.com/alioygur/gores)

## TODO

- [x] write tests
- [ ] add file response

## Contribute

**Use issues for everything**

- Report problems
- Discuss before sending a pull request
- Suggest new features/recipes
- Improve/fix documentation

## Thanks & Authors

I use code/got inspiration from these excellent libraries:

- [labstack/echo](https://github.com/labstack/echo) micro web framework
