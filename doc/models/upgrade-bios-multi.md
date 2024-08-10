
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
    "0000246b-0000-0000-0000-000000000000",
    "0000246c-0000-0000-0000-000000000000"
  ],
  "models": [
    "models0",
    "models1"
  ]
}
```

