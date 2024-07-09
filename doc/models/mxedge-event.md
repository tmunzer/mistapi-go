
# Mxedge Event

## Structure

`MxedgeEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MxclusterId` | `*string` | Optional | - |
| `MxedgeId` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Service` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "mxcluster_id": "2815c917-58e7-472f-a190-bfd44fb58d05",
  "mxedge_id": "00000000-0000-0000-1000-020000dc585c",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "service": "tunterm",
  "timestamp": 1694678225.927,
  "type": "ME_SERVICE_STOPPED"
}
```

