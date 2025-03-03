
# Upgrade Org Devices Version

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeOrgDevicesVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirmwareType` | [`*models.UpgradeOrgDevicesVersionFirmwareTypeEnum`](../../doc/models/upgrade-org-devices-version-firmware-type-enum.md) | Optional | enum: `ap`, `junos` |
| `Force` | `*bool` | Optional | If `firmware_type`==`ap`, set to `true` if upgrade is needed when target version <= running version<br>**Default**: `false` |
| `ModelVersion` | `map[string]string` | Optional | If `firmware_type`==`junos`, used to select different versions for different models (Overrides `version` for the specified models). Property key is the hadware model (e.g. `EX4400-24MP`), Property value is the firmware version (e.g. `23.4R1.9`) |
| `Version` | `*string` | Optional | version of the firmware to deploy |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "force": false,
  "firmware_type": "ap",
  "model_version": {
    "key0": "model_version3"
  },
  "version": "version2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

