package models

import (
    "encoding/json"
)

// PsksImportFile represents a PsksImportFile struct.
type PsksImportFile struct {
    File                 *[]byte        `json:"file,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PsksImportFile.
// It customizes the JSON marshaling process for PsksImportFile objects.
func (p PsksImportFile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PsksImportFile object to a map representation for JSON marshaling.
func (p PsksImportFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.File != nil {
        structMap["file"] = p.File
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PsksImportFile.
// It customizes the JSON unmarshaling process for PsksImportFile objects.
func (p *PsksImportFile) UnmarshalJSON(input []byte) error {
    var temp tempPsksImportFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "file")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.File = temp.File
    return nil
}

// tempPsksImportFile is a temporary struct used for validating the fields of PsksImportFile.
type tempPsksImportFile  struct {
    File *[]byte `json:"file,omitempty"`
}
