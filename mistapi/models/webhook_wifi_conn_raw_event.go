// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// WebhookWifiConnRawEvent represents a WebhookWifiConnRawEvent struct.
type WebhookWifiConnRawEvent struct {
	ApId *string `json:"ap_id,omitempty"`
	// optional, coordinates (if any) of reporting AP (updated once in 60s per client)
	ApLoc         []float64 `json:"ap_loc,omitempty"`
	ClientId      *string   `json:"client_id,omitempty"`
	ConnectedSite *bool     `json:"connected_site,omitempty"`
	// optional, list of specific telemetry packets emited by certain wifi tags (Eg. Centrak)
	ExtendedInfoList []WebhookWifiConnRawEventExtendedInfo `json:"extended_info_list,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	MapId                *uuid.UUID                      `json:"map_id,omitempty"`
	OrgId                *uuid.UUID                      `json:"org_id,omitempty"`
	Packets              []WebhookWifiConnRawEventPacket `json:"packets,omitempty"`
	SiteId               *uuid.UUID                      `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookWifiConnRawEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookWifiConnRawEvent) String() string {
	return fmt.Sprintf(
		"WebhookWifiConnRawEvent[ApId=%v, ApLoc=%v, ClientId=%v, ConnectedSite=%v, ExtendedInfoList=%v, MapId=%v, OrgId=%v, Packets=%v, SiteId=%v, AdditionalProperties=%v]",
		w.ApId, w.ApLoc, w.ClientId, w.ConnectedSite, w.ExtendedInfoList, w.MapId, w.OrgId, w.Packets, w.SiteId, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookWifiConnRawEvent.
// It customizes the JSON marshaling process for WebhookWifiConnRawEvent objects.
func (w WebhookWifiConnRawEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"ap_id", "ap_loc", "client_id", "connected_site", "extended_info_list", "map_id", "org_id", "packets", "site_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookWifiConnRawEvent object to a map representation for JSON marshaling.
func (w WebhookWifiConnRawEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.ApId != nil {
		structMap["ap_id"] = w.ApId
	}
	if w.ApLoc != nil {
		structMap["ap_loc"] = w.ApLoc
	}
	if w.ClientId != nil {
		structMap["client_id"] = w.ClientId
	}
	if w.ConnectedSite != nil {
		structMap["connected_site"] = w.ConnectedSite
	}
	if w.ExtendedInfoList != nil {
		structMap["extended_info_list"] = w.ExtendedInfoList
	}
	if w.MapId != nil {
		structMap["map_id"] = w.MapId
	}
	if w.OrgId != nil {
		structMap["org_id"] = w.OrgId
	}
	if w.Packets != nil {
		structMap["packets"] = w.Packets
	}
	if w.SiteId != nil {
		structMap["site_id"] = w.SiteId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookWifiConnRawEvent.
// It customizes the JSON unmarshaling process for WebhookWifiConnRawEvent objects.
func (w *WebhookWifiConnRawEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookWifiConnRawEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_id", "ap_loc", "client_id", "connected_site", "extended_info_list", "map_id", "org_id", "packets", "site_id")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.ApId = temp.ApId
	w.ApLoc = temp.ApLoc
	w.ClientId = temp.ClientId
	w.ConnectedSite = temp.ConnectedSite
	w.ExtendedInfoList = temp.ExtendedInfoList
	w.MapId = temp.MapId
	w.OrgId = temp.OrgId
	w.Packets = temp.Packets
	w.SiteId = temp.SiteId
	return nil
}

// tempWebhookWifiConnRawEvent is a temporary struct used for validating the fields of WebhookWifiConnRawEvent.
type tempWebhookWifiConnRawEvent struct {
	ApId             *string                               `json:"ap_id,omitempty"`
	ApLoc            []float64                             `json:"ap_loc,omitempty"`
	ClientId         *string                               `json:"client_id,omitempty"`
	ConnectedSite    *bool                                 `json:"connected_site,omitempty"`
	ExtendedInfoList []WebhookWifiConnRawEventExtendedInfo `json:"extended_info_list,omitempty"`
	MapId            *uuid.UUID                            `json:"map_id,omitempty"`
	OrgId            *uuid.UUID                            `json:"org_id,omitempty"`
	Packets          []WebhookWifiConnRawEventPacket       `json:"packets,omitempty"`
	SiteId           *uuid.UUID                            `json:"site_id,omitempty"`
}
