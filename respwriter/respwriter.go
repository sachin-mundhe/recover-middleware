package respwriter

import "net/http"

//ResponseWriter Customised response writer
type ResponseWriter struct {
	http.ResponseWriter
	status int
	writes [][]byte
}

//Write It is responsible for writing data into personal Response writer
func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.writes = append(rw.writes, b)
	return len(b), nil
}

//WriteHeader It is responsible to set StatusCode into personal Response writer
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
}

//Flush It copies data from personal Response writer to http.Response writer
func (rw *ResponseWriter) Flush() {

	if rw.status != 0 {
		rw.ResponseWriter.Header().Set("Content-Type", "text/html")
		rw.ResponseWriter.WriteHeader(rw.status)
	}
	for _, v := range rw.writes {
		rw.ResponseWriter.Write(v)
	}
}
