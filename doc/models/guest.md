
# Guest

Guest

*This model accepts additional fields of type interface{}.*

## Structure

`Guest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessCodeEmail` | `*string` | Optional | If `auth_method`==`email`, the email address where the authorization code has been sent to |
| `ApMac` | `*string` | Optional | MAC Address of the AP the guest was connected to during the registration process |
| `AuthMethod` | `*string` | Optional | Type of guest authorization |
| `Authorized` | `*bool` | Optional | Whether the guest is current authorized<br>**Default**: `true` |
| `AuthorizedExpiringTime` | `*float64` | Optional | When the authorization would expire |
| `AuthorizedTime` | `*float64` | Optional | When the guest was authorized |
| `Company` | `*string` | Optional | Optional, the info provided by user |
| `Email` | `*string` | Optional | Optional, the info provided by user |
| `Field1` | `*string` | Optional | Optional, the info provided by user |
| `Field2` | `*string` | Optional | - |
| `Field3` | `*string` | Optional | - |
| `Field4` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | MAC Address |
| `Minutes` | `*int` | Optional | Authorization duration, in minutes. Default is 1440 minutes (1 day), maximum is 259200 (180 days)<br>**Default**: `1440`<br>**Constraints**: `>= 0`, `<= 259200` |
| `Name` | `*string` | Optional | Optional, the info provided by user |
| `RandomMac` | `*bool` | Optional | If the client is using a randomized MAC Address to connect the SSID |
| `Ssid` | `*string` | Optional | Name of the SSID |
| `WlanId` | `*uuid.UUID` | Optional | ID of the SSID |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "authorized": true,
  "authorized_expiring_time": 1480704955.0,
  "authorized_time": 1480704355,
  "company": "abc",
  "email": "john@abc.com",
  "minutes": 1440,
  "name": "John Smith",
  "ssid": "Guest-SSID",
  "wlan_id": "6748cfa6-4e12-11e6-9188-0242ac110007",
  "access_code_email": "access_code_email4",
  "ap_mac": "ap_mac4",
  "auth_method": "auth_method6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

