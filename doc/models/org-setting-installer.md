
# Org Setting Installer

## Structure

`OrgSettingInstaller`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowAllDevices` | `*bool` | Optional | - |
| `AllowAllSites` | `*bool` | Optional | - |
| `ExtraSiteIds` | `[]uuid.UUID` | Optional | - |
| `GracePeriod` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "allow_all_devices": false,
  "allow_all_sites": false,
  "extra_site_ids": [
    "00001c5e-0000-0000-0000-000000000000"
  ],
  "grace_period": 49.68
}
```

