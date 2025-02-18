
# Switch Search

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Clustered` | `*bool` | Optional | - |
| `EvpnMissingLinks` | `*bool` | Optional | - |
| `EvpntopoId` | `*string` | Optional | - |
| `ExtIp` | `*string` | Optional | - |
| `Hostname` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Ip` | `*string` | Optional | - |
| `LastHostname` | `*string` | Optional | - |
| `LastTroubleCode` | `*string` | Optional | - |
| `LastTroubleTimestamp` | `*int` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Managed` | `*bool` | Optional | - |
| `Model` | `*string` | Optional | - |
| `NumMembers` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Role` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TimeDrifted` | `*bool` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `string` | Required, Constant | Device Type. enum: `switch`<br>**Value**: `"switch"` |
| `Uptime` | `*int` | Optional | - |
| `Version` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "switch",
  "clustered": false,
  "evpn_missing_links": false,
  "evpntopo_id": "evpntopo_id6",
  "ext_ip": "ext_ip0",
  "hostname": [
    "hostname0"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

