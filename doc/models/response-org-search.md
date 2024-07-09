
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
  "end": 57.76,
  "limit": 26,
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
  "start": 13.82,
  "total": 120,
  "next": "next6"
}
```

