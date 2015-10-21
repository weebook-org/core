package main

import (
	"github.com/bmizerany/pat"
	"log"
	"net/http"
	"odos.back/reading"
	"odos.back/submission"
	"odos.back/vote"
)

func main() {
	m := pat.New()
	bind(m)
	http.Handle("/", m)
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Printf("The server stop because of %v", error)
	}
}

func bind(m *pat.PatternServeMux) {
	submission.BindSubmission(m)
	submission.BindComments(m)
	reading.BindReading(m)
	vote.BindVote(m)
}
