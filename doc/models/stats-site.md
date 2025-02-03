
# Stats Site

Site statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsSite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `string` | Required | - |
| `AlarmtemplateId` | `*uuid.UUID` | Required | - |
| `CountryCode` | `string` | Required | - |
| `CreatedTime` | `float64` | Required | When the object has been created, in epoch |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `Lat` | `float64` | Required | - |
| `Latlng` | [`models.LatLng`](../../doc/models/lat-lng.md) | Required | - |
| `Lng` | `float64` | Required | - |
| `ModifiedTime` | `float64` | Required | When the object has been modified for the last time, in epoch |
| `MspId` | `uuid.UUID` | Required | - |
| `Name` | `string` | Required | - |
| `NetworktemplateId` | `*uuid.UUID` | Required | - |
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
| `RftemplateId` | `*uuid.UUID` | Required | - |
| `SecpolicyId` | `models.Optional[uuid.UUID]` | Optional | - |
| `SitegroupIds` | `[]uuid.UUID` | Required | - |
| `Timezone` | `string` | Required | - |
| `Tzoffset` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "address": "address2",
  "alarmtemplate_id": "00000b0a-0000-0000-0000-000000000000",
  "country_code": "country_code6",
  "created_time": 47.86,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "lat": 183.64,
  "latlng": {
    "lat": 37.295833,
    "lng": -122.032946,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "lng": 239.82,
  "modified_time": 31.1,
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
  "name": "name6",
  "networktemplate_id": "000001e2-0000-0000-0000-000000000000",
  "num_ap": 224,
  "num_ap_connected": 10,
  "num_clients": 210,
  "num_devices": 186,
  "num_devices_connected": 214,
  "num_gateway": 34,
  "num_gateway_connected": 146,
  "num_switch": 206,
  "num_switch_connected": 74,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rftemplate_id": "00000e3c-0000-0000-0000-000000000000",
  "sitegroup_ids": [
    "00001aac-0000-0000-0000-000000000000",
    "00001aad-0000-0000-0000-000000000000"
  ],
  "timezone": "timezone6",
  "tzoffset": 240,
  "secpolicy_id": "0000247e-0000-0000-0000-000000000000",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

