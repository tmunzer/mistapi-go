
# Response Site Search

## Structure

`ResponseSiteSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseSiteSearchItem`](../../doc/models/response-site-search-item.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 238,
  "limit": 188,
  "results": [
    {
      "auto_upgrade_enabled": false,
      "auto_upgrade_version": "auto_upgrade_version4",
      "country_code": "country_code6",
      "honeypot_enabled": false,
      "id": "000023ba-0000-0000-0000-000000000000",
      "name": "name6",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 2.64,
      "timezone": "timezone4",
      "vna_enabled": false,
      "wifi_enabled": false
    }
  ],
  "start": 196,
  "total": 26,
  "next": "next6"
}
```

