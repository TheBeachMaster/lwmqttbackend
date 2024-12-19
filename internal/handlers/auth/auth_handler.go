package auth

import "net/http"

type MQTTAuthRouteHandlers interface {
	Authn() http.HandlerFunc
	Authz() http.HandlerFunc
}
