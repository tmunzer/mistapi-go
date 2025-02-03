
# Vrrp Group

Junos VRRP group

*This model accepts additional fields of type interface{}.*

## Structure

`VrrpGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKey` | `*string` | Optional | If `auth_type`==`md5` |
| `AuthPassword` | `*string` | Optional | If `auth_type`==`simple` |
| `AuthType` | [`*models.VrrpGroupAuthTypeEnum`](../../doc/models/vrrp-group-auth-type-enum.md) | Optional | enum: `md5`, `simple`<br>**Default**: `"md5"` |
| `Networks` | [`map[string]models.VrrpGroupNetwork`](../../doc/models/vrrp-group-network.md) | Optional | Property key is the network name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auth_key": "auth-key-1",
  "auth_type": "md5",
  "networks": {
    "data": {
      "ip": "10.182.96.1",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "mgmt": {
      "ip": "10.182.104.1",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "v10": {
      "ip": "10.182.104.129",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "wap": {
      "ip": "10.182.102.1",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "auth_password": "auth_password2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

