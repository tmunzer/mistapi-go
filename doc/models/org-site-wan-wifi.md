
# Org Site Wan Wifi

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 34.8,
  "interval": 62,
  "limit": 18,
  "page": 132,
  "results": [
    {
      "gateway-health": 241.74,
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "application_health": 246.28,
      "num_clients": 68.96,
      "num_gateways": 213.6,
      "wan-link-health": 134.82,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 246.86,
  "total": 112,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

