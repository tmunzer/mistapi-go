package models

import (
    "encoding/json"
)

// WlanDynamicPsk represents a WlanDynamicPsk struct.
// for dynamic PSK where we get per_user PSK from Radius. dynamic_psk allows PSK to be selected at runtime depending on context (wlan/site/user/...) thus following configurations are assumed (currently)
//   * PSK will come from RADIUS server
//   * AP sends client MAC as username ans password (i.e. `enable_mac_auth` is assumed)
//   * AP sends BSSID:SSID as Caller-Station-ID
//   * `auth_servers` is required
//   * PSK will come from cloud WLC if source is cloud_psks
//   * default_psk will be used if cloud WLC is not available
//   * `multi_psk_only` and `psk` is ignored
//   * `pairwise` can only be wpa2-ccmp (for now, wpa3 support on the roadmap)
type WlanDynamicPsk struct {
    // default PSK to use if cloud WLC is not available, 8-63 characters
    DefaultPsk           *string                      `json:"default_psk,omitempty"`
    DefaultVlanId        *WlanDynamicPskDefaultVlanId `json:"default_vlan_id,omitempty"`
    Enabled              *bool                        `json:"enabled,omitempty"`
    // when 11r is enabled, we'll try to use the cached PMK, this can be disabled
    // `false` means auto
    ForceLookup          *bool                        `json:"force_lookup,omitempty"`
    Source               *DynamicPskSourceEnum        `json:"source,omitempty"`
    VlanIds              []WlanDynamicPskVlanIds      `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanDynamicPsk.
// It customizes the JSON marshaling process for WlanDynamicPsk objects.
func (w WlanDynamicPsk) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanDynamicPsk object to a map representation for JSON marshaling.
func (w WlanDynamicPsk) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.DefaultPsk != nil {
        structMap["default_psk"] = w.DefaultPsk
    }
    if w.DefaultVlanId != nil {
        structMap["default_vlan_id"] = w.DefaultVlanId.toMap()
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.ForceLookup != nil {
        structMap["force_lookup"] = w.ForceLookup
    }
    if w.Source != nil {
        structMap["source"] = w.Source
    }
    if w.VlanIds != nil {
        structMap["vlan_ids"] = w.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanDynamicPsk.
// It customizes the JSON unmarshaling process for WlanDynamicPsk objects.
func (w *WlanDynamicPsk) UnmarshalJSON(input []byte) error {
    var temp wlanDynamicPsk
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "default_psk", "default_vlan_id", "enabled", "force_lookup", "source", "vlan_ids")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.DefaultPsk = temp.DefaultPsk
    w.DefaultVlanId = temp.DefaultVlanId
    w.Enabled = temp.Enabled
    w.ForceLookup = temp.ForceLookup
    w.Source = temp.Source
    w.VlanIds = temp.VlanIds
    return nil
}

// wlanDynamicPsk is a temporary struct used for validating the fields of WlanDynamicPsk.
type wlanDynamicPsk  struct {
    DefaultPsk    *string                      `json:"default_psk,omitempty"`
    DefaultVlanId *WlanDynamicPskDefaultVlanId `json:"default_vlan_id,omitempty"`
    Enabled       *bool                        `json:"enabled,omitempty"`
    ForceLookup   *bool                        `json:"force_lookup,omitempty"`
    Source        *DynamicPskSourceEnum        `json:"source,omitempty"`
    VlanIds       []WlanDynamicPskVlanIds      `json:"vlan_ids,omitempty"`
}
