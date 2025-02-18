
# Tacacs

*This model accepts additional fields of type interface{}.*

## Structure

`Tacacs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.TacacsAcctServer`](../../doc/models/tacacs-acct-server.md) | Optional | - |
| `DefaultRole` | [`*models.TacacsDefaultRoleEnum`](../../doc/models/tacacs-default-role-enum.md) | Optional | enum: `admin`, `helpdesk`, `none`, `read`<br>**Default**: `"none"` |
| `Enabled` | `*bool` | Optional | - |
| `Network` | `*string` | Optional | Which network the TACACS server resides |
| `TacplusServers` | [`[]models.TacacsAuthServer`](../../doc/models/tacacs-auth-server.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "default_role": "none",
  "acct_servers": [
    {
      "host": "host4",
      "port": "port4",
      "secret": "secret0",
      "timeout": 254,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "host": "host4",
      "port": "port4",
      "secret": "secret0",
      "timeout": 254,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "enabled": false,
  "network": "network6",
  "tacplus_servers": [
    {
      "host": "host6",
      "port": "port2",
      "secret": "secret2",
      "timeout": 18,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

