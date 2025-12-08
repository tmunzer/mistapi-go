
# Response Org System Events Search

## Structure

`ResponseOrgSystemEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.OrgSystemEvent`](../../doc/models/org-system-event.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 104,
  "limit": 190,
  "results": [
    {
      "change_cat": "admin_action",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "scope": "org",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "type": "delete-wlan",
      "metadata": "metadata0"
    }
  ],
  "start": 62,
  "total": 96,
  "next": "next0"
}
```

