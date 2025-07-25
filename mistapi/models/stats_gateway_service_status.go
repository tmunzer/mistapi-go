// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsGatewayServiceStatus represents a StatsGatewayServiceStatus struct.
type StatsGatewayServiceStatus struct {
    AppidInstallResult    *string                `json:"appid_install_result,omitempty"`
    AppidInstallTimestamp *string                `json:"appid_install_timestamp,omitempty"`
    AppidStatus           *string                `json:"appid_status,omitempty"`
    AppidVersion          *int                   `json:"appid_version,omitempty"`
    EwfStatus             *string                `json:"ewf_status,omitempty"`
    IdpInstallResult      *string                `json:"idp_install_result,omitempty"`
    IdpInstallTimestamp   *string                `json:"idp_install_timestamp,omitempty"`
    IdpPolicy             *string                `json:"idp_policy,omitempty"`
    IdpStatus             *string                `json:"idp_status,omitempty"`
    IdpUpdateTimestamp    *string                `json:"idp_update_timestamp,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsGatewayServiceStatus,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsGatewayServiceStatus) String() string {
    return fmt.Sprintf(
    	"StatsGatewayServiceStatus[AppidInstallResult=%v, AppidInstallTimestamp=%v, AppidStatus=%v, AppidVersion=%v, EwfStatus=%v, IdpInstallResult=%v, IdpInstallTimestamp=%v, IdpPolicy=%v, IdpStatus=%v, IdpUpdateTimestamp=%v, AdditionalProperties=%v]",
    	s.AppidInstallResult, s.AppidInstallTimestamp, s.AppidStatus, s.AppidVersion, s.EwfStatus, s.IdpInstallResult, s.IdpInstallTimestamp, s.IdpPolicy, s.IdpStatus, s.IdpUpdateTimestamp, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsGatewayServiceStatus.
// It customizes the JSON marshaling process for StatsGatewayServiceStatus objects.
func (s StatsGatewayServiceStatus) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "appid_install_result", "appid_install_timestamp", "appid_status", "appid_version", "ewf_status", "idp_install_result", "idp_install_timestamp", "idp_policy", "idp_status", "idp_update_timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsGatewayServiceStatus object to a map representation for JSON marshaling.
func (s StatsGatewayServiceStatus) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AppidInstallResult != nil {
        structMap["appid_install_result"] = s.AppidInstallResult
    }
    if s.AppidInstallTimestamp != nil {
        structMap["appid_install_timestamp"] = s.AppidInstallTimestamp
    }
    if s.AppidStatus != nil {
        structMap["appid_status"] = s.AppidStatus
    }
    if s.AppidVersion != nil {
        structMap["appid_version"] = s.AppidVersion
    }
    if s.EwfStatus != nil {
        structMap["ewf_status"] = s.EwfStatus
    }
    if s.IdpInstallResult != nil {
        structMap["idp_install_result"] = s.IdpInstallResult
    }
    if s.IdpInstallTimestamp != nil {
        structMap["idp_install_timestamp"] = s.IdpInstallTimestamp
    }
    if s.IdpPolicy != nil {
        structMap["idp_policy"] = s.IdpPolicy
    }
    if s.IdpStatus != nil {
        structMap["idp_status"] = s.IdpStatus
    }
    if s.IdpUpdateTimestamp != nil {
        structMap["idp_update_timestamp"] = s.IdpUpdateTimestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGatewayServiceStatus.
// It customizes the JSON unmarshaling process for StatsGatewayServiceStatus objects.
func (s *StatsGatewayServiceStatus) UnmarshalJSON(input []byte) error {
    var temp tempStatsGatewayServiceStatus
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "appid_install_result", "appid_install_timestamp", "appid_status", "appid_version", "ewf_status", "idp_install_result", "idp_install_timestamp", "idp_policy", "idp_status", "idp_update_timestamp")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AppidInstallResult = temp.AppidInstallResult
    s.AppidInstallTimestamp = temp.AppidInstallTimestamp
    s.AppidStatus = temp.AppidStatus
    s.AppidVersion = temp.AppidVersion
    s.EwfStatus = temp.EwfStatus
    s.IdpInstallResult = temp.IdpInstallResult
    s.IdpInstallTimestamp = temp.IdpInstallTimestamp
    s.IdpPolicy = temp.IdpPolicy
    s.IdpStatus = temp.IdpStatus
    s.IdpUpdateTimestamp = temp.IdpUpdateTimestamp
    return nil
}

// tempStatsGatewayServiceStatus is a temporary struct used for validating the fields of StatsGatewayServiceStatus.
type tempStatsGatewayServiceStatus  struct {
    AppidInstallResult    *string `json:"appid_install_result,omitempty"`
    AppidInstallTimestamp *string `json:"appid_install_timestamp,omitempty"`
    AppidStatus           *string `json:"appid_status,omitempty"`
    AppidVersion          *int    `json:"appid_version,omitempty"`
    EwfStatus             *string `json:"ewf_status,omitempty"`
    IdpInstallResult      *string `json:"idp_install_result,omitempty"`
    IdpInstallTimestamp   *string `json:"idp_install_timestamp,omitempty"`
    IdpPolicy             *string `json:"idp_policy,omitempty"`
    IdpStatus             *string `json:"idp_status,omitempty"`
    IdpUpdateTimestamp    *string `json:"idp_update_timestamp,omitempty"`
}
