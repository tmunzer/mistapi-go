
# Gateway Search

## Structure

`GatewaySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExtIp` | `*string` | Optional | - |
| `Hostname` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Ip` | `*string` | Optional | - |
| `LastHostname` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `NumMembers` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ext_ip": "ext_ip2",
  "hostname": [
    "hostname2",
    "hostname1",
    "hostname0"
  ],
  "ip": "ip8",
  "last_hostname": "last_hostname0",
  "mac": "mac8"
}
```

