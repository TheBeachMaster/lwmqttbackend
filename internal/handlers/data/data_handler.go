package data

import "net/http"

type MQTTDataRouteHandler interface {
	Store() http.HandlerFunc
	Default() http.HandlerFunc
}
