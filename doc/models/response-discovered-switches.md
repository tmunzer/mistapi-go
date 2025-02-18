
# Response Discovered Switches

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 93.54,
  "limit": 224,
  "results": [
    {
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "adopted": false,
      "ap_redundancy": {
        "modules": {
          "key0": {
            "num_aps": 2,
            "num_aps_with_switch_redundancy": 254,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        },
        "num_aps": 246,
        "num_aps_with_switch_redundancy": 10,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "aps": [
        {
          "hostname": "hostname0",
          "mac": "mac8",
          "poe_status": false,
          "port": "port4",
          "port_id": "port_id4",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "hostname": "hostname0",
          "mac": "mac8",
          "poe_status": false,
          "port": "port4",
          "port_id": "port_id4",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "chassis_id": [
        "chassis_id0",
        "chassis_id1",
        "chassis_id2"
      ],
      "for_site": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 49.6,
  "total": 130,
  "next": "next4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

