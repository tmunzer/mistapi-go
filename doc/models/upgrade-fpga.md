
# Upgrade Fpga

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeFpga`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed<br>**Default**: `false` |
| `Version` | `*string` | Optional | Specific fpga version |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "reboot": false,
  "version": "REV37",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

