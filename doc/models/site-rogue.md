
# Site Rogue

Rogue site settings

*This model accepts additional fields of type interface{}.*

## Structure

`SiteRogue`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether rogue detection is enabled<br>**Default**: `false` |
| `HoneypotEnabled` | `*bool` | Optional | Whether honeypot detection is enabled<br>**Default**: `false` |
| `MinDuration` | `*int` | Optional | Minimum duration for a bssid to be considered rogue<br>**Default**: `10`<br>**Constraints**: `<= 59` |
| `MinRssi` | `*int` | Optional | Minimum RSSI for an AP to be considered rogue (ignoring APs that’s far away)<br>**Default**: `-80`<br>**Constraints**: `>= -85` |
| `WhitelistedBssids` | `[]string` | Optional | list of BSSIDs to whitelist. Ex: "cc-:8e-:6f-:d4-:bf-:16", "cc-8e-6f-d4-bf-16", "cc-73-*", "cc:82:*" |
| `WhitelistedSsids` | `[]string` | Optional | List of SSIDs to whitelist |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

