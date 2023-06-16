package auth

import "net/http"

type MQTTRouteHandlers interface {
	Authn() http.HandlerFunc
	Authz() http.HandlerFunc
	Store() http.HandlerFunc
	Default() http.HandlerFunc
}
