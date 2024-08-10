
# Event Otherdevice

## Structure

`EventOtherdevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Text` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Vendor` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "5c5b351e13b5",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "text": "Plugged: The Internal 5GB (SIM1) has been inserted into Internal 1.",
  "timestamp": 547235620.89,
  "type": "CELLULAR_EDGE_MODEM_WAN_PLUGGED",
  "vendor": "cradlepoint",
  "device_mac": "device_mac8"
}
```

