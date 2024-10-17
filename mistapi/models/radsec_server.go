package models

import (
	"encoding/json"
)

// RadsecServer represents a RadsecServer struct.
type RadsecServer struct {
	Host                 *string        `json:"host,omitempty"`
	Port                 *int           `json:"port,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RadsecServer.
// It customizes the JSON marshaling process for RadsecServer objects.
func (r RadsecServer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RadsecServer object to a map representation for JSON marshaling.
func (r RadsecServer) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Host != nil {
		structMap["host"] = r.Host
	}
	if r.Port != nil {
		structMap["port"] = r.Port
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RadsecServer.
// It customizes the JSON unmarshaling process for RadsecServer objects.
func (r *RadsecServer) UnmarshalJSON(input []byte) error {
	var temp tempRadsecServer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "host", "port")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Host = temp.Host
	r.Port = temp.Port
	return nil
}

// tempRadsecServer is a temporary struct used for validating the fields of RadsecServer.
type tempRadsecServer struct {
	Host *string `json:"host,omitempty"`
	Port *int    `json:"port,omitempty"`
}
