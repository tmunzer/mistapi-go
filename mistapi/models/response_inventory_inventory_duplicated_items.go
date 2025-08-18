// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseInventoryInventoryDuplicatedItems represents a ResponseInventoryInventoryDuplicatedItems struct.
type ResponseInventoryInventoryDuplicatedItems struct {
	Mac                  string                 `json:"mac"`
	Magic                string                 `json:"magic"`
	Model                string                 `json:"model"`
	Serial               string                 `json:"serial"`
	Type                 string                 `json:"type"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseInventoryInventoryDuplicatedItems,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseInventoryInventoryDuplicatedItems) String() string {
	return fmt.Sprintf(
		"ResponseInventoryInventoryDuplicatedItems[Mac=%v, Magic=%v, Model=%v, Serial=%v, Type=%v, AdditionalProperties=%v]",
		r.Mac, r.Magic, r.Model, r.Serial, r.Type, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseInventoryInventoryDuplicatedItems.
// It customizes the JSON marshaling process for ResponseInventoryInventoryDuplicatedItems objects.
func (r ResponseInventoryInventoryDuplicatedItems) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"mac", "magic", "model", "serial", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseInventoryInventoryDuplicatedItems object to a map representation for JSON marshaling.
func (r ResponseInventoryInventoryDuplicatedItems) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["mac"] = r.Mac
	structMap["magic"] = r.Magic
	structMap["model"] = r.Model
	structMap["serial"] = r.Serial
	structMap["type"] = r.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseInventoryInventoryDuplicatedItems.
// It customizes the JSON unmarshaling process for ResponseInventoryInventoryDuplicatedItems objects.
func (r *ResponseInventoryInventoryDuplicatedItems) UnmarshalJSON(input []byte) error {
	var temp tempResponseInventoryInventoryDuplicatedItems
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "magic", "model", "serial", "type")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Mac = *temp.Mac
	r.Magic = *temp.Magic
	r.Model = *temp.Model
	r.Serial = *temp.Serial
	r.Type = *temp.Type
	return nil
}

// tempResponseInventoryInventoryDuplicatedItems is a temporary struct used for validating the fields of ResponseInventoryInventoryDuplicatedItems.
type tempResponseInventoryInventoryDuplicatedItems struct {
	Mac    *string `json:"mac"`
	Magic  *string `json:"magic"`
	Model  *string `json:"model"`
	Serial *string `json:"serial"`
	Type   *string `json:"type"`
}

func (r *tempResponseInventoryInventoryDuplicatedItems) validate() error {
	var errs []string
	if r.Mac == nil {
		errs = append(errs, "required field `mac` is missing for type `response_inventory_inventory_duplicated_items`")
	}
	if r.Magic == nil {
		errs = append(errs, "required field `magic` is missing for type `response_inventory_inventory_duplicated_items`")
	}
	if r.Model == nil {
		errs = append(errs, "required field `model` is missing for type `response_inventory_inventory_duplicated_items`")
	}
	if r.Serial == nil {
		errs = append(errs, "required field `serial` is missing for type `response_inventory_inventory_duplicated_items`")
	}
	if r.Type == nil {
		errs = append(errs, "required field `type` is missing for type `response_inventory_inventory_duplicated_items`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
