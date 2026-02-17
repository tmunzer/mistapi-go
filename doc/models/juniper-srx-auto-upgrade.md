
# Juniper Srx Auto Upgrade

auto_upgrade device first time it is onboarded

## Structure

`JuniperSrxAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomVersions` | `map[string]string` | Optional | Property key is the SRX Hardware model (e.g. "SRX4600") |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Snapshot` | `*bool` | Optional | **Default**: `false` |
| `Version` | `*string` | Optional | Firmware version to deploy (e.g. 23.4R2-S5.5). Optional, used when custom_versions not specified |

## Example (as JSON)

```json
{
  "enabled": false,
  "snapshot": false,
  "version": "23.4R2-S5.5",
  "custom_versions": {
    "key0": "custom_versions7",
    "key1": "custom_versions8",
    "key2": "custom_versions9"
  }
}
```

