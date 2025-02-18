
# Nac Tag

*This model accepts additional fields of type interface{}.*

## Structure

`NacTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowUsermacOverride` | `*bool` | Optional | Can be set to true to allow the override by usermac result<br>**Default**: `false` |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `EgressVlanNames` | `[]string` | Optional | If `type`==`egress_vlan_names`, list of egress vlans to return |
| `GbpTag` | `*int` | Optional | If `type`==`gbp_tag` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Match` | [`*models.NacTagMatchEnum`](../../doc/models/nac-tag-match-enum.md) | Optional | if `type`==`match`. enum: `cert_cn`, `cert_issuer`, `cert_san`, `cert_serial`, `cert_sub`, `cert_template`, `client_mac`, `idp_role`, `ingress_vlan`, `mdm_status`, `nas_ip`, `radius_group`, `realm`, `ssid`, `user_name`, `usermac_label`<br>**Constraints**: *Minimum Length*: `1` |
| `MatchAll` | `*bool` | Optional | This field is applicable only when `type`==`match`<br><br>* `false`: means it is sufficient to match any of the values (i.e., match-any behavior)<br>* `true`: means all values should be matched (i.e., match-all behavior)<br><br>Currently it makes sense to set this field to `true` only if the `match`==`idp_role` or `match`==`usermac_label`<br>**Default**: `false` |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RadiusAttrs` | `[]string` | Optional | If `type`==`radius_attrs`, user can specify a list of one or more standard attributes in the field "radius_attrs".<br>It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.<br>Note that it is allowed to have more than one radius_attrs in the result of a given rule. |
| `RadiusGroup` | `*string` | Optional | If `type`==`radius_group` |
| `RadiusVendorAttrs` | `[]string` | Optional | If `type`==`radius_vendor_attrs`, user can specify a list of one or more vendor-specific attributes in the field "radius_vendor_attrs".<br>It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.<br>Note that it is allowed to have more than one radius_vendor_attrs in the result of a given rule. |
| `SessionTimeout` | `*int` | Optional | If `type`==`session_timeout, in seconds |
| `Type` | [`models.NacTagTypeEnum`](../../doc/models/nac-tag-type-enum.md) | Required | enum: `egress_vlan_names`, `gbp_tag`, `match`, `radius_attrs`, `radius_group`, `radius_vendor_attrs`, `session_timeout`, `username_attr`, `vlan`<br>**Constraints**: *Minimum Length*: `1` |
| `UsernameAttr` | [`*models.NacTagUsernameAttrEnum`](../../doc/models/nac-tag-username-attr-enum.md) | Optional | enum: `automatic`, `cn`, `dns`, `email`, `upn` |
| `Values` | `[]string` | Optional | If `type`==`match` |
| `Vlan` | `*string` | Optional | If `type`==`vlan` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "allow_usermac_override": false,
  "egress_vlan_names": [
    "1vlan-30",
    "1vlan-20",
    "2-vlan10"
  ],
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "match_all": false,
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "radius_attrs": [
    "Idle-Timeout=600",
    "Termination-Action=RADIUS-Request"
  ],
  "radius_vendor_attrs": [
    "PaloAlto-Admin-Role=superuser",
    "PaloAlto-Panorama-Admin-Role=administrator"
  ],
  "session_timeout": 86000,
  "type": "vlan",
  "created_time": 105.64,
  "gbp_tag": 162,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

