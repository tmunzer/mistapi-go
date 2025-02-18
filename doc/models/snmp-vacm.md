
# Snmp Vacm

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpVacm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Access` | [`[]models.SnmpVacmAccessItem`](../../doc/models/snmp-vacm-access-item.md) | Optional | - |
| `SecurityToGroup` | [`*models.SnmpVacmSecurityToGroup`](../../doc/models/snmp-vacm-security-to-group.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "access": [
    {
      "group_name": "group_name4",
      "prefix_list": [
        {
          "context_prefix": "context_prefix2",
          "notify_view": "notify_view6",
          "read_view": "read_view0",
          "security_level": "none",
          "security_model": "v1",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "context_prefix": "context_prefix2",
          "notify_view": "notify_view6",
          "read_view": "read_view0",
          "security_level": "none",
          "security_model": "v1",
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
    {
      "group_name": "group_name4",
      "prefix_list": [
        {
          "context_prefix": "context_prefix2",
          "notify_view": "notify_view6",
          "read_view": "read_view0",
          "security_level": "none",
          "security_model": "v1",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "context_prefix": "context_prefix2",
          "notify_view": "notify_view6",
          "read_view": "read_view0",
          "security_level": "none",
          "security_model": "v1",
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
    }
  ],
  "security_to_group": {
    "content": [
      {
        "group": "group2",
        "security_name": "security_name6",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "group": "group2",
        "security_name": "security_name6",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "group": "group2",
        "security_name": "security_name6",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ],
    "security_model": "v2c",
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

