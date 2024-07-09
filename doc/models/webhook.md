
# Webhook

## Structure

`Webhook`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `Enabled` | `*bool` | Optional | whether webhook is enabled<br>**Default**: `true` |
| `ForSite` | `*bool` | Optional | - |
| `Headers` | `models.Optional[map[string]string]` | Optional | if `type`=`http-post`, additional custom HTTP headers to add<br>the headers name and value must be string, total bytes of headers name and value must be less than 1000 |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `models.Optional[string]` | Optional | name of the webhook |
| `Oauth2ClientId` | `*string` | Optional | required when `oauth2_grant_type`==`client_credentials` |
| `Oauth2ClientSecret` | `*string` | Optional | required when `oauth2_grant_type`==`client_credentials` |
| `Oauth2GrantType` | [`*models.WebhookOauth2GrantTypeEnum`](../../doc/models/webhook-oauth-2-grant-type-enum.md) | Optional | required when `type`==`oauth2` |
| `Oauth2Password` | `*string` | Optional | required when `oauth2_grant_type`==`password` |
| `Oauth2Scopes` | `[]string` | Optional | required when `type`==`oauth2`, if provided, will be used in the token request |
| `Oauth2TokenUrl` | `*string` | Optional | required when `type`==`oauth2` |
| `Oauth2Username` | `*string` | Optional | required when `oauth2_grant_type`==`password` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Secret` | `models.Optional[string]` | Optional | only if `type`=`http-post`<br>when `secret` is provided, two HTTP headers will be added:<br><br>* X-Mist-Signature-v2: HMAC_SHA256(secret, body)<br>* X-Mist-Signature: HMAC_SHA1(secret, body) |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SplunkToken` | `models.Optional[string]` | Optional | required if `type`=`splunk`<br>If splunk_token is not defined for a type Splunk webhook, it will not send, regardless if the webhook receiver is configured to accept it.' |
| `Topics` | [`[]models.WebhookTopicEnum`](../../doc/models/webhook-topic-enum.md) | Optional | N.B. For org webhooks, only device_events/alarms/audits/client-join/client-sessions/nac-sessions/nac_events topics are supported. |
| `Type` | [`*models.WebhookTypeEnum`](../../doc/models/webhook-type-enum.md) | Optional | **Default**: `"http-post"` |
| `Url` | `models.Optional[string]` | Optional | - |
| `VerifyCert` | `*bool` | Optional | when url uses HTTPS, whether to verify the certificate<br>**Default**: `true` |

## Example (as JSON)

```json
{
  "enabled": true,
  "headers": {
    "x-custom-1": "your_custom_header_value1",
    "x-custom-2": "your_custom_header_value2"
  },
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "http-post",
  "verify_cert": true,
  "created_time": 4.16,
  "for_site": false,
  "id": "0000088a-0000-0000-0000-000000000000"
}
```
