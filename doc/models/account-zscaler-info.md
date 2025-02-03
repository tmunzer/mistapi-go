
# Account Zscaler Info

OAuth linked Zscaler apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountZscalerInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `*string` | Optional | - |
| `PartnerKey` | `*string` | Optional | - |
| `Username` | `*string` | Optional | Customer account user name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cloud_name": "zscalerbeta.net",
  "partner_key": "K35vrZcK3JvrZc",
  "username": "john@nmo.com",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

