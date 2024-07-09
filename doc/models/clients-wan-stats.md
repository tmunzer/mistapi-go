
# Clients Wan Stats

## Structure

`ClientsWanStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `When` | `*string` | Optional | - |
| `Hostname` | `[]string` | Optional | - |
| `Ip` | `[]string` | Optional | - |
| `LastHostname` | `*string` | Optional | - |
| `LastIp` | `*string` | Optional | - |
| `Mfg` | `*string` | Optional | - |
| `Network` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Wcid` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "When": "12/31/2022 23:59:43",
  "last_hostname": "sonoszp",
  "last_ip": "192.168.1.139",
  "mfg": "Sonos",
  "network": "lan",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "wcid": "8bbe7389-212b-c65d-2208-00fab2017936",
  "hostname": [
    "hostname8",
    "hostname7",
    "hostname6"
  ],
  "ip": [
    "ip1",
    "ip2"
  ]
}
```

