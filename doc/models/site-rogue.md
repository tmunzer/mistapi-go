
# Site Rogue

Rogue site settings

## Structure

`SiteRogue`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether or not rogue detection is enabled<br>**Default**: `false` |
| `HoneypotEnabled` | `*bool` | Optional | whether or not honeypot detection is enabled<br>**Default**: `false` |
| `MinDuration` | `*int` | Optional | minimum duration for a bssid to be considered rogue<br>**Default**: `10`<br>**Constraints**: `<= 59` |
| `MinRssi` | `*int` | Optional | minimum RSSI for an AP to be considered rogue (ignoring APs that’s far away)<br>**Default**: `-80`<br>**Constraints**: `>= -85` |
| `WhitelistedBssids` | `[]string` | Optional | list of BSSIDs to whitelist. Ex: "cc-:8e-:6f-:d4-:bf-:16", "cc-8e-6f-d4-bf-16", "cc-73-*", "cc:82:*" |
| `WhitelistedSsids` | `[]string` | Optional | list of SSIDs to whitelist |

## Example (as JSON)

```json
{
  "enabled": false,
  "honeypot_enabled": false,
  "min_duration": 10,
  "min_rssi": -80,
  "whitelisted_bssids": [
    "NeighborSSID"
  ],
  "whitelisted_ssids": [
    "cc:8e:6f:d4:bf:16",
    "cc-8e-6f-d4-bf-16",
    "cc-73-*",
    "cc:82:*"
  ]
}
```

