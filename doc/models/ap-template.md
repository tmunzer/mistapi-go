
# Ap Template

*This model accepts additional fields of type interface{}.*

## Structure

`ApTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMatching` | [`models.ApTemplateMatching`](../../doc/models/ap-template-matching.md) | Required | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Wifi` | [`*models.ApTemplateWifi`](../../doc/models/ap-template-wifi.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          },
          "key1": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        },
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
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
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          },
          "key1": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        },
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
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
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          },
          "key1": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        },
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 236.6,
  "for_site": false,
  "modified_time": 98.36,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

