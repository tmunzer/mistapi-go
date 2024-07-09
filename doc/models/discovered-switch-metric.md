
# Discovered Switch Metric

## Structure

`DiscoveredSwitchMetric`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Adopted` | `*bool` | Optional | - |
| `Aps` | [`[]models.DiscoveredSwitchMetricAp`](../../doc/models/discovered-switch-metric-ap.md) | Optional | - |
| `ChassisId` | `[]string` | Optional | - |
| `Hostname` | `*string` | Optional | - |
| `MgmtAddr` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Scope` | `*string` | Optional | - |
| `Score` | `*int` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SystemDesc` | `*string` | Optional | - |
| `SystemName` | `*string` | Optional | - |
| `Timestamp` | `*int` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Vendor` | `*string` | Optional | - |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "hostname": "SW-HLAB-ea2e00",
  "mgmt_addr": "10.10.20.42",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "scope": "site",
  "score": 100,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "system_desc": "Juniper Networks, Inc. ex4100-f-12p Ethernet Switch, kernel JUNOS 22.4R3.25, Build date: 2024-02-10 00:49:09 UTC Copyright (c) 1996-2024 Juniper Networks, Inc.",
  "system_name": "SW-HLAB-ea2e00",
  "timestamp": 1675193400,
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
    "chassis_id2"
  ]
}
```

