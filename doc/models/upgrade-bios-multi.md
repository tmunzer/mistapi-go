
# Upgrade Bios Multi

## Structure

`UpgradeBiosMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Optional | list of device id to upgrade bios |
| `Models` | `[]string` | Optional | list of device model to upgrade bios |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed<br>**Default**: `false` |
| `Version` | `*string` | Optional | specific bios version |

## Example (as JSON)

```json
{
  "reboot": false,
  "version": "CDEN_P_EX1_00.15.01.00",
  "device_ids": [
    "00000ced-0000-0000-0000-000000000000",
    "00000cee-0000-0000-0000-000000000000",
    "00000cef-0000-0000-0000-000000000000"
  ],
  "models": [
    "models6",
    "models7",
    "models8"
  ]
}
```

