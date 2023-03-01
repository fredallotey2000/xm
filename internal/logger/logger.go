package log

import "net/http"

//Logger interface, can be implemented by with logger library
type Logger interface {
	HttpDebug(r *http.Request, fn string, err error)

	MqDebug(action string, msg interface{})
	MqError(action string, msg interface{}, err error)

	HttpError(r *http.Request, fn string, err error)
	Fatal(fn string, err error)

	Info(fn string)
}
