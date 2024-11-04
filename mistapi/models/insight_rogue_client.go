package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// InsightRogueClient represents a InsightRogueClient struct.
type InsightRogueClient struct {
    Annotation           string         `json:"annotation"`
    ApMac                string         `json:"ap_mac"`
    AvgRssi              float64        `json:"avg_rssi"`
    Band                 string         `json:"band"`
    Bssid                string         `json:"bssid"`
    ClientMac            string         `json:"client_mac"`
    NumAps               int            `json:"num_aps"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InsightRogueClient.
// It customizes the JSON marshaling process for InsightRogueClient objects.
func (i InsightRogueClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InsightRogueClient object to a map representation for JSON marshaling.
func (i InsightRogueClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["annotation"] = i.Annotation
    structMap["ap_mac"] = i.ApMac
    structMap["avg_rssi"] = i.AvgRssi
    structMap["band"] = i.Band
    structMap["bssid"] = i.Bssid
    structMap["client_mac"] = i.ClientMac
    structMap["num_aps"] = i.NumAps
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InsightRogueClient.
// It customizes the JSON unmarshaling process for InsightRogueClient objects.
func (i *InsightRogueClient) UnmarshalJSON(input []byte) error {
    var temp tempInsightRogueClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "annotation", "ap_mac", "avg_rssi", "band", "bssid", "client_mac", "num_aps")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Annotation = *temp.Annotation
    i.ApMac = *temp.ApMac
    i.AvgRssi = *temp.AvgRssi
    i.Band = *temp.Band
    i.Bssid = *temp.Bssid
    i.ClientMac = *temp.ClientMac
    i.NumAps = *temp.NumAps
    return nil
}

// tempInsightRogueClient is a temporary struct used for validating the fields of InsightRogueClient.
type tempInsightRogueClient  struct {
    Annotation *string  `json:"annotation"`
    ApMac      *string  `json:"ap_mac"`
    AvgRssi    *float64 `json:"avg_rssi"`
    Band       *string  `json:"band"`
    Bssid      *string  `json:"bssid"`
    ClientMac  *string  `json:"client_mac"`
    NumAps     *int     `json:"num_aps"`
}

func (i *tempInsightRogueClient) validate() error {
    var errs []string
    if i.Annotation == nil {
        errs = append(errs, "required field `annotation` is missing for type `insight_rogue_client`")
    }
    if i.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `insight_rogue_client`")
    }
    if i.AvgRssi == nil {
        errs = append(errs, "required field `avg_rssi` is missing for type `insight_rogue_client`")
    }
    if i.Band == nil {
        errs = append(errs, "required field `band` is missing for type `insight_rogue_client`")
    }
    if i.Bssid == nil {
        errs = append(errs, "required field `bssid` is missing for type `insight_rogue_client`")
    }
    if i.ClientMac == nil {
        errs = append(errs, "required field `client_mac` is missing for type `insight_rogue_client`")
    }
    if i.NumAps == nil {
        errs = append(errs, "required field `num_aps` is missing for type `insight_rogue_client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
