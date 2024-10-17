package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// ResponseMspInventoryDevice represents a ResponseMspInventoryDevice struct.
type ResponseMspInventoryDevice struct {
	ForSite              *bool          `json:"for_site,omitempty"`
	Mac                  string         `json:"mac"`
	Model                string         `json:"model"`
	OrgId                uuid.UUID      `json:"org_id"`
	Serial               string         `json:"serial"`
	SiteId               uuid.UUID      `json:"site_id"`
	Type                 string         `json:"type"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseMspInventoryDevice.
// It customizes the JSON marshaling process for ResponseMspInventoryDevice objects.
func (r ResponseMspInventoryDevice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseMspInventoryDevice object to a map representation for JSON marshaling.
func (r ResponseMspInventoryDevice) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.ForSite != nil {
		structMap["for_site"] = r.ForSite
	}
	structMap["mac"] = r.Mac
	structMap["model"] = r.Model
	structMap["org_id"] = r.OrgId
	structMap["serial"] = r.Serial
	structMap["site_id"] = r.SiteId
	structMap["type"] = r.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMspInventoryDevice.
// It customizes the JSON unmarshaling process for ResponseMspInventoryDevice objects.
func (r *ResponseMspInventoryDevice) UnmarshalJSON(input []byte) error {
	var temp tempResponseMspInventoryDevice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "for_site", "mac", "model", "org_id", "serial", "site_id", "type")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.ForSite = temp.ForSite
	r.Mac = *temp.Mac
	r.Model = *temp.Model
	r.OrgId = *temp.OrgId
	r.Serial = *temp.Serial
	r.SiteId = *temp.SiteId
	r.Type = *temp.Type
	return nil
}

// tempResponseMspInventoryDevice is a temporary struct used for validating the fields of ResponseMspInventoryDevice.
type tempResponseMspInventoryDevice struct {
	ForSite *bool      `json:"for_site,omitempty"`
	Mac     *string    `json:"mac"`
	Model   *string    `json:"model"`
	OrgId   *uuid.UUID `json:"org_id"`
	Serial  *string    `json:"serial"`
	SiteId  *uuid.UUID `json:"site_id"`
	Type    *string    `json:"type"`
}

func (r *tempResponseMspInventoryDevice) validate() error {
	var errs []string
	if r.Mac == nil {
		errs = append(errs, "required field `mac` is missing for type `response_msp_inventory_device`")
	}
	if r.Model == nil {
		errs = append(errs, "required field `model` is missing for type `response_msp_inventory_device`")
	}
	if r.OrgId == nil {
		errs = append(errs, "required field `org_id` is missing for type `response_msp_inventory_device`")
	}
	if r.Serial == nil {
		errs = append(errs, "required field `serial` is missing for type `response_msp_inventory_device`")
	}
	if r.SiteId == nil {
		errs = append(errs, "required field `site_id` is missing for type `response_msp_inventory_device`")
	}
	if r.Type == nil {
		errs = append(errs, "required field `type` is missing for type `response_msp_inventory_device`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
