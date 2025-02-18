package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// VpnPeerStatSearch represents a VpnPeerStatSearch struct.
type VpnPeerStatSearch struct {
    End                  float64                `json:"end"`
    Limit                int                    `json:"limit"`
    Next                 *string                `json:"next,omitempty"`
    Results              []VpnPeerStat          `json:"results"`
    Start                float64                `json:"start"`
    Total                int                    `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VpnPeerStatSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VpnPeerStatSearch) String() string {
    return fmt.Sprintf(
    	"VpnPeerStatSearch[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	v.End, v.Limit, v.Next, v.Results, v.Start, v.Total, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VpnPeerStatSearch.
// It customizes the JSON marshaling process for VpnPeerStatSearch objects.
func (v VpnPeerStatSearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "end", "limit", "next", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VpnPeerStatSearch object to a map representation for JSON marshaling.
func (v VpnPeerStatSearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
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
    var temp tempVpnPeerStatSearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start", "total")
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

// tempVpnPeerStatSearch is a temporary struct used for validating the fields of VpnPeerStatSearch.
type tempVpnPeerStatSearch  struct {
    End     *float64       `json:"end"`
    Limit   *int           `json:"limit"`
    Next    *string        `json:"next,omitempty"`
    Results *[]VpnPeerStat `json:"results"`
    Start   *float64       `json:"start"`
    Total   *int           `json:"total"`
}

func (v *tempVpnPeerStatSearch) validate() error {
    var errs []string
    if v.End == nil {
        errs = append(errs, "required field `end` is missing for type `vpn_peer_stat_search`")
    }
    if v.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `vpn_peer_stat_search`")
    }
    if v.Results == nil {
        errs = append(errs, "required field `results` is missing for type `vpn_peer_stat_search`")
    }
    if v.Start == nil {
        errs = append(errs, "required field `start` is missing for type `vpn_peer_stat_search`")
    }
    if v.Total == nil {
        errs = append(errs, "required field `total` is missing for type `vpn_peer_stat_search`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
