
# Snmp Usm

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpUsm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EngineId` | `*string` | Optional | required only if `engine_type`==`remote_engine` |
| `EngineType` | [`*models.SnmpUsmEngineTypeEnum`](../../doc/models/snmp-usm-engine-type-enum.md) | Optional | enum: `local_engine`, `remote_engine` |
| `Users` | [`[]models.SnmpUsmpUser`](../../doc/models/snmp-usmp-user.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "engine-id": "00:00:00:0b:00:00:70:10:6f:08:b6:3f",
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
}
```

