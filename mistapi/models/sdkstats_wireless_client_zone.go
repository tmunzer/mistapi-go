package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// SdkstatsWirelessClientZone represents a SdkstatsWirelessClientZone struct.
type SdkstatsWirelessClientZone struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   uuid.UUID              `json:"id"`
    Since                float64                `json:"since"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SdkstatsWirelessClientZone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SdkstatsWirelessClientZone) String() string {
    return fmt.Sprintf(
    	"SdkstatsWirelessClientZone[Id=%v, Since=%v, AdditionalProperties=%v]",
    	s.Id, s.Since, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SdkstatsWirelessClientZone.
// It customizes the JSON marshaling process for SdkstatsWirelessClientZone objects.
func (s SdkstatsWirelessClientZone) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "since"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SdkstatsWirelessClientZone object to a map representation for JSON marshaling.
func (s SdkstatsWirelessClientZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["id"] = s.Id
    structMap["since"] = s.Since
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SdkstatsWirelessClientZone.
// It customizes the JSON unmarshaling process for SdkstatsWirelessClientZone objects.
func (s *SdkstatsWirelessClientZone) UnmarshalJSON(input []byte) error {
    var temp tempSdkstatsWirelessClientZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "since")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = *temp.Id
    s.Since = *temp.Since
    return nil
}

// tempSdkstatsWirelessClientZone is a temporary struct used for validating the fields of SdkstatsWirelessClientZone.
type tempSdkstatsWirelessClientZone  struct {
    Id    *uuid.UUID `json:"id"`
    Since *float64   `json:"since"`
}

func (s *tempSdkstatsWirelessClientZone) validate() error {
    var errs []string
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `sdkstats_wireless_client_zone`")
    }
    if s.Since == nil {
        errs = append(errs, "required field `since` is missing for type `sdkstats_wireless_client_zone`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
