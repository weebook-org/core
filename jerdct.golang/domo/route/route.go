// here is the component which
// describe the routing logic
// to reach REST end points

package route

import (
	"net/http"
	"jerdct.golang/domo/route/light"

)

type EndPoint interface {
	Pattern() string
	ServeHTTP(http.ResponseWriter, *http.Request)
}

var (
	endPoint []EndPoint
)

// build the route
func Build() {
	fillRoute()
	for _,e:= range endPoint {
		http.Handle(e.Pattern(), e)
	}
}

func fillRoute() {
	endPoint = append(endPoint, new(light.Controller))
}



