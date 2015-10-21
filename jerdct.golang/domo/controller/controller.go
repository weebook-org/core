package controller
import "net/http"


type EndPoint interface {
	Pattern() string
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Controller struct{
	base string
	ep []EndPoint
}

type CtrlErr string

func (err CtrlErr) Error() string {
	return err
}

func NewController(base string) *Controller{
	ctrl := &new(Controller)
	ctrl.base = base
	return ctrl
}

func (ctrl *Controller) append(endPoint EndPoint) error {
	if endPoint == nil {
		return CtrlErr("the EndPoint has to be a non nil value")
	}
	ctrl.ep = append(ctrl.ep, endPoint)
	return nil
}



