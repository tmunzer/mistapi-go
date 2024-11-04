package models

import (
    "encoding/json"
)

// ResponseAutoZone represents a ResponseAutoZone struct.
type ResponseAutoZone struct {
    // The status for the auto zones service for a given map. enum:
    // * not_started: The auto zones service has not been run on this map or the results were cleared by the user
    // * in_progress: The auto zones service is currently in progress
    // * awaiting_review: The auto zones service has completed and suggested location zones to be added to the map
    // * error: There was an error with the auto zones service
    Status               *ResponseAutoZoneStatusEnum `json:"status,omitempty"`
    // A list of suggested zones to review and accept for a given map
    Zones                *ResponseAutoZoneZone       `json:"zones,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoZone.
// It customizes the JSON marshaling process for ResponseAutoZone objects.
func (r ResponseAutoZone) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoZone object to a map representation for JSON marshaling.
func (r ResponseAutoZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    if r.Zones != nil {
        structMap["zones"] = r.Zones.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoZone.
// It customizes the JSON unmarshaling process for ResponseAutoZone objects.
func (r *ResponseAutoZone) UnmarshalJSON(input []byte) error {
    var temp tempResponseAutoZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "status", "zones")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Status = temp.Status
    r.Zones = temp.Zones
    return nil
}

// tempResponseAutoZone is a temporary struct used for validating the fields of ResponseAutoZone.
type tempResponseAutoZone  struct {
    Status *ResponseAutoZoneStatusEnum `json:"status,omitempty"`
    Zones  *ResponseAutoZoneZone       `json:"zones,omitempty"`
}
