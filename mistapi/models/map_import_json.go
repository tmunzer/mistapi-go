// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// MapImportJson represents a MapImportJson struct.
type MapImportJson struct {
    ImportAllFloorplans  *bool                       `json:"import_all_floorplans,omitempty"`
    ImportHeight         *bool                       `json:"import_height,omitempty"`
    ImportOrientation    *bool                       `json:"import_orientation,omitempty"`
    // enum: `ekahau`, `ibwave`
    VendorName           MapImportJsonVendorNameEnum `json:"vendor_name"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for MapImportJson,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapImportJson) String() string {
    return fmt.Sprintf(
    	"MapImportJson[ImportAllFloorplans=%v, ImportHeight=%v, ImportOrientation=%v, VendorName=%v, AdditionalProperties=%v]",
    	m.ImportAllFloorplans, m.ImportHeight, m.ImportOrientation, m.VendorName, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapImportJson.
// It customizes the JSON marshaling process for MapImportJson objects.
func (m MapImportJson) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "import_all_floorplans", "import_height", "import_orientation", "vendor_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapImportJson object to a map representation for JSON marshaling.
func (m MapImportJson) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.ImportAllFloorplans != nil {
        structMap["import_all_floorplans"] = m.ImportAllFloorplans
    }
    if m.ImportHeight != nil {
        structMap["import_height"] = m.ImportHeight
    }
    if m.ImportOrientation != nil {
        structMap["import_orientation"] = m.ImportOrientation
    }
    structMap["vendor_name"] = m.VendorName
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapImportJson.
// It customizes the JSON unmarshaling process for MapImportJson objects.
func (m *MapImportJson) UnmarshalJSON(input []byte) error {
    var temp tempMapImportJson
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "import_all_floorplans", "import_height", "import_orientation", "vendor_name")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.ImportAllFloorplans = temp.ImportAllFloorplans
    m.ImportHeight = temp.ImportHeight
    m.ImportOrientation = temp.ImportOrientation
    m.VendorName = *temp.VendorName
    return nil
}

// tempMapImportJson is a temporary struct used for validating the fields of MapImportJson.
type tempMapImportJson  struct {
    ImportAllFloorplans *bool                        `json:"import_all_floorplans,omitempty"`
    ImportHeight        *bool                        `json:"import_height,omitempty"`
    ImportOrientation   *bool                        `json:"import_orientation,omitempty"`
    VendorName          *MapImportJsonVendorNameEnum `json:"vendor_name"`
}

func (m *tempMapImportJson) validate() error {
    var errs []string
    if m.VendorName == nil {
        errs = append(errs, "required field `vendor_name` is missing for type `map_import_json`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
