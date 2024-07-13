package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// VpnPeerStatSearch represents a VpnPeerStatSearch struct.
type VpnPeerStatSearch struct {
    End                  float64        `json:"end"`
    Limit                int            `json:"limit"`
    Next                 *string        `json:"next,omitempty"`
    Results              []VpnPeerStat  `json:"results"`
    Start                float64        `json:"start"`
    Total                int            `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VpnPeerStatSearch.
// It customizes the JSON marshaling process for VpnPeerStatSearch objects.
func (v VpnPeerStatSearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VpnPeerStatSearch object to a map representation for JSON marshaling.
func (v VpnPeerStatSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    structMap["end"] = v.End
    structMap["limit"] = v.Limit
    if v.Next != nil {
        structMap["next"] = v.Next
    }
    structMap["results"] = v.Results
    structMap["start"] = v.Start
    structMap["total"] = v.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VpnPeerStatSearch.
// It customizes the JSON unmarshaling process for VpnPeerStatSearch objects.
func (v *VpnPeerStatSearch) UnmarshalJSON(input []byte) error {
    var temp vpnPeerStatSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "next", "results", "start", "total")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.End = *temp.End
    v.Limit = *temp.Limit
    v.Next = temp.Next
    v.Results = *temp.Results
    v.Start = *temp.Start
    v.Total = *temp.Total
    return nil
}

// vpnPeerStatSearch is a temporary struct used for validating the fields of VpnPeerStatSearch.
type vpnPeerStatSearch  struct {
    End     *float64       `json:"end"`
    Limit   *int           `json:"limit"`
    Next    *string        `json:"next,omitempty"`
    Results *[]VpnPeerStat `json:"results"`
    Start   *float64       `json:"start"`
    Total   *int           `json:"total"`
}

func (v *vpnPeerStatSearch) validate() error {
    var errs []string
    if v.End == nil {
        errs = append(errs, "required field `end` is missing for type `Vpn_Peer_Stat_Search`")
    }
    if v.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `Vpn_Peer_Stat_Search`")
    }
    if v.Results == nil {
        errs = append(errs, "required field `results` is missing for type `Vpn_Peer_Stat_Search`")
    }
    if v.Start == nil {
        errs = append(errs, "required field `start` is missing for type `Vpn_Peer_Stat_Search`")
    }
    if v.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Vpn_Peer_Stat_Search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
