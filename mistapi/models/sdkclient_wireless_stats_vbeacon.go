package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// SdkclientWirelessStatsVbeacon represents a SdkclientWirelessStatsVbeacon struct.
type SdkclientWirelessStatsVbeacon struct {
    Id                   uuid.UUID      `json:"id"`
    Since                float64        `json:"since"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SdkclientWirelessStatsVbeacon.
// It customizes the JSON marshaling process for SdkclientWirelessStatsVbeacon objects.
func (s SdkclientWirelessStatsVbeacon) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SdkclientWirelessStatsVbeacon object to a map representation for JSON marshaling.
func (s SdkclientWirelessStatsVbeacon) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["id"] = s.Id
    structMap["since"] = s.Since
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SdkclientWirelessStatsVbeacon.
// It customizes the JSON unmarshaling process for SdkclientWirelessStatsVbeacon objects.
func (s *SdkclientWirelessStatsVbeacon) UnmarshalJSON(input []byte) error {
    var temp sdkclientWirelessStatsVbeacon
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

// sdkclientWirelessStatsVbeacon is a temporary struct used for validating the fields of SdkclientWirelessStatsVbeacon.
type sdkclientWirelessStatsVbeacon  struct {
    Id    *uuid.UUID `json:"id"`
    Since *float64   `json:"since"`
}

func (s *sdkclientWirelessStatsVbeacon) validate() error {
    var errs []string
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Sdkclient_Wireless_Stats_Vbeacon`")
    }
    if s.Since == nil {
        errs = append(errs, "required field `since` is missing for type `Sdkclient_Wireless_Stats_Vbeacon`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
