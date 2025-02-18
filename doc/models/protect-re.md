
# Protect Re

Restrict inbound-traffic to host
when enabled, all traffic that is not essential to our operation will be dropped
e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works

*This model accepts additional fields of type interface{}.*

## Structure

`ProtectRe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedServices` | [`[]models.ProtectReAllowedServiceEnum`](../../doc/models/protect-re-allowed-service-enum.md) | Optional | Optionally, services we'll allow |
| `Custom` | [`[]models.ProtectReCustom`](../../doc/models/protect-re-custom.md) | Optional | - |
| `Enabled` | `*bool` | Optional | When enabled, all traffic that is not essential to our operation will be dropped<br>e.g. ntp / dns / traffic to mist will be allowed by default<br>if dhcpd is enabled, we'll make sure it works<br>**Default**: `false` |
| `TrustedHosts` | `[]string` | Optional | host/subnets we'll allow traffic to/from |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
      "subnets": [
        "subnets5",
        "subnets6"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "port_range": "port_range6",
      "protocol": "any",
      "subnets": [
        "subnets5",
        "subnets6"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "port_range": "port_range6",
      "protocol": "any",
      "subnets": [
        "subnets5",
        "subnets6"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "trusted_hosts": [
    "trusted_hosts2"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

