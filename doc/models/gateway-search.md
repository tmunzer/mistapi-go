
# Gateway Search

*This model accepts additional fields of type interface{}.*

## Structure

`GatewaySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Clustered` | `*bool` | Optional | - |
| `EvpnMissingLinks` | `*bool` | Optional | - |
| `EvpntopoId` | `*string` | Optional | - |
| `ExtIp` | `*string` | Optional | - |
| `Hostname` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Ip` | `*string` | Optional | - |
| `LastConfigStatus` | `*string` | Optional | - |
| `LastHostname` | `*string` | Optional | - |
| `LastTroubleCode` | `*string` | Optional | - |
| `LastTroubleTimestamp` | `*int` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Managed` | `*bool` | Optional | - |
| `Model` | `*string` | Optional | - |
| `Node` | `*string` | Optional | - |
| `Node0Mac` | `*string` | Optional | - |
| `Node1Mac` | `*string` | Optional | - |
| `NumMembers` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Role` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `T128agentVersion` | `*string` | Optional | - |
| `TimeDrifted` | `*bool` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `string` | Required, Constant | Device Type. enum: `gateway`<br>**Value**: `"gateway"` |
| `Uptime` | `*int` | Optional | - |
| `Version` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "gateway",
  "clustered": false,
  "evpn_missing_links": false,
  "evpntopo_id": "evpntopo_id8",
  "ext_ip": "ext_ip2",
  "hostname": [
    "hostname2",
    "hostname1",
    "hostname0"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

