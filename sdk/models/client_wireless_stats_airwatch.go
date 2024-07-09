package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ClientWirelessStatsAirwatch represents a ClientWirelessStatsAirwatch struct.
// information if airwatch enabled
type ClientWirelessStatsAirwatch struct {
    Authorized           bool           `json:"authorized"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientWirelessStatsAirwatch.
// It customizes the JSON marshaling process for ClientWirelessStatsAirwatch objects.
func (c ClientWirelessStatsAirwatch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClientWirelessStatsAirwatch object to a map representation for JSON marshaling.
func (c ClientWirelessStatsAirwatch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["authorized"] = c.Authorized
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientWirelessStatsAirwatch.
// It customizes the JSON unmarshaling process for ClientWirelessStatsAirwatch objects.
func (c *ClientWirelessStatsAirwatch) UnmarshalJSON(input []byte) error {
    var temp clientWirelessStatsAirwatch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "authorized")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Authorized = *temp.Authorized
    return nil
}

// clientWirelessStatsAirwatch is a temporary struct used for validating the fields of ClientWirelessStatsAirwatch.
type clientWirelessStatsAirwatch  struct {
    Authorized *bool `json:"authorized"`
}

func (c *clientWirelessStatsAirwatch) validate() error {
    var errs []string
    if c.Authorized == nil {
        errs = append(errs, "required field `authorized` is missing for type `Client_Wireless_Stats_Airwatch`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
