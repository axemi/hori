package discord

// s and t are null when op is not 0
type GatewayEvent struct {
	Op       GatewayOpcode `json:"op"`
	Data     interface{}   `json:"d"`
	Sequence *int          `json:"s"`
	Type     *string       `json:"t"`
}

// GatewayOpcode is a custom type to identify only specific allowed integers
type GatewayOpcode int

const (
	DISPATCH              GatewayOpcode = 0
	HEARTBEAT             GatewayOpcode = 1
	IDENTIFY              GatewayOpcode = 2
	PRESENCE_UPDATE       GatewayOpcode = 3
	VOICE_STATE_UPDATE    GatewayOpcode = 4
	RESUME                GatewayOpcode = 6
	RECONNECT             GatewayOpcode = 7
	REQUEST_GUILD_MEMBERS GatewayOpcode = 8
	INVALID_SESSION       GatewayOpcode = 9
	HELLO                 GatewayOpcode = 10
	HEARTBEAT_ACK         GatewayOpcode = 11
)
