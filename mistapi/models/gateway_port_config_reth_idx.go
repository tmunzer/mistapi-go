package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// GatewayPortConfigRethIdx represents a GatewayPortConfigRethIdx struct.
// For SRX only and if HA Mode
type GatewayPortConfigRethIdx struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for GatewayPortConfigRethIdx,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPortConfigRethIdx) String() string {
    return fmt.Sprintf("%v", g.value)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortConfigRethIdx.
// It customizes the JSON marshaling process for GatewayPortConfigRethIdx objects.
func (g GatewayPortConfigRethIdx) MarshalJSON() (
    []byte,
    error) {
    if g.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.GatewayPortConfigRethIdxContainer.From*` functions to initialize the GatewayPortConfigRethIdx object.")
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortConfigRethIdx object to a map representation for JSON marshaling.
func (g *GatewayPortConfigRethIdx) toMap() any {
    switch obj := g.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortConfigRethIdx.
// It customizes the JSON unmarshaling process for GatewayPortConfigRethIdx objects.
func (g *GatewayPortConfigRethIdx) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), false, &g.isNumber),
        NewTypeHolder(new(string), false, &g.isString),
    )
    
    g.value = result
    return err
}

func (g *GatewayPortConfigRethIdx) AsNumber() (
    *int,
    bool) {
    if !g.isNumber {
        return nil, false
    }
    return g.value.(*int), true
}

func (g *GatewayPortConfigRethIdx) AsString() (
    *string,
    bool) {
    if !g.isString {
        return nil, false
    }
    return g.value.(*string), true
}

// internalGatewayPortConfigRethIdx represents a gatewayPortConfigRethIdx struct.
// For SRX only and if HA Mode
type internalGatewayPortConfigRethIdx struct {}

var GatewayPortConfigRethIdxContainer internalGatewayPortConfigRethIdx

// The internalGatewayPortConfigRethIdx instance, wrapping the provided int value.
func (g *internalGatewayPortConfigRethIdx) FromNumber(val int) GatewayPortConfigRethIdx {
    return GatewayPortConfigRethIdx{value: &val}
}

// The internalGatewayPortConfigRethIdx instance, wrapping the provided string value.
func (g *internalGatewayPortConfigRethIdx) FromString(val string) GatewayPortConfigRethIdx {
    return GatewayPortConfigRethIdx{value: &val}
}
