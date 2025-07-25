// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// LatLng represents a LatLng struct.
type LatLng struct {
    Lat                  float64                `json:"lat"`
    Lng                  float64                `json:"lng"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for LatLng,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l LatLng) String() string {
    return fmt.Sprintf(
    	"LatLng[Lat=%v, Lng=%v, AdditionalProperties=%v]",
    	l.Lat, l.Lng, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for LatLng.
// It customizes the JSON marshaling process for LatLng objects.
func (l LatLng) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "lat", "lng"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the LatLng object to a map representation for JSON marshaling.
func (l LatLng) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["lat"] = l.Lat
    structMap["lng"] = l.Lng
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LatLng.
// It customizes the JSON unmarshaling process for LatLng objects.
func (l *LatLng) UnmarshalJSON(input []byte) error {
    var temp tempLatLng
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
    l.AdditionalProperties = additionalProperties
    
    l.Lat = *temp.Lat
    l.Lng = *temp.Lng
    return nil
}

// tempLatLng is a temporary struct used for validating the fields of LatLng.
type tempLatLng  struct {
    Lat *float64 `json:"lat"`
    Lng *float64 `json:"lng"`
}

func (l *tempLatLng) validate() error {
    var errs []string
    if l.Lat == nil {
        errs = append(errs, "required field `lat` is missing for type `lat_lng`")
    }
    if l.Lng == nil {
        errs = append(errs, "required field `lng` is missing for type `lat_lng`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
