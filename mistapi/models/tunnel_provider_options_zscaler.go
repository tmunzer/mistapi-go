package models

import (
    "encoding/json"
)

// TunnelProviderOptionsZscaler represents a TunnelProviderOptionsZscaler struct.
// for zscaler-ipsec and zscaler-gre
type TunnelProviderOptionsZscaler struct {
    AupAcceptanceRequired *bool                                     `json:"aup_acceptance_required,omitempty"`
    // days before AUP is requested again
    AupExpire             *int                                      `json:"aup_expire,omitempty"`
    // proxy HTTPs traffic, requiring Zscaler cert to be installed in browser
    AupSslProxy           *bool                                     `json:"aup_ssl_proxy,omitempty"`
    // the download bandwidth cap of the link, in Mbps
    DownloadMbps          *int                                      `json:"download_mbps,omitempty"`
    // if `use_xff`==`true`, display Acceptable Use Policy (AUP)
    EnableAup             *bool                                     `json:"enable_aup,omitempty"`
    // when `enforce_authentication`==`false`, display caution notification for non-authenticated users
    EnableCaution         *bool                                     `json:"enable_caution,omitempty"`
    EnforceAuthentication *bool                                     `json:"enforce_authentication,omitempty"`
    Name                  *string                                   `json:"name,omitempty"`
    // if `use_xff`==`true`
    SubLocations          []TunnelProviderOptionsZscalerSubLocation `json:"sub_locations,omitempty"`
    // the download bandwidth cap of the link, in Mbps
    UploadMbps            *int                                      `json:"upload_mbps,omitempty"`
    // location uses proxy chaining to forward traffic
    UseXff                *bool                                     `json:"use_xff,omitempty"`
    AdditionalProperties  map[string]interface{}                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TunnelProviderOptionsZscaler.
// It customizes the JSON marshaling process for TunnelProviderOptionsZscaler objects.
func (t TunnelProviderOptionsZscaler) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "aup_acceptance_required", "aup_expire", "aup_ssl_proxy", "download_mbps", "enable_aup", "enable_caution", "enforce_authentication", "name", "sub_locations", "upload_mbps", "use_xff"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelProviderOptionsZscaler object to a map representation for JSON marshaling.
func (t TunnelProviderOptionsZscaler) toMap() map[string]any {
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
    if t.Name != nil {
        structMap["name"] = t.Name
    }
    if t.SubLocations != nil {
        structMap["sub_locations"] = t.SubLocations
    }
    if t.UploadMbps != nil {
        structMap["upload_mbps"] = t.UploadMbps
    }
    if t.UseXff != nil {
        structMap["use_xff"] = t.UseXff
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelProviderOptionsZscaler.
// It customizes the JSON unmarshaling process for TunnelProviderOptionsZscaler objects.
func (t *TunnelProviderOptionsZscaler) UnmarshalJSON(input []byte) error {
    var temp tempTunnelProviderOptionsZscaler
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aup_acceptance_required", "aup_expire", "aup_ssl_proxy", "download_mbps", "enable_aup", "enable_caution", "enforce_authentication", "name", "sub_locations", "upload_mbps", "use_xff")
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
    t.Name = temp.Name
    t.SubLocations = temp.SubLocations
    t.UploadMbps = temp.UploadMbps
    t.UseXff = temp.UseXff
    return nil
}

// tempTunnelProviderOptionsZscaler is a temporary struct used for validating the fields of TunnelProviderOptionsZscaler.
type tempTunnelProviderOptionsZscaler  struct {
    AupAcceptanceRequired *bool                                     `json:"aup_acceptance_required,omitempty"`
    AupExpire             *int                                      `json:"aup_expire,omitempty"`
    AupSslProxy           *bool                                     `json:"aup_ssl_proxy,omitempty"`
    DownloadMbps          *int                                      `json:"download_mbps,omitempty"`
    EnableAup             *bool                                     `json:"enable_aup,omitempty"`
    EnableCaution         *bool                                     `json:"enable_caution,omitempty"`
    EnforceAuthentication *bool                                     `json:"enforce_authentication,omitempty"`
    Name                  *string                                   `json:"name,omitempty"`
    SubLocations          []TunnelProviderOptionsZscalerSubLocation `json:"sub_locations,omitempty"`
    UploadMbps            *int                                      `json:"upload_mbps,omitempty"`
    UseXff                *bool                                     `json:"use_xff,omitempty"`
}
