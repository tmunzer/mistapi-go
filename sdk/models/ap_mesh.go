package models

import (
    "encoding/json"
)

// ApMesh represents a ApMesh struct.
// Mesh AP settings
type ApMesh struct {
    // whether mesh is enabled on this AP
    Enabled              *bool           `json:"enabled,omitempty"`
    // mesh group, base AP(s) will only allow remote AP(s) in the same mesh group to join, 1-9, optional
    Group                Optional[int]   `json:"group"`
    Role                 *ApMeshRoleEnum `json:"role,omitempty"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApMesh.
// It customizes the JSON marshaling process for ApMesh objects.
func (a ApMesh) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApMesh object to a map representation for JSON marshaling.
func (a ApMesh) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Group.IsValueSet() {
        if a.Group.Value() != nil {
            structMap["group"] = a.Group.Value()
        } else {
            structMap["group"] = nil
        }
    }
    if a.Role != nil {
        structMap["role"] = a.Role
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApMesh.
// It customizes the JSON unmarshaling process for ApMesh objects.
func (a *ApMesh) UnmarshalJSON(input []byte) error {
    var temp apMesh
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "group", "role")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Enabled = temp.Enabled
    a.Group = temp.Group
    a.Role = temp.Role
    return nil
}

// apMesh is a temporary struct used for validating the fields of ApMesh.
type apMesh  struct {
    Enabled *bool           `json:"enabled,omitempty"`
    Group   Optional[int]   `json:"group"`
    Role    *ApMeshRoleEnum `json:"role,omitempty"`
}
