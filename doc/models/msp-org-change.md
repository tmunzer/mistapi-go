
# Msp Org Change

*This model accepts additional fields of type interface{}.*

## Structure

`MspOrgChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Op` | [`models.MspOrgChangeOperationEnum`](../../doc/models/msp-org-change-operation-enum.md) | Required | enum: `assign`, `unassign` |
| `OrgIds` | `[]string` | Required | list of org_id |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "op": "assign",
  "org_ids": [
    "org_ids6",
    "org_ids5"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

