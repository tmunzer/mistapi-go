package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CaptureSite represents a CaptureSite struct.
type CaptureSite struct {
	value                  any
	isCaptureClient        bool
	isCaptureGateway       bool
	isCaptureNewAssoc      bool
	isCaptureRadiotap      bool
	isCaptureRadiotapwired bool
	isCaptureScan          bool
	isCaptureSwitch        bool
	isCaptureWired         bool
	isCaptureWireless      bool
}

// String converts the CaptureSite object to a string representation.
func (c CaptureSite) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CaptureSite.
// It customizes the JSON marshaling process for CaptureSite objects.
func (c CaptureSite) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CaptureSiteContainer.From*` functions to initialize the CaptureSite object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CaptureSite object to a map representation for JSON marshaling.
func (c *CaptureSite) toMap() any {
	switch obj := c.value.(type) {
	case *CaptureClient:
		return obj.toMap()
	case *CaptureGateway:
		return obj.toMap()
	case *CaptureNewAssoc:
		return obj.toMap()
	case *CaptureRadiotap:
		return obj.toMap()
	case *CaptureRadiotapwired:
		return obj.toMap()
	case *CaptureScan:
		return obj.toMap()
	case *CaptureSwitch:
		return obj.toMap()
	case *CaptureWired:
		return obj.toMap()
	case *CaptureWireless:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureSite.
// It customizes the JSON unmarshaling process for CaptureSite objects.
func (c *CaptureSite) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&CaptureClient{}, false, &c.isCaptureClient),
		NewTypeHolder(&CaptureGateway{}, false, &c.isCaptureGateway),
		NewTypeHolder(&CaptureNewAssoc{}, false, &c.isCaptureNewAssoc),
		NewTypeHolder(&CaptureRadiotap{}, false, &c.isCaptureRadiotap),
		NewTypeHolder(&CaptureRadiotapwired{}, false, &c.isCaptureRadiotapwired),
		NewTypeHolder(&CaptureScan{}, false, &c.isCaptureScan),
		NewTypeHolder(&CaptureSwitch{}, false, &c.isCaptureSwitch),
		NewTypeHolder(&CaptureWired{}, false, &c.isCaptureWired),
		NewTypeHolder(&CaptureWireless{}, false, &c.isCaptureWireless),
	)

	c.value = result
	return err
}

func (c *CaptureSite) AsCaptureClient() (
	*CaptureClient,
	bool) {
	if !c.isCaptureClient {
		return nil, false
	}
	return c.value.(*CaptureClient), true
}

func (c *CaptureSite) AsCaptureGateway() (
	*CaptureGateway,
	bool) {
	if !c.isCaptureGateway {
		return nil, false
	}
	return c.value.(*CaptureGateway), true
}

func (c *CaptureSite) AsCaptureNewAssoc() (
	*CaptureNewAssoc,
	bool) {
	if !c.isCaptureNewAssoc {
		return nil, false
	}
	return c.value.(*CaptureNewAssoc), true
}

func (c *CaptureSite) AsCaptureRadiotap() (
	*CaptureRadiotap,
	bool) {
	if !c.isCaptureRadiotap {
		return nil, false
	}
	return c.value.(*CaptureRadiotap), true
}

func (c *CaptureSite) AsCaptureRadiotapwired() (
	*CaptureRadiotapwired,
	bool) {
	if !c.isCaptureRadiotapwired {
		return nil, false
	}
	return c.value.(*CaptureRadiotapwired), true
}

func (c *CaptureSite) AsCaptureScan() (
	*CaptureScan,
	bool) {
	if !c.isCaptureScan {
		return nil, false
	}
	return c.value.(*CaptureScan), true
}

func (c *CaptureSite) AsCaptureSwitch() (
	*CaptureSwitch,
	bool) {
	if !c.isCaptureSwitch {
		return nil, false
	}
	return c.value.(*CaptureSwitch), true
}

func (c *CaptureSite) AsCaptureWired() (
	*CaptureWired,
	bool) {
	if !c.isCaptureWired {
		return nil, false
	}
	return c.value.(*CaptureWired), true
}

func (c *CaptureSite) AsCaptureWireless() (
	*CaptureWireless,
	bool) {
	if !c.isCaptureWireless {
		return nil, false
	}
	return c.value.(*CaptureWireless), true
}

// internalCaptureSite represents a captureSite struct.
type internalCaptureSite struct{}

var CaptureSiteContainer internalCaptureSite

// The internalCaptureSite instance, wrapping the provided CaptureClient value.
func (c *internalCaptureSite) FromCaptureClient(val CaptureClient) CaptureSite {
	return CaptureSite{value: &val}
}

// The internalCaptureSite instance, wrapping the provided CaptureGateway value.
func (c *internalCaptureSite) FromCaptureGateway(val CaptureGateway) CaptureSite {
	return CaptureSite{value: &val}
}

// The internalCaptureSite instance, wrapping the provided CaptureNewAssoc value.
func (c *internalCaptureSite) FromCaptureNewAssoc(val CaptureNewAssoc) CaptureSite {
	return CaptureSite{value: &val}
}

// The internalCaptureSite instance, wrapping the provided CaptureRadiotap value.
func (c *internalCaptureSite) FromCaptureRadiotap(val CaptureRadiotap) CaptureSite {
	return CaptureSite{value: &val}
}

// The internalCaptureSite instance, wrapping the provided CaptureRadiotapwired value.
func (c *internalCaptureSite) FromCaptureRadiotapwired(val CaptureRadiotapwired) CaptureSite {
	return CaptureSite{value: &val}
}

// The internalCaptureSite instance, wrapping the provided CaptureScan value.
func (c *internalCaptureSite) FromCaptureScan(val CaptureScan) CaptureSite {
	return CaptureSite{value: &val}
}

// The internalCaptureSite instance, wrapping the provided CaptureSwitch value.
func (c *internalCaptureSite) FromCaptureSwitch(val CaptureSwitch) CaptureSite {
	return CaptureSite{value: &val}
}

// The internalCaptureSite instance, wrapping the provided CaptureWired value.
func (c *internalCaptureSite) FromCaptureWired(val CaptureWired) CaptureSite {
	return CaptureSite{value: &val}
}

// The internalCaptureSite instance, wrapping the provided CaptureWireless value.
func (c *internalCaptureSite) FromCaptureWireless(val CaptureWireless) CaptureSite {
	return CaptureSite{value: &val}
}
