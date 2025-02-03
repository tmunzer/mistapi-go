package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Service represents a Service struct.
// APplications used for the Gateway configurations
type Service struct {
    // If `type`==`custom`, ip subnets (e.g. 10.0.0.0/8)
    Addresses                     []string                   `json:"addresses,omitempty"`
    // When `type`==`app_categories`, list of application categories are available through /api/v1/const/app_categories
    AppCategories                 []string                   `json:"app_categories,omitempty"`
    // When `type`==`app_categories`, list of application categories are available through /api/v1/const/app_subcategories
    AppSubcategories              []string                   `json:"app_subcategories,omitempty"`
    // When `type`==`apps`, list of applications are available through:
    // * /api/v1/const/applications
    // * /api/v1/const/gateway_applications
    // * /insight/top_app_by-bytes?wired=true
    Apps                          []string                   `json:"apps,omitempty"`
    // 0 means unlimited
    ClientLimitDown               *int                       `json:"client_limit_down,omitempty"`
    // 0 means unlimited
    ClientLimitUp                 *int                       `json:"client_limit_up,omitempty"`
    // When the object has been created, in epoch
    CreatedTime                   *float64                   `json:"created_time,omitempty"`
    Description                   *string                    `json:"description,omitempty"`
    // For SSR only, when `traffic_type`==`custom`. 0-63 or variable
    Dscp                          *ServiceDscp               `json:"dscp,omitempty"`
    // enum: `non_revertable`, `none`, `revertable`
    FailoverPolicy                *ServiceFailoverPolicyEnum `json:"failover_policy,omitempty"`
    // If `type`==`custom`, web filtering
    Hostnames                     []string                   `json:"hostnames,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                            *uuid.UUID                 `json:"id,omitempty"`
    // For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable
    MaxJitter                     *ServiceMaxJitter          `json:"max_jitter,omitempty"`
    // For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable
    MaxLatency                    *ServiceMaxLatency         `json:"max_latency,omitempty"`
    // For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-100 or variable
    MaxLoss                       *ServiceMaxLoss            `json:"max_loss,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime                  *float64                   `json:"modified_time,omitempty"`
    Name                          *string                    `json:"name,omitempty"`
    OrgId                         *uuid.UUID                 `json:"org_id,omitempty"`
    // 0 means unlimited
    ServiceLimitDown              *int                       `json:"service_limit_down,omitempty"`
    // 0 means unlimited
    ServiceLimitUp                *int                       `json:"service_limit_up,omitempty"`
    // Whether to enable measure SLE
    SleEnabled                    *bool                      `json:"sle_enabled,omitempty"`
    // When `type`==`custom`, optional, if it doesn't exist, http and https is assumed
    Specs                         []ServiceSpec              `json:"specs,omitempty"`
    SsrRelaxedTcpStateEnforcement *bool                      `json:"ssr_relaxed_tcp_state_enforcement,omitempty"`
    // when `traffic_type`==`custom`. enum: `best_effort`, `high`, `low`, `medium`
    TrafficClass                  *ServiceTrafficClassEnum   `json:"traffic_class,omitempty"`
    // values from `/api/v1/consts/traffic_types`
    TrafficType                   *string                    `json:"traffic_type,omitempty"`
    // enum: `app_categories`, `apps`, `custom`, `urls`
    Type                          *ServiceTypeEnum           `json:"type,omitempty"`
    // When `type`==`urls`, no need for spec as URL can encode the ports being used
    Urls                          []string                   `json:"urls,omitempty"`
    AdditionalProperties          map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for Service,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Service) String() string {
    return fmt.Sprintf(
    	"Service[Addresses=%v, AppCategories=%v, AppSubcategories=%v, Apps=%v, ClientLimitDown=%v, ClientLimitUp=%v, CreatedTime=%v, Description=%v, Dscp=%v, FailoverPolicy=%v, Hostnames=%v, Id=%v, MaxJitter=%v, MaxLatency=%v, MaxLoss=%v, ModifiedTime=%v, Name=%v, OrgId=%v, ServiceLimitDown=%v, ServiceLimitUp=%v, SleEnabled=%v, Specs=%v, SsrRelaxedTcpStateEnforcement=%v, TrafficClass=%v, TrafficType=%v, Type=%v, Urls=%v, AdditionalProperties=%v]",
    	s.Addresses, s.AppCategories, s.AppSubcategories, s.Apps, s.ClientLimitDown, s.ClientLimitUp, s.CreatedTime, s.Description, s.Dscp, s.FailoverPolicy, s.Hostnames, s.Id, s.MaxJitter, s.MaxLatency, s.MaxLoss, s.ModifiedTime, s.Name, s.OrgId, s.ServiceLimitDown, s.ServiceLimitUp, s.SleEnabled, s.Specs, s.SsrRelaxedTcpStateEnforcement, s.TrafficClass, s.TrafficType, s.Type, s.Urls, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Service.
// It customizes the JSON marshaling process for Service objects.
func (s Service) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "addresses", "app_categories", "app_subcategories", "apps", "client_limit_down", "client_limit_up", "created_time", "description", "dscp", "failover_policy", "hostnames", "id", "max_jitter", "max_latency", "max_loss", "modified_time", "name", "org_id", "service_limit_down", "service_limit_up", "sle_enabled", "specs", "ssr_relaxed_tcp_state_enforcement", "traffic_class", "traffic_type", "type", "urls"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Service object to a map representation for JSON marshaling.
func (s Service) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Addresses != nil {
        structMap["addresses"] = s.Addresses
    }
    if s.AppCategories != nil {
        structMap["app_categories"] = s.AppCategories
    }
    if s.AppSubcategories != nil {
        structMap["app_subcategories"] = s.AppSubcategories
    }
    if s.Apps != nil {
        structMap["apps"] = s.Apps
    }
    if s.ClientLimitDown != nil {
        structMap["client_limit_down"] = s.ClientLimitDown
    }
    if s.ClientLimitUp != nil {
        structMap["client_limit_up"] = s.ClientLimitUp
    }
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.Description != nil {
        structMap["description"] = s.Description
    }
    if s.Dscp != nil {
        structMap["dscp"] = s.Dscp.toMap()
    }
    if s.FailoverPolicy != nil {
        structMap["failover_policy"] = s.FailoverPolicy
    }
    if s.Hostnames != nil {
        structMap["hostnames"] = s.Hostnames
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.MaxJitter != nil {
        structMap["max_jitter"] = s.MaxJitter.toMap()
    }
    if s.MaxLatency != nil {
        structMap["max_latency"] = s.MaxLatency.toMap()
    }
    if s.MaxLoss != nil {
        structMap["max_loss"] = s.MaxLoss.toMap()
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.ServiceLimitDown != nil {
        structMap["service_limit_down"] = s.ServiceLimitDown
    }
    if s.ServiceLimitUp != nil {
        structMap["service_limit_up"] = s.ServiceLimitUp
    }
    if s.SleEnabled != nil {
        structMap["sle_enabled"] = s.SleEnabled
    }
    if s.Specs != nil {
        structMap["specs"] = s.Specs
    }
    if s.SsrRelaxedTcpStateEnforcement != nil {
        structMap["ssr_relaxed_tcp_state_enforcement"] = s.SsrRelaxedTcpStateEnforcement
    }
    if s.TrafficClass != nil {
        structMap["traffic_class"] = s.TrafficClass
    }
    if s.TrafficType != nil {
        structMap["traffic_type"] = s.TrafficType
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Urls != nil {
        structMap["urls"] = s.Urls
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Service.
// It customizes the JSON unmarshaling process for Service objects.
func (s *Service) UnmarshalJSON(input []byte) error {
    var temp tempService
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "addresses", "app_categories", "app_subcategories", "apps", "client_limit_down", "client_limit_up", "created_time", "description", "dscp", "failover_policy", "hostnames", "id", "max_jitter", "max_latency", "max_loss", "modified_time", "name", "org_id", "service_limit_down", "service_limit_up", "sle_enabled", "specs", "ssr_relaxed_tcp_state_enforcement", "traffic_class", "traffic_type", "type", "urls")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Addresses = temp.Addresses
    s.AppCategories = temp.AppCategories
    s.AppSubcategories = temp.AppSubcategories
    s.Apps = temp.Apps
    s.ClientLimitDown = temp.ClientLimitDown
    s.ClientLimitUp = temp.ClientLimitUp
    s.CreatedTime = temp.CreatedTime
    s.Description = temp.Description
    s.Dscp = temp.Dscp
    s.FailoverPolicy = temp.FailoverPolicy
    s.Hostnames = temp.Hostnames
    s.Id = temp.Id
    s.MaxJitter = temp.MaxJitter
    s.MaxLatency = temp.MaxLatency
    s.MaxLoss = temp.MaxLoss
    s.ModifiedTime = temp.ModifiedTime
    s.Name = temp.Name
    s.OrgId = temp.OrgId
    s.ServiceLimitDown = temp.ServiceLimitDown
    s.ServiceLimitUp = temp.ServiceLimitUp
    s.SleEnabled = temp.SleEnabled
    s.Specs = temp.Specs
    s.SsrRelaxedTcpStateEnforcement = temp.SsrRelaxedTcpStateEnforcement
    s.TrafficClass = temp.TrafficClass
    s.TrafficType = temp.TrafficType
    s.Type = temp.Type
    s.Urls = temp.Urls
    return nil
}

// tempService is a temporary struct used for validating the fields of Service.
type tempService  struct {
    Addresses                     []string                   `json:"addresses,omitempty"`
    AppCategories                 []string                   `json:"app_categories,omitempty"`
    AppSubcategories              []string                   `json:"app_subcategories,omitempty"`
    Apps                          []string                   `json:"apps,omitempty"`
    ClientLimitDown               *int                       `json:"client_limit_down,omitempty"`
    ClientLimitUp                 *int                       `json:"client_limit_up,omitempty"`
    CreatedTime                   *float64                   `json:"created_time,omitempty"`
    Description                   *string                    `json:"description,omitempty"`
    Dscp                          *ServiceDscp               `json:"dscp,omitempty"`
    FailoverPolicy                *ServiceFailoverPolicyEnum `json:"failover_policy,omitempty"`
    Hostnames                     []string                   `json:"hostnames,omitempty"`
    Id                            *uuid.UUID                 `json:"id,omitempty"`
    MaxJitter                     *ServiceMaxJitter          `json:"max_jitter,omitempty"`
    MaxLatency                    *ServiceMaxLatency         `json:"max_latency,omitempty"`
    MaxLoss                       *ServiceMaxLoss            `json:"max_loss,omitempty"`
    ModifiedTime                  *float64                   `json:"modified_time,omitempty"`
    Name                          *string                    `json:"name,omitempty"`
    OrgId                         *uuid.UUID                 `json:"org_id,omitempty"`
    ServiceLimitDown              *int                       `json:"service_limit_down,omitempty"`
    ServiceLimitUp                *int                       `json:"service_limit_up,omitempty"`
    SleEnabled                    *bool                      `json:"sle_enabled,omitempty"`
    Specs                         []ServiceSpec              `json:"specs,omitempty"`
    SsrRelaxedTcpStateEnforcement *bool                      `json:"ssr_relaxed_tcp_state_enforcement,omitempty"`
    TrafficClass                  *ServiceTrafficClassEnum   `json:"traffic_class,omitempty"`
    TrafficType                   *string                    `json:"traffic_type,omitempty"`
    Type                          *ServiceTypeEnum           `json:"type,omitempty"`
    Urls                          []string                   `json:"urls,omitempty"`
}
