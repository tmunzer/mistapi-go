
# Network Internet Access

Whether this network has direct internet access

## Structure

`NetworkInternetAccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreateSimpleServicePolicy` | `*bool` | Optional | **Default**: `false` |
| `DestinationNat` | [`map[string]models.NetworkInternetAccessDestinationNatProperty`](../../doc/models/network-internet-access-destination-nat-property.md) | Optional | Property key can be an External IP (i.e. "63.16.0.3"), an External IP:Port (i.e. "63.16.0.3:443"), an External Port (i.e. ":443"), an External CIDR (i.e. "63.16.0.0/30"), an External CIDR:Port (i.e. "63.16.0.0/30:443") or a Variable (i.e. "{{myvar}}"). At least one of the `internal_ip` or `port` must be defined |
| `Enabled` | `*bool` | Optional | - |
| `Restricted` | `*bool` | Optional | By default, all access is allowed, to only allow certain traffic, make `restricted`=`true` and define service_policies<br><br>**Default**: `false` |
| `StaticNat` | [`map[string]models.NetworkInternetAccessStaticNatProperty`](../../doc/models/network-internet-access-static-nat-property.md) | Optional | Property key may be an External IP Address (i.e. "63.16.0.3"), a CIDR (i.e. "63.16.0.12/20") or a Variable (i.e. "{{myvar}}") |

## Example (as JSON)

```json
{
  "create_simple_service_policy": false,
  "restricted": false,
  "destination_nat": {
    "key0": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "port": "port4",
      "wan_name": "wan_name0"
    },
    "key1": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "port": "port4",
      "wan_name": "wan_name0"
    }
  },
  "enabled": false,
  "static_nat": {
    "key0": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "wan_name": "wan_name0"
    },
    "key1": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "wan_name": "wan_name0"
    },
    "key2": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "wan_name": "wan_name0"
    }
  }
}
```

