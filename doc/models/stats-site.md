
# Stats Site

Site statistics

## Structure

`StatsSite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `string` | Required | - |
| `AlarmtemplateId` | `uuid.UUID` | Required | - |
| `CountryCode` | `string` | Required | - |
| `CreatedTime` | `float64` | Required | - |
| `Id` | `uuid.UUID` | Required | - |
| `Lat` | `float64` | Required | - |
| `Latlng` | [`models.LatLng`](../../doc/models/lat-lng.md) | Required | - |
| `Lng` | `float64` | Required | - |
| `ModifiedTime` | `float64` | Required | - |
| `MspId` | `string` | Required | - |
| `Name` | `string` | Required | - |
| `NetworktemplateId` | `uuid.UUID` | Required | - |
| `NumAp` | `int` | Required | - |
| `NumApConnected` | `int` | Required | - |
| `NumClients` | `int` | Required | - |
| `NumDevices` | `int` | Required | - |
| `NumDevicesConnected` | `int` | Required | - |
| `NumGateway` | `int` | Required | - |
| `NumGatewayConnected` | `int` | Required | - |
| `NumSwitch` | `int` | Required | - |
| `NumSwitchConnected` | `int` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `RftemplateId` | `uuid.UUID` | Required | - |
| `SecpolicyId` | `*interface{}` | Optional | - |
| `SitegroupIds` | `[]uuid.UUID` | Required | - |
| `Timezone` | `string` | Required | - |
| `Tzoffset` | `int` | Required | - |

## Example (as JSON)

```json
{
  "address": "address4",
  "alarmtemplate_id": "000011b0-0000-0000-0000-000000000000",
  "country_code": "country_code8",
  "created_time": 8.88,
  "id": "00000f12-0000-0000-0000-000000000000",
  "lat": 144.66,
  "latlng": {
    "lat": 37.295833,
    "lng": -122.032946
  },
  "lng": 22.8,
  "modified_time": 70.08,
  "msp_id": "msp_id2",
  "name": "name8",
  "networktemplate_id": "00000888-0000-0000-0000-000000000000",
  "num_ap": 166,
  "num_ap_connected": 208,
  "num_clients": 152,
  "num_devices": 128,
  "num_devices_connected": 16,
  "num_gateway": 232,
  "num_gateway_connected": 52,
  "num_switch": 148,
  "num_switch_connected": 16,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rftemplate_id": "00001f7a-0000-0000-0000-000000000000",
  "sitegroup_ids": [
    "00002152-0000-0000-0000-000000000000",
    "00002153-0000-0000-0000-000000000000"
  ],
  "timezone": "timezone8",
  "tzoffset": 182,
  "secpolicy_id": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

