package routehandler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/sachin-mundhe/recover-middleware/links"
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

				fmt.Fprintf(w, "<h1>Panic: %v</h1> <pre>%s</pre>", err, links.GenerateLinks(string(stack)))
			}
		}()
		// nw := &respwriter.ResponseWriter{ResponseWriter: w}
		app.ServeHTTP(w, r)
		// nw.Flush()
	}

}

// PanicHandler It handles panic route
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("Oh no!")
}

// HomepageHandler It handles Homepage route
func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

// SourceCodeHandler This method handles /sourcecode/ route
func SourceCodeHandler(w http.ResponseWriter, r *http.Request) {

	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr)

	if err != nil {
		line = -1
	}
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	io.Copy(b, file)

	var lines [][2]int
	if line > 1 {
		lines = append(lines, [2]int{line, line})
	}

	lexer := lexers.Get("go")
	iterator, _ := lexer.Tokenise(nil, b.String())
	formatter := html.New(html.HighlightLines(lines), html.WithLineNumbers(), html.TabWidth(6))
	style := styles.Get("github")
	w.Header().Set("Content-Type", "text/html")
	formatter.Format(w, style, iterator)

}
