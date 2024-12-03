
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
* radius_class: thie matches the ATTR-Class(25)
* radius_attr: the values are [ \u201C6=1\u201D, \u201C26=10.2.3.4\u201D ], this support other RADIUS attributes where we know the type
* radius_vendor: the values are [ \u201C14179.10=1\u201D, \u201C14178.16=1.2.3.4\u201D ], this matches vendor attributes and will be dynamically evaluated

*This model accepts additional fields of type interface{}.*

## Structure

`WxlanTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `LastIps` | `[]string` | Optional | - |
| `Mac` | `models.Optional[string]` | Optional | if `type`==`client`, Client MAC Address |
| `Match` | [`*models.WxlanTagMatchEnum`](../../doc/models/wxlan-tag-match-enum.md) | Optional | required if `type`==`match`. enum: `ap_id`, `app`, `asset_mac`, `client_mac`, `hostname`, `ip_range_subnet`, `port`, `psk_name`, `psk_role`, `radius_attr`, `radius_class`, `radius_group`, `radius_username`, `sdkclient_uuid`, `wlan_id` |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | The name |
| `Op` | [`*models.WxlanTagOperationEnum`](../../doc/models/wxlan-tag-operation-enum.md) | Optional | required if `type`==`match`, type of tag (inclusive/exclusive). enum: `in`, `not_in`<br>**Default**: `"in"` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `ResourceMac` | `models.Optional[string]` | Optional | - |
| `Services` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Specs` | [`[]models.WxlanTagSpec`](../../doc/models/wxlan-tag-spec.md) | Optional | if `type`==`spec` |
| `Subnet` | `*string` | Optional | - |
| `Type` | [`models.WxlanTagTypeEnum`](../../doc/models/wxlan-tag-type-enum.md) | Required | enum: `client`, `match`, `resource`, `spec`, `subnet`, `vlan` |
| `Values` | `[]string` | Optional | required if `type`==`match` and<br><br>* `match`==`ap_id`: list of AP IDs<br>* `match`==`app`: list of Application Names<br>* `match`==`asset_mac`: list of Asset MAC Addresses<br>* `match`==`client_mac`: list of Client MAC Addresses<br>* `match`==`hostname`: list of Resources Hostnames<br>* `match`==`ip_range_subnet`: list of IP Addresses and/or CIDRs<br>* `match`==`psk_name`: list of PSK Names<br>* `match`==`psk_role`: list of PSK Roles<br>* `match`==`port`: list of Ports or Port Ranges<br>* `match`==`radius_attr`: list of RADIUS Attributes. The values are [ “6=1”, “26=10.2.3.4” ], this support other RADIUS attributes where we know the type<br>* `match`==`radius_class`: list of RADIUS Classes. This matches the ATTR-Class(25)<br>* `match`==`radius_group`: list of RADIUS Groups. This is a smart tag that matches RADIUS-Filter-ID, Airespace-ACL-Name (VendorID=14179, VendorType=6) / Aruba-User-Role (VendorID=14823, VendorType=1)<br>* `match`==`radius_username`: list of RADIUS Usernames. This matches the ATTR-User-Name(1)<br>* `match`==`sdkclient_uuid`: list of SDK UUIDs<br>* `match`==`wlan_id`: list of WLAN IDs<br><br>**Notes**:<br>Variables are not allowed |
| `VlanId` | [`*models.WxlanTagVlanId`](../../doc/models/containers/wxlan-tag-vlan-id.md) | Optional | if `type`==`vlan_id`, VLAN ID or variable |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name6",
  "op": "in",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "resource",
  "vlan_id": 1055,
  "created_time": 251.96,
  "for_site": false,
  "last_ips": [
    "last_ips0"
  ],
  "mac": "mac0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

