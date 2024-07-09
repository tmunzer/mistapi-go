
# Upgrade Bios

## Structure

`UpgradeBios`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed<br>**Default**: `false` |
| `Version` | `*string` | Optional | specific bios version |

## Example (as JSON)

```json
{
  "reboot": false,
  "version": "CDEN_P_EX1_00.20.01.00"
}
```

