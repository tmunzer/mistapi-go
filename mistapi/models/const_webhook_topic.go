package models

import (
    "encoding/json"
    "fmt"
)

// ConstWebhookTopic represents a ConstWebhookTopic struct.
type ConstWebhookTopic struct {
    // Can be used in org webhooks, optional
    ForOrg               *bool                  `json:"for_org,omitempty"`
    // Supports webhook delivery results /api/v1/:scope/:scope_id/webhooks/:webhook_id/events/search
    HasDeliveryResults   *bool                  `json:"has_delivery_results,omitempty"`
    // Internal topic (not selectable in site/org webhooks)
    Internal             *bool                  `json:"internal,omitempty"`
    // Webhook topic name
    Key                  *string                `json:"key,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstWebhookTopic,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstWebhookTopic) String() string {
    return fmt.Sprintf(
    	"ConstWebhookTopic[ForOrg=%v, HasDeliveryResults=%v, Internal=%v, Key=%v, AdditionalProperties=%v]",
    	c.ForOrg, c.HasDeliveryResults, c.Internal, c.Key, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstWebhookTopic.
// It customizes the JSON marshaling process for ConstWebhookTopic objects.
func (c ConstWebhookTopic) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "for_org", "has_delivery_results", "internal", "key"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstWebhookTopic object to a map representation for JSON marshaling.
func (c ConstWebhookTopic) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ForOrg != nil {
        structMap["for_org"] = c.ForOrg
    }
    if c.HasDeliveryResults != nil {
        structMap["has_delivery_results"] = c.HasDeliveryResults
    }
    if c.Internal != nil {
        structMap["internal"] = c.Internal
    }
    if c.Key != nil {
        structMap["key"] = c.Key
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstWebhookTopic.
// It customizes the JSON unmarshaling process for ConstWebhookTopic objects.
func (c *ConstWebhookTopic) UnmarshalJSON(input []byte) error {
    var temp tempConstWebhookTopic
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "for_org", "has_delivery_results", "internal", "key")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ForOrg = temp.ForOrg
    c.HasDeliveryResults = temp.HasDeliveryResults
    c.Internal = temp.Internal
    c.Key = temp.Key
    return nil
}

// tempConstWebhookTopic is a temporary struct used for validating the fields of ConstWebhookTopic.
type tempConstWebhookTopic  struct {
    ForOrg             *bool   `json:"for_org,omitempty"`
    HasDeliveryResults *bool   `json:"has_delivery_results,omitempty"`
    Internal           *bool   `json:"internal,omitempty"`
    Key                *string `json:"key,omitempty"`
}
