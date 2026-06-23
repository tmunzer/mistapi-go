
# Psk

Personal pre-shared key configuration for WLAN access

*This model accepts additional fields of type interface{}.*

## Structure

`Psk`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminSsoId` | `*string` | Optional, Read-only | sso id for psk created from psk portal |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Email` | `*string` | Optional | Notification recipient email address for PSK creation notification and expiration reminders |
| `ExpireTime` | `models.Optional[int]` | Optional | Expire time for this PSK key (epoch time in seconds). Default `null` (as no expiration) |
| `ExpiryNotificationTime` | `*int` | Optional | Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Mac` | `*string` | Optional | If `usage`==`single`, client MAC address this PSK is bound to; empty when auto-binding is used |
| `Macs` | `[]string` | Optional | If `usage`==`macs`, this list contains client MAC addresses, MAC patterns such as `1122*`, or both. This list is capped at 5000 entries |
| `MaxUsage` | `*int` | Optional | For Org PSK Only. Max concurrent users for this PSK key. Default is 0 (unlimited)<br><br>**Default**: `0` |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the PSK |
| `Note` | `*string` | Optional | Admin note or description stored with the PSK |
| `NotifyExpiry` | `*bool` | Optional | If set to true, reminder notification will be sent when psk is about to expire<br><br>**Default**: `false` |
| `NotifyOnCreateOrEdit` | `*bool` | Optional | If set to true, notification will be sent when psk is created or edited |
| `OldPassphrase` | `*string` | Optional | previous passphrase of the PSK if it has been rotated |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Passphrase` | `string` | Required | PSK passphrase, 8-63 characters or 64 hexadecimal characters<br><br>**Constraints**: *Minimum Length*: `8`, *Maximum Length*: `64` |
| `Role` | `*string` | Optional | Client role applied to users authenticated with this PSK<br><br>**Constraints**: *Minimum Length*: `0`, *Maximum Length*: `32` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Ssid` | `string` | Required | WLAN SSID where this PSK can be used |
| `Usage` | [`*models.PskUsageEnum`](../../doc/models/psk-usage-enum.md) | Optional | enum: `macs`, `multi`, `single`<br><br>**Default**: `"multi"` |
| `VlanId` | [`*models.PskVlanId`](../../doc/models/containers/psk-vlan-id.md) | Optional | VLAN for this PSK key |
| `VlanName` | `*string` | Optional | VLAN name to be assigned. Optional, `vlan_id` takes precedence if both are provided |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    psk := models.Psk{
        AdminSsoId:             models.ToPointer("admin_sso_id4"),
        CreatedTime:            models.ToPointer(float64(36.88)),
        Email:                  models.ToPointer("email8"),
        ExpireTime:             models.NewOptional(models.ToPointer(1614990263)),
        ExpiryNotificationTime: models.ToPointer(222),
        Id:                     models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Macs:                   []string{
            "112233abcedf",
            "aabbcc*",
        },
        MaxUsage:               models.ToPointer(0),
        Name:                   "name8",
        NotifyExpiry:           models.ToPointer(false),
        OrgId:                  models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Passphrase:             "passphrase8",
        SiteId:                 models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Ssid:                   "ssid6",
        Usage:                  models.ToPointer(models.PskUsageEnum_MULTI),
        AdditionalProperties:   map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

