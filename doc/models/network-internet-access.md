
# Network Internet Access

whether this network has direct internet access

*This model accepts additional fields of type interface{}.*

## Structure

`NetworkInternetAccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreateSimpleServicePolicy` | `*bool` | Optional | **Default**: `false` |
| `DestinationNat` | [`map[string]models.NetworkDestinationNatProperty`](../../doc/models/network-destination-nat-property.md) | Optional | Property key may be an IP/Port (i.e. "63.16.0.3:443"), or a port (i.e. ":2222") |
| `Enabled` | `*bool` | Optional | - |
| `Restricted` | `*bool` | Optional | by default, all access is allowed, to only allow certain traffic, make `restricted`=`true` and define service_policies<br>**Default**: `false` |
| `StaticNat` | [`map[string]models.NetworkStaticNatProperty`](../../doc/models/network-static-nat-property.md) | Optional | Property key may be an IP Address (i.e. "172.16.0.1"), and IP Address and Port (i.e. "172.16.0.1:8443") or a CIDR (i.e. "172.16.0.12/20") |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "create_simple_service_policy": false,
  "restricted": false,
  "destination_nat": {
    "key0": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "port": 162,
      "wan_name": "wan_name0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "port": 162,
      "wan_name": "wan_name0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "enabled": false,
  "static_nat": {
    "key0": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "wan_name": "wan_name0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "wan_name": "wan_name0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "wan_name": "wan_name0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

