// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// OrgSiteSleWifiResult represents a OrgSiteSleWifiResult struct.
type OrgSiteSleWifiResult struct {
    ApAvailability       float64                `json:"ap-availability"`
    ApHealth             *float64               `json:"ap-health,omitempty"`
    Capacity             *float64               `json:"capacity,omitempty"`
    Coverage             *float64               `json:"coverage,omitempty"`
    NumAps               *float64               `json:"num_aps,omitempty"`
    NumClients           *float64               `json:"num_clients,omitempty"`
    Roaming              *float64               `json:"roaming,omitempty"`
    SiteId               uuid.UUID              `json:"site_id"`
    SuccessfulConnect    *float64               `json:"successful-connect,omitempty"`
    Throughput           *float64               `json:"throughput,omitempty"`
    TimeToConnect        *float64               `json:"time-to-connect,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSiteSleWifiResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSiteSleWifiResult) String() string {
    return fmt.Sprintf(
    	"OrgSiteSleWifiResult[ApAvailability=%v, ApHealth=%v, Capacity=%v, Coverage=%v, NumAps=%v, NumClients=%v, Roaming=%v, SiteId=%v, SuccessfulConnect=%v, Throughput=%v, TimeToConnect=%v, AdditionalProperties=%v]",
    	o.ApAvailability, o.ApHealth, o.Capacity, o.Coverage, o.NumAps, o.NumClients, o.Roaming, o.SiteId, o.SuccessfulConnect, o.Throughput, o.TimeToConnect, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSiteSleWifiResult.
// It customizes the JSON marshaling process for OrgSiteSleWifiResult objects.
func (o OrgSiteSleWifiResult) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "ap-availability", "ap-health", "capacity", "coverage", "num_aps", "num_clients", "roaming", "site_id", "successful-connect", "throughput", "time-to-connect"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSiteSleWifiResult object to a map representation for JSON marshaling.
func (o OrgSiteSleWifiResult) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["ap-availability"] = o.ApAvailability
    if o.ApHealth != nil {
        structMap["ap-health"] = o.ApHealth
    }
    if o.Capacity != nil {
        structMap["capacity"] = o.Capacity
    }
    if o.Coverage != nil {
        structMap["coverage"] = o.Coverage
    }
    if o.NumAps != nil {
        structMap["num_aps"] = o.NumAps
    }
    if o.NumClients != nil {
        structMap["num_clients"] = o.NumClients
    }
    if o.Roaming != nil {
        structMap["roaming"] = o.Roaming
    }
    structMap["site_id"] = o.SiteId
    if o.SuccessfulConnect != nil {
        structMap["successful-connect"] = o.SuccessfulConnect
    }
    if o.Throughput != nil {
        structMap["throughput"] = o.Throughput
    }
    if o.TimeToConnect != nil {
        structMap["time-to-connect"] = o.TimeToConnect
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSiteSleWifiResult.
// It customizes the JSON unmarshaling process for OrgSiteSleWifiResult objects.
func (o *OrgSiteSleWifiResult) UnmarshalJSON(input []byte) error {
    var temp tempOrgSiteSleWifiResult
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap-availability", "ap-health", "capacity", "coverage", "num_aps", "num_clients", "roaming", "site_id", "successful-connect", "throughput", "time-to-connect")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.ApAvailability = *temp.ApAvailability
    o.ApHealth = temp.ApHealth
    o.Capacity = temp.Capacity
    o.Coverage = temp.Coverage
    o.NumAps = temp.NumAps
    o.NumClients = temp.NumClients
    o.Roaming = temp.Roaming
    o.SiteId = *temp.SiteId
    o.SuccessfulConnect = temp.SuccessfulConnect
    o.Throughput = temp.Throughput
    o.TimeToConnect = temp.TimeToConnect
    return nil
}

// tempOrgSiteSleWifiResult is a temporary struct used for validating the fields of OrgSiteSleWifiResult.
type tempOrgSiteSleWifiResult  struct {
    ApAvailability    *float64   `json:"ap-availability"`
    ApHealth          *float64   `json:"ap-health,omitempty"`
    Capacity          *float64   `json:"capacity,omitempty"`
    Coverage          *float64   `json:"coverage,omitempty"`
    NumAps            *float64   `json:"num_aps,omitempty"`
    NumClients        *float64   `json:"num_clients,omitempty"`
    Roaming           *float64   `json:"roaming,omitempty"`
    SiteId            *uuid.UUID `json:"site_id"`
    SuccessfulConnect *float64   `json:"successful-connect,omitempty"`
    Throughput        *float64   `json:"throughput,omitempty"`
    TimeToConnect     *float64   `json:"time-to-connect,omitempty"`
}

func (o *tempOrgSiteSleWifiResult) validate() error {
    var errs []string
    if o.ApAvailability == nil {
        errs = append(errs, "required field `ap-availability` is missing for type `org_site_sle_wifi_result`")
    }
    if o.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `org_site_sle_wifi_result`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
