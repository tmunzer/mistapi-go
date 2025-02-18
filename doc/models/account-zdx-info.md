
# Account Zdx Info

*This model accepts additional fields of type interface{}.*

## Structure

`AccountZdxInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountId` | `string` | Required | Generated account id (uuid) |
| `CloudName` | `string` | Required | ZDX cloud name |
| `KeyId` | `string` | Required | Customer account API key ID |
| `WebhookToken` | `string` | Required | Bearer token for the webhook url |
| `WebhookUrl` | `string` | Required | Webhook url to notify Mist about a ZDX alert |
| `ZdxOrgId` | `string` | Required | ZDX organization id |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "account_id": "80364a28-7ccc-4746-b110-ecf3dfd3a359",
  "cloud_name": "zdxcloud.net",
  "key_id": "K35vrZcK3JvrZc",
  "webhook_token": "f92591d7-02f0-4e8c-b2eb-033c7902018",
  "webhook_url": "https://webhook.url/xxx",
  "zdx_org_id": "123456",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

