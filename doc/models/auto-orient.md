
# Auto Orient

## Structure

`AutoOrient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForceCollection` | `*bool` | Optional | If `force_collection`==`false`, the API attempts to start auto orientation with existing BLE data.<br>If `force_collection`==`true`, the API attempts to start BLE orchestration.<br>**Default**: `false` |
| `Macs` | `[]string` | Optional | list of device macs |

## Example (as JSON)

```json
{
  "force_collection": false,
  "macs": [
    "macs3",
    "macs2",
    "macs1"
  ]
}
```

