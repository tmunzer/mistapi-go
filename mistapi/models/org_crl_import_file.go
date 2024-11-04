package models

import (
    "encoding/json"
)

// OrgCrlImportFile represents a OrgCrlImportFile struct.
type OrgCrlImportFile struct {
    // a PEM or DER formatted CRL file
    File                 *[]byte        `json:"file,omitempty"`
    // a JSON string with "name" field for CRL file issuer (optional)
    Json                 *string        `json:"json,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgCrlImportFile.
// It customizes the JSON marshaling process for OrgCrlImportFile objects.
func (o OrgCrlImportFile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgCrlImportFile object to a map representation for JSON marshaling.
func (o OrgCrlImportFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.File != nil {
        structMap["file"] = o.File
    }
    if o.Json != nil {
        structMap["json"] = o.Json
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgCrlImportFile.
// It customizes the JSON unmarshaling process for OrgCrlImportFile objects.
func (o *OrgCrlImportFile) UnmarshalJSON(input []byte) error {
    var temp tempOrgCrlImportFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "file", "json")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.File = temp.File
    o.Json = temp.Json
    return nil
}

// tempOrgCrlImportFile is a temporary struct used for validating the fields of OrgCrlImportFile.
type tempOrgCrlImportFile  struct {
    File *[]byte `json:"file,omitempty"`
    Json *string `json:"json,omitempty"`
}
