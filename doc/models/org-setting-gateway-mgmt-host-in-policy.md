
# Org Setting Gateway Mgmt Host in Policy

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingGatewayMgmtHostInPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Tenants` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name2",
  "tenants": [
    "tenants3"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

