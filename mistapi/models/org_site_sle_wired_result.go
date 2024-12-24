package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// OrgSiteSleWiredResult represents a OrgSiteSleWiredResult struct.
type OrgSiteSleWiredResult struct {
    NumClients           *float64               `json:"num_clients,omitempty"`
    NumSwitches          *float64               `json:"num_switches,omitempty"`
    SiteId               uuid.UUID              `json:"site_id"`
    SwitchBandwidth      *float64               `json:"switch-bandwidth,omitempty"`
    SwitchHealth         float64                `json:"switch-health"`
    SwitchThroughput     *float64               `json:"switch-throughput,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSiteSleWiredResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSiteSleWiredResult) String() string {
    return fmt.Sprintf(
    	"OrgSiteSleWiredResult[NumClients=%v, NumSwitches=%v, SiteId=%v, SwitchBandwidth=%v, SwitchHealth=%v, SwitchThroughput=%v, AdditionalProperties=%v]",
    	o.NumClients, o.NumSwitches, o.SiteId, o.SwitchBandwidth, o.SwitchHealth, o.SwitchThroughput, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSiteSleWiredResult.
// It customizes the JSON marshaling process for OrgSiteSleWiredResult objects.
func (o OrgSiteSleWiredResult) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "num_clients", "num_switches", "site_id", "switch-bandwidth", "switch-health", "switch-throughput"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSiteSleWiredResult object to a map representation for JSON marshaling.
func (o OrgSiteSleWiredResult) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.NumClients != nil {
        structMap["num_clients"] = o.NumClients
    }
    if o.NumSwitches != nil {
        structMap["num_switches"] = o.NumSwitches
    }
    structMap["site_id"] = o.SiteId
    if o.SwitchBandwidth != nil {
        structMap["switch-bandwidth"] = o.SwitchBandwidth
    }
    structMap["switch-health"] = o.SwitchHealth
    if o.SwitchThroughput != nil {
        structMap["switch-throughput"] = o.SwitchThroughput
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_clients", "num_switches", "site_id", "switch-bandwidth", "switch-health", "switch-throughput")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.NumClients = temp.NumClients
    o.NumSwitches = temp.NumSwitches
    o.SiteId = *temp.SiteId
    o.SwitchBandwidth = temp.SwitchBandwidth
    o.SwitchHealth = *temp.SwitchHealth
    o.SwitchThroughput = temp.SwitchThroughput
    return nil
}

// tempOrgSiteSleWiredResult is a temporary struct used for validating the fields of OrgSiteSleWiredResult.
type tempOrgSiteSleWiredResult  struct {
    NumClients       *float64   `json:"num_clients,omitempty"`
    NumSwitches      *float64   `json:"num_switches,omitempty"`
    SiteId           *uuid.UUID `json:"site_id"`
    SwitchBandwidth  *float64   `json:"switch-bandwidth,omitempty"`
    SwitchHealth     *float64   `json:"switch-health"`
    SwitchThroughput *float64   `json:"switch-throughput,omitempty"`
}

func (o *tempOrgSiteSleWiredResult) validate() error {
    var errs []string
    if o.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `org_site_sle_wired_result`")
    }
    if o.SwitchHealth == nil {
        errs = append(errs, "required field `switch-health` is missing for type `org_site_sle_wired_result`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
