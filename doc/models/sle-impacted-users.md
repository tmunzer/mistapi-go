
# Sle Impacted Users

## Structure

`SleImpactedUsers`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `string` | Required | - |
| `Clients` | [`[]models.SleImpactedUsersClient`](../../doc/models/sle-impacted-users-client.md) | Optional | **Constraints**: *Unique Items Required* |
| `End` | `float64` | Required | - |
| `Failure` | `string` | Required | - |
| `Limit` | `float64` | Required | - |
| `Metric` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Page` | `float64` | Required | - |
| `Start` | `float64` | Required | - |
| `TotalCount` | `float64` | Required | - |
| `Users` | [`[]models.SleImpactedUsersUser`](../../doc/models/sle-impacted-users-user.md) | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "classifier": "classifier6",
  "clients": [
    {
      "degraded": 166.58,
      "duration": 39.64,
      "gateways": [
        {
          "chassis_mac": "chassis_mac4",
          "gateway_mac": "gateway_mac2",
          "gateway_name": "gateway_name0",
          "interfaces": [
            "interfaces5",
            "interfaces6"
          ]
        }
      ],
      "mac": "mac2",
      "name": "name8"
    }
  ],
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
      "device_type": "device_type4"
    },
    {
      "ap_mac": "ap_mac8",
      "ap_name": "ap_name4",
      "degraded": 137.76,
      "device_os": "device_os6",
      "device_type": "device_type4"
    },
    {
      "ap_mac": "ap_mac8",
      "ap_name": "ap_name4",
      "degraded": 137.76,
      "device_os": "device_os6",
      "device_type": "device_type4"
    }
  ]
}
```

