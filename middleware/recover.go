package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

//Recover Middleware function
func Recover(app http.Handler, isDev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// log.Println(err)
				stack := debug.Stack()
				if !isDev {
					http.Error(w, "Something went wrong :-(", http.StatusInternalServerError)
					return
				}
				// w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>Panic: %v</h1> <pre>%s</pre>", err, string(stack))
			}
		}()
		// nw := &respwriter.ResponseWriter{ResponseWriter: w}
		app.ServeHTTP(w, r)
		// nw.Flush()
	}

}
