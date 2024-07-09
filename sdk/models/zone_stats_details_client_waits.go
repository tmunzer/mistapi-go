package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ZoneStatsDetailsClientWaits represents a ZoneStatsDetailsClientWaits struct.
// client wait time right now
type ZoneStatsDetailsClientWaits struct {
    // average wait time in seconds
    Avg                  int            `json:"avg"`
    // longest wait time in seconds
    Max                  int            `json:"max"`
    // shortest wait time in seconds
    Min                  int            `json:"min"`
    // 95th percentile of all the wait time(s)
    P95                  int            `json:"p95"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ZoneStatsDetailsClientWaits.
// It customizes the JSON marshaling process for ZoneStatsDetailsClientWaits objects.
func (z ZoneStatsDetailsClientWaits) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(z.toMap())
}

// toMap converts the ZoneStatsDetailsClientWaits object to a map representation for JSON marshaling.
func (z ZoneStatsDetailsClientWaits) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, z.AdditionalProperties)
    structMap["avg"] = z.Avg
    structMap["max"] = z.Max
    structMap["min"] = z.Min
    structMap["p95"] = z.P95
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ZoneStatsDetailsClientWaits.
// It customizes the JSON unmarshaling process for ZoneStatsDetailsClientWaits objects.
func (z *ZoneStatsDetailsClientWaits) UnmarshalJSON(input []byte) error {
    var temp zoneStatsDetailsClientWaits
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "avg", "max", "min", "p95")
    if err != nil {
    	return err
    }
    
    z.AdditionalProperties = additionalProperties
    z.Avg = *temp.Avg
    z.Max = *temp.Max
    z.Min = *temp.Min
    z.P95 = *temp.P95
    return nil
}

// zoneStatsDetailsClientWaits is a temporary struct used for validating the fields of ZoneStatsDetailsClientWaits.
type zoneStatsDetailsClientWaits  struct {
    Avg *int `json:"avg"`
    Max *int `json:"max"`
    Min *int `json:"min"`
    P95 *int `json:"p95"`
}

func (z *zoneStatsDetailsClientWaits) validate() error {
    var errs []string
    if z.Avg == nil {
        errs = append(errs, "required field `avg` is missing for type `Zone_Stats_Details_Client_Waits`")
    }
    if z.Max == nil {
        errs = append(errs, "required field `max` is missing for type `Zone_Stats_Details_Client_Waits`")
    }
    if z.Min == nil {
        errs = append(errs, "required field `min` is missing for type `Zone_Stats_Details_Client_Waits`")
    }
    if z.P95 == nil {
        errs = append(errs, "required field `p95` is missing for type `Zone_Stats_Details_Client_Waits`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
