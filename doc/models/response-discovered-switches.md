
# Response Discovered Switches

## Structure

`ResponseDiscoveredSwitches`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.DiscoveredSwitch`](../../doc/models/discovered-switch.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 55.42,
  "limit": 4,
  "results": [
    {
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "adopted": false,
      "ap_redundancy": {
        "modules": {
          "key0": {
            "num_aps": 2,
            "num_aps_with_switch_redundancy": 254
          }
        },
        "num_aps": 246,
        "num_aps_with_switch_redundancy": 10
      },
      "aps": [
        {
          "hostname": "hostname0",
          "mac": "mac8",
          "poe_status": false,
          "port": "port4",
          "port_id": "port_id4"
        },
        {
          "hostname": "hostname0",
          "mac": "mac8",
          "poe_status": false,
          "port": "port4",
          "port_id": "port_id4"
        }
      ],
      "chassis_id": [
        "chassis_id0",
        "chassis_id1",
        "chassis_id2"
      ],
      "for_site": false
    }
  ],
  "start": 11.48,
  "total": 98,
  "next": "next2"
}
```

