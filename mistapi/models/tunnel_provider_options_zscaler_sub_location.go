package models

import (
    "encoding/json"
)

// TunnelProviderOptionsZscalerSubLocation represents a TunnelProviderOptionsZscalerSubLocation struct.
type TunnelProviderOptionsZscalerSubLocation struct {
    AupAcceptanceRequired *bool                  `json:"aup_acceptance_required,omitempty"`
    // days before AUP is requested again
    AupExpire             *int                   `json:"aup_expire,omitempty"`
    // proxy HTTPs traffic, requiring Zscaler cert to be installed in browser
    AupSslProxy           *bool                  `json:"aup_ssl_proxy,omitempty"`
    // the download bandwidth cap of the link, in Mbps
    DownloadMbps          *int                   `json:"download_mbps,omitempty"`
    // if `use_xff`==`true`, display Acceptable Use Policy (AUP)
    EnableAup             *bool                  `json:"enable_aup,omitempty"`
    // when `enforce_authentication`==`false`, display caution notification for non-authenticated users
    EnableCaution         *bool                  `json:"enable_caution,omitempty"`
    EnforceAuthentication *bool                  `json:"enforce_authentication,omitempty"`
    Subnets               []string               `json:"subnets,omitempty"`
    // the download bandwidth cap of the link, in Mbps
    UploadMbps            *int                   `json:"upload_mbps,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TunnelProviderOptionsZscalerSubLocation.
// It customizes the JSON marshaling process for TunnelProviderOptionsZscalerSubLocation objects.
func (t TunnelProviderOptionsZscalerSubLocation) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "aup_acceptance_required", "aup_expire", "aup_ssl_proxy", "download_mbps", "enable_aup", "enable_caution", "enforce_authentication", "subnets", "upload_mbps"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelProviderOptionsZscalerSubLocation object to a map representation for JSON marshaling.
func (t TunnelProviderOptionsZscalerSubLocation) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.AupAcceptanceRequired != nil {
        structMap["aup_acceptance_required"] = t.AupAcceptanceRequired
    }
    if t.AupExpire != nil {
        structMap["aup_expire"] = t.AupExpire
    }
    if t.AupSslProxy != nil {
        structMap["aup_ssl_proxy"] = t.AupSslProxy
    }
    if t.DownloadMbps != nil {
        structMap["download_mbps"] = t.DownloadMbps
    }
    if t.EnableAup != nil {
        structMap["enable_aup"] = t.EnableAup
    }
    if t.EnableCaution != nil {
        structMap["enable_caution"] = t.EnableCaution
    }
    if t.EnforceAuthentication != nil {
        structMap["enforce_authentication"] = t.EnforceAuthentication
    }
    if t.Subnets != nil {
        structMap["subnets"] = t.Subnets
    }
    if t.UploadMbps != nil {
        structMap["upload_mbps"] = t.UploadMbps
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelProviderOptionsZscalerSubLocation.
// It customizes the JSON unmarshaling process for TunnelProviderOptionsZscalerSubLocation objects.
func (t *TunnelProviderOptionsZscalerSubLocation) UnmarshalJSON(input []byte) error {
    var temp tempTunnelProviderOptionsZscalerSubLocation
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aup_acceptance_required", "aup_expire", "aup_ssl_proxy", "download_mbps", "enable_aup", "enable_caution", "enforce_authentication", "subnets", "upload_mbps")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.AupAcceptanceRequired = temp.AupAcceptanceRequired
    t.AupExpire = temp.AupExpire
    t.AupSslProxy = temp.AupSslProxy
    t.DownloadMbps = temp.DownloadMbps
    t.EnableAup = temp.EnableAup
    t.EnableCaution = temp.EnableCaution
    t.EnforceAuthentication = temp.EnforceAuthentication
    t.Subnets = temp.Subnets
    t.UploadMbps = temp.UploadMbps
    return nil
}

// tempTunnelProviderOptionsZscalerSubLocation is a temporary struct used for validating the fields of TunnelProviderOptionsZscalerSubLocation.
type tempTunnelProviderOptionsZscalerSubLocation  struct {
    AupAcceptanceRequired *bool    `json:"aup_acceptance_required,omitempty"`
    AupExpire             *int     `json:"aup_expire,omitempty"`
    AupSslProxy           *bool    `json:"aup_ssl_proxy,omitempty"`
    DownloadMbps          *int     `json:"download_mbps,omitempty"`
    EnableAup             *bool    `json:"enable_aup,omitempty"`
    EnableCaution         *bool    `json:"enable_caution,omitempty"`
    EnforceAuthentication *bool    `json:"enforce_authentication,omitempty"`
    Subnets               []string `json:"subnets,omitempty"`
    UploadMbps            *int     `json:"upload_mbps,omitempty"`
}
