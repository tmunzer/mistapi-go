
# Account Zdx Config

OAuth linked ZDX apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountZdxConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `*string` | Optional | ZDX cloud name. Refer https://help.zscaler.com/zdx/getting-started-zdx-api for ZDX cloud name<br>**Default**: `"zdxcloud.net"` |
| `KeyId` | `string` | Required | Customer account API key ID |
| `KeySecret` | `string` | Required | Customer account API key Secret |
| `ZdxOrgId` | `string` | Required | ZDX organization id |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cloud_name": "zdxcloud.net",
  "key_id": "K35vrZcK3JvrZc",
  "key_secret": "K35vrZcK3JvrZcjjswpp2eii1oo100",
  "zdx_org_id": "123456",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

