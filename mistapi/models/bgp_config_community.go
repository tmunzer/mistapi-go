package models

import (
    "encoding/json"
)

// BgpConfigCommunity represents a BgpConfigCommunity struct.
type BgpConfigCommunity struct {
    Id                   *string        `json:"id,omitempty"`
    LocalPreference      *int           `json:"local_preference,omitempty"`
    VpnName              *string        `json:"vpn_name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BgpConfigCommunity.
// It customizes the JSON marshaling process for BgpConfigCommunity objects.
func (b BgpConfigCommunity) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BgpConfigCommunity object to a map representation for JSON marshaling.
func (b BgpConfigCommunity) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Id != nil {
        structMap["id"] = b.Id
    }
    if b.LocalPreference != nil {
        structMap["local_preference"] = b.LocalPreference
    }
    if b.VpnName != nil {
        structMap["vpn_name"] = b.VpnName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BgpConfigCommunity.
// It customizes the JSON unmarshaling process for BgpConfigCommunity objects.
func (b *BgpConfigCommunity) UnmarshalJSON(input []byte) error {
    var temp tempBgpConfigCommunity
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "local_preference", "vpn_name")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.Id = temp.Id
    b.LocalPreference = temp.LocalPreference
    b.VpnName = temp.VpnName
    return nil
}

// tempBgpConfigCommunity is a temporary struct used for validating the fields of BgpConfigCommunity.
type tempBgpConfigCommunity  struct {
    Id              *string `json:"id,omitempty"`
    LocalPreference *int    `json:"local_preference,omitempty"`
    VpnName         *string `json:"vpn_name,omitempty"`
}
