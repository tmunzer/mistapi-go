
# Sle Impacted Client Gateway

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedClientGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChassisMac` | `*string` | Optional | - |
| `GatewayMac` | `*string` | Optional | - |
| `GatewayName` | `*string` | Optional | - |
| `Interfaces` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "chassis_mac": "chassis_mac4",
  "gateway_mac": "gateway_mac0",
  "gateway_name": "gateway_name8",
  "interfaces": [
    "interfaces3",
    "interfaces4"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

