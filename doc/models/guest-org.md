
# Guest Org

Guest

*This model accepts additional fields of type interface{}.*

## Structure

`GuestOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessCodeEmail` | `*string` | Optional | if `auth_method`==`email`, the email address where the authorization code has been sent to |
| `AllowWlanIdRoam` | `*bool` | Optional | based on the WLAN portal configuration (field `allow_wlan_id_roam`), if the user is also authorized on other Guest WLANs of the same Org without reauthentication |
| `ApMac` | `*string` | Optional | the MAC Address of the AP the guest was connected to during the registration process |
| `AuthMethod` | `*string` | Optional | type of guest authorization |
| `Authorized` | `*bool` | Optional | whether the guest is current authorized<br>**Default**: `true` |
| `AuthorizedExpiringTime` | `*float64` | Optional | when the authorization would expire |
| `AuthorizedTime` | `*float64` | Optional | when the guest was authorized |
| `Company` | `*string` | Optional | optional, the info provided by user |
| `CrossSite` | `*bool` | Optional | based on the WLAN portal configuration (field `cross_site`), if the user is also authorized on other sites (same `wlan.ssid`) of the same Org without reauthentication |
| `Email` | `*string` | Optional | optional, the info provided by user |
| `Field1` | `*string` | Optional | optional, the info provided by user |
| `Field2` | `*string` | Optional | - |
| `Field3` | `*string` | Optional | - |
| `Field4` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | mac |
| `Minutes` | `*int` | Optional | Authorization duration, in minutes. Default is 1440 minutes (1 day), maximum is 259200 (180 days)<br>**Default**: `1440`<br>**Constraints**: `>= 0`, `<= 259200` |
| `Name` | `*string` | Optional | optional, the info provided by user |
| `RandomMac` | `*bool` | Optional | if the client is using a randomized MAC Address to connect the SSID |
| `Ssid` | `*string` | Optional | name of the SSID |
| `WlanId` | `uuid.UUID` | Required | ID of the WLAN |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "authorized": true,
  "authorized_expiring_time": 1480704955,
  "authorized_time": 1480704355,
  "company": "abc",
  "email": "john@abc.com",
  "minutes": 1440,
  "name": "John Smith",
  "ssid": "Guest-SSID",
  "wlan_id": "6748cfa6-4e12-11e6-9188-0242ac110007",
  "access_code_email": "access_code_email8",
  "allow_wlan_id_roam": false,
  "ap_mac": "ap_mac8",
  "auth_method": "auth_method0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

