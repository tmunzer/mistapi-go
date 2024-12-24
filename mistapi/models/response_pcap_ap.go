package models

import (
    "encoding/json"
    "fmt"
)

// ResponsePcapAp represents a ResponsePcapAp struct.
type ResponsePcapAp struct {
    Band                 *int                   `json:"band,omitempty"`
    Bandwidth            *int                   `json:"bandwidth,omitempty"`
    Channel              *int                   `json:"channel,omitempty"`
    TcpdumpExpresssion   Optional[string]       `json:"tcpdump_expresssion"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePcapAp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePcapAp) String() string {
    return fmt.Sprintf(
    	"ResponsePcapAp[Band=%v, Bandwidth=%v, Channel=%v, TcpdumpExpresssion=%v, AdditionalProperties=%v]",
    	r.Band, r.Bandwidth, r.Channel, r.TcpdumpExpresssion, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapAp.
// It customizes the JSON marshaling process for ResponsePcapAp objects.
func (r ResponsePcapAp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "band", "bandwidth", "channel", "tcpdump_expresssion"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapAp object to a map representation for JSON marshaling.
func (r ResponsePcapAp) toMap() map[string]any {
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
    if r.TcpdumpExpresssion.IsValueSet() {
        if r.TcpdumpExpresssion.Value() != nil {
            structMap["tcpdump_expresssion"] = r.TcpdumpExpresssion.Value()
        } else {
            structMap["tcpdump_expresssion"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePcapAp.
// It customizes the JSON unmarshaling process for ResponsePcapAp objects.
func (r *ResponsePcapAp) UnmarshalJSON(input []byte) error {
    var temp tempResponsePcapAp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "bandwidth", "channel", "tcpdump_expresssion")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Band = temp.Band
    r.Bandwidth = temp.Bandwidth
    r.Channel = temp.Channel
    r.TcpdumpExpresssion = temp.TcpdumpExpresssion
    return nil
}

// tempResponsePcapAp is a temporary struct used for validating the fields of ResponsePcapAp.
type tempResponsePcapAp  struct {
    Band               *int             `json:"band,omitempty"`
    Bandwidth          *int             `json:"bandwidth,omitempty"`
    Channel            *int             `json:"channel,omitempty"`
    TcpdumpExpresssion Optional[string] `json:"tcpdump_expresssion"`
}
