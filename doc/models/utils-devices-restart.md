
# Utils Devices Restart

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsDevicesRestart`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `*int` | Optional | Optional for VC member<br>**Constraints**: `>= 0`, `<= 9` |
| `Node` | [`*models.UtilsDevicesRestartNodeEnum`](../../doc/models/utils-devices-restart-node-enum.md) | Optional | only for SRX/SSR: if node is not present, both nodes are restarted. For other devices: node should not be present |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "member": 184,
  "node": "node0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

