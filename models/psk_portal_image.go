package models

import (
    "encoding/json"
)

// PskPortalImage represents a PskPortalImage struct.
type PskPortalImage struct {
    // Binary file
    File                 *[]byte        `json:"file,omitempty"`
    // JSON string describing the upload
    Json                 *string        `json:"json,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PskPortalImage.
// It customizes the JSON marshaling process for PskPortalImage objects.
func (p PskPortalImage) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PskPortalImage object to a map representation for JSON marshaling.
func (p PskPortalImage) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.File != nil {
        structMap["file"] = p.File
    }
    if p.Json != nil {
        structMap["json"] = p.Json
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskPortalImage.
// It customizes the JSON unmarshaling process for PskPortalImage objects.
func (p *PskPortalImage) UnmarshalJSON(input []byte) error {
    var temp pskPortalImage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "file", "json")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.File = temp.File
    p.Json = temp.Json
    return nil
}

// pskPortalImage is a temporary struct used for validating the fields of PskPortalImage.
type pskPortalImage  struct {
    File *[]byte `json:"file,omitempty"`
    Json *string `json:"json,omitempty"`
}
