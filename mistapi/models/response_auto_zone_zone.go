package models

import (
    "encoding/json"
    "fmt"
)

// ResponseAutoZoneZone represents a ResponseAutoZoneZone struct.
// A list of suggested zones to review and accept for a given map
type ResponseAutoZoneZone struct {
    // The name of the suggested zone
    Name                 *string                      `json:"name,omitempty"`
    // A list of of points comprising the zones map location in pixels
    Vertices             []ResponseAutoZoneZoneVertex `json:"vertices,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoZoneZone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoZoneZone) String() string {
    return fmt.Sprintf(
    	"ResponseAutoZoneZone[Name=%v, Vertices=%v, AdditionalProperties=%v]",
    	r.Name, r.Vertices, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoZoneZone.
// It customizes the JSON marshaling process for ResponseAutoZoneZone objects.
func (r ResponseAutoZoneZone) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "name", "vertices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoZoneZone object to a map representation for JSON marshaling.
func (r ResponseAutoZoneZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Name != nil {
        structMap["name"] = r.Name
    }
    if r.Vertices != nil {
        structMap["vertices"] = r.Vertices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoZoneZone.
// It customizes the JSON unmarshaling process for ResponseAutoZoneZone objects.
func (r *ResponseAutoZoneZone) UnmarshalJSON(input []byte) error {
    var temp tempResponseAutoZoneZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "vertices")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Name = temp.Name
    r.Vertices = temp.Vertices
    return nil
}

// tempResponseAutoZoneZone is a temporary struct used for validating the fields of ResponseAutoZoneZone.
type tempResponseAutoZoneZone  struct {
    Name     *string                      `json:"name,omitempty"`
    Vertices []ResponseAutoZoneZoneVertex `json:"vertices,omitempty"`
}
