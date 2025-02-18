package models

import (
    "encoding/json"
    "fmt"
)

// MapOrgImportFile represents a MapOrgImportFile struct.
type MapOrgImportFile struct {
    // Whether to auto assign device to deviceprofile by name
    AutoDeviceprofileAssignment *bool                  `json:"auto_deviceprofile_assignment,omitempty"`
    // CSV file for ap name mapping, optional
    Csv                         *[]byte                `json:"csv,omitempty"`
    // Ekahau or ibwave file
    File                        *[]byte                `json:"file,omitempty"`
    Json                        *MapOrgImportFileJson  `json:"json,omitempty"`
    AdditionalProperties        map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapOrgImportFile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapOrgImportFile) String() string {
    return fmt.Sprintf(
    	"MapOrgImportFile[AutoDeviceprofileAssignment=%v, Csv=%v, File=%v, Json=%v, AdditionalProperties=%v]",
    	m.AutoDeviceprofileAssignment, m.Csv, m.File, m.Json, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapOrgImportFile.
// It customizes the JSON marshaling process for MapOrgImportFile objects.
func (m MapOrgImportFile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "auto_deviceprofile_assignment", "csv", "file", "json"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapOrgImportFile object to a map representation for JSON marshaling.
func (m MapOrgImportFile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.AutoDeviceprofileAssignment != nil {
        structMap["auto_deviceprofile_assignment"] = m.AutoDeviceprofileAssignment
    }
    if m.Csv != nil {
        structMap["csv"] = m.Csv
    }
    if m.File != nil {
        structMap["file"] = m.File
    }
    if m.Json != nil {
        structMap["json"] = m.Json.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapOrgImportFile.
// It customizes the JSON unmarshaling process for MapOrgImportFile objects.
func (m *MapOrgImportFile) UnmarshalJSON(input []byte) error {
    var temp tempMapOrgImportFile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_deviceprofile_assignment", "csv", "file", "json")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.AutoDeviceprofileAssignment = temp.AutoDeviceprofileAssignment
    m.Csv = temp.Csv
    m.File = temp.File
    m.Json = temp.Json
    return nil
}

// tempMapOrgImportFile is a temporary struct used for validating the fields of MapOrgImportFile.
type tempMapOrgImportFile  struct {
    AutoDeviceprofileAssignment *bool                 `json:"auto_deviceprofile_assignment,omitempty"`
    Csv                         *[]byte               `json:"csv,omitempty"`
    File                        *[]byte               `json:"file,omitempty"`
    Json                        *MapOrgImportFileJson `json:"json,omitempty"`
}
