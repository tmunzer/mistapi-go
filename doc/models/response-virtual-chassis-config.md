
# Response Virtual Chassis Config

## Structure

`ResponseVirtualChassisConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | [`*models.VirtualChassisConfig`](../../doc/models/virtual-chassis-config.md) | Optional | Virtual Chassis |

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
        "vc_role": "linecard"
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
        "vc_role": "linecard"
      }
    ],
    "preprovisioned": false
  }
}
```

