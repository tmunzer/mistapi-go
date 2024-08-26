package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// SdkstatsWirelessClientVbeacon represents a SdkstatsWirelessClientVbeacon struct.
type SdkstatsWirelessClientVbeacon struct {
    Id                   uuid.UUID      `json:"id"`
    Since                float64        `json:"since"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SdkstatsWirelessClientVbeacon.
// It customizes the JSON marshaling process for SdkstatsWirelessClientVbeacon objects.
func (s SdkstatsWirelessClientVbeacon) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SdkstatsWirelessClientVbeacon object to a map representation for JSON marshaling.
func (s SdkstatsWirelessClientVbeacon) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["id"] = s.Id
    structMap["since"] = s.Since
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SdkstatsWirelessClientVbeacon.
// It customizes the JSON unmarshaling process for SdkstatsWirelessClientVbeacon objects.
func (s *SdkstatsWirelessClientVbeacon) UnmarshalJSON(input []byte) error {
    var temp tempSdkstatsWirelessClientVbeacon
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

// tempSdkstatsWirelessClientVbeacon is a temporary struct used for validating the fields of SdkstatsWirelessClientVbeacon.
type tempSdkstatsWirelessClientVbeacon  struct {
    Id    *uuid.UUID `json:"id"`
    Since *float64   `json:"since"`
}

func (s *tempSdkstatsWirelessClientVbeacon) validate() error {
    var errs []string
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `sdkstats_wireless_client_vbeacon`")
    }
    if s.Since == nil {
        errs = append(errs, "required field `since` is missing for type `sdkstats_wireless_client_vbeacon`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
