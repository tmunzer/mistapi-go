
# Tacacs

## Structure

`Tacacs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.TacacsAcctServer`](../../doc/models/tacacs-acct-server.md) | Optional | - |
| `DefaultRole` | [`*models.TacacsDefaultRoleEnum`](../../doc/models/tacacs-default-role-enum.md) | Optional | enum: `admin`, `helpdesk`, `none`, `read`<br>**Default**: `"none"` |
| `Enabled` | `*bool` | Optional | - |
| `Network` | `*string` | Optional | which network the TACACS server resides |
| `TacplusServers` | [`[]models.TacacsAuthServer`](../../doc/models/tacacs-auth-server.md) | Optional | - |

## Example (as JSON)

```json
{
  "default_role": "none",
  "acct_servers": [
    {
      "host": "host4",
      "port": "port4",
      "secret": "secret0",
      "timeout": 254
    },
    {
      "host": "host4",
      "port": "port4",
      "secret": "secret0",
      "timeout": 254
    }
  ],
  "enabled": false,
  "network": "network6",
  "tacplus_servers": [
    {
      "host": "host6",
      "port": "port2",
      "secret": "secret2",
      "timeout": 18
    }
  ]
}
```

