package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingCloudshark represents a OrgSettingCloudshark struct.
type OrgSettingCloudshark struct {
    Apitoken             *string                `json:"apitoken,omitempty"`
    // If using CS Enteprise
    Url                  *string                `json:"url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingCloudshark,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingCloudshark) String() string {
    return fmt.Sprintf(
    	"OrgSettingCloudshark[Apitoken=%v, Url=%v, AdditionalProperties=%v]",
    	o.Apitoken, o.Url, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingCloudshark.
// It customizes the JSON marshaling process for OrgSettingCloudshark objects.
func (o OrgSettingCloudshark) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "apitoken", "url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingCloudshark object to a map representation for JSON marshaling.
func (o OrgSettingCloudshark) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Apitoken != nil {
        structMap["apitoken"] = o.Apitoken
    }
    if o.Url != nil {
        structMap["url"] = o.Url
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingCloudshark.
// It customizes the JSON unmarshaling process for OrgSettingCloudshark objects.
func (o *OrgSettingCloudshark) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingCloudshark
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "apitoken", "url")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Apitoken = temp.Apitoken
    o.Url = temp.Url
    return nil
}

// tempOrgSettingCloudshark is a temporary struct used for validating the fields of OrgSettingCloudshark.
type tempOrgSettingCloudshark  struct {
    Apitoken *string `json:"apitoken,omitempty"`
    Url      *string `json:"url,omitempty"`
}
