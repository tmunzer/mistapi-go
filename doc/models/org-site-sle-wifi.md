
# Org Site Sle Wifi

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 108.82,
  "interval": 40,
  "limit": 216,
  "page": 154,
  "results": [
    {
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "ap-availability": 204.4,
      "ap-health": 236.64,
      "capacity": 106.72,
      "coverage": 128.26,
      "num_aps": 247.16,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 64.88,
  "total": 122,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

