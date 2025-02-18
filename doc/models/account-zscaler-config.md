
# Account Zscaler Config

OAuth linked Zscaler apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountZscalerConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `string` | Required | - |
| `PartnerKey` | `string` | Required | - |
| `Password` | `string` | Required | Customer account password |
| `Username` | `string` | Required | Customer account user name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cloud_name": "zscalerbeta.net",
  "partner_key": "K35vrZcK3JvrZc",
  "password": "password",
  "username": "john@nmo.com",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

