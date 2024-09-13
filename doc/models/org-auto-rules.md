
# Org Auto Rules

auto_rules in org settings

## Structure

`OrgAutoRules`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreateNewSiteIfNeeded` | `*bool` | Optional | if `src`==`geoip`. By default, a claimed device only gets assigned if the site exists to auto-create the site, enable this<br>**Default**: `false` |
| `Expression` | `models.Optional[string]` | Optional | if `src`==`name`, `src`==`lldp_system_name`,  `src`==`dns_suffix`<br>"[0:3]"            // "abcdef" -> "abc"<br>"split(.)[1]"      // "a.b.c" -> "b"<br>"split(-)[1][0:3]" // "a1234-b5678-c90" -> "b56"' |
| `GatewaytemplateId` | `*string` | Optional | if `src`==`geoip` and `create_new_site_if_needed`==`true`. If a gateway template is desired for this newly created site |
| `MatchCountry` | `*string` | Optional | if `src`==`geoip` |
| `MatchDeviceType` | [`*models.OrgAutoRulesMatchDeviceTypeEnum`](../../doc/models/org-auto-rules-match-device-type-enum.md) | Optional | optional/additional filter. enum: `ap`, `gateway`, `other`, `switch`<br>**Default**: `"ap"` |
| `MatchModel` | `*string` | Optional | optional/additional filter |
| `Model` | `*string` | Optional | if `src`==`model` |
| `Prefix` | `models.Optional[string]` | Optional | if `src`==`name` |
| `Src` | [`models.OrgAutoRulesSrcEnum`](../../doc/models/org-auto-rules-src-enum.md) | Required | enum: `dns_suffix`, `geoip`, `lldp_port_desc`, `lldp_system_name`, `model`, `name`, `subnet` |
| `Subnet` | `*string` | Optional | if `src`==`subnet` |
| `Suffix` | `models.Optional[string]` | Optional | if `src`==`name` |
| `Value` | `*string` | Optional | if<br><br>* `src`==`model`<br>* `src`==`geoip: site name for the device to be assigned to ("city" / "city+country" / ...) |

## Example (as JSON)

```json
{
  "create_new_site_if_needed": false,
  "expression": "split(.)[1]",
  "match_device_type": "ap",
  "prefix": "XX-",
  "src": "lldp_port_desc",
  "suffix": "-YY",
  "gatewaytemplate_id": "gatewaytemplate_id2",
  "match_country": "match_country6"
}
```

