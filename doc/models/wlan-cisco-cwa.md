
# Wlan Cisco Cwa

Cisco CWA (central web authentication) required RADIUS with COA in order to work. See CWA: https://www.cisco.com/c/en/us/support/docs/security/identity-services-engine/115732-central-web-auth-00.html

*This model accepts additional fields of type interface{}.*

## Structure

`WlanCiscoCwa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedHostnames` | `[]string` | Optional | List of hostnames without http(s):// (matched by substring) |
| `AllowedSubnets` | `[]string` | Optional | List of CIDRs |
| `BlockedSubnets` | `[]string` | Optional | List of blocked CIDRs |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

