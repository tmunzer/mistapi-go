
# Sle Impacted Users Client

## Structure

`SleImpactedUsersClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*float64` | Optional | - |
| `Duration` | `*float64` | Optional | - |
| `Gateways` | [`[]models.SleImpactedClientGateway`](../../doc/models/sle-impacted-client-gateway.md) | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `SrcIp` | `*string` | Optional | - |
| `Total` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": 51.78,
  "duration": 180.84,
  "gateways": [
    {
      "chassis_mac": "chassis_mac4",
      "gateway_mac": "gateway_mac2",
      "gateway_name": "gateway_name0",
      "interfaces": [
        "interfaces5",
        "interfaces6"
      ]
    },
    {
      "chassis_mac": "chassis_mac4",
      "gateway_mac": "gateway_mac2",
      "gateway_name": "gateway_name0",
      "interfaces": [
        "interfaces5",
        "interfaces6"
      ]
    },
    {
      "chassis_mac": "chassis_mac4",
      "gateway_mac": "gateway_mac2",
      "gateway_name": "gateway_name0",
      "interfaces": [
        "interfaces5",
        "interfaces6"
      ]
    }
  ],
  "mac": "mac2",
  "name": "name8"
}
```

