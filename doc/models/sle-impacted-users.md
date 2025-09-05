
# Sle Impacted Users

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedUsers`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `string` | Required | - |
| `End` | `float64` | Required | - |
| `Failure` | `string` | Required | - |
| `Limit` | `float64` | Required | - |
| `Metric` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Page` | `float64` | Required | - |
| `Start` | `float64` | Required | - |
| `TotalCount` | `float64` | Required | - |
| `Users` | [`[]models.SleImpactedUsersUser`](../../doc/models/sle-impacted-users-user.md) | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "classifier": "classifier6",
  "end": 53.18,
  "failure": "failure4",
  "limit": 196.84,
  "metric": "metric8",
  "page": 69.98,
  "start": 9.24,
  "total_count": 123.36,
  "users": [
    {
      "ap_mac": "ap_mac8",
      "ap_name": "ap_name4",
      "degraded": 137.76,
      "device_os": "device_os6",
      "device_type": "device_type4",
      "duration": 10.82,
      "mac": "mac0",
      "name": "name6",
      "ssid": "ssid4",
      "total": 90.24,
      "wlan_id": "wlan_id8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

