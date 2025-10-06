package handler

import (
	"encoding/json"
	"os"
)

type ControlPlane struct {
	Identifier string `json:"identifier"`
}

type DiscoveryResponse[T any] struct {
	VersionInfo  string       `json:"version_info"`
	Resources    []T          `json:"resources"`
	TypeURL      string       `json:"type_url"`
	Nonce        string       `json:"nonce"`
	ControlPlane ControlPlane `json:"control_plane"`
}


func GetJson[T any](path string) (T, error) {
	file, err := os.Open(path)
	if err != nil {
		return *new(T), err
	}
	defer file.Close()

	var result T
	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return *new(T), err
	}
	return result, nil
}
