package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// WebhookDelivery represents a WebhookDelivery struct.
type WebhookDelivery struct {
    // error message, if there is one
    Error                *string                    `json:"error,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                 `json:"id,omitempty"`
    OrgId                *uuid.UUID                 `json:"org_id,omitempty"`
    // HTTP request headers
    ReqHeaders           *string                    `json:"req_headers,omitempty"`
    // HTTP request payload
    ReqPayload           *string                    `json:"req_payload,omitempty"`
    // HTTP request URL
    ReqUrl               *string                    `json:"req_url,omitempty"`
    // HTTP response body
    RespBody             *string                    `json:"resp_body,omitempty"`
    // HTTP response headers
    RespHeaders          *string                    `json:"resp_headers,omitempty"`
    SiteId               *uuid.UUID                 `json:"site_id,omitempty"`
    // webhook delivery status. enum: `failure`, `success`
    Status               *WebhookDeliveryStatusEnum `json:"status,omitempty"`
    StatusCode           *int                       `json:"status_code,omitempty"`
    Timestamp            *float64                   `json:"timestamp,omitempty"`
    // webhook topic. enum: `alarms`, `audits`, `device-updowns`, `occupancy-alerts`, `ping`
    Topic                *WebhookDeliveryTopicEnum  `json:"topic,omitempty"`
    WebhookId            *uuid.UUID                 `json:"webhook_id,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookDelivery.
// It customizes the JSON marshaling process for WebhookDelivery objects.
func (w WebhookDelivery) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookDelivery object to a map representation for JSON marshaling.
func (w WebhookDelivery) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Error != nil {
        structMap["error"] = w.Error
    }
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.ReqHeaders != nil {
        structMap["req_headers"] = w.ReqHeaders
    }
    if w.ReqPayload != nil {
        structMap["req_payload"] = w.ReqPayload
    }
    if w.ReqUrl != nil {
        structMap["req_url"] = w.ReqUrl
    }
    if w.RespBody != nil {
        structMap["resp_body"] = w.RespBody
    }
    if w.RespHeaders != nil {
        structMap["resp_headers"] = w.RespHeaders
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.Status != nil {
        structMap["status"] = w.Status
    }
    if w.StatusCode != nil {
        structMap["status_code"] = w.StatusCode
    }
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    if w.Topic != nil {
        structMap["topic"] = w.Topic
    }
    if w.WebhookId != nil {
        structMap["webhook_id"] = w.WebhookId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookDelivery.
// It customizes the JSON unmarshaling process for WebhookDelivery objects.
func (w *WebhookDelivery) UnmarshalJSON(input []byte) error {
    var temp tempWebhookDelivery
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "error", "id", "org_id", "req_headers", "req_payload", "req_url", "resp_body", "resp_headers", "site_id", "status", "status_code", "timestamp", "topic", "webhook_id")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Error = temp.Error
    w.Id = temp.Id
    w.OrgId = temp.OrgId
    w.ReqHeaders = temp.ReqHeaders
    w.ReqPayload = temp.ReqPayload
    w.ReqUrl = temp.ReqUrl
    w.RespBody = temp.RespBody
    w.RespHeaders = temp.RespHeaders
    w.SiteId = temp.SiteId
    w.Status = temp.Status
    w.StatusCode = temp.StatusCode
    w.Timestamp = temp.Timestamp
    w.Topic = temp.Topic
    w.WebhookId = temp.WebhookId
    return nil
}

// tempWebhookDelivery is a temporary struct used for validating the fields of WebhookDelivery.
type tempWebhookDelivery  struct {
    Error       *string                    `json:"error,omitempty"`
    Id          *uuid.UUID                 `json:"id,omitempty"`
    OrgId       *uuid.UUID                 `json:"org_id,omitempty"`
    ReqHeaders  *string                    `json:"req_headers,omitempty"`
    ReqPayload  *string                    `json:"req_payload,omitempty"`
    ReqUrl      *string                    `json:"req_url,omitempty"`
    RespBody    *string                    `json:"resp_body,omitempty"`
    RespHeaders *string                    `json:"resp_headers,omitempty"`
    SiteId      *uuid.UUID                 `json:"site_id,omitempty"`
    Status      *WebhookDeliveryStatusEnum `json:"status,omitempty"`
    StatusCode  *int                       `json:"status_code,omitempty"`
    Timestamp   *float64                   `json:"timestamp,omitempty"`
    Topic       *WebhookDeliveryTopicEnum  `json:"topic,omitempty"`
    WebhookId   *uuid.UUID                 `json:"webhook_id,omitempty"`
}
