
# Org Setting Auto Device Naming Rule

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingAutoDeviceNamingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Expression` | `*string` | Optional | "[0:3]"            // "abcdef" -> "abc"  <br>"split(.)[1]"      // "a.b.c" -> "b"  <br>"split(-)[1][0:3]" // "a1234-b5678-c90" -> "b56"' |
| `MatchDevice` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br>**Default**: `"ap"` |
| `Prefix` | `*string` | Optional | Prefix to append to the device name |
| `Src` | [`*models.OrgSettingAutoDeviceNamingRuleSrcEnum`](../../doc/models/org-setting-auto-device-naming-rule-src-enum.md) | Optional | enum: `lldp_port_desc`, `mac` |
| `Suffix` | `*string` | Optional | Suffix to append to the device name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "expression": "split(.)[1]",
  "match_device": "ap",
  "prefix": "prefix6",
  "src": "lldp_port_desc",
  "suffix": "suffix2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

