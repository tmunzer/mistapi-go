
# Ap Template

## Structure

`ApTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMatching` | [`models.ApTemplateMatching`](../../doc/models/ap-template-matching.md) | Required | - |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Wifi` | [`*models.ApTemplateWifi`](../../doc/models/ap-template-wifi.md) | Optional | - |

## Example (as JSON)

```json
{
  "ap_matching": {
    "enabled": false,
    "rules": [
      {
        "match_model": "match_model0",
        "name": "name8",
        "port_config": {
          "key0": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false
          },
          "key1": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false
          }
        }
      },
      {
        "match_model": "match_model0",
        "name": "name8",
        "port_config": {
          "key0": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false
          },
          "key1": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false
          }
        }
      },
      {
        "match_model": "match_model0",
        "name": "name8",
        "port_config": {
          "key0": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false
          },
          "key1": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false
          }
        }
      }
    ]
  },
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 236.6,
  "for_site": false,
  "id": "00002666-0000-0000-0000-000000000000",
  "modified_time": 98.36,
  "org_id": "0000002e-0000-0000-0000-000000000000"
}
```

