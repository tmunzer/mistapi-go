package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ConstInsightMetricsPropertyExample represents a ConstInsightMetricsPropertyExample struct.
type ConstInsightMetricsPropertyExample struct {
    value       any
    isNumber    bool
    isPrecision bool
    isString    bool
    isBoolean   bool
}

// String converts the ConstInsightMetricsPropertyExample object to a string representation.
func (c ConstInsightMetricsPropertyExample) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsPropertyExample.
// It customizes the JSON marshaling process for ConstInsightMetricsPropertyExample objects.
func (c ConstInsightMetricsPropertyExample) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ConstInsightMetricsPropertyExampleContainer.From*` functions to initialize the ConstInsightMetricsPropertyExample object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsPropertyExample object to a map representation for JSON marshaling.
func (c *ConstInsightMetricsPropertyExample) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *float64:
        return *obj
    case *string:
        return *obj
    case *bool:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstInsightMetricsPropertyExample.
// It customizes the JSON unmarshaling process for ConstInsightMetricsPropertyExample objects.
func (c *ConstInsightMetricsPropertyExample) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(new(int), true, &c.isNumber),
        NewTypeHolder(new(float64), true, &c.isPrecision),
        NewTypeHolder(new(string), true, &c.isString),
        NewTypeHolder(new(bool), false, &c.isBoolean),
    )
    
    c.value = result
    return err
}

func (c *ConstInsightMetricsPropertyExample) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    if c.value == nil {
        return nil, true
    }
    return c.value.(*int), true
}

func (c *ConstInsightMetricsPropertyExample) AsPrecision() (
    *float64,
    bool) {
    if !c.isPrecision {
        return nil, false
    }
    if c.value == nil {
        return nil, true
    }
    return c.value.(*float64), true
}

func (c *ConstInsightMetricsPropertyExample) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    if c.value == nil {
        return nil, true
    }
    return c.value.(*string), true
}

func (c *ConstInsightMetricsPropertyExample) AsBoolean() (
    *bool,
    bool) {
    if !c.isBoolean {
        return nil, false
    }
    return c.value.(*bool), true
}

// internalConstInsightMetricsPropertyExample represents a constInsightMetricsPropertyExample struct.
type internalConstInsightMetricsPropertyExample struct {}

var ConstInsightMetricsPropertyExampleContainer internalConstInsightMetricsPropertyExample

// The internalConstInsightMetricsPropertyExample instance, wrapping the provided int value.
func (c *internalConstInsightMetricsPropertyExample) FromNumber(val *int) ConstInsightMetricsPropertyExample {
    return ConstInsightMetricsPropertyExample{value: &val}
}

// The internalConstInsightMetricsPropertyExample instance, wrapping the provided float64 value.
func (c *internalConstInsightMetricsPropertyExample) FromPrecision(val *float64) ConstInsightMetricsPropertyExample {
    return ConstInsightMetricsPropertyExample{value: &val}
}

// The internalConstInsightMetricsPropertyExample instance, wrapping the provided string value.
func (c *internalConstInsightMetricsPropertyExample) FromString(val *string) ConstInsightMetricsPropertyExample {
    return ConstInsightMetricsPropertyExample{value: &val}
}

// The internalConstInsightMetricsPropertyExample instance, wrapping the provided bool value.
func (c *internalConstInsightMetricsPropertyExample) FromBoolean(val bool) ConstInsightMetricsPropertyExample {
    return ConstInsightMetricsPropertyExample{value: &val}
}
