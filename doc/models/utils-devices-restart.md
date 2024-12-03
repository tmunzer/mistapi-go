
# Utils Devices Restart

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsDevicesRestart`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `*string` | Optional | optional for VC member |
| `Node` | `*string` | Optional | only for SSR: if node is not present, both nodes are restarted<br>for other devices: node should not be present |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "member": "member2",
  "node": "node4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

