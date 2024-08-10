
# Other Device Update Multi

## Structure

`OtherDeviceUpdateMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Macs` | `[]string` | Optional | The mac address of the peer device. |
| `Op` | [`models.OtherDeviceUpdateOperationEnum`](../../doc/models/other-device-update-operation-enum.md) | Required | The operation being performed. enum: `assign`, `unassign` |
| `SiteId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "macs": [
    "macs3",
    "macs2"
  ],
  "op": "assign",
  "site_id": "000007ac-0000-0000-0000-000000000000"
}
```

