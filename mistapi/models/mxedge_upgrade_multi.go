package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// MxedgeUpgradeMulti represents a MxedgeUpgradeMulti struct.
type MxedgeUpgradeMulti struct {
	// whether downgrade is allowed when running version is higher than expected version for each service
	AllowDowngrades *MxedgeUpgradeMultiAllowDowngrades `json:"allow_downgrades,omitempty"`
	// upgrade channel to follow. enum: `alpha`, `beta`, `stable`
	Channel *MxedgeUpgradeChannelEnum `json:"channel,omitempty"`
	// distro upgrade, optional, to specific codename (e.g. bullseye) with highest qualified versions
	Distro *string `json:"distro,omitempty"`
	// list of mxedge IDs to upgrade. If not specified, it means all the org mxedges.
	MxedgeIds []uuid.UUID `json:"mxedge_ids"`
	// enum:
	// * `big_bang`: upgrade all at once
	// * `serial`: one at a time. enum: `big_bang`, `serial`'
	Strategy *MxedgeUpgradeStrategyEnum `json:"strategy,omitempty"`
	// version to upgrade for each service, `current` / `latest` / `default` / specific version (e.g. `2.5.100`).\nIgnored if distro upgrade, `tunterm`, `radsecproxy`, `mxagent`, `mxocproxy`, `mxdas` or `mxnacedge`
	Versions             *MxedgeUpgradeVersion `json:"versions,omitempty"`
	AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeUpgradeMulti.
// It customizes the JSON marshaling process for MxedgeUpgradeMulti objects.
func (m MxedgeUpgradeMulti) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MxedgeUpgradeMulti object to a map representation for JSON marshaling.
func (m MxedgeUpgradeMulti) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, m.AdditionalProperties)
	if m.AllowDowngrades != nil {
		structMap["allow_downgrades"] = m.AllowDowngrades.toMap()
	}
	if m.Channel != nil {
		structMap["channel"] = m.Channel
	}
	if m.Distro != nil {
		structMap["distro"] = m.Distro
	}
	structMap["mxedge_ids"] = m.MxedgeIds
	if m.Strategy != nil {
		structMap["strategy"] = m.Strategy
	}
	if m.Versions != nil {
		structMap["versions"] = m.Versions.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeUpgradeMulti.
// It customizes the JSON unmarshaling process for MxedgeUpgradeMulti objects.
func (m *MxedgeUpgradeMulti) UnmarshalJSON(input []byte) error {
	var temp tempMxedgeUpgradeMulti
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "allow_downgrades", "channel", "distro", "mxedge_ids", "strategy", "versions")
	if err != nil {
		return err
	}

	m.AdditionalProperties = additionalProperties
	m.AllowDowngrades = temp.AllowDowngrades
	m.Channel = temp.Channel
	m.Distro = temp.Distro
	m.MxedgeIds = *temp.MxedgeIds
	m.Strategy = temp.Strategy
	m.Versions = temp.Versions
	return nil
}

// tempMxedgeUpgradeMulti is a temporary struct used for validating the fields of MxedgeUpgradeMulti.
type tempMxedgeUpgradeMulti struct {
	AllowDowngrades *MxedgeUpgradeMultiAllowDowngrades `json:"allow_downgrades,omitempty"`
	Channel         *MxedgeUpgradeChannelEnum          `json:"channel,omitempty"`
	Distro          *string                            `json:"distro,omitempty"`
	MxedgeIds       *[]uuid.UUID                       `json:"mxedge_ids"`
	Strategy        *MxedgeUpgradeStrategyEnum         `json:"strategy,omitempty"`
	Versions        *MxedgeUpgradeVersion              `json:"versions,omitempty"`
}

func (m *tempMxedgeUpgradeMulti) validate() error {
	var errs []string
	if m.MxedgeIds == nil {
		errs = append(errs, "required field `mxedge_ids` is missing for type `mxedge_upgrade_multi`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
