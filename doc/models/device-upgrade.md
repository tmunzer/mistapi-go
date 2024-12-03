
# Device Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed (Available on Junos OS devices)<br>**Default**: `false` |
| `RebootAt` | `*int` | Optional | reboot start time in epoch |
| `Snapshot` | `*bool` | Optional | Perform recovery snapshot after device is rebooted (Available on Junos OS devices)<br>**Default**: `false` |
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

