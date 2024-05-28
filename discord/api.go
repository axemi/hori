package discord

import (
	"encoding/json"
	"io"
	"net/http"
)

const DISCORD_API_URL = "https://discord.com/api/v10"

type GatewayEndpoint struct {
	URL string `json:"url"`
}

func GetGateway(apiUrl string) (string, error) {
	resp, err := http.Get(apiUrl + "/gateway")
	if err != nil {
		return "", err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	gateway := &GatewayEndpoint{}
	err = json.Unmarshal(bytes, gateway)
	if err != nil {
		return "", err
	}
	return gateway.URL, nil
}
