package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// AssetOfInterest represents a AssetOfInterest struct.
type AssetOfInterest struct {
    ApMac                *string                `json:"ap_mac,omitempty"`
    Beam                 *float64               `json:"beam,omitempty"`
    By                   *string                `json:"by,omitempty"`
    CurrSite             *string                `json:"curr_site,omitempty"`
    DeviceName           *string                `json:"device_name,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    LastSeen             *float64               `json:"last_seen,omitempty"`
    Mac                  *string                `json:"mac,omitempty"`
    Manufacture          *string                `json:"manufacture,omitempty"`
    MapId                *string                `json:"map_id,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Rssi                 *float64               `json:"rssi,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AssetOfInterest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AssetOfInterest) String() string {
    return fmt.Sprintf(
    	"AssetOfInterest[ApMac=%v, Beam=%v, By=%v, CurrSite=%v, DeviceName=%v, Id=%v, LastSeen=%v, Mac=%v, Manufacture=%v, MapId=%v, Name=%v, Rssi=%v, AdditionalProperties=%v]",
    	a.ApMac, a.Beam, a.By, a.CurrSite, a.DeviceName, a.Id, a.LastSeen, a.Mac, a.Manufacture, a.MapId, a.Name, a.Rssi, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AssetOfInterest.
// It customizes the JSON marshaling process for AssetOfInterest objects.
func (a AssetOfInterest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "ap_mac", "beam", "by", "curr_site", "device_name", "id", "last_seen", "mac", "manufacture", "map_id", "name", "rssi"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AssetOfInterest object to a map representation for JSON marshaling.
func (a AssetOfInterest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ApMac != nil {
        structMap["ap_mac"] = a.ApMac
    }
    if a.Beam != nil {
        structMap["beam"] = a.Beam
    }
    if a.By != nil {
        structMap["by"] = a.By
    }
    if a.CurrSite != nil {
        structMap["curr_site"] = a.CurrSite
    }
    if a.DeviceName != nil {
        structMap["device_name"] = a.DeviceName
    }
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.LastSeen != nil {
        structMap["last_seen"] = a.LastSeen
    }
    if a.Mac != nil {
        structMap["mac"] = a.Mac
    }
    if a.Manufacture != nil {
        structMap["manufacture"] = a.Manufacture
    }
    if a.MapId != nil {
        structMap["map_id"] = a.MapId
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.Rssi != nil {
        structMap["rssi"] = a.Rssi
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AssetOfInterest.
// It customizes the JSON unmarshaling process for AssetOfInterest objects.
func (a *AssetOfInterest) UnmarshalJSON(input []byte) error {
    var temp tempAssetOfInterest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "beam", "by", "curr_site", "device_name", "id", "last_seen", "mac", "manufacture", "map_id", "name", "rssi")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.ApMac = temp.ApMac
    a.Beam = temp.Beam
    a.By = temp.By
    a.CurrSite = temp.CurrSite
    a.DeviceName = temp.DeviceName
    a.Id = temp.Id
    a.LastSeen = temp.LastSeen
    a.Mac = temp.Mac
    a.Manufacture = temp.Manufacture
    a.MapId = temp.MapId
    a.Name = temp.Name
    a.Rssi = temp.Rssi
    return nil
}

// tempAssetOfInterest is a temporary struct used for validating the fields of AssetOfInterest.
type tempAssetOfInterest  struct {
    ApMac       *string    `json:"ap_mac,omitempty"`
    Beam        *float64   `json:"beam,omitempty"`
    By          *string    `json:"by,omitempty"`
    CurrSite    *string    `json:"curr_site,omitempty"`
    DeviceName  *string    `json:"device_name,omitempty"`
    Id          *uuid.UUID `json:"id,omitempty"`
    LastSeen    *float64   `json:"last_seen,omitempty"`
    Mac         *string    `json:"mac,omitempty"`
    Manufacture *string    `json:"manufacture,omitempty"`
    MapId       *string    `json:"map_id,omitempty"`
    Name        *string    `json:"name,omitempty"`
    Rssi        *float64   `json:"rssi,omitempty"`
}
