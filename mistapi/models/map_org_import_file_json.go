package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// MapOrgImportFileJson represents a MapOrgImportFileJson struct.
type MapOrgImportFileJson struct {
    ImportAllFloorplans  *bool                              `json:"import_all_floorplans,omitempty"`
    ImportHeight         *bool                              `json:"import_height,omitempty"`
    ImportOrientation    *bool                              `json:"import_orientation,omitempty"`
    SiteId               *uuid.UUID                         `json:"site_id,omitempty"`
    // enum: `ekahau`, `ibwave`
    VendorName           MapOrgImportFileJsonVendorNameEnum `json:"vendor_name"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for MapOrgImportFileJson,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapOrgImportFileJson) String() string {
    return fmt.Sprintf(
    	"MapOrgImportFileJson[ImportAllFloorplans=%v, ImportHeight=%v, ImportOrientation=%v, SiteId=%v, VendorName=%v, AdditionalProperties=%v]",
    	m.ImportAllFloorplans, m.ImportHeight, m.ImportOrientation, m.SiteId, m.VendorName, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapOrgImportFileJson.
// It customizes the JSON marshaling process for MapOrgImportFileJson objects.
func (m MapOrgImportFileJson) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "import_all_floorplans", "import_height", "import_orientation", "site_id", "vendor_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MapOrgImportFileJson object to a map representation for JSON marshaling.
func (m MapOrgImportFileJson) toMap() map[string]any {
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
    if m.SiteId != nil {
        structMap["site_id"] = m.SiteId
    }
    structMap["vendor_name"] = m.VendorName
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapOrgImportFileJson.
// It customizes the JSON unmarshaling process for MapOrgImportFileJson objects.
func (m *MapOrgImportFileJson) UnmarshalJSON(input []byte) error {
    var temp tempMapOrgImportFileJson
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "import_all_floorplans", "import_height", "import_orientation", "site_id", "vendor_name")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.ImportAllFloorplans = temp.ImportAllFloorplans
    m.ImportHeight = temp.ImportHeight
    m.ImportOrientation = temp.ImportOrientation
    m.SiteId = temp.SiteId
    m.VendorName = *temp.VendorName
    return nil
}

// tempMapOrgImportFileJson is a temporary struct used for validating the fields of MapOrgImportFileJson.
type tempMapOrgImportFileJson  struct {
    ImportAllFloorplans *bool                               `json:"import_all_floorplans,omitempty"`
    ImportHeight        *bool                               `json:"import_height,omitempty"`
    ImportOrientation   *bool                               `json:"import_orientation,omitempty"`
    SiteId              *uuid.UUID                          `json:"site_id,omitempty"`
    VendorName          *MapOrgImportFileJsonVendorNameEnum `json:"vendor_name"`
}

func (m *tempMapOrgImportFileJson) validate() error {
    var errs []string
    if m.VendorName == nil {
        errs = append(errs, "required field `vendor_name` is missing for type `map_org_import_file_json`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
