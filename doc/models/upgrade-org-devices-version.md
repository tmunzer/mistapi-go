
# Upgrade Org Devices Version

Target firmware version entry for an organization upgrade request

## Structure

`UpgradeOrgDevicesVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirmwareType` | [`*models.UpgradeOrgDevicesVersionFirmwareTypeEnum`](../../doc/models/upgrade-org-devices-version-firmware-type-enum.md) | Optional | Firmware family this version entry applies to. enum: `ap`, `junos` |
| `Force` | `*bool` | Optional | If `firmware_type`==`ap`, set to `true` if upgrade is needed when target version <= running version<br><br>**Default**: `false` |
| `ModelVersion` | `map[string]string` | Optional | If `firmware_type`==`junos`, used to select different versions for different models (Overrides `version` for the specified models). Property key is the hadware model (e.g. `EX4400-24MP`), Property value is the firmware version (e.g. `23.4R1.9`) |
| `Version` | `*string` | Optional | Firmware version to deploy for this entry |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeOrgDevicesVersion := models.UpgradeOrgDevicesVersion{
        FirmwareType:         models.ToPointer(models.UpgradeOrgDevicesVersionFirmwareTypeEnum_AP),
        Force:                models.ToPointer(false),
        ModelVersion:         map[string]string{
            "key0": "model_version3",
            "key1": "model_version2",
        },
        Version:              models.ToPointer("version6"),
    }

}
```

