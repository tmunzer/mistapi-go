package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// NacRuleMatching represents a NacRuleMatching struct.
type NacRuleMatching struct {
    // enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `peap-tls`, `psk`
    AuthType             *NacAuthTypeEnum              `json:"auth_type,omitempty"`
    // List of client device families to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed family values
    Family               []string                      `json:"family,omitempty"`
    // List of client device models to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed model values
    Mfg                  []string                      `json:"mfg,omitempty"`
    // List of client device manufacturers to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed mfg values
    Model                []string                      `json:"model,omitempty"`
    Nactags              []string                      `json:"nactags,omitempty"`
    // List of client device os types to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed os_type values
    OsType               []string                      `json:"os_type,omitempty"`
    PortTypes            []NacRuleMatchingPortTypeEnum `json:"port_types,omitempty"`
    // List of site ids to match
    SiteIds              []uuid.UUID                   `json:"site_ids,omitempty"`
    // List of sitegroup ids to match
    SitegroupIds         []uuid.UUID                   `json:"sitegroup_ids,omitempty"`
    // List of vendors to match
    Vendor               []string                      `json:"vendor,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for NacRuleMatching,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacRuleMatching) String() string {
    return fmt.Sprintf(
    	"NacRuleMatching[AuthType=%v, Family=%v, Mfg=%v, Model=%v, Nactags=%v, OsType=%v, PortTypes=%v, SiteIds=%v, SitegroupIds=%v, Vendor=%v, AdditionalProperties=%v]",
    	n.AuthType, n.Family, n.Mfg, n.Model, n.Nactags, n.OsType, n.PortTypes, n.SiteIds, n.SitegroupIds, n.Vendor, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacRuleMatching.
// It customizes the JSON marshaling process for NacRuleMatching objects.
func (n NacRuleMatching) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "auth_type", "family", "mfg", "model", "nactags", "os_type", "port_types", "site_ids", "sitegroup_ids", "vendor"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NacRuleMatching object to a map representation for JSON marshaling.
func (n NacRuleMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.AuthType != nil {
        structMap["auth_type"] = n.AuthType
    }
    if n.Family != nil {
        structMap["family"] = n.Family
    }
    if n.Mfg != nil {
        structMap["mfg"] = n.Mfg
    }
    if n.Model != nil {
        structMap["model"] = n.Model
    }
    if n.Nactags != nil {
        structMap["nactags"] = n.Nactags
    }
    if n.OsType != nil {
        structMap["os_type"] = n.OsType
    }
    if n.PortTypes != nil {
        structMap["port_types"] = n.PortTypes
    }
    if n.SiteIds != nil {
        structMap["site_ids"] = n.SiteIds
    }
    if n.SitegroupIds != nil {
        structMap["sitegroup_ids"] = n.SitegroupIds
    }
    if n.Vendor != nil {
        structMap["vendor"] = n.Vendor
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacRuleMatching.
// It customizes the JSON unmarshaling process for NacRuleMatching objects.
func (n *NacRuleMatching) UnmarshalJSON(input []byte) error {
    var temp tempNacRuleMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_type", "family", "mfg", "model", "nactags", "os_type", "port_types", "site_ids", "sitegroup_ids", "vendor")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.AuthType = temp.AuthType
    n.Family = temp.Family
    n.Mfg = temp.Mfg
    n.Model = temp.Model
    n.Nactags = temp.Nactags
    n.OsType = temp.OsType
    n.PortTypes = temp.PortTypes
    n.SiteIds = temp.SiteIds
    n.SitegroupIds = temp.SitegroupIds
    n.Vendor = temp.Vendor
    return nil
}

// tempNacRuleMatching is a temporary struct used for validating the fields of NacRuleMatching.
type tempNacRuleMatching  struct {
    AuthType     *NacAuthTypeEnum              `json:"auth_type,omitempty"`
    Family       []string                      `json:"family,omitempty"`
    Mfg          []string                      `json:"mfg,omitempty"`
    Model        []string                      `json:"model,omitempty"`
    Nactags      []string                      `json:"nactags,omitempty"`
    OsType       []string                      `json:"os_type,omitempty"`
    PortTypes    []NacRuleMatchingPortTypeEnum `json:"port_types,omitempty"`
    SiteIds      []uuid.UUID                   `json:"site_ids,omitempty"`
    SitegroupIds []uuid.UUID                   `json:"sitegroup_ids,omitempty"`
    Vendor       []string                      `json:"vendor,omitempty"`
}
