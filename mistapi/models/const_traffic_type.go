package models

import (
    "encoding/json"
    "fmt"
)

// ConstTrafficType represents a ConstTrafficType struct.
type ConstTrafficType struct {
    Display              *string                `json:"display,omitempty"`
    Dscp                 *int                   `json:"dscp,omitempty"`
    FailoverPolicy       *string                `json:"failover_policy,omitempty"`
    MaxJitter            *int                   `json:"max_jitter,omitempty"`
    MaxLatency           *int                   `json:"max_latency,omitempty"`
    MaxLoss              *int                   `json:"max_loss,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    TrafficClass         *string                `json:"traffic_class,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstTrafficType,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstTrafficType) String() string {
    return fmt.Sprintf(
    	"ConstTrafficType[Display=%v, Dscp=%v, FailoverPolicy=%v, MaxJitter=%v, MaxLatency=%v, MaxLoss=%v, Name=%v, TrafficClass=%v, AdditionalProperties=%v]",
    	c.Display, c.Dscp, c.FailoverPolicy, c.MaxJitter, c.MaxLatency, c.MaxLoss, c.Name, c.TrafficClass, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstTrafficType.
// It customizes the JSON marshaling process for ConstTrafficType objects.
func (c ConstTrafficType) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "display", "dscp", "failover_policy", "max_jitter", "max_latency", "max_loss", "name", "traffic_class"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstTrafficType object to a map representation for JSON marshaling.
func (c ConstTrafficType) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Display != nil {
        structMap["display"] = c.Display
    }
    if c.Dscp != nil {
        structMap["dscp"] = c.Dscp
    }
    if c.FailoverPolicy != nil {
        structMap["failover_policy"] = c.FailoverPolicy
    }
    if c.MaxJitter != nil {
        structMap["max_jitter"] = c.MaxJitter
    }
    if c.MaxLatency != nil {
        structMap["max_latency"] = c.MaxLatency
    }
    if c.MaxLoss != nil {
        structMap["max_loss"] = c.MaxLoss
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.TrafficClass != nil {
        structMap["traffic_class"] = c.TrafficClass
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstTrafficType.
// It customizes the JSON unmarshaling process for ConstTrafficType objects.
func (c *ConstTrafficType) UnmarshalJSON(input []byte) error {
    var temp tempConstTrafficType
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "display", "dscp", "failover_policy", "max_jitter", "max_latency", "max_loss", "name", "traffic_class")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Display = temp.Display
    c.Dscp = temp.Dscp
    c.FailoverPolicy = temp.FailoverPolicy
    c.MaxJitter = temp.MaxJitter
    c.MaxLatency = temp.MaxLatency
    c.MaxLoss = temp.MaxLoss
    c.Name = temp.Name
    c.TrafficClass = temp.TrafficClass
    return nil
}

// tempConstTrafficType is a temporary struct used for validating the fields of ConstTrafficType.
type tempConstTrafficType  struct {
    Display        *string `json:"display,omitempty"`
    Dscp           *int    `json:"dscp,omitempty"`
    FailoverPolicy *string `json:"failover_policy,omitempty"`
    MaxJitter      *int    `json:"max_jitter,omitempty"`
    MaxLatency     *int    `json:"max_latency,omitempty"`
    MaxLoss        *int    `json:"max_loss,omitempty"`
    Name           *string `json:"name,omitempty"`
    TrafficClass   *string `json:"traffic_class,omitempty"`
}
