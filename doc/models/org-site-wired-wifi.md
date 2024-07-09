
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
| `Results` | [`[]models.OrgSiteSleWiredResult`](../../doc/models/org-site-sle-wired-result.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 201.46,
  "interval": 88,
  "limit": 248,
  "page": 106,
  "results": [
    {
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "num_clients": 68.96,
      "num_switches": 67.86,
      "switch_health": 68.36,
      "switch_stc": 162.54,
      "switch_throughput": 20.2
    }
  ],
  "start": 157.52,
  "total": 86
}
```

