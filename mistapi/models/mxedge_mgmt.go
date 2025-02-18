package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeMgmt represents a MxedgeMgmt struct.
type MxedgeMgmt struct {
    FipsEnabled          *bool                     `json:"fips_enabled,omitempty"`
    MistPassword         *string                   `json:"mist_password,omitempty"`
    // enum: `dhcp`, `disabled`, `static`
    OobIpType            *MxedgeMgmtOobIpTypeEnum  `json:"oob_ip_type,omitempty"`
    // enum: `autoconf`, `dhcp`, `disabled`, `static`
    OobIpType6           *MxedgeMgmtOobIpType6Enum `json:"oob_ip_type6,omitempty"`
    RootPassword         *string                   `json:"root_password,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeMgmt,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeMgmt) String() string {
    return fmt.Sprintf(
    	"MxedgeMgmt[FipsEnabled=%v, MistPassword=%v, OobIpType=%v, OobIpType6=%v, RootPassword=%v, AdditionalProperties=%v]",
    	m.FipsEnabled, m.MistPassword, m.OobIpType, m.OobIpType6, m.RootPassword, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeMgmt.
// It customizes the JSON marshaling process for MxedgeMgmt objects.
func (m MxedgeMgmt) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "fips_enabled", "mist_password", "oob_ip_type", "oob_ip_type6", "root_password"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeMgmt object to a map representation for JSON marshaling.
func (m MxedgeMgmt) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.FipsEnabled != nil {
        structMap["fips_enabled"] = m.FipsEnabled
    }
    if m.MistPassword != nil {
        structMap["mist_password"] = m.MistPassword
    }
    if m.OobIpType != nil {
        structMap["oob_ip_type"] = m.OobIpType
    }
    if m.OobIpType6 != nil {
        structMap["oob_ip_type6"] = m.OobIpType6
    }
    if m.RootPassword != nil {
        structMap["root_password"] = m.RootPassword
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeMgmt.
// It customizes the JSON unmarshaling process for MxedgeMgmt objects.
func (m *MxedgeMgmt) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeMgmt
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "fips_enabled", "mist_password", "oob_ip_type", "oob_ip_type6", "root_password")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.FipsEnabled = temp.FipsEnabled
    m.MistPassword = temp.MistPassword
    m.OobIpType = temp.OobIpType
    m.OobIpType6 = temp.OobIpType6
    m.RootPassword = temp.RootPassword
    return nil
}

// tempMxedgeMgmt is a temporary struct used for validating the fields of MxedgeMgmt.
type tempMxedgeMgmt  struct {
    FipsEnabled  *bool                     `json:"fips_enabled,omitempty"`
    MistPassword *string                   `json:"mist_password,omitempty"`
    OobIpType    *MxedgeMgmtOobIpTypeEnum  `json:"oob_ip_type,omitempty"`
    OobIpType6   *MxedgeMgmtOobIpType6Enum `json:"oob_ip_type6,omitempty"`
    RootPassword *string                   `json:"root_password,omitempty"`
}
