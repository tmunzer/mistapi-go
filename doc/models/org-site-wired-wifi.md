
# Org Site Wired Wifi

## Structure

`OrgSiteWiredWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Interval` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Page` | `int` | Required | - |
| `Results` | [`[]models.OrgSiteSleWiredResult`](../../doc/models/org-site-sle-wired-result.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 152.44,
  "interval": 50,
  "limit": 226,
  "page": 112,
  "results": [
    {
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "switch-health": 218.74,
      "num_clients": 68.96,
      "num_switches": 67.86,
      "switch-bandwidth": 71.16,
      "switch-throughput": 157.64
    }
  ],
  "start": 108.5,
  "total": 132
}
```

