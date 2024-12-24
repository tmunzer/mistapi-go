package models

import (
    "encoding/json"
    "fmt"
)

// MapSiteImportFile represents a MapSiteImportFile struct.
type MapSiteImportFile struct {
    // whether to auto assign device to deviceprofile by name
    AutoDeviceprofileAssignment *bool                  `json:"auto_deviceprofile_assignment,omitempty"`
    // csv file for ap name mapping, optional
    Csv                         *[]byte                `json:"csv,omitempty"`
    // ekahau or ibwave file
    File                        *[]byte                `json:"file,omitempty"`
    Json                        *MapImportJson         `json:"json,omitempty"`
    AdditionalProperties        map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapSiteImportFile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapSiteImportFile) String() string {
    return fmt.Sprintf(
    	"MapSiteImportFile[AutoDeviceprofileAssignment=%v, Csv=%v, File=%v, Json=%v, AdditionalProperties=%v]",
    	m.AutoDeviceprofileAssignment, m.Csv, m.File, m.Json, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapSiteImportFile.
// It customizes the JSON marshaling process for MapSiteImportFile objects.
func (m MapSiteImportFile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "auto_deviceprofile_assignment", "csv", "file", "json"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapSiteImportFile object to a map representation for JSON marshaling.
func (m MapSiteImportFile) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for MapSiteImportFile.
// It customizes the JSON unmarshaling process for MapSiteImportFile objects.
func (m *MapSiteImportFile) UnmarshalJSON(input []byte) error {
    var temp tempMapSiteImportFile
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

// tempMapSiteImportFile is a temporary struct used for validating the fields of MapSiteImportFile.
type tempMapSiteImportFile  struct {
    AutoDeviceprofileAssignment *bool          `json:"auto_deviceprofile_assignment,omitempty"`
    Csv                         *[]byte        `json:"csv,omitempty"`
    File                        *[]byte        `json:"file,omitempty"`
    Json                        *MapImportJson `json:"json,omitempty"`
}
