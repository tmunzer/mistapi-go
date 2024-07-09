
# Sle Impacted Users User

## Structure

`SleImpactedUsersUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `ApName` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Degraded` | `float64` | Required | - |
| `DeviceOs` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `DeviceType` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Duration` | `float64` | Required | - |
| `Mac` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Ssid` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | - |
| `WlanId` | `string` | Required | **Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "ap_mac": "ap_mac2",
  "ap_name": "ap_name2",
  "degraded": 38.3,
  "device_os": "device_os0",
  "device_type": "device_type0",
  "duration": 167.36,
  "mac": "mac4",
  "name": "name0",
  "ssid": "ssid2",
  "total": 189.7,
  "wlan_id": "wlan_id2"
}
```

