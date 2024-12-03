
# Switch Stp Config

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchStpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BridgePriority` | `*string` | Optional | Switch STP priority: from `0k` to `15k`<br>**Default**: `"8k"` |
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

