
# Webhook Guest Authorizations Event

## Structure

`WebhookGuestAuthorizationsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | mac address of the AP the guest is connected to |
| `AuthMethod` | `*string` | Optional | authentication method used |
| `AuthorizedExpiringTime` | `*int` | Optional | expiry time for guest |
| `AuthorizedTime` | `*int` | Optional | time of authorization of guest |
| `Carrier` | `*string` | Optional | carrier used when authentication by free cell provider |
| `Client` | `*string` | Optional | client mac |
| `Company` | `*string` | Optional | guest company |
| `Email` | `*string` | Optional | guest email |
| `Field1` | `*string` | Optional | field1 value |
| `Field2` | `*string` | Optional | field2 value |
| `Field3` | `*string` | Optional | field3 value |
| `Field4` | `*string` | Optional | field4 value |
| `Mobile` | `*string` | Optional | guest mobile number |
| `Name` | `*string` | Optional | guest name |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SmsGateway` | `*string` | Optional | sms gateway used via text auth paid service |
| `SponsorEmail` | `*string` | Optional | guest sponsor email |
| `Ssid` | `*string` | Optional | ssid |
| `WlanId` | `*string` | Optional | wlan id |

## Example (as JSON)

```json
{
  "ap": "5c5b350e55c8",
  "auth_method": "passphrase",
  "authorized_expiring_time": 1677076639,
  "authorized_time": 1677076519,
  "carrier": "docomo",
  "client": "ac2316eca70a",
  "company": "MIST",
  "email": "abcd@abcd.com",
  "field1": "field1 value",
  "field2": "field2 value",
  "field3": "field3 value",
  "field4": "field4 value",
  "mobile": "+0123456789",
  "name": "Dr Strange",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "sms_gateway": "Telstra",
  "sponsor_email": "sponsor@gmail.com",
  "ssid": "Portal Auth",
  "wlan_id": "7681be9a-044a-4622-90cf-3accde5ad853"
}
```

