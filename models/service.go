package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// Service represents a Service struct.
// Applications used for the Gateway configurations
type Service struct {
    // if `type`==`custom`, ip subnets
    Addresses                     []string                   `json:"addresses,omitempty"`
    // when `type`==`app_categories`
    // list of application categories are available through /api/v1/const/app_categories
    AppCategories                 []string                   `json:"app_categories,omitempty"`
    // when `type`==`app_categories`
    // list of application categories are available through /api/v1/const/app_subcategories
    AppSubcategories              []string                   `json:"app_subcategories,omitempty"`
    // when `type`==`apps`
    // list of applications are available through:
    //   - /api/v1/const/applications,
    //   - /api/v1/const/gateway_applications
    //   - /insight/top_app_by-bytes?wired=true
    Apps                          []string                   `json:"apps,omitempty"`
    CreatedTime                   *float64                   `json:"created_time,omitempty"`
    Description                   *string                    `json:"description,omitempty"`
    // when `traffic_type`==`custom`
    Dscp                          *int                       `json:"dscp,omitempty"`
    FailoverPolicy                *ServiceFailoverPolicyEnum `json:"failover_policy,omitempty"`
    // if `type`==`custom`, web filtering
    Hostnames                     []string                   `json:"hostnames,omitempty"`
    Id                            *uuid.UUID                 `json:"id,omitempty"`
    // when `traffic_type`==`custom`, for uplink selection
    MaxJitter                     *int                       `json:"max_jitter,omitempty"`
    // when `traffic_type`==`custom`, for uplink selection
    MaxLatency                    *int                       `json:"max_latency,omitempty"`
    // when `traffic_type`==`custom`, for uplink selection
    MaxLoss                       *int                       `json:"max_loss,omitempty"`
    ModifiedTime                  *float64                   `json:"modified_time,omitempty"`
    Name                          *string                    `json:"name,omitempty"`
    OrgId                         *uuid.UUID                 `json:"org_id,omitempty"`
    // whether to enable measure SLE
    SleEnabled                    *bool                      `json:"sle_enabled,omitempty"`
    Specs                         []ServiceSpec              `json:"specs,omitempty"`
    SsrRelaxedTcpStateEnforcement *bool                      `json:"ssr_relaxed_tcp_state_enforcement,omitempty"`
    // when `traffic_type`==`custom`
    TrafficClass                  *ServiceTrafficClassEnum   `json:"traffic_class,omitempty"`
    // values from `/api/v1/consts/traffic_types`
    // * when `type`==`apps`, we'll choose traffic_type automatically
    // * when `type`==`addresses` or `type`==`hostnames`, you can provide your own settings (optional)
    TrafficType                   *string                    `json:"traffic_type,omitempty"`
    Type                          *ServiceTypeEnum           `json:"type,omitempty"`
    // when `type`==`urls
    // no need for spec as URL can encode the ports being used`
    Urls                          []string                   `json:"urls,omitempty"`
    AdditionalProperties          map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Service.
// It customizes the JSON marshaling process for Service objects.
func (s Service) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Service object to a map representation for JSON marshaling.
func (s Service) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.Description != nil {
        structMap["description"] = s.Description
    }
    if s.Dscp != nil {
        structMap["dscp"] = s.Dscp
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
        structMap["max_jitter"] = s.MaxJitter
    }
    if s.MaxLatency != nil {
        structMap["max_latency"] = s.MaxLatency
    }
    if s.MaxLoss != nil {
        structMap["max_loss"] = s.MaxLoss
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
    var temp service
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "addresses", "app_categories", "app_subcategories", "apps", "created_time", "description", "dscp", "failover_policy", "hostnames", "id", "max_jitter", "max_latency", "max_loss", "modified_time", "name", "org_id", "sle_enabled", "specs", "ssr_relaxed_tcp_state_enforcement", "traffic_class", "traffic_type", "type", "urls")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Addresses = temp.Addresses
    s.AppCategories = temp.AppCategories
    s.AppSubcategories = temp.AppSubcategories
    s.Apps = temp.Apps
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
    s.SleEnabled = temp.SleEnabled
    s.Specs = temp.Specs
    s.SsrRelaxedTcpStateEnforcement = temp.SsrRelaxedTcpStateEnforcement
    s.TrafficClass = temp.TrafficClass
    s.TrafficType = temp.TrafficType
    s.Type = temp.Type
    s.Urls = temp.Urls
    return nil
}

// service is a temporary struct used for validating the fields of Service.
type service  struct {
    Addresses                     []string                   `json:"addresses,omitempty"`
    AppCategories                 []string                   `json:"app_categories,omitempty"`
    AppSubcategories              []string                   `json:"app_subcategories,omitempty"`
    Apps                          []string                   `json:"apps,omitempty"`
    CreatedTime                   *float64                   `json:"created_time,omitempty"`
    Description                   *string                    `json:"description,omitempty"`
    Dscp                          *int                       `json:"dscp,omitempty"`
    FailoverPolicy                *ServiceFailoverPolicyEnum `json:"failover_policy,omitempty"`
    Hostnames                     []string                   `json:"hostnames,omitempty"`
    Id                            *uuid.UUID                 `json:"id,omitempty"`
    MaxJitter                     *int                       `json:"max_jitter,omitempty"`
    MaxLatency                    *int                       `json:"max_latency,omitempty"`
    MaxLoss                       *int                       `json:"max_loss,omitempty"`
    ModifiedTime                  *float64                   `json:"modified_time,omitempty"`
    Name                          *string                    `json:"name,omitempty"`
    OrgId                         *uuid.UUID                 `json:"org_id,omitempty"`
    SleEnabled                    *bool                      `json:"sle_enabled,omitempty"`
    Specs                         []ServiceSpec              `json:"specs,omitempty"`
    SsrRelaxedTcpStateEnforcement *bool                      `json:"ssr_relaxed_tcp_state_enforcement,omitempty"`
    TrafficClass                  *ServiceTrafficClassEnum   `json:"traffic_class,omitempty"`
    TrafficType                   *string                    `json:"traffic_type,omitempty"`
    Type                          *ServiceTypeEnum           `json:"type,omitempty"`
    Urls                          []string                   `json:"urls,omitempty"`
}
