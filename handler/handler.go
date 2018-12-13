package handler

import (
	"fmt"
	"net/http"
)

func Dispatch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	fmt.Fprint(w, "Header:\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "  %q: %q\n", k, v)
	}
	fmt.Fprintf(w, "Host: %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %q\n", r.RemoteAddr)
}
