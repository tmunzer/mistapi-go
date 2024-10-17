package models

import (
	"encoding/json"
)

// MxedgeUpgradeMultiAllowDowngrades represents a MxedgeUpgradeMultiAllowDowngrades struct.
// whether downgrade is allowed when running version is higher than expected version for each service
type MxedgeUpgradeMultiAllowDowngrades struct {
	Mxagent              *bool          `json:"mxagent,omitempty"`
	Mxdas                *bool          `json:"mxdas,omitempty"`
	Mxocproxy            *bool          `json:"mxocproxy,omitempty"`
	Radsecproxy          *bool          `json:"radsecproxy,omitempty"`
	Tunterm              *bool          `json:"tunterm,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeUpgradeMultiAllowDowngrades.
// It customizes the JSON marshaling process for MxedgeUpgradeMultiAllowDowngrades objects.
func (m MxedgeUpgradeMultiAllowDowngrades) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeUpgradeMultiAllowDowngrades object to a map representation for JSON marshaling.
func (m MxedgeUpgradeMultiAllowDowngrades) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Mxagent != nil {
		structMap["mxagent"] = m.Mxagent
	}
	if m.Mxdas != nil {
		structMap["mxdas"] = m.Mxdas
	}
	if m.Mxocproxy != nil {
		structMap["mxocproxy"] = m.Mxocproxy
	}
	if m.Radsecproxy != nil {
		structMap["radsecproxy"] = m.Radsecproxy
	}
	if m.Tunterm != nil {
		structMap["tunterm"] = m.Tunterm
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeUpgradeMultiAllowDowngrades.
// It customizes the JSON unmarshaling process for MxedgeUpgradeMultiAllowDowngrades objects.
func (m *MxedgeUpgradeMultiAllowDowngrades) UnmarshalJSON(input []byte) error {
	var temp tempMxedgeUpgradeMultiAllowDowngrades
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "mxagent", "mxdas", "mxocproxy", "radsecproxy", "tunterm")
	if err != nil {
		return err
	}

	m.AdditionalProperties = additionalProperties
	m.Mxagent = temp.Mxagent
	m.Mxdas = temp.Mxdas
	m.Mxocproxy = temp.Mxocproxy
	m.Radsecproxy = temp.Radsecproxy
	m.Tunterm = temp.Tunterm
	return nil
}

// tempMxedgeUpgradeMultiAllowDowngrades is a temporary struct used for validating the fields of MxedgeUpgradeMultiAllowDowngrades.
type tempMxedgeUpgradeMultiAllowDowngrades struct {
	Mxagent     *bool `json:"mxagent,omitempty"`
	Mxdas       *bool `json:"mxdas,omitempty"`
	Mxocproxy   *bool `json:"mxocproxy,omitempty"`
	Radsecproxy *bool `json:"radsecproxy,omitempty"`
	Tunterm     *bool `json:"tunterm,omitempty"`
}
