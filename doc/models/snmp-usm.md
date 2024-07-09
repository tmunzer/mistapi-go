
# Snmp Usm

## Structure

`SnmpUsm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EngineId` | `*string` | Optional | required only if `engine_type`==`remote_engine` |
| `EngineType` | [`*models.SnmpUsmEngineTypeEnum`](../../doc/models/snmp-usm-engine-type-enum.md) | Optional | - |
| `Users` | [`[]models.SnmpUsmpUser`](../../doc/models/snmp-usmp-user.md) | Optional | - |

## Example (as JSON)

```json
{
  "engine-id": "00:00:00:0b:00:00:70:10:6f:08:b6:3f",
  "engine_type": "remote_engine",
  "users": [
    {
      "authentication_password": "authentication_password0",
      "authentication_type": "authentication_sha512",
      "encryption_password": "encryption_password4",
      "encryption_type": "privacy-aes128",
      "name": "name6"
    },
    {
      "authentication_password": "authentication_password0",
      "authentication_type": "authentication_sha512",
      "encryption_password": "encryption_password4",
      "encryption_type": "privacy-aes128",
      "name": "name6"
    }
  ]
}
```

