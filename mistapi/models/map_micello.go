package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// MapMicello represents a MapMicello struct.
type MapMicello struct {
    // the account key that has access to the map
    AccountKey           string         `json:"account_key"`
    // micello floor/level id
    DefaultLevelId       int            `json:"default_level_id"`
    // micello map id
    MapId                uuid.UUID      `json:"map_id"`
    // the vendor ‘micello’. enum: `micello`
    VendorName           string         `json:"vendor_name"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MapMicello.
// It customizes the JSON marshaling process for MapMicello objects.
func (m MapMicello) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MapMicello object to a map representation for JSON marshaling.
func (m MapMicello) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["account_key"] = m.AccountKey
    structMap["default_level_id"] = m.DefaultLevelId
    structMap["map_id"] = m.MapId
    structMap["vendor_name"] = m.VendorName
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapMicello.
// It customizes the JSON unmarshaling process for MapMicello objects.
func (m *MapMicello) UnmarshalJSON(input []byte) error {
    var temp mapMicello
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "account_key", "default_level_id", "map_id", "vendor_name")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.AccountKey = *temp.AccountKey
    m.DefaultLevelId = *temp.DefaultLevelId
    m.MapId = *temp.MapId
    m.VendorName = *temp.VendorName
    return nil
}

// mapMicello is a temporary struct used for validating the fields of MapMicello.
type mapMicello  struct {
    AccountKey     *string    `json:"account_key"`
    DefaultLevelId *int       `json:"default_level_id"`
    MapId          *uuid.UUID `json:"map_id"`
    VendorName     *string    `json:"vendor_name"`
}

func (m *mapMicello) validate() error {
    var errs []string
    if m.AccountKey == nil {
        errs = append(errs, "required field `account_key` is missing for type `Map_Micello`")
    }
    if m.DefaultLevelId == nil {
        errs = append(errs, "required field `default_level_id` is missing for type `Map_Micello`")
    }
    if m.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `Map_Micello`")
    }
    if m.VendorName == nil {
        errs = append(errs, "required field `vendor_name` is missing for type `Map_Micello`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
