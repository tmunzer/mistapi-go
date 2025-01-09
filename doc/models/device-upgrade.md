
# Device Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reboot` | `*bool` | Optional | For Junos devices only (APs are automatically rebooted). Reboot device immediately after upgrade is completed<br>**Default**: `false` |
| `RebootAt` | `*int` | Optional | For Junos devices only and if `reboot`==`true`. Reboot start time in epoch seconds, default is `start_time` |
| `Snapshot` | `*bool` | Optional | For Junos devices only. Perform recovery snapshot after device is rebooted<br>**Default**: `false` |
| `StartTime` | `*float64` | Optional | firmware download start time in epoch |
| `Version` | `string` | Required | specific version / `stable`, default is to use the latest<br>**Default**: `"stable"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "reboot": false,
  "snapshot": false,
  "version": "stable",
  "reboot_at": 212,
  "start_time": 187.14,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

