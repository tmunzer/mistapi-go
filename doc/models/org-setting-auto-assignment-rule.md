
# Org Setting Auto Assignment Rule

Auto_rules in org settings

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingAutoAssignmentRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreateNewSiteIfNeeded` | `*bool` | Optional | If `src`==`geoip`. By default, a claimed device only gets assigned if the site exists to auto-create the site, enable this<br><br>**Default**: `false` |
| `Expression` | `models.Optional[string]` | Optional | If `src`==`name`, `src`==`lldp_system_name`,  `src`==`dns_suffix`  <br>"[0:3]"            // "abcdef" -> "abc"  <br>"split(.)[1]"      // "a.b.c" -> "b"  <br>"split(-)[1][0:3]" // "a1234-b5678-c90" -> "b56"' |
| `GatewaytemplateId` | `*string` | Optional | If `src`==`geoip` and `create_new_site_if_needed`==`true`. If a gateway template is desired for this newly created site |
| `MatchCountry` | `*string` | Optional | If `src`==`geoip` |
| `MatchDeviceType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `MatchModel` | `*string` | Optional | Optional/additional filter |
| `Model` | `*string` | Optional | If `src`==`model` |
| `Prefix` | `models.Optional[string]` | Optional | If `src`==`name` |
| `Src` | [`models.OrgSettingAutoSiteAssignmentSrcEnum`](../../doc/models/org-setting-auto-site-assignment-src-enum.md) | Required | enum: `ext_ip`, `dns_suffix`, `geoip`, `lldp_port_desc`, `lldp_system_name`, `model`, `name`, `subnet` |
| `Subnet` | `*string` | Optional | If `src`==`subnet` or `ext_ip`==`ext_ip` |
| `Suffix` | `models.Optional[string]` | Optional | If `src`==`name` |
| `Value` | `*string` | Optional | If<br><br>* `src`==`ext_ip`, `src`==`subnet` or `src`==`model`, the site name<br>* `src`==`geoip`: site name for the device to be assigned to (\"city\" / \"city+country\" / ...)" |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "create_new_site_if_needed": false,
  "expression": "split(.)[1]",
  "match_device_type": "ap",
  "prefix": "XX-",
  "src": "ext_ip",
  "suffix": "-YY",
  "gatewaytemplate_id": "gatewaytemplate_id2",
  "match_country": "match_country6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

