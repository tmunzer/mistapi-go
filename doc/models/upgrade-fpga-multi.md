
# Upgrade Fpga Multi

## Structure

`UpgradeFpgaMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Optional | list of device id to upgrade bios |
| `Models` | `[]string` | Optional | list of device model to upgrade bios |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed<br>**Default**: `false` |
| `Version` | `*string` | Optional | specific FPGA version |

## Example (as JSON)

```json
{
  "reboot": false,
  "version": "REV37",
  "device_ids": [
    "0000176b-0000-0000-0000-000000000000",
    "0000176c-0000-0000-0000-000000000000"
  ],
  "models": [
    "models2",
    "models3"
  ]
}
```

