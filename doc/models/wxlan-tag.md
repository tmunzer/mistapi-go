
# Wxlan Tag

WxLAN Tag

* type:
  ** client: created manually (e.g. on wireless client table, when they spot a device of interest, they can create a wxlan tag for it
  ** resource: created automatically when we discover a network resource
  ** subnet: create automatically when a subnet is discovered

* match:
  ** wlan_id, ap_id: values are a list of Wlan / Device ids
  ** client_mac: values are a list of MAC addresses

* radius_group: this is a smart tag that matches RADIUS-Filter-ID, Airespace-ACL-Name (VendorID=14179, VendorType=6) / Aruba-User-Role (VendorID=14823, VendorType=1)

* radius_username: this matches the ATTR-User-Name(1)

* radius_class: thie matches the ATTR-Class(25)

* radius_attr: the values are [ “6=1”, “26=10.2.3.4” ], this support other RADIUS attributes where we know the type

* radius_vendor: the values are [ “14179.10=1”, “14178.16=1.2.3.4” ], this matches vendor attributes and will be dynamically evaluated

## Structure

`WxlanTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `LastIps` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Mac` | `models.Optional[string]` | Optional | - |
| `Match` | [`*models.WxlanTagMatchEnum`](../../doc/models/wxlan-tag-match-enum.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | The name |
| `Op` | [`*models.WxlanTagOperationEnum`](../../doc/models/wxlan-tag-operation-enum.md) | Optional | **Default**: `"in"` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `ResourceMac` | `models.Optional[string]` | Optional | - |
| `Services` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Specs` | [`[]models.WxlanTagSpec`](../../doc/models/wxlan-tag-spec.md) | Optional | if `type`==`specs` |
| `Subnet` | `*string` | Optional | - |
| `Type` | [`models.WxlanTagTypeEnum`](../../doc/models/wxlan-tag-type-enum.md) | Required | - |
| `Values` | `[]string` | Optional | if `type`!=`vlan_id` and `type`!=`specs`, list of values to match |
| `VlanId` | `*int` | Optional | if `type`==`vlan_id` |

## Example (as JSON)

```json
{
  "name": "name4",
  "op": "in",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "match",
  "vlan_id": 1055,
  "created_time": 172.84,
  "for_site": false,
  "id": "0000235e-0000-0000-0000-000000000000",
  "last_ips": [
    "last_ips2",
    "last_ips1",
    "last_ips0"
  ],
  "mac": "mac8"
}
```

