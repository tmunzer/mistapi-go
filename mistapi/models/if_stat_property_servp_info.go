package models

import (
    "encoding/json"
)

// IfStatPropertyServpInfo represents a IfStatPropertyServpInfo struct.
type IfStatPropertyServpInfo struct {
    Asn                  *string        `json:"asn,omitempty"`
    City                 *string        `json:"city,omitempty"`
    CountryCode          *string        `json:"country_code,omitempty"`
    Latitude             *float64       `json:"latitude,omitempty"`
    Longitude            *float64       `json:"longitude,omitempty"`
    Org                  *string        `json:"org,omitempty"`
    RegionCode           *string        `json:"region_code,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IfStatPropertyServpInfo.
// It customizes the JSON marshaling process for IfStatPropertyServpInfo objects.
func (i IfStatPropertyServpInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the IfStatPropertyServpInfo object to a map representation for JSON marshaling.
func (i IfStatPropertyServpInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Asn != nil {
        structMap["asn"] = i.Asn
    }
    if i.City != nil {
        structMap["city"] = i.City
    }
    if i.CountryCode != nil {
        structMap["country_code"] = i.CountryCode
    }
    if i.Latitude != nil {
        structMap["latitude"] = i.Latitude
    }
    if i.Longitude != nil {
        structMap["longitude"] = i.Longitude
    }
    if i.Org != nil {
        structMap["org"] = i.Org
    }
    if i.RegionCode != nil {
        structMap["region_code"] = i.RegionCode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IfStatPropertyServpInfo.
// It customizes the JSON unmarshaling process for IfStatPropertyServpInfo objects.
func (i *IfStatPropertyServpInfo) UnmarshalJSON(input []byte) error {
    var temp tempIfStatPropertyServpInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "asn", "city", "country_code", "latitude", "longitude", "org", "region_code")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Asn = temp.Asn
    i.City = temp.City
    i.CountryCode = temp.CountryCode
    i.Latitude = temp.Latitude
    i.Longitude = temp.Longitude
    i.Org = temp.Org
    i.RegionCode = temp.RegionCode
    return nil
}

// tempIfStatPropertyServpInfo is a temporary struct used for validating the fields of IfStatPropertyServpInfo.
type tempIfStatPropertyServpInfo  struct {
    Asn         *string  `json:"asn,omitempty"`
    City        *string  `json:"city,omitempty"`
    CountryCode *string  `json:"country_code,omitempty"`
    Latitude    *float64 `json:"latitude,omitempty"`
    Longitude   *float64 `json:"longitude,omitempty"`
    Org         *string  `json:"org,omitempty"`
    RegionCode  *string  `json:"region_code,omitempty"`
}
