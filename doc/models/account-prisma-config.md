
# Account Prisma Config

OAuth linked CrowdStrike apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountPrismaConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account api client ID |
| `ClientSecret` | `string` | Required | Customer account api client Secret |
| `TsgId` | `string` | Required | Prisma Tenant Service Group id |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_id": "client_id2",
  "client_secret": "client_secret8",
  "tsg_id": "tsg_id6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

