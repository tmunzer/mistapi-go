
# Org Site Wan Wifi

## Structure

`OrgSiteWanWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Interval` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Page` | `int` | Required | - |
| `Results` | [`[]models.OrgSiteSleWanResult`](../../doc/models/org-site-sle-wan-result.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 61.44,
  "interval": 166,
  "limit": 86,
  "page": 28,
  "results": [
    {
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "application_health": 246.28,
      "gateway-health": 241.74,
      "num_clients": 68.96,
      "num_gateways": 213.6,
      "wan-link-health": 134.82
    }
  ],
  "start": 17.5,
  "total": 248
}
```

