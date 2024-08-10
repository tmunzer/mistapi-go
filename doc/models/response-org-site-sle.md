
# Response Org Site Sle

## Structure

`ResponseOrgSiteSle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | - |
| `Interval` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Results` | [`[]models.OrgSiteSleWifiResult1`](../../doc/models/org-site-sle-wifi-result-1.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `*float64` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 78.54,
  "interval": 84,
  "limit": 252,
  "page": 110,
  "results": [
    {
      "ap-availability": 204.4,
      "ap-health": 236.64,
      "capacity": 106.72,
      "coverage": 128.26,
      "num_aps": 247.16,
      "site_id": "00001420-0000-0000-0000-000000000000"
    },
    {
      "ap-availability": 204.4,
      "ap-health": 236.64,
      "capacity": 106.72,
      "coverage": 128.26,
      "num_aps": 247.16,
      "site_id": "00001420-0000-0000-0000-000000000000"
    }
  ]
}
```

