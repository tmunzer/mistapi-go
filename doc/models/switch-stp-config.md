
# Switch Stp Config

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchStpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BridgePriority` | `*string` | Optional | Switch STP priority. Range [0, 4k, 8k.. 60k] in steps of 4k. Bridge priority applies to both VSTP and RSTP.<br>**Default**: `"32k"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bridge_priority": "40k",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

