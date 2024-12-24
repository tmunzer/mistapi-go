package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CaptureOrg represents a CaptureOrg struct.
type CaptureOrg struct {
    value           any
    isCaptureMxedge bool
}

// String implements the fmt.Stringer interface for CaptureOrg,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureOrg) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CaptureOrg.
// It customizes the JSON marshaling process for CaptureOrg objects.
func (c CaptureOrg) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CaptureOrgContainer.From*` functions to initialize the CaptureOrg object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureOrg object to a map representation for JSON marshaling.
func (c *CaptureOrg) toMap() any {
    switch obj := c.value.(type) {
    case *CaptureMxedge:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureOrg.
// It customizes the JSON unmarshaling process for CaptureOrg objects.
func (c *CaptureOrg) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&CaptureMxedge{}, false, &c.isCaptureMxedge),
    )
    
    c.value = result
    return err
}

func (c *CaptureOrg) AsCaptureMxedge() (
    *CaptureMxedge,
    bool) {
    if !c.isCaptureMxedge {
        return nil, false
    }
    return c.value.(*CaptureMxedge), true
}

// internalCaptureOrg represents a captureOrg struct.
type internalCaptureOrg struct {}

var CaptureOrgContainer internalCaptureOrg

// The internalCaptureOrg instance, wrapping the provided CaptureMxedge value.
func (c *internalCaptureOrg) FromCaptureMxedge(val CaptureMxedge) CaptureOrg {
    return CaptureOrg{value: &val}
}
