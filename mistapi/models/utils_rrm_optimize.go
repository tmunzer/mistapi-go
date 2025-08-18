// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// UtilsRrmOptimize represents a UtilsRrmOptimize struct.
type UtilsRrmOptimize struct {
	// List of bands
	Bands []string `json:"bands"`
	// Targeting AP (neighbor APs may get changed, too), default is empty for ALL APs
	Macs []string `json:"macs,omitempty"`
	// Only changing TX Power (will not disconnect clients)
	TxpowerOnly          *bool                  `json:"txpower_only,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsRrmOptimize,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsRrmOptimize) String() string {
	return fmt.Sprintf(
		"UtilsRrmOptimize[Bands=%v, Macs=%v, TxpowerOnly=%v, AdditionalProperties=%v]",
		u.Bands, u.Macs, u.TxpowerOnly, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsRrmOptimize.
// It customizes the JSON marshaling process for UtilsRrmOptimize objects.
func (u UtilsRrmOptimize) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"bands", "macs", "txpower_only"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsRrmOptimize object to a map representation for JSON marshaling.
func (u UtilsRrmOptimize) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
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
	var temp tempUtilsRrmOptimize
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bands", "macs", "txpower_only")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Bands = *temp.Bands
	u.Macs = temp.Macs
	u.TxpowerOnly = temp.TxpowerOnly
	return nil
}

// tempUtilsRrmOptimize is a temporary struct used for validating the fields of UtilsRrmOptimize.
type tempUtilsRrmOptimize struct {
	Bands       *[]string `json:"bands"`
	Macs        []string  `json:"macs,omitempty"`
	TxpowerOnly *bool     `json:"txpower_only,omitempty"`
}

func (u *tempUtilsRrmOptimize) validate() error {
	var errs []string
	if u.Bands == nil {
		errs = append(errs, "required field `bands` is missing for type `utils_rrm_optimize`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
