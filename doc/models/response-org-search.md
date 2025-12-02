
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
      "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
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

