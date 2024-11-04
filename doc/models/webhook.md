
# Webhook

## Structure

`Webhook`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `Enabled` | `*bool` | Optional | whether webhook is enabled<br>**Default**: `true` |
| `ForSite` | `*bool` | Optional | - |
| `Headers` | `models.Optional[map[string]string]` | Optional | if `type`=`http-post`, additional custom HTTP headers to add<br>the headers name and value must be string, total bytes of headers name and value must be less than 1000 |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `models.Optional[string]` | Optional | name of the webhook |
| `Oauth2ClientId` | `*string` | Optional | required when `oauth2_grant_type`==`client_credentials` |
| `Oauth2ClientSecret` | `*string` | Optional | required when `oauth2_grant_type`==`client_credentials` |
| `Oauth2GrantType` | [`*models.WebhookOauth2GrantTypeEnum`](../../doc/models/webhook-oauth-2-grant-type-enum.md) | Optional | required when `type`==`oauth2`. enum: `client_credentials`, `password` |
| `Oauth2Password` | `*string` | Optional | required when `oauth2_grant_type`==`password` |
| `Oauth2Scopes` | `[]string` | Optional | required when `type`==`oauth2`, if provided, will be used in the token request |
| `Oauth2TokenUrl` | `*string` | Optional | required when `type`==`oauth2` |
| `Oauth2Username` | `*string` | Optional | required when `oauth2_grant_type`==`password` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Secret` | `models.Optional[string]` | Optional | only if `type`=`http-post`<br><br>when `secret` is provided, two  HTTP headers will be added:<br><br>* X-Mist-Signature-v2: HMAC_SHA256(secret, body)<br>* X-Mist-Signature: HMAC_SHA1(secret, body) |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SplunkToken` | `models.Optional[string]` | Optional | required if `type`=`splunk`<br>If splunk_token is not defined for a type Splunk webhook, it will not send, regardless if the webhook receiver is configured to accept it.' |
| `Topics` | `[]string` | Optional | List of supported webhook topics available with the API Call [List Webhook Topics]($e/Constants%20Definitions/listWebhookTopics) |
| `Type` | [`*models.WebhookTypeEnum`](../../doc/models/webhook-type-enum.md) | Optional | enum: `aws-sns`, `google-pubsub`, `http-post`, `oauth2`, `splunk`<br>**Default**: `"http-post"` |
| `Url` | `*string` | Optional | - |
| `VerifyCert` | `*bool` | Optional | when url uses HTTPS, whether to verify the certificate<br>**Default**: `true` |

## Example (as JSON)

```json
{
  "enabled": true,
  "headers": {
    "x-custom-1": "your_custom_header_value1",
    "x-custom-2": "your_custom_header_value2"
  },
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "http-post",
  "verify_cert": true,
  "created_time": 238.32,
  "for_site": false
}
```

