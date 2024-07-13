package models

import (
    "encoding/json"
)

// GatewayStatsServiceStatus represents a GatewayStatsServiceStatus struct.
type GatewayStatsServiceStatus struct {
    AppidInstallResult    *string        `json:"appid_install_result,omitempty"`
    AppidInstallTimestamp *string        `json:"appid_install_timestamp,omitempty"`
    AppidStatus           *string        `json:"appid_status,omitempty"`
    AppidVersion          *int           `json:"appid_version,omitempty"`
    EwfStatus             *string        `json:"ewf_status,omitempty"`
    IdpInstallResult      *string        `json:"idp_install_result,omitempty"`
    IdpInstallTimestamp   *string        `json:"idp_install_timestamp,omitempty"`
    IdpPolicy             *string        `json:"idp_policy,omitempty"`
    IdpStatus             *string        `json:"idp_status,omitempty"`
    IdpUpdateTimestamp    *string        `json:"idp_update_timestamp,omitempty"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayStatsServiceStatus.
// It customizes the JSON marshaling process for GatewayStatsServiceStatus objects.
func (g GatewayStatsServiceStatus) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayStatsServiceStatus object to a map representation for JSON marshaling.
func (g GatewayStatsServiceStatus) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.AppidInstallResult != nil {
        structMap["appid_install_result"] = g.AppidInstallResult
    }
    if g.AppidInstallTimestamp != nil {
        structMap["appid_install_timestamp"] = g.AppidInstallTimestamp
    }
    if g.AppidStatus != nil {
        structMap["appid_status"] = g.AppidStatus
    }
    if g.AppidVersion != nil {
        structMap["appid_version"] = g.AppidVersion
    }
    if g.EwfStatus != nil {
        structMap["ewf_status"] = g.EwfStatus
    }
    if g.IdpInstallResult != nil {
        structMap["idp_install_result"] = g.IdpInstallResult
    }
    if g.IdpInstallTimestamp != nil {
        structMap["idp_install_timestamp"] = g.IdpInstallTimestamp
    }
    if g.IdpPolicy != nil {
        structMap["idp_policy"] = g.IdpPolicy
    }
    if g.IdpStatus != nil {
        structMap["idp_status"] = g.IdpStatus
    }
    if g.IdpUpdateTimestamp != nil {
        structMap["idp_update_timestamp"] = g.IdpUpdateTimestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayStatsServiceStatus.
// It customizes the JSON unmarshaling process for GatewayStatsServiceStatus objects.
func (g *GatewayStatsServiceStatus) UnmarshalJSON(input []byte) error {
    var temp gatewayStatsServiceStatus
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "appid_install_result", "appid_install_timestamp", "appid_status", "appid_version", "ewf_status", "idp_install_result", "idp_install_timestamp", "idp_policy", "idp_status", "idp_update_timestamp")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.AppidInstallResult = temp.AppidInstallResult
    g.AppidInstallTimestamp = temp.AppidInstallTimestamp
    g.AppidStatus = temp.AppidStatus
    g.AppidVersion = temp.AppidVersion
    g.EwfStatus = temp.EwfStatus
    g.IdpInstallResult = temp.IdpInstallResult
    g.IdpInstallTimestamp = temp.IdpInstallTimestamp
    g.IdpPolicy = temp.IdpPolicy
    g.IdpStatus = temp.IdpStatus
    g.IdpUpdateTimestamp = temp.IdpUpdateTimestamp
    return nil
}

// gatewayStatsServiceStatus is a temporary struct used for validating the fields of GatewayStatsServiceStatus.
type gatewayStatsServiceStatus  struct {
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
