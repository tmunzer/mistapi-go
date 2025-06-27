
# Tunnel Provider Options Prisma

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelProviderOptionsPrisma`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ServiceAccountName` | `*string` | Optional | For prisma-ipsec, service account name to used for tunnel auto provisioning |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "service_account_name": "sa1@1823425211",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

