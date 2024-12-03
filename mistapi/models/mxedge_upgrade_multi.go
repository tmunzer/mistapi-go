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
    AllowDowngrades      *MxedgeUpgradeMultiAllowDowngrades `json:"allow_downgrades,omitempty"`
    // Only if `strategy`==`canary`. Phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. default is [1, 10, 50, 100]
    CanaryPhases         []int                              `json:"canary_phases,omitempty"`
    // upgrade channel to follow. enum: `alpha`, `beta`, `stable`
    Channel              *MxedgeUpgradeChannelEnum          `json:"channel,omitempty"`
    // distro upgrade, optional, to specific codename (e.g. bullseye) with highest qualified versions
    Distro               *string                            `json:"distro,omitempty"`
    // Failure threshold before we stop the upgrade and mark it as failed
    MaxFailurePercentage *int                               `json:"max_failure_percentage,omitempty"`
    // list of mxedge IDs to upgrade. If not specified, it means all the org mxedges.
    MxedgeIds            []uuid.UUID                        `json:"mxedge_ids"`
    // upgrade start time in epoch seconds, default is now
    StartTime            *float64                           `json:"start_time,omitempty"`
    // enum:
    // * `big_bang`: upgrade all at once, no orchestration
    // * `serial`: one at a time'
    // * `canary`: upgrade in phases
    Strategy             *MxedgeUpgradeStrategyEnum         `json:"strategy,omitempty"`
    // version to upgrade for each service, `current` / `latest` / `default` / specific version (e.g. `2.5.100`).\nIgnored if distro upgrade, `tunterm`, `radsecproxy`, `mxagent`, `mxocproxy`, `mxdas` or `mxnacedge`
    Versions             *MxedgeUpgradeVersion              `json:"versions,omitempty"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeUpgradeMulti.
// It customizes the JSON marshaling process for MxedgeUpgradeMulti objects.
func (m MxedgeUpgradeMulti) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "allow_downgrades", "canary_phases", "channel", "distro", "max_failure_percentage", "mxedge_ids", "start_time", "strategy", "versions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeUpgradeMulti object to a map representation for JSON marshaling.
func (m MxedgeUpgradeMulti) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.AllowDowngrades != nil {
        structMap["allow_downgrades"] = m.AllowDowngrades.toMap()
    }
    if m.CanaryPhases != nil {
        structMap["canary_phases"] = m.CanaryPhases
    }
    if m.Channel != nil {
        structMap["channel"] = m.Channel
    }
    if m.Distro != nil {
        structMap["distro"] = m.Distro
    }
    if m.MaxFailurePercentage != nil {
        structMap["max_failure_percentage"] = m.MaxFailurePercentage
    }
    structMap["mxedge_ids"] = m.MxedgeIds
    if m.StartTime != nil {
        structMap["start_time"] = m.StartTime
    }
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allow_downgrades", "canary_phases", "channel", "distro", "max_failure_percentage", "mxedge_ids", "start_time", "strategy", "versions")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.AllowDowngrades = temp.AllowDowngrades
    m.CanaryPhases = temp.CanaryPhases
    m.Channel = temp.Channel
    m.Distro = temp.Distro
    m.MaxFailurePercentage = temp.MaxFailurePercentage
    m.MxedgeIds = *temp.MxedgeIds
    m.StartTime = temp.StartTime
    m.Strategy = temp.Strategy
    m.Versions = temp.Versions
    return nil
}

// tempMxedgeUpgradeMulti is a temporary struct used for validating the fields of MxedgeUpgradeMulti.
type tempMxedgeUpgradeMulti  struct {
    AllowDowngrades      *MxedgeUpgradeMultiAllowDowngrades `json:"allow_downgrades,omitempty"`
    CanaryPhases         []int                              `json:"canary_phases,omitempty"`
    Channel              *MxedgeUpgradeChannelEnum          `json:"channel,omitempty"`
    Distro               *string                            `json:"distro,omitempty"`
    MaxFailurePercentage *int                               `json:"max_failure_percentage,omitempty"`
    MxedgeIds            *[]uuid.UUID                       `json:"mxedge_ids"`
    StartTime            *float64                           `json:"start_time,omitempty"`
    Strategy             *MxedgeUpgradeStrategyEnum         `json:"strategy,omitempty"`
    Versions             *MxedgeUpgradeVersion              `json:"versions,omitempty"`
}

func (m *tempMxedgeUpgradeMulti) validate() error {
    var errs []string
    if m.MxedgeIds == nil {
        errs = append(errs, "required field `mxedge_ids` is missing for type `mxedge_upgrade_multi`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
