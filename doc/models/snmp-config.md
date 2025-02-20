
# Snmp Config

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientList` | [`[]models.SnmpConfigClientList`](../../doc/models/snmp-config-client-list.md) | Optional | - |
| `Contact` | `*string` | Optional | - |
| `Description` | `*string` | Optional | - |
| `Enabled` | `*bool` | Optional | **Default**: `true` |
| `EngineId` | `*string` | Optional | **Constraints**: *Maximum Length*: `27` |
| `Location` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Network` | `*string` | Optional | **Default**: `"default"` |
| `TrapGroups` | [`[]models.SnmpConfigTrapGroup`](../../doc/models/snmp-config-trap-group.md) | Optional | - |
| `V2cConfig` | [`[]models.SnmpConfigV2cConfig`](../../doc/models/snmp-config-v2-c-config.md) | Optional | - |
| `V3Config` | [`*models.Snmpv3Config`](../../doc/models/snmpv-3-config.md) | Optional | - |
| `Views` | [`[]models.SnmpConfigView`](../../doc/models/snmp-config-view.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "contact": "cns@juniper.net",
  "description": "Juniper QFX Series Switch - 1K_5LA",
  "enabled": true,
  "location": "Las Vegas, NV",
  "name": "TGH-1K-QFX10K",
  "network": "default",
  "client_list": [
    {
      "client_list_name": "client_list_name2",
      "clients": [
        "clients4"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "client_list_name": "client_list_name2",
      "clients": [
        "clients4"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "client_list_name": "client_list_name2",
      "clients": [
        "clients4"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "engine_id": "engine_id6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

