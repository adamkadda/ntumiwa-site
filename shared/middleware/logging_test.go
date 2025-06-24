package middleware

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging_Success(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := log.New(buf, "", 0) // no prefix, no flags

	teapot := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte("I'm a little teapot~"))
	})

	/*
		Funny syntax.
		In case you get confused later,
		these are equivalent:

		middleware := Logging(logger)
		handler := middleware(teapot)

					---

		handler := Loggin(logger)(teapot)
	*/

	handler := Logging(logger)(teapot)

	r := httptest.NewRequest("GET", "/tea", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	logOutput := buf.String()

	fmt.Print(logOutput)

	assert.Regexp(t, `GET /tea \d+(\.\d+)?(ns|µs|ms|s)\s*$`, logOutput)
}
