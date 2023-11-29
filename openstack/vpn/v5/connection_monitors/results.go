package connection_monitors

type ConnectionMonitor struct {
	ID                 string                  `json:"id"`
	RequestId          string                  `json:"request_id"`
	ConnectionMonitors []ConnectionMonitorInfo `json:"connection_monitors"`
}

type ConnectionMonitorInfo struct {
	ID              string `json:"id"`
	Status          string `json:"status"`
	VpnConnectionId string `json:"vpn_connection_id"`
	Type            string `json:"type"`
	SourceIp        string `json:"source_ip"`
	DestinationIp   string `json:"destination_ip"`
	ProtoType       string `json:"proto_type"`
}

type ListResp struct {
	ConnectionMonitors []ConnectionMonitor `json:"connection_monitors"`
	TotalCount         int                 `json:"total_count"`
	RequestId          string              `json:"request_id"`
}
