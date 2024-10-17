package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// OrgSiteSleWiredResult represents a OrgSiteSleWiredResult struct.
type OrgSiteSleWiredResult struct {
	NumClients           *float64       `json:"num_clients,omitempty"`
	NumSwitches          *float64       `json:"num_switches,omitempty"`
	SiteId               uuid.UUID      `json:"site_id"`
	SwitchHealth         *float64       `json:"switch_health,omitempty"`
	SwitchStc            *float64       `json:"switch_stc,omitempty"`
	SwitchThroughput     *float64       `json:"switch_throughput,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSiteSleWiredResult.
// It customizes the JSON marshaling process for OrgSiteSleWiredResult objects.
func (o OrgSiteSleWiredResult) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSiteSleWiredResult object to a map representation for JSON marshaling.
func (o OrgSiteSleWiredResult) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.NumClients != nil {
		structMap["num_clients"] = o.NumClients
	}
	if o.NumSwitches != nil {
		structMap["num_switches"] = o.NumSwitches
	}
	structMap["site_id"] = o.SiteId
	if o.SwitchHealth != nil {
		structMap["switch_health"] = o.SwitchHealth
	}
	if o.SwitchStc != nil {
		structMap["switch_stc"] = o.SwitchStc
	}
	if o.SwitchThroughput != nil {
		structMap["switch_throughput"] = o.SwitchThroughput
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSiteSleWiredResult.
// It customizes the JSON unmarshaling process for OrgSiteSleWiredResult objects.
func (o *OrgSiteSleWiredResult) UnmarshalJSON(input []byte) error {
	var temp tempOrgSiteSleWiredResult
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "num_clients", "num_switches", "site_id", "switch_health", "switch_stc", "switch_throughput")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.NumClients = temp.NumClients
	o.NumSwitches = temp.NumSwitches
	o.SiteId = *temp.SiteId
	o.SwitchHealth = temp.SwitchHealth
	o.SwitchStc = temp.SwitchStc
	o.SwitchThroughput = temp.SwitchThroughput
	return nil
}

// tempOrgSiteSleWiredResult is a temporary struct used for validating the fields of OrgSiteSleWiredResult.
type tempOrgSiteSleWiredResult struct {
	NumClients       *float64   `json:"num_clients,omitempty"`
	NumSwitches      *float64   `json:"num_switches,omitempty"`
	SiteId           *uuid.UUID `json:"site_id"`
	SwitchHealth     *float64   `json:"switch_health,omitempty"`
	SwitchStc        *float64   `json:"switch_stc,omitempty"`
	SwitchThroughput *float64   `json:"switch_throughput,omitempty"`
}

func (o *tempOrgSiteSleWiredResult) validate() error {
	var errs []string
	if o.SiteId == nil {
		errs = append(errs, "required field `site_id` is missing for type `org_site_sle_wired_result`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
