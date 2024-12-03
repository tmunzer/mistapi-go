
# Snmpv 3 Config

*This model accepts additional fields of type interface{}.*

## Structure

`Snmpv3Config`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Notify` | [`[]models.Snmpv3ConfigNotifyItems`](../../doc/models/snmpv-3-config-notify-items.md) | Optional | - |
| `NotifyFilter` | [`[]models.Snmpv3ConfigNotifyFilterItem`](../../doc/models/snmpv-3-config-notify-filter-item.md) | Optional | - |
| `TargetAddress` | [`[]models.Snmpv3ConfigTargetAddressItem`](../../doc/models/snmpv-3-config-target-address-item.md) | Optional | - |
| `TargetParameters` | [`[]models.Snmpv3ConfigTargetParam`](../../doc/models/snmpv-3-config-target-param.md) | Optional | - |
| `Usm` | [`*models.SnmpUsm`](../../doc/models/snmp-usm.md) | Optional | - |
| `Vacm` | [`*models.SnmpVacm`](../../doc/models/snmp-vacm.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "notify": [
    {
      "name": "name0",
      "tag": "tag4",
      "type": "inform",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "name": "name0",
      "tag": "tag4",
      "type": "inform",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "notify_filter": [
    {
      "contents": [
        {
          "include": false,
          "oid": "oid4",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "profile_name": "profile_name2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "target_address": [
    {
      "address": "address8",
      "address_mask": "address_mask0",
      "port": 198,
      "tag_list": "tag_list4",
      "target_address_name": "target_address_name6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "address": "address8",
      "address_mask": "address_mask0",
      "port": 198,
      "tag_list": "tag_list4",
      "target_address_name": "target_address_name6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "target_parameters": [
    {
      "message_processing_model": "v3",
      "name": "name4",
      "notify_filter": "notify_filter6",
      "security_level": "none",
      "security_model": "usm",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "message_processing_model": "v3",
      "name": "name4",
      "notify_filter": "notify_filter6",
      "security_level": "none",
      "security_model": "usm",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "usm": {
    "engine-id": "engine-id8",
    "engine_type": "local_engine",
    "users": [
      {
        "authentication_password": "authentication_password0",
        "authentication_type": "authentication_sha384",
        "encryption_password": "encryption_password4",
        "encryption_type": "privacy-3des",
        "name": "name6",
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

