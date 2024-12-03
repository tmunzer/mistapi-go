
# Upgrade Fpga Multi

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeFpgaMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Optional | list of device id to upgrade bios |
| `Models` | `[]string` | Optional | list of device model to upgrade bios |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed<br>**Default**: `false` |
| `Version` | `*string` | Optional | specific FPGA version |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "reboot": false,
  "version": "REV37",
  "device_ids": [
    "000020d7-0000-0000-0000-000000000000",
    "000020d8-0000-0000-0000-000000000000"
  ],
  "models": [
    "models4",
    "models5"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

