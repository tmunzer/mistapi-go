
# Org Auto Rules

auto_rules in org settings

## Structure

`OrgAutoRules`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Expression` | `models.Optional[string]` | Optional | "[0:3]"            // "abcdef" -> "abc"<br>"split(.)[1]"      // "a.b.c" -> "b"<br>"split(-)[1][0:3]" // "a1234-b5678-c90" -> "b56" |
| `MatchDeviceType` | [`*models.OrgAutoRulesMatchDeviceTypeEnum`](../../doc/models/org-auto-rules-match-device-type-enum.md) | Optional | optional/additional filter. enum: `ap`, `gateway`, `other`, `switch`<br>**Default**: `"ap"` |
| `MatchModel` | `*string` | Optional | optional/additional filter |
| `Model` | `*string` | Optional | - |
| `Prefix` | `models.Optional[string]` | Optional | - |
| `Src` | [`models.OrgAutoRulesSrcEnum`](../../doc/models/org-auto-rules-src-enum.md) | Required | enum: `dns_suffix`, `lldp_port_desc`, `lldp_system_name`, `model`, `name`, `subnet` |
| `Subnet` | `*string` | Optional | - |
| `Suffix` | `models.Optional[string]` | Optional | - |
| `Value` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "expression": "split(.)[1]",
  "match_device_type": "ap",
  "prefix": "XX-",
  "src": "lldp_system_name",
  "suffix": "-YY",
  "match_model": "match_model8",
  "model": "model6"
}
```

