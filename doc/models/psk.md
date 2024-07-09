
# Psk

PSK

## Structure

`Psk`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminSsoId` | `*string` | Optional | sso id for psk created from psk portal |
| `CreatedTime` | `*float64` | Optional | - |
| `Email` | `*string` | Optional | email to send psk expiring notifications to |
| `ExpireTime` | `models.Optional[int]` | Optional | Expire time for this PSK key (epoch time in seconds). Default `null` (as no expiration) |
| `ExpiryNotificationTime` | `*int` | Optional | Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire |
| `Id` | `*uuid.UUID` | Optional | - |
| `Mac` | `*string` | Optional | if `usage`==`single`, the mac that this PSK ties to, empty if `auto-binding` |
| `Macs` | `[]string` | Optional | if `usage`==`macs`, this list contains N number of client mac addresses or mac patterns(11:22:*) or both. This list is capped at 5000 |
| `MaxUsage` | `*int` | Optional | For Org PSK Only. Max concurrent users for this PSK key. Default is 0 (unlimited)<br>**Default**: `0` |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | - |
| `Note` | `*string` | Optional | - |
| `NotifyExpiry` | `*bool` | Optional | If set to true, reminder notification will be sent when psk is about to expire<br>**Default**: `false` |
| `NotifyOnCreateOrEdit` | `*bool` | Optional | If set to true, notification will be sent when psk is created or edited |
| `OldPassphrase` | `*string` | Optional | previous passphrase of the PSK if it has been rotated |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Passphrase` | `string` | Required | passphrase of the PSK (8-63 character or 64 in hex)<br>**Constraints**: *Minimum Length*: `8`, *Maximum Length*: `64` |
| `Role` | `*string` | Optional | **Constraints**: *Minimum Length*: `0`, *Maximum Length*: `32` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `string` | Required | SSID this PSK should be applicable to |
| `Usage` | [`*models.PskUsageEnum`](../../doc/models/psk-usage-enum.md) | Optional | **Default**: `"multi"` |
| `VlanId` | `*int` | Optional | VLAN for this PSK key |

## Example (as JSON)

```json
{
  "expire_time": 1614990263,
  "macs": [
    "11:22:33:44:55:66",
    "aa:bb:",
    "53"
  ],
  "max_usage": 0,
  "name": "name2",
  "notify_expiry": false,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "passphrase": "passphrase2",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "ssid0",
  "usage": "multi",
  "admin_sso_id": "admin_sso_id8",
  "created_time": 144.92,
  "email": "email4",
  "expiry_notification_time": 18
}
```
