package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WlanAuth represents a WlanAuth struct.
// Authentication wlan settings
type WlanAuth struct {
    // SAE anti-clogging token threshold
    AnticlogThreshold    *int                       `json:"anticlog_threshold,omitempty"`
    // Whether to trigger EAP reauth when the session ends
    EapReauth            *bool                      `json:"eap_reauth,omitempty"`
    // Whether to enable MAC Auth, uses the same auth_servers
    EnableMacAuth        *bool                      `json:"enable_mac_auth,omitempty"`
    // When `type`==`wep`
    KeyIdx               *int                       `json:"key_idx,omitempty"`
    // When type=wep, four 10-character or 26-character hex string, null can be used. All keys, if provided, have to be in the same length
    Keys                 []string                   `json:"keys,omitempty"`
    // When `type`==`psk`, whether to only use multi_psk
    MultiPskOnly         *bool                      `json:"multi_psk_only,omitempty"`
    // if `type`==`open`. enum: `disabled`, `enabled` (means transition mode), `required`
    Owe                  *WlanAuthOweEnum           `json:"owe,omitempty"`
    // When `type`=`psk` or `type`=`eap`, one or more of `wpa1-ccmp`, `wpa1-tkip`, `wpa2-ccmp`, `wpa2-tkip`, `wpa3`
    Pairwise             []WlanAuthPairwiseItemEnum `json:"pairwise,omitempty"`
    // When `multi_psk_only`==`true`, whether private wlan is enabled
    PrivateWlan          *bool                      `json:"private_wlan,omitempty"`
    // When `type`==`psk`, 8-64 characters, or 64 hex characters
    Psk                  Optional[string]           `json:"psk"`
    // enum: `eap`, `eap192`, `open`, `psk`, `psk-tkip`, `psk-wpa2-tkip`, `wep`
    Type                 WlanAuthTypeEnum           `json:"type"`
    // Enable WEP as secondary auth
    WepAsSecondaryAuth   *bool                      `json:"wep_as_secondary_auth,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for WlanAuth,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanAuth) String() string {
    return fmt.Sprintf(
    	"WlanAuth[AnticlogThreshold=%v, EapReauth=%v, EnableMacAuth=%v, KeyIdx=%v, Keys=%v, MultiPskOnly=%v, Owe=%v, Pairwise=%v, PrivateWlan=%v, Psk=%v, Type=%v, WepAsSecondaryAuth=%v, AdditionalProperties=%v]",
    	w.AnticlogThreshold, w.EapReauth, w.EnableMacAuth, w.KeyIdx, w.Keys, w.MultiPskOnly, w.Owe, w.Pairwise, w.PrivateWlan, w.Psk, w.Type, w.WepAsSecondaryAuth, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanAuth.
// It customizes the JSON marshaling process for WlanAuth objects.
func (w WlanAuth) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "anticlog_threshold", "eap_reauth", "enable_mac_auth", "key_idx", "keys", "multi_psk_only", "owe", "pairwise", "private_wlan", "psk", "type", "wep_as_secondary_auth"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanAuth object to a map representation for JSON marshaling.
func (w WlanAuth) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AnticlogThreshold != nil {
        structMap["anticlog_threshold"] = w.AnticlogThreshold
    }
    if w.EapReauth != nil {
        structMap["eap_reauth"] = w.EapReauth
    }
    if w.EnableMacAuth != nil {
        structMap["enable_mac_auth"] = w.EnableMacAuth
    }
    if w.KeyIdx != nil {
        structMap["key_idx"] = w.KeyIdx
    }
    if w.Keys != nil {
        structMap["keys"] = w.Keys
    }
    if w.MultiPskOnly != nil {
        structMap["multi_psk_only"] = w.MultiPskOnly
    }
    if w.Owe != nil {
        structMap["owe"] = w.Owe
    }
    if w.Pairwise != nil {
        structMap["pairwise"] = w.Pairwise
    }
    if w.PrivateWlan != nil {
        structMap["private_wlan"] = w.PrivateWlan
    }
    if w.Psk.IsValueSet() {
        if w.Psk.Value() != nil {
            structMap["psk"] = w.Psk.Value()
        } else {
            structMap["psk"] = nil
        }
    }
    structMap["type"] = w.Type
    if w.WepAsSecondaryAuth != nil {
        structMap["wep_as_secondary_auth"] = w.WepAsSecondaryAuth
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanAuth.
// It customizes the JSON unmarshaling process for WlanAuth objects.
func (w *WlanAuth) UnmarshalJSON(input []byte) error {
    var temp tempWlanAuth
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "anticlog_threshold", "eap_reauth", "enable_mac_auth", "key_idx", "keys", "multi_psk_only", "owe", "pairwise", "private_wlan", "psk", "type", "wep_as_secondary_auth")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.AnticlogThreshold = temp.AnticlogThreshold
    w.EapReauth = temp.EapReauth
    w.EnableMacAuth = temp.EnableMacAuth
    w.KeyIdx = temp.KeyIdx
    w.Keys = temp.Keys
    w.MultiPskOnly = temp.MultiPskOnly
    w.Owe = temp.Owe
    w.Pairwise = temp.Pairwise
    w.PrivateWlan = temp.PrivateWlan
    w.Psk = temp.Psk
    w.Type = *temp.Type
    w.WepAsSecondaryAuth = temp.WepAsSecondaryAuth
    return nil
}

// tempWlanAuth is a temporary struct used for validating the fields of WlanAuth.
type tempWlanAuth  struct {
    AnticlogThreshold  *int                       `json:"anticlog_threshold,omitempty"`
    EapReauth          *bool                      `json:"eap_reauth,omitempty"`
    EnableMacAuth      *bool                      `json:"enable_mac_auth,omitempty"`
    KeyIdx             *int                       `json:"key_idx,omitempty"`
    Keys               []string                   `json:"keys,omitempty"`
    MultiPskOnly       *bool                      `json:"multi_psk_only,omitempty"`
    Owe                *WlanAuthOweEnum           `json:"owe,omitempty"`
    Pairwise           []WlanAuthPairwiseItemEnum `json:"pairwise,omitempty"`
    PrivateWlan        *bool                      `json:"private_wlan,omitempty"`
    Psk                Optional[string]           `json:"psk"`
    Type               *WlanAuthTypeEnum          `json:"type"`
    WepAsSecondaryAuth *bool                      `json:"wep_as_secondary_auth,omitempty"`
}

func (w *tempWlanAuth) validate() error {
    var errs []string
    if w.Type == nil {
        errs = append(errs, "required field `type` is missing for type `wlan_auth`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
