package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ResponseVirtualChassisConfig represents a ResponseVirtualChassisConfig struct.
type ResponseVirtualChassisConfig struct {
    ConfigType           *string                     `json:"config_type,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID                  `json:"id,omitempty"`
    Locating             *bool                       `json:"locating,omitempty"`
    Members              []StatsSwitchModuleStatItem `json:"members,omitempty"`
    Model                *string                     `json:"model,omitempty"`
    OrgId                *uuid.UUID                  `json:"org_id,omitempty"`
    Serial               *string                     `json:"serial,omitempty"`
    SiteId               *uuid.UUID                  `json:"site_id,omitempty"`
    Status               *string                     `json:"status,omitempty"`
    VcMac                *string                     `json:"vc_mac,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseVirtualChassisConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseVirtualChassisConfig) String() string {
    return fmt.Sprintf(
    	"ResponseVirtualChassisConfig[ConfigType=%v, Id=%v, Locating=%v, Members=%v, Model=%v, OrgId=%v, Serial=%v, SiteId=%v, Status=%v, VcMac=%v, AdditionalProperties=%v]",
    	r.ConfigType, r.Id, r.Locating, r.Members, r.Model, r.OrgId, r.Serial, r.SiteId, r.Status, r.VcMac, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseVirtualChassisConfig.
// It customizes the JSON marshaling process for ResponseVirtualChassisConfig objects.
func (r ResponseVirtualChassisConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "config_type", "id", "locating", "members", "model", "org_id", "serial", "site_id", "status", "vc_mac"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseVirtualChassisConfig object to a map representation for JSON marshaling.
func (r ResponseVirtualChassisConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ConfigType != nil {
        structMap["config_type"] = r.ConfigType
    }
    if r.Id != nil {
        structMap["id"] = r.Id
    }
    if r.Locating != nil {
        structMap["locating"] = r.Locating
    }
    if r.Members != nil {
        structMap["members"] = r.Members
    }
    if r.Model != nil {
        structMap["model"] = r.Model
    }
    if r.OrgId != nil {
        structMap["org_id"] = r.OrgId
    }
    if r.Serial != nil {
        structMap["serial"] = r.Serial
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    if r.VcMac != nil {
        structMap["vc_mac"] = r.VcMac
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseVirtualChassisConfig.
// It customizes the JSON unmarshaling process for ResponseVirtualChassisConfig objects.
func (r *ResponseVirtualChassisConfig) UnmarshalJSON(input []byte) error {
    var temp tempResponseVirtualChassisConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "config_type", "id", "locating", "members", "model", "org_id", "serial", "site_id", "status", "vc_mac")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ConfigType = temp.ConfigType
    r.Id = temp.Id
    r.Locating = temp.Locating
    r.Members = temp.Members
    r.Model = temp.Model
    r.OrgId = temp.OrgId
    r.Serial = temp.Serial
    r.SiteId = temp.SiteId
    r.Status = temp.Status
    r.VcMac = temp.VcMac
    return nil
}

// tempResponseVirtualChassisConfig is a temporary struct used for validating the fields of ResponseVirtualChassisConfig.
type tempResponseVirtualChassisConfig  struct {
    ConfigType *string                     `json:"config_type,omitempty"`
    Id         *uuid.UUID                  `json:"id,omitempty"`
    Locating   *bool                       `json:"locating,omitempty"`
    Members    []StatsSwitchModuleStatItem `json:"members,omitempty"`
    Model      *string                     `json:"model,omitempty"`
    OrgId      *uuid.UUID                  `json:"org_id,omitempty"`
    Serial     *string                     `json:"serial,omitempty"`
    SiteId     *uuid.UUID                  `json:"site_id,omitempty"`
    Status     *string                     `json:"status,omitempty"`
    VcMac      *string                     `json:"vc_mac,omitempty"`
}
