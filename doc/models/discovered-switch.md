
# Discovered Switch

## Structure

`DiscoveredSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Adopted` | `*bool` | Optional | - |
| `ApRedundancy` | [`*models.ApRedundancy`](../../doc/models/ap-redundancy.md) | Optional | - |
| `Aps` | [`[]models.DiscoveredSwitchAp`](../../doc/models/discovered-switch-ap.md) | Optional | **Constraints**: *Unique Items Required* |
| `ChassisId` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `ForSite` | `*bool` | Optional | - |
| `MgmtAddr` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SystemDesc` | `*string` | Optional | - |
| `SystemName` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Vendor` | `*string` | Optional | - |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
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
      "inactive_wired_vlans": [
        168,
        169,
        170
      ],
      "mac": "mac8",
      "poe_status": false,
      "port": "port4"
    },
    {
      "hostname": "hostname0",
      "inactive_wired_vlans": [
        168,
        169,
        170
      ],
      "mac": "mac8",
      "poe_status": false,
      "port": "port4"
    },
    {
      "hostname": "hostname0",
      "inactive_wired_vlans": [
        168,
        169,
        170
      ],
      "mac": "mac8",
      "poe_status": false,
      "port": "port4"
    }
  ],
  "chassis_id": [
    "chassis_id2"
  ],
  "for_site": false
}
```

