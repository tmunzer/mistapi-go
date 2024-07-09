
# Sle Impacted Clients Client

## Structure

`SleImpactedClientsClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*int` | Optional | - |
| `Duration` | `*int` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Switches` | [`[]models.SleImpactedClientsClientSwitch`](../../doc/models/sle-impacted-clients-client-switch.md) | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": 100,
  "duration": 206,
  "mac": "mac4",
  "name": "name0",
  "switches": [
    {
      "interfaces": [
        "interfaces9"
      ],
      "switch_mac": "switch_mac6",
      "switch_name": "switch_name0"
    }
  ]
}
```

