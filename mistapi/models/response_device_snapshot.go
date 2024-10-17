package models

import (
	"encoding/json"
)

// ResponseDeviceSnapshot represents a ResponseDeviceSnapshot struct.
type ResponseDeviceSnapshot struct {
	// the internal status id
	StatusId *string `json:"status_id,omitempty"`
	// enum: `error`, `inprogress`, `scheduled`, `starting`, `success`
	Staus                *ResponseDeviceSnapshotStatusEnum `json:"staus,omitempty"`
	Timestamp            *float64                          `json:"timestamp,omitempty"`
	AdditionalProperties map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceSnapshot.
// It customizes the JSON marshaling process for ResponseDeviceSnapshot objects.
func (r ResponseDeviceSnapshot) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceSnapshot object to a map representation for JSON marshaling.
func (r ResponseDeviceSnapshot) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.StatusId != nil {
		structMap["status_id"] = r.StatusId
	}
	if r.Staus != nil {
		structMap["staus"] = r.Staus
	}
	if r.Timestamp != nil {
		structMap["timestamp"] = r.Timestamp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceSnapshot.
// It customizes the JSON unmarshaling process for ResponseDeviceSnapshot objects.
func (r *ResponseDeviceSnapshot) UnmarshalJSON(input []byte) error {
	var temp tempResponseDeviceSnapshot
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "status_id", "staus", "timestamp")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.StatusId = temp.StatusId
	r.Staus = temp.Staus
	r.Timestamp = temp.Timestamp
	return nil
}

// tempResponseDeviceSnapshot is a temporary struct used for validating the fields of ResponseDeviceSnapshot.
type tempResponseDeviceSnapshot struct {
	StatusId  *string                           `json:"status_id,omitempty"`
	Staus     *ResponseDeviceSnapshotStatusEnum `json:"staus,omitempty"`
	Timestamp *float64                          `json:"timestamp,omitempty"`
}
