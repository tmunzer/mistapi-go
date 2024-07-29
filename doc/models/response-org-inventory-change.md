
# Response Org Inventory Change

## Structure

`ResponseOrgInventoryChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Op` | [`models.ResponseOrgInventoryChangeOpEnum`](../../doc/models/response-org-inventory-change-op-enum.md) | Required | enum: `assign`, `delete`, `downgrade_to_jsi`, `unassign`, `upgrade_to_mist` |
| `Reason` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Success` | `[]string` | Required | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "error": [
    "error3"
  ],
  "op": "delete",
  "reason": [
    "reason4",
    "reason3",
    "reason2"
  ],
  "success": [
    "success2",
    "success3"
  ]
}
```

