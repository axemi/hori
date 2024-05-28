package discord

import (
	"testing"
)

func TestReceiveGatewayUrl(t *testing.T) {
	want, err := GetGateway(DISCORD_API_URL)
	if err != nil {
		t.Fatal(err)
	}
	if want != "" {
		t.Log(want)
	}
}
