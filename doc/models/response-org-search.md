
# Response Org Search

## Structure

`ResponseOrgSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseOrgSearchItem`](../../doc/models/response-org-search-item.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 15.98,
  "limit": 108,
  "results": [
    {
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "msp_id": "000021b6-0000-0000-0000-000000000000",
      "name": "name6",
      "num_aps": 140,
      "num_gateways": 112,
      "num_sites": 50
    }
  ],
  "start": 228.04,
  "total": 202,
  "next": "next2"
}
```

