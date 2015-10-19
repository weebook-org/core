package services

import (
	"net/http"
	"odos.back/services/submission"
)


func BindServices() {
	subService := submission.NewService()
	for _,e:= range subService.EndPoints() {
		http.Handle(e.Pattern(), e)
	}
}
