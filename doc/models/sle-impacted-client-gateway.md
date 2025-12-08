
# Sle Impacted Client Gateway

## Structure

`SleImpactedClientGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChassisMac` | `*string` | Optional | - |
| `GatewayMac` | `*string` | Optional | - |
| `GatewayName` | `*string` | Optional | - |
| `Interfaces` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "chassis_mac": "chassis_mac4",
  "gateway_mac": "gateway_mac0",
  "gateway_name": "gateway_name8",
  "interfaces": [
    "interfaces3",
    "interfaces4"
  ]
}
```

