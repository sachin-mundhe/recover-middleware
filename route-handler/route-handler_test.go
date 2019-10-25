package routehandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSourceCodeHandler(t *testing.T) {

	testTable := []struct {
		testCaseName string
		url          string
		status       int
	}{
		{
			testCaseName: "TC1",
			url:          "localhost:3000/sourcecode/?line=24&path=/home/gslab/Desktop/InstError.txt",
			status:       200,
		}, {
			testCaseName: "TC2",
			url:          "localhost:3000/sourcecode/?line=ewr&path=/home/gslab/Desktop/InstError.txt",
			status:       200,
		},
		{
			testCaseName: "TC3",
			url:          "localhost:3000/sourcecode/?line=24&path=/hom/Desktop/InstError.txt",
			status:       500,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.testCaseName, func(ts *testing.T) {

			req, err := http.NewRequest("GET", tc.url, nil)

			if err != nil {
				t.Error("Error occured while making new request", err)
			}

			rec := httptest.NewRecorder()
			SourceCodeHandler(rec, req)
			res := rec.Result()
			if res.StatusCode != tc.status {
				ts.Error("Expected statuscode:", tc.status, "and Got:", res.StatusCode)
			}
		})
	}

}

func TestHomepageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3000/", nil)

	if err != nil {
		t.Error("Error occured while making new request", err)
	}

	rec := httptest.NewRecorder()
	HomepageHandler(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Error("Expected statuscode:", http.StatusOK, "and Got:", res.StatusCode)
	}

}

func TestMiddleware(t *testing.T) {
	handler := http.HandlerFunc(PanicHandler)
	executeRequest("Get", "/", Recover(handler, true))
}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Output: ", string(bodyBytes))
	return res, err
}

//isDev == false
func TestMiddleware1(t *testing.T) {
	handler := http.HandlerFunc(PanicHandler)
	executeRequest("Get", "/", Recover(handler, false))
}
