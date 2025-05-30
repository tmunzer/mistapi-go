
# Wlan Hotspot 20

Hostspot 2.0 wlan settings

*This model accepts additional fields of type interface{}.*

## Structure

`WlanHotspot20`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DomainName` | `[]string` | Optional | - |
| `Enabled` | `*bool` | Optional | Whether to enable hotspot 2.0 config |
| `NaiRealms` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Operators` | [`[]models.WlanHotspot20OperatorsItemEnum`](../../doc/models/wlan-hotspot-20-operators-item-enum.md) | Optional | List of operators to support |
| `Rcoi` | `[]string` | Optional | - |
| `VenueName` | `*string` | Optional | Venue name, default is site name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
    "nai_realms5",
    "nai_realms4",
    "nai_realms3"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

