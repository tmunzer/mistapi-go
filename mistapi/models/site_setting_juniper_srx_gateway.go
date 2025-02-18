package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingJuniperSrxGateway represents a SiteSettingJuniperSrxGateway struct.
type SiteSettingJuniperSrxGateway struct {
    ApiKey               *string                `json:"api_key,omitempty"`
    ApiUrl               *string                `json:"api_url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingJuniperSrxGateway,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingJuniperSrxGateway) String() string {
    return fmt.Sprintf(
    	"SiteSettingJuniperSrxGateway[ApiKey=%v, ApiUrl=%v, AdditionalProperties=%v]",
    	s.ApiKey, s.ApiUrl, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingJuniperSrxGateway.
// It customizes the JSON marshaling process for SiteSettingJuniperSrxGateway objects.
func (s SiteSettingJuniperSrxGateway) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "api_key", "api_url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingJuniperSrxGateway object to a map representation for JSON marshaling.
func (s SiteSettingJuniperSrxGateway) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ApiKey != nil {
        structMap["api_key"] = s.ApiKey
    }
    if s.ApiUrl != nil {
        structMap["api_url"] = s.ApiUrl
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingJuniperSrxGateway.
// It customizes the JSON unmarshaling process for SiteSettingJuniperSrxGateway objects.
func (s *SiteSettingJuniperSrxGateway) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingJuniperSrxGateway
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "api_key", "api_url")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ApiKey = temp.ApiKey
    s.ApiUrl = temp.ApiUrl
    return nil
}

// tempSiteSettingJuniperSrxGateway is a temporary struct used for validating the fields of SiteSettingJuniperSrxGateway.
type tempSiteSettingJuniperSrxGateway  struct {
    ApiKey *string `json:"api_key,omitempty"`
    ApiUrl *string `json:"api_url,omitempty"`
}
