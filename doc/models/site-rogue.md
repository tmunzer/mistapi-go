
# Site Rogue

Rogue site settings

*This model accepts additional fields of type interface{}.*

## Structure

`SiteRogue`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedVlanIds` | `[]int` | Optional | list of VLAN IDs on which rogue APs are ignored<br><br>**Constraints**: `>= 0`, `<= 4096` |
| `Enabled` | `*bool` | Optional | Whether rogue detection is enabled<br><br>**Default**: `false` |
| `HoneypotEnabled` | `*bool` | Optional | Whether honeypot detection is enabled<br><br>**Default**: `false` |
| `MinDuration` | `*int` | Optional | Minimum duration for a bssid to be considered neighbor<br><br>**Default**: `10`<br><br>**Constraints**: `<= 59` |
| `MinRogueDuration` | `*int` | Optional | Minimum duration for a bssid to be considered rogue<br><br>**Default**: `10`<br><br>**Constraints**: `<= 59` |
| `MinRogueRssi` | `*int` | Optional | Minimum RSSI for an AP to be considered rogue<br><br>**Default**: `-80`<br><br>**Constraints**: `>= -85` |
| `MinRssi` | `*int` | Optional | Minimum RSSI for an AP to be considered neighbor (ignoring APs that’s far away)<br><br>**Default**: `-80`<br><br>**Constraints**: `>= -85` |
| `WhitelistedBssids` | `[]string` | Optional | list of BSSIDs to whitelist. Ex: "cc-:8e-:6f-:d4-:bf-:16", "cc-8e-6f-d4-bf-16", "cc-73-*", "cc:82:*" |
| `WhitelistedSsids` | `[]string` | Optional | List of SSIDs to whitelist |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "honeypot_enabled": false,
  "min_duration": 10,
  "min_rogue_duration": 10,
  "min_rogue_rssi": -80,
  "min_rssi": -80,
  "whitelisted_bssids": [
    "NeighborSSID"
  ],
  "whitelisted_ssids": [
    "cc:8e:6f:d4:bf:16",
    "cc-8e-6f-d4-bf-16",
    "cc-73-*",
    "cc:82:*"
  ],
  "allowed_vlan_ids": [
    19
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

