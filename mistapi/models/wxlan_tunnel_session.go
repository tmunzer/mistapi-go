package models

import (
    "encoding/json"
)

// WxlanTunnelSession represents a WxlanTunnelSession struct.
type WxlanTunnelSession struct {
    // if `use_ap_as_session_ids`==`true`, only apmac is supported right now. This is the name WLAN should use for wxtunnel_remote_id
    ApAsSessionId        *string                          `json:"ap_as_session_id,omitempty"`
    // optional, user-specified string for display purpose
    Comment              *string                          `json:"comment,omitempty"`
    EnableCookie         *bool                            `json:"enable_cookie,omitempty"`
    // enum: `ethernet`, `vlan`
    Ethertype            *WxlanTunnelSessionEthertypeEnum `json:"ethertype,omitempty"`
    // 1-2147483647
    LocalSessionId       *int                             `json:"local_session_id,omitempty"`
    // optional. Enables the pseudo 802.1ad QinQ mode where the AP device drops the outer vlan tag (QinQ). This mode is useful when tunneling Mist AP’s to some aggregation routers.
    Pseudo8021adEnabled  *bool                            `json:"pseudo_802.1ad_enabled,omitempty"`
    // remote-id of the session, has to be unique in the same tunnel
    RemoteId             *string                          `json:"remote_id,omitempty"`
    // 1-2147483647
    RemoteSessionId      *int                             `json:"remote_session_id,omitempty"`
    // whether to use AP (last 4 bytes of MAC currently) as session ids
    UseApAsSessionIds    *bool                            `json:"use_ap_as_session_ids,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WxlanTunnelSession.
// It customizes the JSON marshaling process for WxlanTunnelSession objects.
func (w WxlanTunnelSession) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WxlanTunnelSession object to a map representation for JSON marshaling.
func (w WxlanTunnelSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.ApAsSessionId != nil {
        structMap["ap_as_session_id"] = w.ApAsSessionId
    }
    if w.Comment != nil {
        structMap["comment"] = w.Comment
    }
    if w.EnableCookie != nil {
        structMap["enable_cookie"] = w.EnableCookie
    }
    if w.Ethertype != nil {
        structMap["ethertype"] = w.Ethertype
    }
    if w.LocalSessionId != nil {
        structMap["local_session_id"] = w.LocalSessionId
    }
    if w.Pseudo8021adEnabled != nil {
        structMap["pseudo_802.1ad_enabled"] = w.Pseudo8021adEnabled
    }
    if w.RemoteId != nil {
        structMap["remote_id"] = w.RemoteId
    }
    if w.RemoteSessionId != nil {
        structMap["remote_session_id"] = w.RemoteSessionId
    }
    if w.UseApAsSessionIds != nil {
        structMap["use_ap_as_session_ids"] = w.UseApAsSessionIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxlanTunnelSession.
// It customizes the JSON unmarshaling process for WxlanTunnelSession objects.
func (w *WxlanTunnelSession) UnmarshalJSON(input []byte) error {
    var temp tempWxlanTunnelSession
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_as_session_id", "comment", "enable_cookie", "ethertype", "local_session_id", "pseudo_802.1ad_enabled", "remote_id", "remote_session_id", "use_ap_as_session_ids")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.ApAsSessionId = temp.ApAsSessionId
    w.Comment = temp.Comment
    w.EnableCookie = temp.EnableCookie
    w.Ethertype = temp.Ethertype
    w.LocalSessionId = temp.LocalSessionId
    w.Pseudo8021adEnabled = temp.Pseudo8021adEnabled
    w.RemoteId = temp.RemoteId
    w.RemoteSessionId = temp.RemoteSessionId
    w.UseApAsSessionIds = temp.UseApAsSessionIds
    return nil
}

// tempWxlanTunnelSession is a temporary struct used for validating the fields of WxlanTunnelSession.
type tempWxlanTunnelSession  struct {
    ApAsSessionId       *string                          `json:"ap_as_session_id,omitempty"`
    Comment             *string                          `json:"comment,omitempty"`
    EnableCookie        *bool                            `json:"enable_cookie,omitempty"`
    Ethertype           *WxlanTunnelSessionEthertypeEnum `json:"ethertype,omitempty"`
    LocalSessionId      *int                             `json:"local_session_id,omitempty"`
    Pseudo8021adEnabled *bool                            `json:"pseudo_802.1ad_enabled,omitempty"`
    RemoteId            *string                          `json:"remote_id,omitempty"`
    RemoteSessionId     *int                             `json:"remote_session_id,omitempty"`
    UseApAsSessionIds   *bool                            `json:"use_ap_as_session_ids,omitempty"`
}
