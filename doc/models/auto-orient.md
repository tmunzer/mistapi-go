
# Auto Orient

*This model accepts additional fields of type interface{}.*

## Structure

`AutoOrient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForceCollection` | `*bool` | Optional | If `force_collection`==`false`, the API attempts to start auto orientation with existing BLE data.<br>If `force_collection`==`true`, the API attempts to start BLE orchestration.<br>**Default**: `false` |
| `Macs` | `[]string` | Optional | list of device macs |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "force_collection": false,
  "macs": [
    "macs3",
    "macs4",
    "macs5"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

