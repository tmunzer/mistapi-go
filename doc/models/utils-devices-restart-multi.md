
# Utils Devices Restart Multi

## Structure

`UtilsDevicesRestartMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Required | - |
| `Node` | `*string` | Optional | only for SSR: if node is not present, both nodes are restarted<br>for other devices: node should not be present |

## Example (as JSON)

```json
{
  "device_ids": [
    "00001261-0000-0000-0000-000000000000",
    "00001262-0000-0000-0000-000000000000"
  ],
  "node": "node6"
}
```

