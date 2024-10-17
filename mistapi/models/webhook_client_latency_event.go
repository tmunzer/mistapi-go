package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// WebhookClientLatencyEvent represents a WebhookClientLatencyEvent struct.
type WebhookClientLatencyEvent struct {
	AvgAuth              *float64       `json:"avg_auth,omitempty"`
	AvgDhcp              *float64       `json:"avg_dhcp,omitempty"`
	AvgDns               *float64       `json:"avg_dns,omitempty"`
	MaxAuth              *float64       `json:"max_auth,omitempty"`
	MaxDhcp              *float64       `json:"max_dhcp,omitempty"`
	MaxDns               *float64       `json:"max_dns,omitempty"`
	MinAuth              *float64       `json:"min_auth,omitempty"`
	MinDhcp              *float64       `json:"min_dhcp,omitempty"`
	MinDns               *float64       `json:"min_dns,omitempty"`
	OrgId                *uuid.UUID     `json:"org_id,omitempty"`
	SiteId               *uuid.UUID     `json:"site_id,omitempty"`
	Timestamp            *int           `json:"timestamp,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientLatencyEvent.
// It customizes the JSON marshaling process for WebhookClientLatencyEvent objects.
func (w WebhookClientLatencyEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientLatencyEvent object to a map representation for JSON marshaling.
func (w WebhookClientLatencyEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
	if w.AvgAuth != nil {
		structMap["avg_auth"] = w.AvgAuth
	}
	if w.AvgDhcp != nil {
		structMap["avg_dhcp"] = w.AvgDhcp
	}
	if w.AvgDns != nil {
		structMap["avg_dns"] = w.AvgDns
	}
	if w.MaxAuth != nil {
		structMap["max_auth"] = w.MaxAuth
	}
	if w.MaxDhcp != nil {
		structMap["max_dhcp"] = w.MaxDhcp
	}
	if w.MaxDns != nil {
		structMap["max_dns"] = w.MaxDns
	}
	if w.MinAuth != nil {
		structMap["min_auth"] = w.MinAuth
	}
	if w.MinDhcp != nil {
		structMap["min_dhcp"] = w.MinDhcp
	}
	if w.MinDns != nil {
		structMap["min_dns"] = w.MinDns
	}
	if w.OrgId != nil {
		structMap["org_id"] = w.OrgId
	}
	if w.SiteId != nil {
		structMap["site_id"] = w.SiteId
	}
	if w.Timestamp != nil {
		structMap["timestamp"] = w.Timestamp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientLatencyEvent.
// It customizes the JSON unmarshaling process for WebhookClientLatencyEvent objects.
func (w *WebhookClientLatencyEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookClientLatencyEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "avg_auth", "avg_dhcp", "avg_dns", "max_auth", "max_dhcp", "max_dns", "min_auth", "min_dhcp", "min_dns", "org_id", "site_id", "timestamp")
	if err != nil {
		return err
	}

	w.AdditionalProperties = additionalProperties
	w.AvgAuth = temp.AvgAuth
	w.AvgDhcp = temp.AvgDhcp
	w.AvgDns = temp.AvgDns
	w.MaxAuth = temp.MaxAuth
	w.MaxDhcp = temp.MaxDhcp
	w.MaxDns = temp.MaxDns
	w.MinAuth = temp.MinAuth
	w.MinDhcp = temp.MinDhcp
	w.MinDns = temp.MinDns
	w.OrgId = temp.OrgId
	w.SiteId = temp.SiteId
	w.Timestamp = temp.Timestamp
	return nil
}

// tempWebhookClientLatencyEvent is a temporary struct used for validating the fields of WebhookClientLatencyEvent.
type tempWebhookClientLatencyEvent struct {
	AvgAuth   *float64   `json:"avg_auth,omitempty"`
	AvgDhcp   *float64   `json:"avg_dhcp,omitempty"`
	AvgDns    *float64   `json:"avg_dns,omitempty"`
	MaxAuth   *float64   `json:"max_auth,omitempty"`
	MaxDhcp   *float64   `json:"max_dhcp,omitempty"`
	MaxDns    *float64   `json:"max_dns,omitempty"`
	MinAuth   *float64   `json:"min_auth,omitempty"`
	MinDhcp   *float64   `json:"min_dhcp,omitempty"`
	MinDns    *float64   `json:"min_dns,omitempty"`
	OrgId     *uuid.UUID `json:"org_id,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	Timestamp *int       `json:"timestamp,omitempty"`
}
