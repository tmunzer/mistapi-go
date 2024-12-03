
# Const Country

*This model accepts additional fields of type interface{}.*

## Structure

`ConstCountry`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Alpha2` | `string` | Required | country code, in two-character |
| `Certified` | `bool` | Required | - |
| `Name` | `string` | Required | - |
| `Numeric` | `float64` | Required | country code, ISO 3166-1 numeric |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "alpha2": "FR",
  "certified": true,
  "name": "France",
  "numeric": 250.0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

