package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WebhookOccupancyAlertsEvent represents a WebhookOccupancyAlertsEvent struct.
type WebhookOccupancyAlertsEvent struct {
    // list of occupancy alerts for non-compliance zones within the site detected around the same time
    AlertEvents          []WebhookOccupancyAlertsEventAlertEventsItems `json:"alert_events,omitempty"`
    ForSite              *bool                                         `json:"for_site,omitempty"`
    SiteId               uuid.UUID                                     `json:"site_id"`
    SiteName             string                                        `json:"site_name"`
    AdditionalProperties map[string]any                                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookOccupancyAlertsEvent.
// It customizes the JSON marshaling process for WebhookOccupancyAlertsEvent objects.
func (w WebhookOccupancyAlertsEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookOccupancyAlertsEvent object to a map representation for JSON marshaling.
func (w WebhookOccupancyAlertsEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AlertEvents != nil {
        structMap["alert_events"] = w.AlertEvents
    }
    if w.ForSite != nil {
        structMap["for_site"] = w.ForSite
    }
    structMap["site_id"] = w.SiteId
    structMap["site_name"] = w.SiteName
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookOccupancyAlertsEvent.
// It customizes the JSON unmarshaling process for WebhookOccupancyAlertsEvent objects.
func (w *WebhookOccupancyAlertsEvent) UnmarshalJSON(input []byte) error {
    var temp webhookOccupancyAlertsEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "alert_events", "for_site", "site_id", "site_name")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.AlertEvents = temp.AlertEvents
    w.ForSite = temp.ForSite
    w.SiteId = *temp.SiteId
    w.SiteName = *temp.SiteName
    return nil
}

// webhookOccupancyAlertsEvent is a temporary struct used for validating the fields of WebhookOccupancyAlertsEvent.
type webhookOccupancyAlertsEvent  struct {
    AlertEvents []WebhookOccupancyAlertsEventAlertEventsItems `json:"alert_events,omitempty"`
    ForSite     *bool                                         `json:"for_site,omitempty"`
    SiteId      *uuid.UUID                                    `json:"site_id"`
    SiteName    *string                                       `json:"site_name"`
}

func (w *webhookOccupancyAlertsEvent) validate() error {
    var errs []string
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `Webhook_Occupancy_Alerts_Event`")
    }
    if w.SiteName == nil {
        errs = append(errs, "required field `site_name` is missing for type `Webhook_Occupancy_Alerts_Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
