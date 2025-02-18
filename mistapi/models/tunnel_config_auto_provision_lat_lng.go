package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// TunnelConfigAutoProvisionLatLng represents a TunnelConfigAutoProvisionLatLng struct.
// API override for POP selection
type TunnelConfigAutoProvisionLatLng struct {
    Lat                  float64                `json:"lat"`
    Lng                  float64                `json:"lng"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfigAutoProvisionLatLng,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfigAutoProvisionLatLng) String() string {
    return fmt.Sprintf(
    	"TunnelConfigAutoProvisionLatLng[Lat=%v, Lng=%v, AdditionalProperties=%v]",
    	t.Lat, t.Lng, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigAutoProvisionLatLng.
// It customizes the JSON marshaling process for TunnelConfigAutoProvisionLatLng objects.
func (t TunnelConfigAutoProvisionLatLng) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "lat", "lng"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigAutoProvisionLatLng object to a map representation for JSON marshaling.
func (t TunnelConfigAutoProvisionLatLng) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    structMap["lat"] = t.Lat
    structMap["lng"] = t.Lng
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigAutoProvisionLatLng.
// It customizes the JSON unmarshaling process for TunnelConfigAutoProvisionLatLng objects.
func (t *TunnelConfigAutoProvisionLatLng) UnmarshalJSON(input []byte) error {
    var temp tempTunnelConfigAutoProvisionLatLng
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "lat", "lng")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.Lat = *temp.Lat
    t.Lng = *temp.Lng
    return nil
}

// tempTunnelConfigAutoProvisionLatLng is a temporary struct used for validating the fields of TunnelConfigAutoProvisionLatLng.
type tempTunnelConfigAutoProvisionLatLng  struct {
    Lat *float64 `json:"lat"`
    Lng *float64 `json:"lng"`
}

func (t *tempTunnelConfigAutoProvisionLatLng) validate() error {
    var errs []string
    if t.Lat == nil {
        errs = append(errs, "required field `lat` is missing for type `tunnel_config_auto_provision_lat_lng`")
    }
    if t.Lng == nil {
        errs = append(errs, "required field `lng` is missing for type `tunnel_config_auto_provision_lat_lng`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
