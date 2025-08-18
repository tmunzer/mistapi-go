// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// ResponseMxedgeUpgrade represents a ResponseMxedgeUpgrade struct.
type ResponseMxedgeUpgrade struct {
	Channel string                      `json:"channel"`
	Counts  MxedgeUpgradeResponseCounts `json:"counts"`
	// Unique ID of the object instance in the Mist Organization
	Id                   uuid.UUID              `json:"id"`
	Status               string                 `json:"status"`
	Strategy             string                 `json:"strategy"`
	Versions             interface{}            `json:"versions"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseMxedgeUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseMxedgeUpgrade) String() string {
	return fmt.Sprintf(
		"ResponseMxedgeUpgrade[Channel=%v, Counts=%v, Id=%v, Status=%v, Strategy=%v, Versions=%v, AdditionalProperties=%v]",
		r.Channel, r.Counts, r.Id, r.Status, r.Strategy, r.Versions, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseMxedgeUpgrade.
// It customizes the JSON marshaling process for ResponseMxedgeUpgrade objects.
func (r ResponseMxedgeUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"channel", "counts", "id", "status", "strategy", "versions"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseMxedgeUpgrade object to a map representation for JSON marshaling.
func (r ResponseMxedgeUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["channel"] = r.Channel
	structMap["counts"] = r.Counts.toMap()
	structMap["id"] = r.Id
	structMap["status"] = r.Status
	structMap["strategy"] = r.Strategy
	structMap["versions"] = r.Versions
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMxedgeUpgrade.
// It customizes the JSON unmarshaling process for ResponseMxedgeUpgrade objects.
func (r *ResponseMxedgeUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempResponseMxedgeUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "counts", "id", "status", "strategy", "versions")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Channel = *temp.Channel
	r.Counts = *temp.Counts
	r.Id = *temp.Id
	r.Status = *temp.Status
	r.Strategy = *temp.Strategy
	r.Versions = *temp.Versions
	return nil
}

// tempResponseMxedgeUpgrade is a temporary struct used for validating the fields of ResponseMxedgeUpgrade.
type tempResponseMxedgeUpgrade struct {
	Channel  *string                      `json:"channel"`
	Counts   *MxedgeUpgradeResponseCounts `json:"counts"`
	Id       *uuid.UUID                   `json:"id"`
	Status   *string                      `json:"status"`
	Strategy *string                      `json:"strategy"`
	Versions *interface{}                 `json:"versions"`
}

func (r *tempResponseMxedgeUpgrade) validate() error {
	var errs []string
	if r.Channel == nil {
		errs = append(errs, "required field `channel` is missing for type `response_mxedge_upgrade`")
	}
	if r.Counts == nil {
		errs = append(errs, "required field `counts` is missing for type `response_mxedge_upgrade`")
	}
	if r.Id == nil {
		errs = append(errs, "required field `id` is missing for type `response_mxedge_upgrade`")
	}
	if r.Status == nil {
		errs = append(errs, "required field `status` is missing for type `response_mxedge_upgrade`")
	}
	if r.Strategy == nil {
		errs = append(errs, "required field `strategy` is missing for type `response_mxedge_upgrade`")
	}
	if r.Versions == nil {
		errs = append(errs, "required field `versions` is missing for type `response_mxedge_upgrade`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
