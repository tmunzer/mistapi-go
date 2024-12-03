
# Mxedges Assign

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgesAssign`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MxedgeIds` | `[]uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mxedge_ids": [
    "00001ba8-0000-0000-0000-000000000000",
    "00001ba9-0000-0000-0000-000000000000",
    "00001baa-0000-0000-0000-000000000000"
  ],
  "site_id": "43e9c864-a7e4-4310-8031-d9817d2c5a43",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

