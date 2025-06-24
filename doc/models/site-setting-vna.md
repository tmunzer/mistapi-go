
# Site Setting Vna

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingVna`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Enable Virtual Network Assistant (using SUB-VNA license). This applied to AP / Switch / Gateway<br><br>**Default**: `false` |
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

