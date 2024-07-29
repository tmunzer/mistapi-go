
# Msp Org Change

## Structure

`MspOrgChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Op` | [`models.MspOrgChangeOperationEnum`](../../doc/models/msp-org-change-operation-enum.md) | Required | enum: `assign`, `unassign` |
| `OrgIds` | `[]string` | Required | list of org_id |

## Example (as JSON)

```json
{
  "op": "assign",
  "org_ids": [
    "org_ids0"
  ]
}
```

