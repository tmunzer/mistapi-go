
# Wlan Hotspot 20

hostspot 2.0 wlan settings

## Structure

`WlanHotspot20`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DomainName` | `[]string` | Optional | - |
| `Enabled` | `*bool` | Optional | whether to enable hotspot 2.0 config |
| `NaiRealms` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Operators` | [`[]models.WlanHotspot20OperatorsItemEnum`](../../doc/models/wlan-hotspot-20-operators-item-enum.md) | Optional | list of operators to support |
| `Rcoi` | `[]string` | Optional | - |
| `VenueName` | `*string` | Optional | venue name, default is site name |

## Example (as JSON)

```json
{
  "domain_name": [
    "mist.com"
  ],
  "operators": [
    "google",
    "att"
  ],
  "rcoi": [
    "5A03BA0000"
  ],
  "venue_name": "some_name",
  "enabled": false,
  "nai_realms": [
    "nai_realms5"
  ]
}
```

