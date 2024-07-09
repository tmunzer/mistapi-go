
# Switch Mgmt

## Structure

`SwitchMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigRevert` | `*int` | Optional | **Default**: `10` |
| `ProtectRe` | [`*models.ProtectRe`](../../doc/models/protect-re.md) | Optional | restrict inbound-traffic to host<br>when enabled, all traffic that is not essential to our operation will be dropped<br>e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works |
| `RootPassword` | `*string` | Optional | - |
| `Tacacs` | [`*models.Tacacs`](../../doc/models/tacacs.md) | Optional | - |

## Example (as JSON)

```json
{
  "config_revert": 10,
  "protect_re": {
    "allowed_services": [
      "allowed_services2",
      "allowed_services3",
      "allowed_services4"
    ],
    "custom": [
      {
        "port_range": "port_range6",
        "protocol": "tcp",
        "subnet": [
          "subnet3",
          "subnet4"
        ]
      },
      {
        "port_range": "port_range6",
        "protocol": "tcp",
        "subnet": [
          "subnet3",
          "subnet4"
        ]
      },
      {
        "port_range": "port_range6",
        "protocol": "tcp",
        "subnet": [
          "subnet3",
          "subnet4"
        ]
      }
    ],
    "enabled": false,
    "trusted_hosts": [
      "trusted_hosts2"
    ]
  },
  "root_password": "root_password8",
  "tacacs": {
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
    "default_role": "read",
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
}
```

