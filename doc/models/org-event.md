
# Org Event

## Structure

`OrgEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Text` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "text": "authentication failed, API key invalid",
  "timestamp": 1685658478.0,
  "type": "CRADLEPOINT_SYNC_FAILED"
}
```

