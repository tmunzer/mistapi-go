package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Webhook represents a Webhook struct.
type Webhook struct {
    // When the object has been created, in epoch
    CreatedTime          *float64                    `json:"created_time,omitempty"`
    // Whether webhook is enabled
    Enabled              *bool                       `json:"enabled,omitempty"`
    ForSite              *bool                       `json:"for_site,omitempty"`
    // If `type`=`http-post`, additional custom HTTP headers to add. The headers name and value must be string, total bytes of headers name and value must be less than 1000
    Headers              Optional[map[string]string] `json:"headers"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                  `json:"id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                    `json:"modified_time,omitempty"`
    // Name of the webhook
    Name                 Optional[string]            `json:"name"`
    // Required when `oauth2_grant_type`==`client_credentials`
    Oauth2ClientId       *string                     `json:"oauth2_client_id,omitempty"`
    // Required when `oauth2_grant_type`==`client_credentials`
    Oauth2ClientSecret   *string                     `json:"oauth2_client_secret,omitempty"`
    // required when `type`==`oauth2`. enum: `client_credentials`, `password`
    Oauth2GrantType      *WebhookOauth2GrantTypeEnum `json:"oauth2_grant_type,omitempty"`
    // Required when `oauth2_grant_type`==`password`
    Oauth2Password       *string                     `json:"oauth2_password,omitempty"`
    // Required when `type`==`oauth2`, if provided, will be used in the token request
    Oauth2Scopes         []string                    `json:"oauth2_scopes,omitempty"`
    // Required when `type`==`oauth2`
    Oauth2TokenUrl       *string                     `json:"oauth2_token_url,omitempty"`
    // Required when `oauth2_grant_type`==`password`
    Oauth2Username       *string                     `json:"oauth2_username,omitempty"`
    OrgId                *uuid.UUID                  `json:"org_id,omitempty"`
    // Only if `type`=`http-post`
    // when `secret` is provided, two  HTTP headers will be added:
    // * X-Mist-Signature-v2: HMAC_SHA256(secret, body)
    // * X-Mist-Signature: HMAC_SHA1(secret, body)
    Secret               Optional[string]            `json:"secret"`
    SiteId               *uuid.UUID                  `json:"site_id,omitempty"`
    // Required if `type`=`splunk`. If splunk_token is not defined for a type Splunk webhook, it will not send, regardless if the webhook receiver is configured to accept it.'
    SplunkToken          Optional[string]            `json:"splunk_token"`
    // List of supported webhook topics available with the API Call [List Webhook Topics]($e/Constants%20Definitions/listWebhookTopics)
    Topics               []string                    `json:"topics,omitempty"`
    // enum: `aws-sns`, `google-pubsub`, `http-post`, `oauth2`, `splunk`
    Type                 *WebhookTypeEnum            `json:"type,omitempty"`
    Url                  *string                     `json:"url,omitempty"`
    // When url uses HTTPS, whether to verify the certificate
    VerifyCert           *bool                       `json:"verify_cert,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for Webhook,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w Webhook) String() string {
    return fmt.Sprintf(
    	"Webhook[CreatedTime=%v, Enabled=%v, ForSite=%v, Headers=%v, Id=%v, ModifiedTime=%v, Name=%v, Oauth2ClientId=%v, Oauth2ClientSecret=%v, Oauth2GrantType=%v, Oauth2Password=%v, Oauth2Scopes=%v, Oauth2TokenUrl=%v, Oauth2Username=%v, OrgId=%v, Secret=%v, SiteId=%v, SplunkToken=%v, Topics=%v, Type=%v, Url=%v, VerifyCert=%v, AdditionalProperties=%v]",
    	w.CreatedTime, w.Enabled, w.ForSite, w.Headers, w.Id, w.ModifiedTime, w.Name, w.Oauth2ClientId, w.Oauth2ClientSecret, w.Oauth2GrantType, w.Oauth2Password, w.Oauth2Scopes, w.Oauth2TokenUrl, w.Oauth2Username, w.OrgId, w.Secret, w.SiteId, w.SplunkToken, w.Topics, w.Type, w.Url, w.VerifyCert, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Webhook.
// It customizes the JSON marshaling process for Webhook objects.
func (w Webhook) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "created_time", "enabled", "for_site", "headers", "id", "modified_time", "name", "oauth2_client_id", "oauth2_client_secret", "oauth2_grant_type", "oauth2_password", "oauth2_scopes", "oauth2_token_url", "oauth2_username", "org_id", "secret", "site_id", "splunk_token", "topics", "type", "url", "verify_cert"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the Webhook object to a map representation for JSON marshaling.
func (w Webhook) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.CreatedTime != nil {
        structMap["created_time"] = w.CreatedTime
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.ForSite != nil {
        structMap["for_site"] = w.ForSite
    }
    if w.Headers.IsValueSet() {
        if w.Headers.Value() != nil {
            structMap["headers"] = w.Headers.Value()
        } else {
            structMap["headers"] = nil
        }
    }
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.ModifiedTime != nil {
        structMap["modified_time"] = w.ModifiedTime
    }
    if w.Name.IsValueSet() {
        if w.Name.Value() != nil {
            structMap["name"] = w.Name.Value()
        } else {
            structMap["name"] = nil
        }
    }
    if w.Oauth2ClientId != nil {
        structMap["oauth2_client_id"] = w.Oauth2ClientId
    }
    if w.Oauth2ClientSecret != nil {
        structMap["oauth2_client_secret"] = w.Oauth2ClientSecret
    }
    if w.Oauth2GrantType != nil {
        structMap["oauth2_grant_type"] = w.Oauth2GrantType
    }
    if w.Oauth2Password != nil {
        structMap["oauth2_password"] = w.Oauth2Password
    }
    if w.Oauth2Scopes != nil {
        structMap["oauth2_scopes"] = w.Oauth2Scopes
    }
    if w.Oauth2TokenUrl != nil {
        structMap["oauth2_token_url"] = w.Oauth2TokenUrl
    }
    if w.Oauth2Username != nil {
        structMap["oauth2_username"] = w.Oauth2Username
    }
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.Secret.IsValueSet() {
        if w.Secret.Value() != nil {
            structMap["secret"] = w.Secret.Value()
        } else {
            structMap["secret"] = nil
        }
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.SplunkToken.IsValueSet() {
        if w.SplunkToken.Value() != nil {
            structMap["splunk_token"] = w.SplunkToken.Value()
        } else {
            structMap["splunk_token"] = nil
        }
    }
    if w.Topics != nil {
        structMap["topics"] = w.Topics
    }
    if w.Type != nil {
        structMap["type"] = w.Type
    }
    if w.Url != nil {
        structMap["url"] = w.Url
    }
    if w.VerifyCert != nil {
        structMap["verify_cert"] = w.VerifyCert
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Webhook.
// It customizes the JSON unmarshaling process for Webhook objects.
func (w *Webhook) UnmarshalJSON(input []byte) error {
    var temp tempWebhook
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "enabled", "for_site", "headers", "id", "modified_time", "name", "oauth2_client_id", "oauth2_client_secret", "oauth2_grant_type", "oauth2_password", "oauth2_scopes", "oauth2_token_url", "oauth2_username", "org_id", "secret", "site_id", "splunk_token", "topics", "type", "url", "verify_cert")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.CreatedTime = temp.CreatedTime
    w.Enabled = temp.Enabled
    w.ForSite = temp.ForSite
    w.Headers = temp.Headers
    w.Id = temp.Id
    w.ModifiedTime = temp.ModifiedTime
    w.Name = temp.Name
    w.Oauth2ClientId = temp.Oauth2ClientId
    w.Oauth2ClientSecret = temp.Oauth2ClientSecret
    w.Oauth2GrantType = temp.Oauth2GrantType
    w.Oauth2Password = temp.Oauth2Password
    w.Oauth2Scopes = temp.Oauth2Scopes
    w.Oauth2TokenUrl = temp.Oauth2TokenUrl
    w.Oauth2Username = temp.Oauth2Username
    w.OrgId = temp.OrgId
    w.Secret = temp.Secret
    w.SiteId = temp.SiteId
    w.SplunkToken = temp.SplunkToken
    w.Topics = temp.Topics
    w.Type = temp.Type
    w.Url = temp.Url
    w.VerifyCert = temp.VerifyCert
    return nil
}

// tempWebhook is a temporary struct used for validating the fields of Webhook.
type tempWebhook  struct {
    CreatedTime        *float64                    `json:"created_time,omitempty"`
    Enabled            *bool                       `json:"enabled,omitempty"`
    ForSite            *bool                       `json:"for_site,omitempty"`
    Headers            Optional[map[string]string] `json:"headers"`
    Id                 *uuid.UUID                  `json:"id,omitempty"`
    ModifiedTime       *float64                    `json:"modified_time,omitempty"`
    Name               Optional[string]            `json:"name"`
    Oauth2ClientId     *string                     `json:"oauth2_client_id,omitempty"`
    Oauth2ClientSecret *string                     `json:"oauth2_client_secret,omitempty"`
    Oauth2GrantType    *WebhookOauth2GrantTypeEnum `json:"oauth2_grant_type,omitempty"`
    Oauth2Password     *string                     `json:"oauth2_password,omitempty"`
    Oauth2Scopes       []string                    `json:"oauth2_scopes,omitempty"`
    Oauth2TokenUrl     *string                     `json:"oauth2_token_url,omitempty"`
    Oauth2Username     *string                     `json:"oauth2_username,omitempty"`
    OrgId              *uuid.UUID                  `json:"org_id,omitempty"`
    Secret             Optional[string]            `json:"secret"`
    SiteId             *uuid.UUID                  `json:"site_id,omitempty"`
    SplunkToken        Optional[string]            `json:"splunk_token"`
    Topics             []string                    `json:"topics,omitempty"`
    Type               *WebhookTypeEnum            `json:"type,omitempty"`
    Url                *string                     `json:"url,omitempty"`
    VerifyCert         *bool                       `json:"verify_cert,omitempty"`
}
