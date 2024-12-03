
# Wlan Datarates

data rates wlan settings

*This model accepts additional fields of type interface{}.*

## Structure

`WlanDatarates`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ht` | `models.Optional[string]` | Optional | if `template`==`custom`. MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 00ff 00f0 001f limits HT rates to MCS 0-7 for 1 stream, MCS 4-7 for 2 stream (i.e. MCS 12-15), MCS 1-5 for 3 stream (i.e. MCS 16-20) |
| `Legacy` | [`[]models.WlanDataratesLegacyItemEnum`](../../doc/models/wlan-datarates-legacy-item-enum.md) | Optional | if `template`==`custom`. List of supported rates (IE=1) and extended supported rates (IE=50) for custom template, append ‘b’ at the end to indicate a rate being basic/mandatory. If `template`==`custom` is configured and legacy does not define at least one basic rate, it will use `no-legacy` default values |
| `MinRssi` | `*int` | Optional | Minimum RSSI for client to connect, 0 means not enforcing |
| `Template` | [`models.Optional[models.WlanDataratesTemplateEnum]`](../../doc/models/wlan-datarates-template-enum.md) | Optional | Data Rates template to apply. enum:<br><br>* `no-legacy`: no 11b<br>* `compatible`: all, like before, default setting that Broadcom/Atheros used<br>* `legacy-only`: disable 802.11n and 802.11ac<br>* `high-density`: no 11b, no low rates<br>* `custom`: user defined |
| `Vht` | `models.Optional[string]` | Optional | if `template`==`custom`. MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 03ff 01ff 00ff limits VHT rates to MCS 0-9 for 1 stream, MCS 0-8 for 2 streams, and MCS 0-7 for 3 streams. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ht": "00ff00ff00ff",
  "legacy": [
    "6",
    "9",
    "12",
    "18",
    "24b",
    "36",
    "48",
    "54"
  ],
  "min_rssi": -70,
  "vht": "03ff03ff03ff01ff",
  "template": "compatible",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

