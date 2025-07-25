// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// MxedgeUpgradeVersion represents a MxedgeUpgradeVersion struct.
// Version to upgrade for each service, `current` / `latest` / `default` / specific version (e.g. `2.5.100`).\nIgnored if distro upgrade, `tunterm`, `radsecproxy`, `mxagent`, `mxocproxy`, `mxdas` or `mxnacedge`
type MxedgeUpgradeVersion struct {
    Mxagent              string                 `json:"mxagent"`
    Mxdas                *string                `json:"mxdas,omitempty"`
    Mxocproxy            *string                `json:"mxocproxy,omitempty"`
    Radsecproxy          *string                `json:"radsecproxy,omitempty"`
    Tunterm              string                 `json:"tunterm"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeUpgradeVersion,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeUpgradeVersion) String() string {
    return fmt.Sprintf(
    	"MxedgeUpgradeVersion[Mxagent=%v, Mxdas=%v, Mxocproxy=%v, Radsecproxy=%v, Tunterm=%v, AdditionalProperties=%v]",
    	m.Mxagent, m.Mxdas, m.Mxocproxy, m.Radsecproxy, m.Tunterm, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeUpgradeVersion.
// It customizes the JSON marshaling process for MxedgeUpgradeVersion objects.
func (m MxedgeUpgradeVersion) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "mxagent", "mxdas", "mxocproxy", "radsecproxy", "tunterm"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeUpgradeVersion object to a map representation for JSON marshaling.
func (m MxedgeUpgradeVersion) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["mxagent"] = m.Mxagent
    if m.Mxdas != nil {
        structMap["mxdas"] = m.Mxdas
    }
    if m.Mxocproxy != nil {
        structMap["mxocproxy"] = m.Mxocproxy
    }
    if m.Radsecproxy != nil {
        structMap["radsecproxy"] = m.Radsecproxy
    }
    structMap["tunterm"] = m.Tunterm
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeUpgradeVersion.
// It customizes the JSON unmarshaling process for MxedgeUpgradeVersion objects.
func (m *MxedgeUpgradeVersion) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeUpgradeVersion
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mxagent", "mxdas", "mxocproxy", "radsecproxy", "tunterm")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Mxagent = *temp.Mxagent
    m.Mxdas = temp.Mxdas
    m.Mxocproxy = temp.Mxocproxy
    m.Radsecproxy = temp.Radsecproxy
    m.Tunterm = *temp.Tunterm
    return nil
}

// tempMxedgeUpgradeVersion is a temporary struct used for validating the fields of MxedgeUpgradeVersion.
type tempMxedgeUpgradeVersion  struct {
    Mxagent     *string `json:"mxagent"`
    Mxdas       *string `json:"mxdas,omitempty"`
    Mxocproxy   *string `json:"mxocproxy,omitempty"`
    Radsecproxy *string `json:"radsecproxy,omitempty"`
    Tunterm     *string `json:"tunterm"`
}

func (m *tempMxedgeUpgradeVersion) validate() error {
    var errs []string
    if m.Mxagent == nil {
        errs = append(errs, "required field `mxagent` is missing for type `mxedge_upgrade_version`")
    }
    if m.Tunterm == nil {
        errs = append(errs, "required field `tunterm` is missing for type `mxedge_upgrade_version`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
