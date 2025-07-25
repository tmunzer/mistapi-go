// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SsrUpgrade represents a SsrUpgrade struct.
type SsrUpgrade struct {
    // upgrade channel to follow. enum: `alpha`, `beta`, `stable`
    Channel              *SsrUpgradeChannelEnum `json:"channel,omitempty"`
    // Reboot start time in epoch seconds, default is start_time, -1 disables reboot
    RebootAt             *int                   `json:"reboot_at,omitempty"`
    // 128T firmware download start time in epoch seconds, default is now, -1 disables download
    StartTime            *int                   `json:"start_time,omitempty"`
    // 128T firmware version to upgrade (e.g. 5.3.0-93)
    Version              string                 `json:"version"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SsrUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SsrUpgrade) String() string {
    return fmt.Sprintf(
    	"SsrUpgrade[Channel=%v, RebootAt=%v, StartTime=%v, Version=%v, AdditionalProperties=%v]",
    	s.Channel, s.RebootAt, s.StartTime, s.Version, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SsrUpgrade.
// It customizes the JSON marshaling process for SsrUpgrade objects.
func (s SsrUpgrade) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "channel", "reboot_at", "start_time", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SsrUpgrade object to a map representation for JSON marshaling.
func (s SsrUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Channel != nil {
        structMap["channel"] = s.Channel
    }
    if s.RebootAt != nil {
        structMap["reboot_at"] = s.RebootAt
    }
    if s.StartTime != nil {
        structMap["start_time"] = s.StartTime
    }
    structMap["version"] = s.Version
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsrUpgrade.
// It customizes the JSON unmarshaling process for SsrUpgrade objects.
func (s *SsrUpgrade) UnmarshalJSON(input []byte) error {
    var temp tempSsrUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "reboot_at", "start_time", "version")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Channel = temp.Channel
    s.RebootAt = temp.RebootAt
    s.StartTime = temp.StartTime
    s.Version = *temp.Version
    return nil
}

// tempSsrUpgrade is a temporary struct used for validating the fields of SsrUpgrade.
type tempSsrUpgrade  struct {
    Channel   *SsrUpgradeChannelEnum `json:"channel,omitempty"`
    RebootAt  *int                   `json:"reboot_at,omitempty"`
    StartTime *int                   `json:"start_time,omitempty"`
    Version   *string                `json:"version"`
}

func (s *tempSsrUpgrade) validate() error {
    var errs []string
    if s.Version == nil {
        errs = append(errs, "required field `version` is missing for type `ssr_upgrade`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
