package reading

import (
	"github.com/bmizerany/pat"
	"net/http"
)

const (
	readingInfo string = "/reading/infos"
)

func BindReading(m *pat.PatternServeMux) {
	m.Get(readingInfo, http.HandlerFunc(infoReading))
}

func infoReading(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world reading !"))
}
