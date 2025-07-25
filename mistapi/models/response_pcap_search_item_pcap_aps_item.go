// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponsePcapSearchItemPcapApsItem represents a ResponsePcapSearchItemPcapApsItem struct.
type ResponsePcapSearchItemPcapApsItem struct {
    Band                 *string                `json:"band,omitempty"`
    Bandwidth            *string                `json:"bandwidth,omitempty"`
    Channel              *int                   `json:"channel,omitempty"`
    TcpdumpExpression    Optional[string]       `json:"tcpdump_expression"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePcapSearchItemPcapApsItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePcapSearchItemPcapApsItem) String() string {
    return fmt.Sprintf(
    	"ResponsePcapSearchItemPcapApsItem[Band=%v, Bandwidth=%v, Channel=%v, TcpdumpExpression=%v, AdditionalProperties=%v]",
    	r.Band, r.Bandwidth, r.Channel, r.TcpdumpExpression, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapSearchItemPcapApsItem.
// It customizes the JSON marshaling process for ResponsePcapSearchItemPcapApsItem objects.
func (r ResponsePcapSearchItemPcapApsItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "band", "bandwidth", "channel", "tcpdump_expression"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapSearchItemPcapApsItem object to a map representation for JSON marshaling.
func (r ResponsePcapSearchItemPcapApsItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Band != nil {
        structMap["band"] = r.Band
    }
    if r.Bandwidth != nil {
        structMap["bandwidth"] = r.Bandwidth
    }
    if r.Channel != nil {
        structMap["channel"] = r.Channel
    }
    if r.TcpdumpExpression.IsValueSet() {
        if r.TcpdumpExpression.Value() != nil {
            structMap["tcpdump_expression"] = r.TcpdumpExpression.Value()
        } else {
            structMap["tcpdump_expression"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePcapSearchItemPcapApsItem.
// It customizes the JSON unmarshaling process for ResponsePcapSearchItemPcapApsItem objects.
func (r *ResponsePcapSearchItemPcapApsItem) UnmarshalJSON(input []byte) error {
    var temp tempResponsePcapSearchItemPcapApsItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "bandwidth", "channel", "tcpdump_expression")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Band = temp.Band
    r.Bandwidth = temp.Bandwidth
    r.Channel = temp.Channel
    r.TcpdumpExpression = temp.TcpdumpExpression
    return nil
}

// tempResponsePcapSearchItemPcapApsItem is a temporary struct used for validating the fields of ResponsePcapSearchItemPcapApsItem.
type tempResponsePcapSearchItemPcapApsItem  struct {
    Band              *string          `json:"band,omitempty"`
    Bandwidth         *string          `json:"bandwidth,omitempty"`
    Channel           *int             `json:"channel,omitempty"`
    TcpdumpExpression Optional[string] `json:"tcpdump_expression"`
}
