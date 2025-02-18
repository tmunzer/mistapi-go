
# Response Org Inventory Change

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseOrgInventoryChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Op` | [`models.ResponseOrgInventoryChangeOpEnum`](../../doc/models/response-org-inventory-change-op-enum.md) | Required | enum: `assign`, `delete`, `downgrade_to_jsi`, `unassign`, `upgrade_to_mist` |
| `Reason` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `Success` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "error": [
    "error3",
    "error4",
    "error5"
  ],
  "op": "delete",
  "reason": [
    "reason4",
    "reason3"
  ],
  "success": [
    "success2"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

