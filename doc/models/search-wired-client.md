
# Search Wired Client

## Structure

`SearchWiredClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.WiredClientResponse`](../../doc/models/wired-client-response.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 89.68,
  "limit": 162,
  "results": [
    {
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "device_mac": [
        "device_mac7"
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
        "ip7",
        "ip8"
      ],
      "mac": "mac0"
    }
  ],
  "start": 45.74,
  "total": 0,
  "next": "next2"
}
```

