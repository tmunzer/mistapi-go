
# Response Virtual Chassis Config

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseVirtualChassisConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | [`*models.VirtualChassisConfig`](../../doc/models/virtual-chassis-config.md) | Optional | Virtual Chassis |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": {
    "locating": false,
    "members": [
      {
        "locating": false,
        "mac": "mac2",
        "member": 176,
        "vc_ports": [
          "vc_ports2",
          "vc_ports3",
          "vc_ports4"
        ],
        "vc_role": "master",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "locating": false,
        "mac": "mac2",
        "member": 176,
        "vc_ports": [
          "vc_ports2",
          "vc_ports3",
          "vc_ports4"
        ],
        "vc_role": "master",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ],
    "preprovisioned": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

