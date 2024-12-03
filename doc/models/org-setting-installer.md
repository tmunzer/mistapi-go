
# Org Setting Installer

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingInstaller`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowAllDevices` | `*bool` | Optional | - |
| `AllowAllSites` | `*bool` | Optional | - |
| `ExtraSiteIds` | `[]uuid.UUID` | Optional | - |
| `GracePeriod` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "allow_all_devices": false,
  "allow_all_sites": false,
  "extra_site_ids": [
    "00000448-0000-0000-0000-000000000000",
    "00000449-0000-0000-0000-000000000000",
    "0000044a-0000-0000-0000-000000000000"
  ],
  "grace_period": 126,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

