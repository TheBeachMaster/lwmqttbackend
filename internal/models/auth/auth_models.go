package auth

type AuthenticateDeviceInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthorizationHTTPRequestInfo struct {
	Username string `json:"username"`
	Topic    string `json:"topic"`
	Action   string `json:"action"`
	DeviceId string `json:"clientid"`
}

type AuthNResponse struct {
	Result      string `json:"result"`
	IsSuperUser bool   `json:"is_superuser"`
}

type AuthZResponse struct {
	Result string `json:"result"`
}
