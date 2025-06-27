
# Account Prisma Config

OAuth linked CrowdStrike apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountPrismaConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoProbeSubnet` | `*string` | Optional | Required If `enable_probe`==`true`. This field will accept an IPv4 cidr and an IP address will be picked from this range to be used as tunnel probe source ip address and as well as BGP neighbour IP address. The subnet should be big enough for num_devices * num_tunnel * 2 |
| `ClientId` | `string` | Required | Customer account api client ID |
| `ClientSecret` | `string` | Required | Customer account api client Secret |
| `EnableProbe` | `*bool` | Optional | To enable/disable tunnel probe<br><br>**Default**: `false` |
| `TsgId` | `string` | Required | Prisma Tenant Service Group id |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auto_probe_subnet": "11.0.0.0/8",
  "client_id": "client_id2",
  "client_secret": "client_secret8",
  "enable_probe": false,
  "tsg_id": "tsg_id6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

