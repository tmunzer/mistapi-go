package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// SdkclientWirelessStatsZone represents a SdkclientWirelessStatsZone struct.
type SdkclientWirelessStatsZone struct {
    Id                   uuid.UUID      `json:"id"`
    Since                float64        `json:"since"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SdkclientWirelessStatsZone.
// It customizes the JSON marshaling process for SdkclientWirelessStatsZone objects.
func (s SdkclientWirelessStatsZone) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SdkclientWirelessStatsZone object to a map representation for JSON marshaling.
func (s SdkclientWirelessStatsZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["id"] = s.Id
    structMap["since"] = s.Since
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SdkclientWirelessStatsZone.
// It customizes the JSON unmarshaling process for SdkclientWirelessStatsZone objects.
func (s *SdkclientWirelessStatsZone) UnmarshalJSON(input []byte) error {
    var temp sdkclientWirelessStatsZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "since")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = *temp.Id
    s.Since = *temp.Since
    return nil
}

// sdkclientWirelessStatsZone is a temporary struct used for validating the fields of SdkclientWirelessStatsZone.
type sdkclientWirelessStatsZone  struct {
    Id    *uuid.UUID `json:"id"`
    Since *float64   `json:"since"`
}

func (s *sdkclientWirelessStatsZone) validate() error {
    var errs []string
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Sdkclient_Wireless_Stats_Zone`")
    }
    if s.Since == nil {
        errs = append(errs, "required field `since` is missing for type `Sdkclient_Wireless_Stats_Zone`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
