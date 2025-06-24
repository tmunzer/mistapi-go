
# Upgrade Bios

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeBios`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed<br><br>**Default**: `false` |
| `Version` | `*string` | Optional | Specific bios version |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "reboot": false,
  "version": "CDEN_P_EX1_00.20.01.00",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

