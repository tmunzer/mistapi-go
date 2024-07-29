
# Protect Re

restrict inbound-traffic to host
when enabled, all traffic that is not essential to our operation will be dropped
e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works

## Structure

`ProtectRe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedServices` | `[]string` | Optional | optionally, services we'll allow |
| `Custom` | [`[]models.ProtectReCustom`](../../doc/models/protect-re-custom.md) | Optional | - |
| `Enabled` | `*bool` | Optional | when enabled, all traffic that is not essential to our operation will be dropped<br>e.g. ntp / dns / traffic to mist will be allowed by default<br>if dhcpd is enabled, we'll make sure it works<br>**Default**: `false` |
| `TrustedHosts` | `[]string` | Optional | host/subnets we'll allow traffic to/from |

## Example (as JSON)

```json
{
  "allowed_services": [
    "icmp",
    "ssh"
  ],
  "enabled": false,
  "custom": [
    {
      "port_range": "port_range6",
      "protocol": "any",
      "subnet": [
        "subnet3",
        "subnet4"
      ]
    },
    {
      "port_range": "port_range6",
      "protocol": "any",
      "subnet": [
        "subnet3",
        "subnet4"
      ]
    }
  ],
  "trusted_hosts": [
    "trusted_hosts0",
    "trusted_hosts1"
  ]
}
```

