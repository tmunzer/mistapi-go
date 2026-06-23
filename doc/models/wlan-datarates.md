
# Wlan Datarates

Data rates wlan settings

## Structure

`WlanDatarates`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Eht` | `models.Optional[string]` | Optional | If `template`==`custom`. EHT MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit) |
| `He` | `models.Optional[string]` | Optional | If `template`==`custom`. HE MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit |
| `Ht` | `models.Optional[string]` | Optional | If `template`==`custom`. MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 00ff 00f0 001f limits HT rates to MCS 0-7 for 1 stream, MCS 4-7 for 2 stream (i.e. MCS 12-15), MCS 1-5 for 3 stream (i.e. MCS 16-20) |
| `Legacy` | [`[]models.WlanDataratesLegacyItemEnum`](../../doc/models/wlan-datarates-legacy-item-enum.md) | Optional | If `template`==`custom`. List of supported rates (IE=1) and extended supported rates (IE=50) for custom template, append ‘b’ at the end to indicate a rate being basic/mandatory. If `template`==`custom` is configured and legacy does not define at least one basic rate, it will use `no-legacy` default values |
| `MinRssi` | `*int` | Optional | Minimum RSSI for client to connect, 0 means not enforcing<br><br>**Default**: `0` |
| `Template` | [`models.Optional[models.WlanDataratesTemplateEnum]`](../../doc/models/wlan-datarates-template-enum.md) | Optional | Data Rates template to apply. enum:<br><br>* `no-legacy`: no 11b<br>* `compatible`: all, like before, default setting that Broadcom/Atheros used<br>* `legacy-only`: disable 802.11n and 802.11ac<br>* `high-density`: no 11b, no low rates<br>* `custom`: user defined<br><br>**Default**: `"compatible"` |
| `Vht` | `models.Optional[string]` | Optional | If `template`==`custom`. MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 03ff 01ff 00ff limits VHT rates to MCS 0-9 for 1 stream, MCS 0-8 for 2 streams, and MCS 0-7 for 3 streams. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanDatarates := models.WlanDatarates{
        Eht:                  models.NewOptional(models.ToPointer("3fff0fff0fff03ff")),
        He:                   models.NewOptional(models.ToPointer("0fff0fff0fff0fff")),
        Ht:                   models.NewOptional(models.ToPointer("00ff00ff00ff")),
        Legacy:               []models.WlanDataratesLegacyItemEnum{
            models.WlanDataratesLegacyItemEnum_ENUM6,
            models.WlanDataratesLegacyItemEnum_ENUM9,
            models.WlanDataratesLegacyItemEnum_ENUM12,
            models.WlanDataratesLegacyItemEnum_ENUM18,
            models.WlanDataratesLegacyItemEnum_ENUM24B,
            models.WlanDataratesLegacyItemEnum_ENUM36,
            models.WlanDataratesLegacyItemEnum_ENUM48,
            models.WlanDataratesLegacyItemEnum_ENUM54,
        },
        MinRssi:              models.ToPointer(-70),
        Template:             models.NewOptional(models.ToPointer(models.WlanDataratesTemplateEnum_COMPATIBLE)),
        Vht:                  models.NewOptional(models.ToPointer("03ff03ff03ff01ff")),
    }

}
```

