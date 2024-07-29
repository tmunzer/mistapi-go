package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UtilsClearBgp represents a UtilsClearBgp struct.
type UtilsClearBgp struct {
    // neighbor ip-address or 'all'
    Neighbor             string                `json:"neighbor"`
    // enum: `hard`, `in`, `out`, `soft`
    Type                 UtilsClearBgpTypeEnum `json:"type"`
    // vrf name
    Vrf                  *string               `json:"vrf,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsClearBgp.
// It customizes the JSON marshaling process for UtilsClearBgp objects.
func (u UtilsClearBgp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsClearBgp object to a map representation for JSON marshaling.
func (u UtilsClearBgp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["neighbor"] = u.Neighbor
    structMap["type"] = u.Type
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsClearBgp.
// It customizes the JSON unmarshaling process for UtilsClearBgp objects.
func (u *UtilsClearBgp) UnmarshalJSON(input []byte) error {
    var temp utilsClearBgp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "neighbor", "type", "vrf")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Neighbor = *temp.Neighbor
    u.Type = *temp.Type
    u.Vrf = temp.Vrf
    return nil
}

// utilsClearBgp is a temporary struct used for validating the fields of UtilsClearBgp.
type utilsClearBgp  struct {
    Neighbor *string                `json:"neighbor"`
    Type     *UtilsClearBgpTypeEnum `json:"type"`
    Vrf      *string                `json:"vrf,omitempty"`
}

func (u *utilsClearBgp) validate() error {
    var errs []string
    if u.Neighbor == nil {
        errs = append(errs, "required field `neighbor` is missing for type `Utils_Clear_Bgp`")
    }
    if u.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Utils_Clear_Bgp`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
