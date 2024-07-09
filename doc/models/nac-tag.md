
# Nac Tag

## Structure

`NacTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowUsermacOverride` | `*bool` | Optional | can be set to true to allow the override by usermac result<br>**Default**: `false` |
| `CreatedTime` | `*float64` | Optional | - |
| `EgressVlanNames` | `[]string` | Optional | if `type`==`egress_vlan_names`, list of egress vlans to return |
| `GbpTag` | `*int` | Optional | if `type`==`gbp_tag` |
| `Id` | `*uuid.UUID` | Optional | - |
| `Match` | [`*models.NacTagMatchEnum`](../../doc/models/nac-tag-match-enum.md) | Optional | if `type`==`match`<br>**Constraints**: *Minimum Length*: `1` |
| `MatchAll` | `*bool` | Optional | This field is applicable only when `type`==`match`<br><br>* `false`: means it is sufficient to match any of the values (i.e., match-any behavior)<br>* `true`: means all values should be matched (i.e., match-all behavior)<br><br>Currently it makes sense to set this field to `true` only if the `match`==`idp_role` or `match`==`usermac_label`<br>**Default**: `false` |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RadiusAttrs` | `[]string` | Optional | if `type`==`radius_attrs`, user can specify a list of one or more standard attributes in the field "radius_attrs".<br>It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.<br>Note that it is allowed to have more than one radius_attrs in the result of a given rule. |
| `RadiusGroup` | `*string` | Optional | if `type`==`radius_group` |
| `RadiusVendorAttrs` | `[]string` | Optional | if `type`==`radius_vendor_attrs`, user can specify a list of one or more vendor-specific attributes in the field "radius_vendor_attrs".<br>It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.<br>Note that it is allowed to have more than one radius_vendor_attrs in the result of a given rule. |
| `SessionTimeout` | `*int` | Optional | if `type`==`session_timeout, in seconds |
| `Type` | [`models.NacTagTypeEnum`](../../doc/models/nac-tag-type-enum.md) | Required | **Constraints**: *Minimum Length*: `1` |
| `Values` | `[]string` | Optional | if `type`==`match` |
| `Vlan` | `*string` | Optional | if `type`==`vlan` |

## Example (as JSON)

```json
{
  "allow_usermac_override": false,
  "egress_vlan_names": [
    "1vlan-30",
    "1vlan-20",
    "2-vlan10"
  ],
  "match_all": false,
  "name": "name8",
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
  "type": "radius_attrs",
  "created_time": 67.68,
  "gbp_tag": 206,
  "id": "0000102a-0000-0000-0000-000000000000"
}
```

