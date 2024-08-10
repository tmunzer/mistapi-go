package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ClientStats represents a ClientStats struct.
type ClientStats struct {
    value                        any
    isArrayOfClientWirelessStats bool
    isStatsWiredClient           bool
}

// String converts the ClientStats object to a string representation.
func (c ClientStats) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ClientStats.
// It customizes the JSON marshaling process for ClientStats objects.
func (c ClientStats) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ClientStatsContainer.From*` functions to initialize the ClientStats object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ClientStats object to a map representation for JSON marshaling.
func (c *ClientStats) toMap() any {
    switch obj := c.value.(type) {
    case *[]ClientWirelessStats:
        return *obj
    case *StatsWiredClient:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientStats.
// It customizes the JSON unmarshaling process for ClientStats objects.
func (c *ClientStats) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&[]ClientWirelessStats{}, false, &c.isArrayOfClientWirelessStats),
        NewTypeHolder(&StatsWiredClient{}, false, &c.isStatsWiredClient),
    )
    
    c.value = result
    return err
}

func (c *ClientStats) AsArrayOfClientWirelessStats() (
    *[]ClientWirelessStats,
    bool) {
    if !c.isArrayOfClientWirelessStats {
        return nil, false
    }
    return c.value.(*[]ClientWirelessStats), true
}

func (c *ClientStats) AsStatsWiredClient() (
    *StatsWiredClient,
    bool) {
    if !c.isStatsWiredClient {
        return nil, false
    }
    return c.value.(*StatsWiredClient), true
}

// internalClientStats represents a clientStats struct.
type internalClientStats struct {}

var ClientStatsContainer internalClientStats

// The internalClientStats instance, wrapping the provided []ClientWirelessStats value.
func (c *internalClientStats) FromArrayOfClientWirelessStats(val []ClientWirelessStats) ClientStats {
    return ClientStats{value: &val}
}

// The internalClientStats instance, wrapping the provided StatsWiredClient value.
func (c *internalClientStats) FromStatsWiredClient(val StatsWiredClient) ClientStats {
    return ClientStats{value: &val}
}
