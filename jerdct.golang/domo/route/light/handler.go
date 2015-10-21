package light

import(
	"net/http"
	"html"
	"fmt"
)


type Controller struct{
	Base string
}


func (ctrl *Controller) Pattern() string {
	return "/light"
}

func (ctrl *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}



