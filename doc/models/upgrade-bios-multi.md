
# Upgrade Bios Multi

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeBiosMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Optional | List of device id to upgrade bios |
| `Models` | `[]string` | Optional | List of device model to upgrade bios |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed<br>**Default**: `false` |
| `Version` | `*string` | Optional | Specific bios version |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

