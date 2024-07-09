
# Snmpv 3 Config Target Address Item

## Structure

`Snmpv3ConfigTargetAddressItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `*string` | Optional | - |
| `AddressMask` | `*string` | Optional | - |
| `Port` | `*int` | Optional | **Default**: `161` |
| `TagList` | `*string` | Optional | <refer to notify tag, can be multiple with blank |
| `TargetAddressName` | `*string` | Optional | - |
| `TargetParameters` | `*string` | Optional | refer to notify target parameters name |

## Example (as JSON)

```json
{
  "address": "address",
  "address_mask": "address_mask",
  "port": 161,
  "target_address_name": "target_address_name",
  "tag_list": "tag_list2"
}
```
