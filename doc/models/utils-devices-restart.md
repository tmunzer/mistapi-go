
# Utils Devices Restart

## Structure

`UtilsDevicesRestart`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `*string` | Optional | optional for VC member |
| `Node` | `*string` | Optional | only for SSR: if node is not present, both nodes are restarted<br>for other devices: node should not be present |

## Example (as JSON)

```json
{
  "member": "member2",
  "node": "node4"
}
```

