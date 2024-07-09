
# Wired Client Response

## Structure

`WiredClientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `[]string` | Optional | - |
| `DeviceMacPort` | [`[]models.WiredClientResponseDeviceMacPortItem`](../../doc/models/wired-client-response-device-mac-port-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Ip` | `[]string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortId` | `[]string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Vlan` | `[]int` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "device_mac": [
    "device_mac1"
  ],
  "device_mac_port": [
    {
      "device_mac": "device_mac8",
      "ip": "ip8",
      "port_id": "port_id4",
      "port_parent": "port_parent6",
      "start": "start8"
    },
    {
      "device_mac": "device_mac8",
      "ip": "ip8",
      "port_id": "port_id4",
      "port_parent": "port_parent6",
      "start": "start8"
    }
  ],
  "ip": [
    "ip1",
    "ip2"
  ],
  "mac": "mac4"
}
```

