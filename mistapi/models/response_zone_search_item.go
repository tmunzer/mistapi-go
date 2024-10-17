package models

import (
	"encoding/json"
)

// ResponseZoneSearchItem represents a ResponseZoneSearchItem struct.
type ResponseZoneSearchItem struct {
	Enter                *int           `json:"enter,omitempty"`
	Scope                *string        `json:"scope,omitempty"`
	Timestamp            *int           `json:"timestamp,omitempty"`
	User                 *string        `json:"user,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseZoneSearchItem.
// It customizes the JSON marshaling process for ResponseZoneSearchItem objects.
func (r ResponseZoneSearchItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseZoneSearchItem object to a map representation for JSON marshaling.
func (r ResponseZoneSearchItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Enter != nil {
		structMap["enter"] = r.Enter
	}
	if r.Scope != nil {
		structMap["scope"] = r.Scope
	}
	if r.Timestamp != nil {
		structMap["timestamp"] = r.Timestamp
	}
	if r.User != nil {
		structMap["user"] = r.User
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseZoneSearchItem.
// It customizes the JSON unmarshaling process for ResponseZoneSearchItem objects.
func (r *ResponseZoneSearchItem) UnmarshalJSON(input []byte) error {
	var temp tempResponseZoneSearchItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enter", "scope", "timestamp", "user")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Enter = temp.Enter
	r.Scope = temp.Scope
	r.Timestamp = temp.Timestamp
	r.User = temp.User
	return nil
}

// tempResponseZoneSearchItem is a temporary struct used for validating the fields of ResponseZoneSearchItem.
type tempResponseZoneSearchItem struct {
	Enter     *int    `json:"enter,omitempty"`
	Scope     *string `json:"scope,omitempty"`
	Timestamp *int    `json:"timestamp,omitempty"`
	User      *string `json:"user,omitempty"`
}
