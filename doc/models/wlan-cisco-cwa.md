
# Wlan Cisco Cwa

Cisco CWA (central web authentication) required RADIUS with COA in order to work. See CWA: https://www.cisco.com/c/en/us/support/docs/security/identity-services-engine/115732-central-web-auth-00.html

## Structure

`WlanCiscoCwa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedHostnames` | `[]string` | Optional | list of hostnames without http(s):// (matched by substring) |
| `AllowedSubnets` | `[]string` | Optional | list of CIDRs |
| `BlockedSubnets` | `[]string` | Optional | list of blocked CIDRs |
| `Enabled` | `*bool` | Optional | **Default**: `false` |

## Example (as JSON)

```json
{
  "enabled": false,
  "allowed_hostnames": [
    "allowed_hostnames0"
  ],
  "allowed_subnets": [
    "allowed_subnets6",
    "allowed_subnets7"
  ],
  "blocked_subnets": [
    "blocked_subnets4",
    "blocked_subnets5",
    "blocked_subnets6"
  ]
}
```

