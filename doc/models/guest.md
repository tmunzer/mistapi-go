
# Guest

Guest

## Structure

`Guest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Authorized` | `*bool` | Optional | whether the guest is current authorized |
| `AuthorizedExpiringTime` | `*float64` | Optional | when the authorization would expire |
| `AuthorizedTime` | `*float64` | Optional | when the guest was authorized |
| `Company` | `*string` | Optional | optional, the info provided by user |
| `Email` | `*string` | Optional | optional, the info provided by user |
| `Field1` | `*string` | Optional | optional, the info provided by user |
| `Field2` | `*string` | Optional | - |
| `Field3` | `*string` | Optional | - |
| `Field4` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | mac |
| `Minutes` | `*int` | Optional | minutes, the maximum is 259200 (180 days) |
| `Name` | `*string` | Optional | optional, the info provided by user |
| `Ssid` | `*string` | Optional | - |
| `WlanId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "authorized_expiring_time": 1480704955.0,
  "authorized_time": 1480704355.0,
  "company": "abc",
  "email": "john@abc.com",
  "name": "John Smith",
  "ssid": "Guest-SSID",
  "wlan_id": "6748cfa6-4e12-11e6-9188-0242ac110007",
  "authorized": false
}
```

