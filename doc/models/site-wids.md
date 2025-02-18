
# Site Wids

WIDS site settings

*This model accepts additional fields of type interface{}.*

## Structure

`SiteWids`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RepeatedAuthFailures` | [`*models.SiteWidsRepeatedAuthFailures`](../../doc/models/site-wids-repeated-auth-failures.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "repeated_auth_failures": {
    "duration": 58,
    "threshold": 170,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

