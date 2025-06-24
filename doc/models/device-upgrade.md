
# Device Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reboot` | `*bool` | Optional | For Switches and Gateways only (APs are automatically rebooted). Reboot device immediately after upgrade is completed<br><br>**Default**: `false` |
| `RebootAt` | `*int` | Optional | For Switches and Gateways only and if `reboot`==`true`. Reboot start time in epoch seconds, default is `start_time` |
| `Snapshot` | `*bool` | Optional | For Junos devices only. Perform recovery snapshot after device is rebooted<br><br>**Default**: `false` |
| `StartTime` | `*int` | Optional | Firmware download start time in epoch |
| `Version` | `string` | Required | Specific version / `stable`, default is to use the latest<br><br>**Default**: `"stable"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "reboot": false,
  "snapshot": false,
  "version": "stable",
  "reboot_at": 212,
  "start_time": 26,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

