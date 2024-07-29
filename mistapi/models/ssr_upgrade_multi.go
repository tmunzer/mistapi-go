package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// SsrUpgradeMulti represents a SsrUpgradeMulti struct.
type SsrUpgradeMulti struct {
    // upgrade channel to follow. enum: `alpha`, `beta`, `stable`
    Channel              *SsrUpgradeChannelEnum  `json:"channel,omitempty"`
    // list of 128T device IDs to upgrade
    DeviceIds            []uuid.UUID             `json:"device_ids"`
    // reboot start time in epoch seconds, default is start_time, -1 disables reboot
    RebootAt             *int                    `json:"reboot_at,omitempty"`
    // 128T firmware download start time in epoch seconds, default is now, -1 disables download
    StartTime            *int                    `json:"start_time,omitempty"`
    // enum:
    //   * `big_bang`: upgrade all at once
    //   * `serial`: one at a time
    Strategy             *SsrUpgradeStrategyEnum `json:"strategy,omitempty"`
    // 128T firmware version to upgrade (e.g. 5.3.0-93)
    Version              *string                 `json:"version,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SsrUpgradeMulti.
// It customizes the JSON marshaling process for SsrUpgradeMulti objects.
func (s SsrUpgradeMulti) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SsrUpgradeMulti object to a map representation for JSON marshaling.
func (s SsrUpgradeMulti) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Channel != nil {
        structMap["channel"] = s.Channel
    }
    structMap["device_ids"] = s.DeviceIds
    if s.RebootAt != nil {
        structMap["reboot_at"] = s.RebootAt
    }
    if s.StartTime != nil {
        structMap["start_time"] = s.StartTime
    }
    if s.Strategy != nil {
        structMap["strategy"] = s.Strategy
    }
    if s.Version != nil {
        structMap["version"] = s.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsrUpgradeMulti.
// It customizes the JSON unmarshaling process for SsrUpgradeMulti objects.
func (s *SsrUpgradeMulti) UnmarshalJSON(input []byte) error {
    var temp ssrUpgradeMulti
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "channel", "device_ids", "reboot_at", "start_time", "strategy", "version")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Channel = temp.Channel
    s.DeviceIds = *temp.DeviceIds
    s.RebootAt = temp.RebootAt
    s.StartTime = temp.StartTime
    s.Strategy = temp.Strategy
    s.Version = temp.Version
    return nil
}

// ssrUpgradeMulti is a temporary struct used for validating the fields of SsrUpgradeMulti.
type ssrUpgradeMulti  struct {
    Channel   *SsrUpgradeChannelEnum  `json:"channel,omitempty"`
    DeviceIds *[]uuid.UUID            `json:"device_ids"`
    RebootAt  *int                    `json:"reboot_at,omitempty"`
    StartTime *int                    `json:"start_time,omitempty"`
    Strategy  *SsrUpgradeStrategyEnum `json:"strategy,omitempty"`
    Version   *string                 `json:"version,omitempty"`
}

func (s *ssrUpgradeMulti) validate() error {
    var errs []string
    if s.DeviceIds == nil {
        errs = append(errs, "required field `device_ids` is missing for type `Ssr_Upgrade_Multi`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
