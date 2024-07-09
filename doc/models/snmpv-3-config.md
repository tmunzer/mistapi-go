
# Snmpv 3 Config

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

## Example (as JSON)

```json
{
  "notify": [
    {
      "name": "name0",
      "tag": "tag4",
      "type": "trap"
    },
    {
      "name": "name0",
      "tag": "tag4",
      "type": "trap"
    },
    {
      "name": "name0",
      "tag": "tag4",
      "type": "trap"
    }
  ],
  "notify_filter": [
    {
      "contents": [
        {
          "include": false,
          "oid": "oid4"
        }
      ],
      "profile_name": "profile_name2"
    },
    {
      "contents": [
        {
          "include": false,
          "oid": "oid4"
        }
      ],
      "profile_name": "profile_name2"
    }
  ],
  "target_address": [
    {
      "address": "address8",
      "address_mask": "address_mask0",
      "port": 198,
      "tag_list": "tag_list4",
      "target_address_name": "target_address_name6"
    }
  ],
  "target_parameters": [
    {
      "message_processing_model": "v3",
      "name": "name4",
      "notify_filter": "notify_filter6",
      "security_level": "none",
      "security_model": "usm"
    },
    {
      "message_processing_model": "v3",
      "name": "name4",
      "notify_filter": "notify_filter6",
      "security_level": "none",
      "security_model": "usm"
    },
    {
      "message_processing_model": "v3",
      "name": "name4",
      "notify_filter": "notify_filter6",
      "security_level": "none",
      "security_model": "usm"
    }
  ],
  "usm": {
    "engine-id": "engine-id8",
    "engine_type": "remote_engine",
    "users": [
      {
        "authentication_password": "authentication_password0",
        "authentication_type": "authentication_sha512",
        "encryption_password": "encryption_password4",
        "encryption_type": "privacy-aes128",
        "name": "name6"
      }
    ]
  }
}
```
