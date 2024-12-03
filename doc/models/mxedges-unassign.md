
# Mxedges Unassign

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgesUnassign`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MxedgeIds` | `[]uuid.UUID` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mxedge_ids": [
    "000001ac-0000-0000-0000-000000000000",
    "000001ab-0000-0000-0000-000000000000",
    "000001aa-0000-0000-0000-000000000000"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

