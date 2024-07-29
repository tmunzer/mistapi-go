package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MapImportJson represents a MapImportJson struct.
type MapImportJson struct {
    ImportAllFloorplans  *bool                       `json:"import_all_floorplans,omitempty"`
    ImportHeight         *bool                       `json:"import_height,omitempty"`
    ImportOrientation    *bool                       `json:"import_orientation,omitempty"`
    // enum: `ekahau`, `ibwave`
    VendorName           MapImportJsonVendorNameEnum `json:"vendor_name"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapImportJson.
// It customizes the JSON marshaling process for MapImportJson objects.
func (m MapImportJson) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapImportJson object to a map representation for JSON marshaling.
func (m MapImportJson) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp mapImportJson
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "import_all_floorplans", "import_height", "import_orientation", "vendor_name")
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

// mapImportJson is a temporary struct used for validating the fields of MapImportJson.
type mapImportJson  struct {
    ImportAllFloorplans *bool                        `json:"import_all_floorplans,omitempty"`
    ImportHeight        *bool                        `json:"import_height,omitempty"`
    ImportOrientation   *bool                        `json:"import_orientation,omitempty"`
    VendorName          *MapImportJsonVendorNameEnum `json:"vendor_name"`
}

func (m *mapImportJson) validate() error {
    var errs []string
    if m.VendorName == nil {
        errs = append(errs, "required field `vendor_name` is missing for type `Map_Import_Json`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
