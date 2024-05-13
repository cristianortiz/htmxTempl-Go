package handlers

import (
	"net/http"

	foo "github.com/cristianortiz/htmxTempl-Go/views"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	//rendering foo.index templ component and return it to the client as response
	return Render(w, r, foo.Index())
}
