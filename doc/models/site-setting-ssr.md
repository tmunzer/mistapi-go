
# Site Setting Ssr

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingSsr`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConductorHosts` | `[]string` | Optional | - |
| `DisableStats` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "conductor_hosts": [
    "1.1.1.1",
    "2.2.2.2"
  ],
  "disable_stats": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

