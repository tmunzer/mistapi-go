package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ResponseOrgSiteSle represents a ResponseOrgSiteSle struct.
type ResponseOrgSiteSle struct {
	value              any
	isOrgSiteSleWifi   bool
	isOrgSiteWiredWifi bool
	isOrgSiteWanWifi   bool
}

// String converts the ResponseOrgSiteSle object to a string representation.
func (r ResponseOrgSiteSle) String() string {
	if bytes, err := json.Marshal(r.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgSiteSle.
// It customizes the JSON marshaling process for ResponseOrgSiteSle objects.
func (r ResponseOrgSiteSle) MarshalJSON() (
	[]byte,
	error) {
	if r.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseOrgSiteSleContainer.From*` functions to initialize the ResponseOrgSiteSle object.")
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgSiteSle object to a map representation for JSON marshaling.
func (r *ResponseOrgSiteSle) toMap() any {
	switch obj := r.value.(type) {
	case *OrgSiteSleWifi:
		return obj.toMap()
	case *OrgSiteWiredWifi:
		return obj.toMap()
	case *OrgSiteWanWifi:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgSiteSle.
// It customizes the JSON unmarshaling process for ResponseOrgSiteSle objects.
func (r *ResponseOrgSiteSle) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&OrgSiteSleWifi{}, false, &r.isOrgSiteSleWifi),
		NewTypeHolder(&OrgSiteWiredWifi{}, false, &r.isOrgSiteWiredWifi),
		NewTypeHolder(&OrgSiteWanWifi{}, false, &r.isOrgSiteWanWifi),
	)

	r.value = result
	return err
}

func (r *ResponseOrgSiteSle) AsOrgSiteSleWifi() (
	*OrgSiteSleWifi,
	bool) {
	if !r.isOrgSiteSleWifi {
		return nil, false
	}
	return r.value.(*OrgSiteSleWifi), true
}

func (r *ResponseOrgSiteSle) AsOrgSiteWiredWifi() (
	*OrgSiteWiredWifi,
	bool) {
	if !r.isOrgSiteWiredWifi {
		return nil, false
	}
	return r.value.(*OrgSiteWiredWifi), true
}

func (r *ResponseOrgSiteSle) AsOrgSiteWanWifi() (
	*OrgSiteWanWifi,
	bool) {
	if !r.isOrgSiteWanWifi {
		return nil, false
	}
	return r.value.(*OrgSiteWanWifi), true
}

// internalResponseOrgSiteSle represents a responseOrgSiteSle struct.
type internalResponseOrgSiteSle struct{}

var ResponseOrgSiteSleContainer internalResponseOrgSiteSle

// The internalResponseOrgSiteSle instance, wrapping the provided OrgSiteSleWifi value.
func (r *internalResponseOrgSiteSle) FromOrgSiteSleWifi(val OrgSiteSleWifi) ResponseOrgSiteSle {
	return ResponseOrgSiteSle{value: &val}
}

// The internalResponseOrgSiteSle instance, wrapping the provided OrgSiteWiredWifi value.
func (r *internalResponseOrgSiteSle) FromOrgSiteWiredWifi(val OrgSiteWiredWifi) ResponseOrgSiteSle {
	return ResponseOrgSiteSle{value: &val}
}

// The internalResponseOrgSiteSle instance, wrapping the provided OrgSiteWanWifi value.
func (r *internalResponseOrgSiteSle) FromOrgSiteWanWifi(val OrgSiteWanWifi) ResponseOrgSiteSle {
	return ResponseOrgSiteSle{value: &val}
}
