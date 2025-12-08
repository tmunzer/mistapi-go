
# Response Discovered Switch Metrics

## Structure

`ResponseDiscoveredSwitchMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.DiscoveredSwitchMetric`](../../doc/models/discovered-switch-metric.md) | Required | - |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 222.06,
  "limit": 236,
  "results": [
    {
      "hostname": "SW-HLAB-ea2e00",
      "mgmt_addr": "10.10.20.42",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "scope": "site",
      "score": 100,
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "system_desc": "Juniper Networks, Inc. ex4100-f-12p Ethernet Switch, kernel JUNOS 22.4R3.25, Build date: 2024-02-10 00:49:09 UTC Copyright (c) 1996-2024 Juniper Networks, Inc.",
      "system_name": "SW-HLAB-ea2e00",
      "type": "inactive_wired_vlans",
      "adopted": false,
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
      ]
    }
  ],
  "start": 178.12,
  "total": 182,
  "next": "next6"
}
```

