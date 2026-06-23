
# Org Setting Mist Nac Fingerprinting

Client fingerprinting settings used for Mist NAC policy enforcement

## Structure

`OrgSettingMistNacFingerprinting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | enable/disable writes to NAC DDB fingerprint table<br><br>**Default**: `false` |
| `GenerateCoa` | `*bool` | Optional | enable/disable CoA triggers on fingerprint change for wired clients, always port-bounce<br><br>**Default**: `false` |
| `GenerateWirelessCoa` | `*bool` | Optional | enable/disable CoA triggers on fingerprint change for wireless clients<br><br>**Default**: `false` |
| `WirelessCoaType` | [`*models.OrgSettingMistNacFingerprintingWirelessCoaEnum`](../../doc/models/org-setting-mist-nac-fingerprinting-wireless-coa-enum.md) | Optional | Change of Authorization action sent to wireless clients when fingerprints change. enum: `reauth`, `disconnect` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingMistNacFingerprinting := models.OrgSettingMistNacFingerprinting{
        Enabled:              models.ToPointer(false),
        GenerateCoa:          models.ToPointer(false),
        GenerateWirelessCoa:  models.ToPointer(false),
        WirelessCoaType:      models.ToPointer(models.OrgSettingMistNacFingerprintingWirelessCoaEnum_REAUTH),
    }

}
```

