package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// MxclusterNacClientIp represents a MxclusterNacClientIp struct.
type MxclusterNacClientIp struct {
    // if different from above
    Secret               *string                       `json:"secret,omitempty"`
    // present only for 3rd party clients
    SiteId               *uuid.UUID                    `json:"site_id,omitempty"`
    // convention to be followed is : "<vendor>-<variant>"
    // <variant> could be an os/platform/model/company
    // for ex: for cisco vendor, there could variants wrt os (such as ios, nxos etc), platforms (asa etc), or acquired companies (such as meraki, airnonet) etc.
    Vendor               *MxclusterNacClientVendorEnum `json:"vendor,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxclusterNacClientIp.
// It customizes the JSON marshaling process for MxclusterNacClientIp objects.
func (m MxclusterNacClientIp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxclusterNacClientIp object to a map representation for JSON marshaling.
func (m MxclusterNacClientIp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Secret != nil {
        structMap["secret"] = m.Secret
    }
    if m.SiteId != nil {
        structMap["site_id"] = m.SiteId
    }
    if m.Vendor != nil {
        structMap["vendor"] = m.Vendor
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterNacClientIp.
// It customizes the JSON unmarshaling process for MxclusterNacClientIp objects.
func (m *MxclusterNacClientIp) UnmarshalJSON(input []byte) error {
    var temp mxclusterNacClientIp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "secret", "site_id", "vendor")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Secret = temp.Secret
    m.SiteId = temp.SiteId
    m.Vendor = temp.Vendor
    return nil
}

// mxclusterNacClientIp is a temporary struct used for validating the fields of MxclusterNacClientIp.
type mxclusterNacClientIp  struct {
    Secret *string                       `json:"secret,omitempty"`
    SiteId *uuid.UUID                    `json:"site_id,omitempty"`
    Vendor *MxclusterNacClientVendorEnum `json:"vendor,omitempty"`
}