
# Wlan Auth

WLAN client authentication settings

## Structure

`WlanAuth`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AnticlogThreshold` | `*int` | Optional | SAE anti-clogging token threshold<br><br>**Default**: `16`<br><br>**Constraints**: `>= 16`, `<= 32` |
| `EapReauth` | `*bool` | Optional | Whether to trigger EAP reauth when the session ends<br><br>**Default**: `false` |
| `EnableBeaconProtection` | `*bool` | Optional | Enable Beacon Protection; default is false for better compatibility<br><br>**Default**: `false` |
| `EnableGcmp256` | `*bool` | Optional | Enable GCMP-256 encryption suite; default is false for better compatibility<br><br>**Default**: `false` |
| `EnableMacAuth` | `*bool` | Optional | Whether to enable MAC Auth, uses the same auth_servers<br><br>**Default**: `false` |
| `KeyIdx` | `*int` | Optional | When `type`==`wep`, index of the WEP key used as the default transmit key<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1`, `<= 4` |
| `Keys` | `[]string` | Optional | When type=wep, four 10-character or 26-character hex string, null can be used. All keys, if provided, have to be in the same length |
| `MultiPskOnly` | `*bool` | Optional | When `type`==`psk`, whether to only use multi_psk<br><br>**Default**: `false` |
| `Owe` | [`*models.WlanAuthOweEnum`](../../doc/models/wlan-auth-owe-enum.md) | Optional | if `type`==`open`. enum: `disabled`, `enabled` (means transition mode), `required`<br><br>**Default**: `"disabled"` |
| `Pairwise` | [`[]models.WlanAuthPairwiseItemEnum`](../../doc/models/wlan-auth-pairwise-item-enum.md) | Optional | When `type`=`psk` or `type`=`eap`, one or more of `wpa1-ccmp`, `wpa1-tkip`, `wpa2-ccmp`, `wpa2-tkip`, `wpa3` |
| `PrivateWlan` | `*bool` | Optional | When `multi_psk_only`==`true`, whether private wlan is enabled<br><br>**Default**: `false` |
| `Psk` | `models.Optional[string]` | Optional | When `type`==`psk`, 8-64 characters, or 64 hex characters<br><br>**Constraints**: *Minimum Length*: `8`, *Maximum Length*: `64` |
| `Type` | [`models.WlanAuthTypeEnum`](../../doc/models/wlan-auth-type-enum.md) | Required | enum: `eap`, `eap192`, `open`, `psk`, `psk-tkip`, `psk-wpa2-tkip`, `wep`<br><br>**Default**: `"open"` |
| `WepAsSecondaryAuth` | `*bool` | Optional | Enable WEP as secondary auth<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanAuth := models.WlanAuth{
        AnticlogThreshold:      models.ToPointer(16),
        EapReauth:              models.ToPointer(false),
        EnableBeaconProtection: models.ToPointer(false),
        EnableGcmp256:          models.ToPointer(false),
        EnableMacAuth:          models.ToPointer(false),
        KeyIdx:                 models.ToPointer(1),
        Keys:                   nil,
        MultiPskOnly:           models.ToPointer(false),
        Owe:                    models.ToPointer(models.WlanAuthOweEnum_DISABLED),
        PrivateWlan:            models.ToPointer(false),
        Psk:                    models.NewOptional(models.ToPointer("foryoureyesonly")),
        Type:                   models.WlanAuthTypeEnum_PSK,
        WepAsSecondaryAuth:     models.ToPointer(false),
    }

}
```

