
# Response Org Inventory Change

## Structure

`ResponseOrgInventoryChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Op` | [`models.ResponseOrgInventoryChangeOpEnum`](../../doc/models/response-org-inventory-change-op-enum.md) | Required | - |
| `Reason` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Success` | `[]string` | Required | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "error": [
    "error3"
  ],
  "op": "unassign",
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

