package models

import (
    "encoding/json"
    "fmt"
)

// OrgCrlImportFile represents a OrgCrlImportFile struct.
type OrgCrlImportFile struct {
    // a PEM or DER formatted CRL file
    File                 *[]byte                `json:"file,omitempty"`
    // a JSON string with "name" field for CRL file issuer (optional)
    Json                 *string                `json:"json,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgCrlImportFile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgCrlImportFile) String() string {
    return fmt.Sprintf(
    	"OrgCrlImportFile[File=%v, Json=%v, AdditionalProperties=%v]",
    	o.File, o.Json, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgCrlImportFile.
// It customizes the JSON marshaling process for OrgCrlImportFile objects.
func (o OrgCrlImportFile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "file", "json"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgCrlImportFile object to a map representation for JSON marshaling.
func (o OrgCrlImportFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "file", "json")
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
