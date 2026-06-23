// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseEventsDevices represents a ResponseEventsDevices struct.
// Paginated response for site device event search results
type ResponseEventsDevices struct {
	// Epoch timestamp for the end of the device event search window
	End int `json:"end"`
	// Maximum number of device event records returned in this page
	Limit int `json:"limit"`
	// Pagination cursor or URL for retrieving the next page of device event records
	Next *string `json:"next,omitempty"`
	// List of device event payloads
	Results []DeviceEvent `json:"results"`
	// Epoch timestamp for the start of the device event search window
	Start int `json:"start"`
	// Number of device event records matching the search filters across all pages
	Total                int                    `json:"total"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseEventsDevices,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseEventsDevices) String() string {
	return fmt.Sprintf(
		"ResponseEventsDevices[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Next, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseEventsDevices.
// It customizes the JSON marshaling process for ResponseEventsDevices objects.
func (r ResponseEventsDevices) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "next", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseEventsDevices object to a map representation for JSON marshaling.
func (r ResponseEventsDevices) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["end"] = r.End
	structMap["limit"] = r.Limit
	if r.Next != nil {
		structMap["next"] = r.Next
	}
	structMap["results"] = r.Results
	structMap["start"] = r.Start
	structMap["total"] = r.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseEventsDevices.
// It customizes the JSON unmarshaling process for ResponseEventsDevices objects.
func (r *ResponseEventsDevices) UnmarshalJSON(input []byte) error {
	var temp tempResponseEventsDevices
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
	r.AdditionalProperties = additionalProperties

	r.End = *temp.End
	r.Limit = *temp.Limit
	r.Next = temp.Next
	r.Results = *temp.Results
	r.Start = *temp.Start
	r.Total = *temp.Total
	return nil
}

// tempResponseEventsDevices is a temporary struct used for validating the fields of ResponseEventsDevices.
type tempResponseEventsDevices struct {
	End     *int           `json:"end"`
	Limit   *int           `json:"limit"`
	Next    *string        `json:"next,omitempty"`
	Results *[]DeviceEvent `json:"results"`
	Start   *int           `json:"start"`
	Total   *int           `json:"total"`
}

func (r *tempResponseEventsDevices) validate() error {
	var errs []string
	if r.End == nil {
		errs = append(errs, "required field `end` is missing for type `response_events_devices`")
	}
	if r.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `response_events_devices`")
	}
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_events_devices`")
	}
	if r.Start == nil {
		errs = append(errs, "required field `start` is missing for type `response_events_devices`")
	}
	if r.Total == nil {
		errs = append(errs, "required field `total` is missing for type `response_events_devices`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
