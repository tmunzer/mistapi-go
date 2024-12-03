
# Site Setting Analytic

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingAnalytic`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | enable Advanced Analytic feature (using SUB-ANA license)<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

