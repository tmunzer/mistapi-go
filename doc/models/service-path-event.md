
# Service Path Event

*This model accepts additional fields of type interface{}.*

## Structure

`ServicePathEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Policy` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Text` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Version` | `*string` | Optional | - |
| `VpnName` | `*string` | Optional | - |
| `VpnPath` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "90ec7734b374",
  "model": "SSR120",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "policy": "INTERNET",
  "port_id": "ge-1/0/6",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "text": "Peer Path Down",
  "timestamp": 1697037328.651775,
  "type": "GW_SERVICE_PATH_REMOVE",
  "version": "6.1.5-14.lts",
  "vpn_name": "Syracuse_HUB",
  "vpn_path": "Syracuse_HUB-Wan0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

