
# Wlan Auth

authentication wlan settings

## Structure

`WlanAuth`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AnticlogThreshold` | `*int` | Optional | SAE anti-clogging token threshold<br>**Default**: `16`<br>**Constraints**: `>= 16`, `<= 32` |
| `EapReauth` | `*bool` | Optional | whether to trigger EAP reauth when the session ends<br>**Default**: `false` |
| `EnableMacAuth` | `*bool` | Optional | whether to enable MAC Auth, uses the same auth_servers<br>**Default**: `false` |
| `KeyIdx` | `*int` | Optional | when type=wep<br>**Default**: `1`<br>**Constraints**: `>= 1`, `<= 4` |
| `Keys` | `[]string` | Optional | when type=wep, four 10-character or 26-character hex string, null can be used. All keys, if provided, have to be in the same length |
| `MultiPskOnly` | `*bool` | Optional | whether to only use multi_psk<br>**Default**: `false` |
| `Owe` | [`*models.WlanAuthOweEnum`](../../doc/models/wlan-auth-owe-enum.md) | Optional | `enabled` means transition mode<br>**Default**: `"disabled"` |
| `Pairwise` | [`[]models.WlanAuthPairwiseItemEnum`](../../doc/models/wlan-auth-pairwise-item-enum.md) | Optional | when type=psk / eap, one or more of wpa2-ccmp / wpa1-tkip / wpa1-ccmp / wpa2-tkip |
| `PrivateWlan` | `*bool` | Optional | whether private wlan is enabled. only applicable to multi_psk mode<br>**Default**: `false` |
| `Psk` | `models.Optional[string]` | Optional | when type=psk, 8-64 characters, or 64 hex characters<br>**Constraints**: *Minimum Length*: `8`, *Maximum Length*: `64` |
| `Type` | [`models.WlanAuthTypeEnum`](../../doc/models/wlan-auth-type-enum.md) | Required | **Default**: `"open"` |
| `WepAsSecondaryAuth` | `*bool` | Optional | enable WEP as secondary auth<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "anticlog_threshold": 16,
  "eap_reauth": false,
  "enable_mac_auth": false,
  "key_idx": 1,
  "keys": [
    "keys8",
    "keys9"
  ],
  "multi_psk_only": false,
  "owe": "disabled",
  "private_wlan": false,
  "psk": "foryoureyesonly",
  "type": "psk",
  "wep_as_secondary_auth": false
}
```

