package models

import (
    "encoding/json"
)

// ResponsePcapSearchItemPcapApsItem represents a ResponsePcapSearchItemPcapApsItem struct.
type ResponsePcapSearchItemPcapApsItem struct {
    Band                 *string          `json:"band,omitempty"`
    Bandwith             *string          `json:"bandwith,omitempty"`
    Channel              *int             `json:"channel,omitempty"`
    TcpdumpExpression    Optional[string] `json:"tcpdump_expression"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapSearchItemPcapApsItem.
// It customizes the JSON marshaling process for ResponsePcapSearchItemPcapApsItem objects.
func (r ResponsePcapSearchItemPcapApsItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapSearchItemPcapApsItem object to a map representation for JSON marshaling.
func (r ResponsePcapSearchItemPcapApsItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Band != nil {
        structMap["band"] = r.Band
    }
    if r.Bandwith != nil {
        structMap["bandwith"] = r.Bandwith
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "band", "bandwith", "channel", "tcpdump_expression")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Band = temp.Band
    r.Bandwith = temp.Bandwith
    r.Channel = temp.Channel
    r.TcpdumpExpression = temp.TcpdumpExpression
    return nil
}

// tempResponsePcapSearchItemPcapApsItem is a temporary struct used for validating the fields of ResponsePcapSearchItemPcapApsItem.
type tempResponsePcapSearchItemPcapApsItem  struct {
    Band              *string          `json:"band,omitempty"`
    Bandwith          *string          `json:"bandwith,omitempty"`
    Channel           *int             `json:"channel,omitempty"`
    TcpdumpExpression Optional[string] `json:"tcpdump_expression"`
}
