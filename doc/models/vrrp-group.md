
# Vrrp Group

Junos VRRP group

## Structure

`VrrpGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKey` | `*string` | Optional | if `auth_type`==`md5` |
| `AuthPassword` | `*string` | Optional | if `auth_type`==`simple` |
| `AuthType` | [`*models.VrrpGroupAuthTypeEnum`](../../doc/models/vrrp-group-auth-type-enum.md) | Optional | **Default**: `"md5"` |
| `Networks` | [`map[string]models.VrrpGroupNetwork`](../../doc/models/vrrp-group-network.md) | Optional | Property key is the network name |

## Example (as JSON)

```json
{
  "auth_key": "auth-key-1",
  "auth_type": "md5",
  "networks": {
    "data": {
      "ip": "10.182.96.1"
    },
    "mgmt": {
      "ip": "10.182.104.1"
    },
    "v10": {
      "ip": "10.182.104.129"
    },
    "wap": {
      "ip": "10.182.102.1"
    }
  },
  "auth_password": "auth_password0"
}
```

