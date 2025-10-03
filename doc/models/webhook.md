
# Webhook

*This model accepts additional fields of type interface{}.*

## Structure

`Webhook`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetfilterIds` | `[]uuid.UUID` | Optional | Only if `type`==`asset-raw-rssi`. List of ids to associated asset filters. These filters will be applied to messages routed to a filtered-asset-rssi webhook |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Enabled` | `*bool` | Optional | Whether webhook is enabled<br><br>**Default**: `true` |
| `ForSite` | `*bool` | Optional | - |
| `Headers` | `models.Optional[map[string]string]` | Optional | If `type`=`http-post`, additional custom HTTP headers to add. The headers name and value must be string, total bytes of headers name and value must be less than 1000 |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `models.Optional[string]` | Optional | Name of the webhook |
| `Oauth2ClientId` | `*string` | Optional | Required when `oauth2_grant_type`==`client_credentials` |
| `Oauth2ClientSecret` | `*string` | Optional | Required when `oauth2_grant_type`==`client_credentials` |
| `Oauth2GrantType` | [`*models.WebhookOauth2GrantTypeEnum`](../../doc/models/webhook-oauth-2-grant-type-enum.md) | Optional | required when `type`==`oauth2`. enum: `client_credentials`, `password` |
| `Oauth2Password` | `*string` | Optional | Required when `oauth2_grant_type`==`password` |
| `Oauth2Scopes` | `[]string` | Optional | Required when `type`==`oauth2`, if provided, will be used in the token request |
| `Oauth2TokenUrl` | `*string` | Optional | Required when `type`==`oauth2` |
| `Oauth2Username` | `*string` | Optional | Required when `oauth2_grant_type`==`password` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Secret` | `models.Optional[string]` | Optional | Only if `type`=`http-post`<br><br>when `secret` is provided, two  HTTP headers will be added:<br><br>* X-Mist-Signature-v2: HMAC_SHA256(secret, body)<br>* X-Mist-Signature: HMAC_SHA1(secret, body) |
| `SingleEventPerMessage` | `*bool` | Optional | Some solutions may not be able to parse multiple events from a single message (e.g. IBM Qradar, DSM). When set to `true`, only a single event will be sent per message. this feature is only available on certain topics (see [List Webhook Topics](/#operations/listWebhookTopics))<br><br>**Default**: `false` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SplunkToken` | `models.Optional[string]` | Optional | Required if `type`=`splunk`. If splunk_token is not defined for a type Splunk webhook, it will not send, regardless if the webhook receiver is configured to accept it. |
| `Topics` | `[]string` | Optional | List of supported webhook topics available with the API Call [List Webhook Topics](/#operations/listWebhookTopics) |
| `Type` | [`*models.WebhookTypeEnum`](../../doc/models/webhook-type-enum.md) | Optional | enum: `aws-sns`, `google-pubsub`, `http-post`, `oauth2`, `splunk`<br><br>**Default**: `"http-post"` |
| `Url` | `*string` | Optional | - |
| `VerifyCert` | `*bool` | Optional | When url uses HTTPS, whether to verify the certificate<br><br>**Default**: `true` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "single_event_per_message": false,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "http-post",
  "verify_cert": true,
  "assetfilter_ids": [
    "00001203-0000-0000-0000-000000000000",
    "00001204-0000-0000-0000-000000000000"
  ],
  "created_time": 238.32,
  "for_site": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

