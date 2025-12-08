
# Search Wired Client

## Structure

`SearchWiredClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.WiredClientResponse`](../../doc/models/wired-client-response.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 74.88,
  "limit": 150,
  "results": [
    {
      "auth_method": "mac_auth",
      "auth_state": "authenticated",
      "dhcp_client_identifier": "MAC address 00155df6d500",
      "dhcp_fqdn": "ITS-VMMT0-D1N02.mgthub.local",
      "dhcp_hostname": "ITS-VMMT0-D1N02",
      "dhcp_request_params": "1 3 6 15 31 33 43 44 46 47 119 121 249 252",
      "dhcp_vendor_class_identifier": "MSFT 5.0",
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
      ]
    }
  ],
  "start": 30.94,
  "total": 56,
  "next": "next8"
}
```

