
# Sle Impact Summary

## Structure

`SleImpactSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | [`[]models.SleImpactSummaryApItem`](../../doc/models/sle-impact-summary-ap-item.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Band` | [`[]models.SleImpactSummaryBandItem`](../../doc/models/sle-impact-summary-band-item.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Classifier` | `string` | Required | - |
| `DeviceOs` | [`[]models.SleImpactSummaryDeviceOsItem`](../../doc/models/sle-impact-summary-device-os-item.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `DeviceType` | [`[]models.SleImpactSummaryDeviceTypeItem`](../../doc/models/sle-impact-summary-device-type-item.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `End` | `float64` | Required | - |
| `Failure` | `string` | Required | - |
| `Metric` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Start` | `float64` | Required | - |
| `Wlan` | [`[]models.SleImpactSummaryWlanItem`](../../doc/models/sle-impact-summary-wlan-item.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "ap": [
    {
      "ap_mac": "ap_mac6",
      "degraded": 149.14,
      "duration": 22.2,
      "name": "name4",
      "total": 78.86
    }
  ],
  "band": [
    {
      "band": "band4",
      "degraded": 250.12,
      "duration": 123.18,
      "name": "name2",
      "total": 22.12
    }
  ],
  "classifier": "classifier2",
  "device_os": [
    {
      "degraded": 94.1,
      "device_os": "device_os0",
      "duration": 223.16,
      "name": "name0",
      "total": 122.1
    }
  ],
  "device_type": [
    {
      "degraded": 131.9,
      "device_type": "device_type0",
      "duration": 4.96,
      "name": "name0",
      "total": 159.9
    }
  ],
  "end": 58.6,
  "failure": "failure4",
  "metric": "metric6",
  "start": 14.66,
  "wlan": [
    {
      "degraded": 169.9,
      "duration": 42.96,
      "name": "name0",
      "total": 197.9,
      "wlan_id": "wlan_id2"
    }
  ]
}
```

