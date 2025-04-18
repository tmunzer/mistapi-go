
# Stats Wan Client

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWanClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DhcpExpireTime` | `*float64` | Optional | - |
| `DhcpStartTime` | `*float64` | Optional | - |
| `Hostname` | `[]string` | Optional | - |
| `Ip` | `[]string` | Optional | - |
| `IpSrc` | `*string` | Optional | - |
| `LastHostname` | `*string` | Optional | - |
| `LastIp` | `*string` | Optional | - |
| `Mfg` | `*string` | Optional | - |
| `Network` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Wcid` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ip_src": "dhcp",
  "last_hostname": "sonoszp",
  "last_ip": "192.168.1.139",
  "mfg": "Sonos",
  "network": "lan",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "wcid": "8bbe7389-212b-c65d-2208-00fab2017936",
  "dhcp_expire_time": 131.84,
  "dhcp_start_time": 238.48,
  "hostname": [
    "hostname6"
  ],
  "ip": [
    "ip9"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

