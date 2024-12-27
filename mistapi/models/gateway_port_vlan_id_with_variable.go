package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// GatewayPortVlanIdWithVariable represents a GatewayPortVlanIdWithVariable struct.
// if WAN interface is on a VLAN. Can be the VLAN ID (i.e. "10") or a Variable (i.e. "{{myvar}}")
type GatewayPortVlanIdWithVariable struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for GatewayPortVlanIdWithVariable,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPortVlanIdWithVariable) String() string {
    return fmt.Sprintf("%v", g.value)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortVlanIdWithVariable.
// It customizes the JSON marshaling process for GatewayPortVlanIdWithVariable objects.
func (g GatewayPortVlanIdWithVariable) MarshalJSON() (
    []byte,
    error) {
    if g.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.GatewayPortVlanIdWithVariableContainer.From*` functions to initialize the GatewayPortVlanIdWithVariable object.")
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortVlanIdWithVariable object to a map representation for JSON marshaling.
func (g *GatewayPortVlanIdWithVariable) toMap() any {
    switch obj := g.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortVlanIdWithVariable.
// It customizes the JSON unmarshaling process for GatewayPortVlanIdWithVariable objects.
func (g *GatewayPortVlanIdWithVariable) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &g.isString),
        NewTypeHolder(new(int), false, &g.isNumber),
    )
    
    g.value = result
    return err
}

func (g *GatewayPortVlanIdWithVariable) AsString() (
    *string,
    bool) {
    if !g.isString {
        return nil, false
    }
    return g.value.(*string), true
}

func (g *GatewayPortVlanIdWithVariable) AsNumber() (
    *int,
    bool) {
    if !g.isNumber {
        return nil, false
    }
    return g.value.(*int), true
}

// internalGatewayPortVlanIdWithVariable represents a gatewayPortVlanIdWithVariable struct.
// if WAN interface is on a VLAN. Can be the VLAN ID (i.e. "10") or a Variable (i.e. "{{myvar}}")
type internalGatewayPortVlanIdWithVariable struct {}

var GatewayPortVlanIdWithVariableContainer internalGatewayPortVlanIdWithVariable

// The internalGatewayPortVlanIdWithVariable instance, wrapping the provided string value.
func (g *internalGatewayPortVlanIdWithVariable) FromString(val string) GatewayPortVlanIdWithVariable {
    return GatewayPortVlanIdWithVariable{value: &val}
}

// The internalGatewayPortVlanIdWithVariable instance, wrapping the provided int value.
func (g *internalGatewayPortVlanIdWithVariable) FromNumber(val int) GatewayPortVlanIdWithVariable {
    return GatewayPortVlanIdWithVariable{value: &val}
}
