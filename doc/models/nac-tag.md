
# Nac Tag

NAC tag used as a rule-matching classifier or as a result attribute for allowed users

*This model accepts additional fields of type interface{}.*

## Structure

`NacTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowUsermacOverride` | `*bool` | Optional | Whether usermac result values can override this NAC tag when the result type is also supported by usermac<br><br>**Default**: `false` |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `EgressVlanNames` | `[]string` | Optional | If `type`==`egress_vlan_names`, list of egress vlans to return |
| `GbpTag` | [`*models.NacTagGbpTag`](../../doc/models/containers/nac-tag-gbp-tag.md) | Optional | If `type`==`gbp_tag`, GBP tag value returned by the NAC rule |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Match` | [`*models.NacTagMatchEnum`](../../doc/models/nac-tag-match-enum.md) | Optional | if `type`==`match`. enum: `cert_cn`, `cert_eku`, `cert_issuer`, `cert_san`, `cert_serial`, `cert_sub`, `cert_template`, `client_mac`, `edr_status`, `gbp_tag`, `hostname`, `idp_role`, `ingress_vlan`, `mdm_status`, `nas_ip`, `radius_group`, `realm`, `ssid`, `user_name`, `usermac_label`<br><br>**Constraints**: *Minimum Length*: `1` |
| `MatchAll` | `*bool` | Optional | This field is applicable only when `type`==`match`<br><br>* `false`: means it is sufficient to match any of the values (i.e., match-any behavior)<br>* `true`: means all values should be matched (i.e., match-all behavior)<br><br>Currently it makes sense to set this field to `true` only if the `match`==`idp_role`, `match`==`usermac_label` and `edr_status`<br><br>**Default**: `false` |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `NacportalId` | `*uuid.UUID` | Optional | If `type`==`redirect_nacportal_id`, the ID of the NAC portal to redirect to |
| `Name` | `string` | Required | Human-readable name of the NAC tag<br><br>**Constraints**: *Minimum Length*: `1` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `RadiusAttrs` | `[]string` | Optional | If `type`==`radius_attrs`, user can specify a list of one or more standard attributes in the field "radius_attrs".<br>It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.<br>Note that it is allowed to have more than one radius_attrs in the result of a given rule. |
| `RadiusGroup` | `*string` | Optional | If `type`==`radius_group`, RADIUS group value returned by the NAC rule |
| `RadiusVendorAttrs` | `[]string` | Optional | If `type`==`radius_vendor_attrs`, user can specify a list of one or more vendor-specific attributes in the field "radius_vendor_attrs".<br>It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected.<br>Note that it is allowed to have more than one radius_vendor_attrs in the result of a given rule. |
| `SessionTimeout` | `*int` | Optional | If `type`==`session_timeout`, session timeout returned by the NAC rule, in seconds |
| `Type` | [`models.NacTagTypeEnum`](../../doc/models/nac-tag-type-enum.md) | Required | enum: `egress_vlan_names`, `gbp_tag`, `match`, `radius_attrs`, `radius_group`, `radius_vendor_attrs`, `redirect_nacportal_id`, `session_timeout`, `username_attr`, `vlan`<br><br>**Constraints**: *Minimum Length*: `1` |
| `UsernameAttr` | [`*models.NacTagUsernameAttrEnum`](../../doc/models/nac-tag-username-attr-enum.md) | Optional | enum: `automatic`, `cn`, `dns`, `email`, `upn` |
| `Values` | `[]string` | Optional | If `type`==`match`, attribute values used by the NAC tag matcher |
| `Vlan` | `*string` | Optional | If `type`==`vlan`, VLAN name or ID returned by the NAC rule |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    nacTag := models.NacTag{
        AllowUsermacOverride: models.ToPointer(false),
        CreatedTime:          models.ToPointer(float64(218.38)),
        EgressVlanNames:      []string{
            "1vlan-30",
            "1vlan-20",
            "2-vlan10",
        },
        GbpTag:               models.ToPointer(models.NacTagGbpTagContainer.FromNumber(94)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MatchAll:             models.ToPointer(false),
        NacportalId:          models.ToPointer(uuid.MustParse("1e970fec-0a7a-4d73-a472-3ef3b6a456aa")),
        Name:                 "name8",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RadiusAttrs:          []string{
            "Idle-Timeout=600",
            "Termination-Action=RADIUS-Request",
        },
        RadiusVendorAttrs:    []string{
            "PaloAlto-Admin-Role=superuser",
            "PaloAlto-Panorama-Admin-Role=administrator",
        },
        SessionTimeout:       models.ToPointer(86000),
        Type:                 models.NacTagTypeEnum_MATCH,
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

