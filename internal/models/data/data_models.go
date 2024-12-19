package data

type MQTTMessage struct {
	Topic             string `json:"topic"`
	Message           string `json:"payload"`
	ClientId          string `json:"clientid"`
	IP                string `json:"peerhost"`
	MQTTMessageId     string `json:"id"`
	QoS               int    `json:"qos"`
	MQTTTimestamp     int64  `json:"timestamp"`
	MessageHeader     string `json:"headers"`
	PublishProperties string `json:"pub_props"`
	PublishReceivedAt int64  `json:"publish_received_at"`
	BrokerNodeName    string `json:"node"`
	PublishFlags      string `json:"flags,omitempty"`
}

type HealthStatus struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}
