package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UtilsRrmOptimize represents a UtilsRrmOptimize struct.
type UtilsRrmOptimize struct {
    // list of bands
    Bands                []string       `json:"bands"`
    // targeting AP (neighbor APs may get changed, too), default is empty for ALL APs
    Macs                 []string       `json:"macs,omitempty"`
    // only changng TX Power (will not disconnect clients)
    TxpowerOnly          *bool          `json:"txpower_only,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsRrmOptimize.
// It customizes the JSON marshaling process for UtilsRrmOptimize objects.
func (u UtilsRrmOptimize) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsRrmOptimize object to a map representation for JSON marshaling.
func (u UtilsRrmOptimize) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["bands"] = u.Bands
    if u.Macs != nil {
        structMap["macs"] = u.Macs
    }
    if u.TxpowerOnly != nil {
        structMap["txpower_only"] = u.TxpowerOnly
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsRrmOptimize.
// It customizes the JSON unmarshaling process for UtilsRrmOptimize objects.
func (u *UtilsRrmOptimize) UnmarshalJSON(input []byte) error {
    var temp utilsRrmOptimize
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bands", "macs", "txpower_only")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Bands = *temp.Bands
    u.Macs = temp.Macs
    u.TxpowerOnly = temp.TxpowerOnly
    return nil
}

// utilsRrmOptimize is a temporary struct used for validating the fields of UtilsRrmOptimize.
type utilsRrmOptimize  struct {
    Bands       *[]string `json:"bands"`
    Macs        []string  `json:"macs,omitempty"`
    TxpowerOnly *bool     `json:"txpower_only,omitempty"`
}

func (u *utilsRrmOptimize) validate() error {
    var errs []string
    if u.Bands == nil {
        errs = append(errs, "required field `bands` is missing for type `Utils_Rrm_Optimize`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
