// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// WayfindingImportJson represents a WayfindingImportJson struct.
type WayfindingImportJson struct {
    value           any
    isMapJibestream bool
    isMapMicello    bool
}

// String implements the fmt.Stringer interface for WayfindingImportJson,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WayfindingImportJson) String() string {
    return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for WayfindingImportJson.
// It customizes the JSON marshaling process for WayfindingImportJson objects.
func (w WayfindingImportJson) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WayfindingImportJsonContainer.From*` functions to initialize the WayfindingImportJson object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WayfindingImportJson object to a map representation for JSON marshaling.
func (w *WayfindingImportJson) toMap() any {
    switch obj := w.value.(type) {
    case *MapJibestream:
        return obj.toMap()
    case *MapMicello:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WayfindingImportJson.
// It customizes the JSON unmarshaling process for WayfindingImportJson objects.
func (w *WayfindingImportJson) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&MapJibestream{}, false, &w.isMapJibestream),
        NewTypeHolder(&MapMicello{}, false, &w.isMapMicello),
    )
    
    w.value = result
    return err
}

func (w *WayfindingImportJson) AsMapJibestream() (
    *MapJibestream,
    bool) {
    if !w.isMapJibestream {
        return nil, false
    }
    return w.value.(*MapJibestream), true
}

func (w *WayfindingImportJson) AsMapMicello() (
    *MapMicello,
    bool) {
    if !w.isMapMicello {
        return nil, false
    }
    return w.value.(*MapMicello), true
}

// internalWayfindingImportJson represents a wayfindingImportJson struct.
type internalWayfindingImportJson struct {}

var WayfindingImportJsonContainer internalWayfindingImportJson

// The internalWayfindingImportJson instance, wrapping the provided MapJibestream value.
func (w *internalWayfindingImportJson) FromMapJibestream(val MapJibestream) WayfindingImportJson {
    return WayfindingImportJson{value: &val}
}

// The internalWayfindingImportJson instance, wrapping the provided MapMicello value.
func (w *internalWayfindingImportJson) FromMapMicello(val MapMicello) WayfindingImportJson {
    return WayfindingImportJson{value: &val}
}
