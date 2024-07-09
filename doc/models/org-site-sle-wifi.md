
# Org Site Sle Wifi

## Structure

`OrgSiteSleWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Interval` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Page` | `int` | Required | - |
| `Results` | [`[]models.OrgSiteSleWifiResult`](../../doc/models/org-site-sle-wifi-result.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 64.22,
  "interval": 188,
  "limit": 108,
  "page": 6,
  "results": [
    {
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "ap-availability": 204.4,
      "ap-health": 236.64,
      "capacity": 106.72,
      "coverage": 128.26,
      "num_aps": 247.16
    }
  ],
  "start": 20.28,
  "total": 14
}
```

