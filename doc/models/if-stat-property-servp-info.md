
# If Stat Property Servp Info

*This model accepts additional fields of type interface{}.*

## Structure

`IfStatPropertyServpInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Asn` | `*string` | Optional | - |
| `City` | `*string` | Optional | - |
| `CountryCode` | `*string` | Optional | - |
| `Latitude` | `*float64` | Optional | - |
| `Longitude` | `*float64` | Optional | - |
| `Org` | `*string` | Optional | - |
| `RegionCode` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "asn": "asn0",
  "city": "city4",
  "country_code": "country_code4",
  "latitude": 48.18,
  "longitude": 45.98,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

