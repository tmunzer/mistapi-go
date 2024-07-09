
# Upgrade Fpga

## Structure

`UpgradeFpga`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed<br>**Default**: `false` |
| `Version` | `*string` | Optional | specific fpga version |

## Example (as JSON)

```json
{
  "reboot": false,
  "version": "REV37"
}
```

