
# Utils Devices Restart Multi

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsDevicesRestartMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Required | - |
| `Node` | `*string` | Optional | only for SSR: if node is not present, both nodes are restarted<br>for other devices: node should not be present |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "device_ids": [
    "00001261-0000-0000-0000-000000000000",
    "00001262-0000-0000-0000-000000000000"
  ],
  "node": "node6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

