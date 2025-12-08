
# Snmp Usm

## Structure

`SnmpUsm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EngineType` | [`*models.SnmpUsmEngineTypeEnum`](../../doc/models/snmp-usm-engine-type-enum.md) | Optional | enum: `local_engine`, `remote_engine` |
| `RemoteEngineId` | `*string` | Optional | Required only if `engine_type`==`remote_engine` |
| `Users` | [`[]models.SnmpUsmUser`](../../doc/models/snmp-usm-user.md) | Optional | - |

## Example (as JSON)

```json
{
  "remote_engine_id": "00:00:00:0b:00:00:70:10:6f:08:b6:3f",
  "engine_type": "local_engine",
  "users": [
    {
      "authentication_password": "authentication_password0",
      "authentication_type": "authentication-sha384",
      "encryption_password": "encryption_password4",
      "encryption_type": "privacy-3des",
      "name": "name6"
    }
  ]
}
```

