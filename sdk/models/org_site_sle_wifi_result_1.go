package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// OrgSiteSleWifiResult1 represents a OrgSiteSleWifiResult1 struct.
type OrgSiteSleWifiResult1 struct {
    ApAvailability       *float64       `json:"ap-availability,omitempty"`
    ApHealth             *float64       `json:"ap-health,omitempty"`
    Capacity             *float64       `json:"capacity,omitempty"`
    Coverage             *float64       `json:"coverage,omitempty"`
    NumAps               *float64       `json:"num_aps,omitempty"`
    NumClients           *float64       `json:"num_clients,omitempty"`
    Roaming              *float64       `json:"roaming,omitempty"`
    SiteId               uuid.UUID      `json:"site_id"`
    SuccessfulConnect    *float64       `json:"successful-connect,omitempty"`
    Throughput           *float64       `json:"throughput,omitempty"`
    TimeToConnect        *float64       `json:"time-to-connect,omitempty"`
    NumSwitches          *float64       `json:"num_switches,omitempty"`
    SwitchHealth         *float64       `json:"switch_health,omitempty"`
    SwitchStc            *float64       `json:"switch_stc,omitempty"`
    SwitchThroughput     *float64       `json:"switch_throughput,omitempty"`
    ApplicationHealth    *float64       `json:"application_health,omitempty"`
    GatewayHealth        *float64       `json:"gateway-health,omitempty"`
    NumGateways          *float64       `json:"num_gateways,omitempty"`
    WanLinkHealth        *float64       `json:"wan-link-health,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSiteSleWifiResult1.
// It customizes the JSON marshaling process for OrgSiteSleWifiResult1 objects.
func (o OrgSiteSleWifiResult1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSiteSleWifiResult1 object to a map representation for JSON marshaling.
func (o OrgSiteSleWifiResult1) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.ApAvailability != nil {
        structMap["ap-availability"] = o.ApAvailability
    }
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
    if o.NumSwitches != nil {
        structMap["num_switches"] = o.NumSwitches
    }
    if o.SwitchHealth != nil {
        structMap["switch_health"] = o.SwitchHealth
    }
    if o.SwitchStc != nil {
        structMap["switch_stc"] = o.SwitchStc
    }
    if o.SwitchThroughput != nil {
        structMap["switch_throughput"] = o.SwitchThroughput
    }
    if o.ApplicationHealth != nil {
        structMap["application_health"] = o.ApplicationHealth
    }
    if o.GatewayHealth != nil {
        structMap["gateway-health"] = o.GatewayHealth
    }
    if o.NumGateways != nil {
        structMap["num_gateways"] = o.NumGateways
    }
    if o.WanLinkHealth != nil {
        structMap["wan-link-health"] = o.WanLinkHealth
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSiteSleWifiResult1.
// It customizes the JSON unmarshaling process for OrgSiteSleWifiResult1 objects.
func (o *OrgSiteSleWifiResult1) UnmarshalJSON(input []byte) error {
    var temp orgSiteSleWifiResult1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap-availability", "ap-health", "capacity", "coverage", "num_aps", "num_clients", "roaming", "site_id", "successful-connect", "throughput", "time-to-connect", "num_switches", "switch_health", "switch_stc", "switch_throughput", "application_health", "gateway-health", "num_gateways", "wan-link-health")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.ApAvailability = temp.ApAvailability
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
    o.NumSwitches = temp.NumSwitches
    o.SwitchHealth = temp.SwitchHealth
    o.SwitchStc = temp.SwitchStc
    o.SwitchThroughput = temp.SwitchThroughput
    o.ApplicationHealth = temp.ApplicationHealth
    o.GatewayHealth = temp.GatewayHealth
    o.NumGateways = temp.NumGateways
    o.WanLinkHealth = temp.WanLinkHealth
    return nil
}

// orgSiteSleWifiResult1 is a temporary struct used for validating the fields of OrgSiteSleWifiResult1.
type orgSiteSleWifiResult1  struct {
    ApAvailability    *float64   `json:"ap-availability,omitempty"`
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
    NumSwitches       *float64   `json:"num_switches,omitempty"`
    SwitchHealth      *float64   `json:"switch_health,omitempty"`
    SwitchStc         *float64   `json:"switch_stc,omitempty"`
    SwitchThroughput  *float64   `json:"switch_throughput,omitempty"`
    ApplicationHealth *float64   `json:"application_health,omitempty"`
    GatewayHealth     *float64   `json:"gateway-health,omitempty"`
    NumGateways       *float64   `json:"num_gateways,omitempty"`
    WanLinkHealth     *float64   `json:"wan-link-health,omitempty"`
}

func (o *orgSiteSleWifiResult1) validate() error {
    var errs []string
    if o.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `Org_Site_Sle_Wifi_Result1`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
