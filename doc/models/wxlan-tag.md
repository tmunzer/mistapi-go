
# Wxlan Tag

WxLAN Tag

* type:
  * client: created manually (e.g. on wireless client table, when they spot a device of interest, they can create a wxlan tag for it
  * resource: created automatically when we discover a network resource
  * subnet: create automatically when a subnet is discovered
* match:
  * wlan_id, ap_id: values are a list of Wlan / Device ids
  * client_mac: values are a list of MAC addresses
* radius_group: this is a smart tag that matches RADIUS-Filter-ID, Airespace-ACL-Name (VendorID=14179, VendorType=6) / Aruba-User-Role (VendorID=14823, VendorType=1)
* radius_username: this matches the ATTR-User-Name(1)
* radius_class: the matches the ATTR-Class(25)
* radius_attr: the values are [ "6=1" , "26=10.2.3.4" ], this support other RADIUS attributes where we know the type
* radius_vendor: the values are [ "14179.10=1" , "14178.16=1.2.3.4" ], this matches vendor attributes and will be dynamically evaluated

*This model accepts additional fields of type interface{}.*

## Structure

`WxlanTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether this WxLAN tag is scoped to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `LastIps` | `[]string` | Optional, Read-only | Most recent IP addresses observed for a WxLAN tag |
| `Mac` | `models.Optional[string]` | Optional | If `type`==`client`, Client MAC address |
| `Match` | [`*models.WxlanTagMatchEnum`](../../doc/models/wxlan-tag-match-enum.md) | Optional | required if `type`==`match`. enum: `ap_id`, `app`, `asset_mac`, `client_mac`, `hostname`, `ip_range_subnet`, `port`, `psk_name`, `psk_role`, `radius_attr`, `radius_class`, `radius_group`, `radius_username`, `sdkclient_uuid`, `wlan_id` |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the WxLAN tag |
| `Op` | [`*models.WxlanTagOperationEnum`](../../doc/models/wxlan-tag-operation-enum.md) | Optional | required if `type`==`match`, type of tag (inclusive/exclusive). enum: `in`, `not_in`<br><br>**Default**: `"in"` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `ResourceMac` | `models.Optional[string]` | Optional | MAC address of the discovered resource associated with this tag, when applicable |
| `Services` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Specs` | [`[]models.WxlanTagSpec`](../../doc/models/wxlan-tag-spec.md) | Optional | Traffic match specifications used when `type`==`spec` |
| `Subnet` | `*string` | Optional | CIDR subnet associated with this WxLAN tag when `type`==`subnet` |
| `Type` | [`models.WxlanTagTypeEnum`](../../doc/models/wxlan-tag-type-enum.md) | Required | enum: `client`, `match`, `resource`, `spec`, `subnet`, `vlan` |
| `Values` | `[]string` | Optional | Required if `type`==`match` and<br><br>* `match`==`ap_id`: list of AP IDs<br>* `match`==`app`: list of Application Names<br>* `match`==`asset_mac`: list of Asset MAC addresses<br>* `match`==`client_mac`: list of Client MAC addresses<br>* `match`==`hostname`: list of Resources Hostnames<br>* `match`==`ip_range_subnet`: list of IP addresses and/or CIDRs<br>* `match`==`psk_name`: list of PSK Names<br>* `match`==`psk_role`: list of PSK Roles<br>* `match`==`port`: list of Ports or Port Ranges<br>* `match`==`radius_attr`: list of RADIUS Attributes. The values are [ "6=1", "26=10.2.3.4" ], this support other RADIUS attributes where we know the type<br>* `match`==`radius_class`: list of RADIUS Classes. This matches the ATTR-Class(25)<br>* `match`==`radius_group`: list of RADIUS Groups. This is a smart tag that matches RADIUS-Filter-ID, Airespace-ACL-Name (VendorID=14179, VendorType=6) / Aruba-User-Role (VendorID=14823, VendorType=1)<br>* `match`==`radius_username`: list of RADIUS Usernames. This matches the ATTR-User-Name(1)<br>* `match`==`sdkclient_uuid`: list of SDK UUIDs<br>* `match`==`wlan_id`: list of WLAN IDs<br><br>**Notes**:<br>Variables are not allowed |
| `VlanId` | [`*models.WxlanTagVlanId`](../../doc/models/containers/wxlan-tag-vlan-id.md) | Optional | If `type`==`vlan_id`, VLAN ID or variable |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    wxlanTag := models.WxlanTag{
        CreatedTime:          models.ToPointer(float64(178.7)),
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        LastIps:              []string{
            "last_ips6",
        },
        Mac:                  models.NewOptional(models.ToPointer("mac4")),
        Name:                 "name0",
        Op:                   models.ToPointer(models.WxlanTagOperationEnum_IN),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                 models.WxlanTagTypeEnum_RESOURCE,
        VlanId:               models.ToPointer(),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

