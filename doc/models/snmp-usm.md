
# Snmp Usm

SNMPv3 User-based Security Model configuration

## Structure

`SnmpUsm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EngineType` | [`*models.SnmpUsmEngineTypeEnum`](../../doc/models/snmp-usm-engine-type-enum.md) | Optional | SNMP engine type used for this USM configuration. enum: `local_engine`, `remote_engine` |
| `RemoteEngineId` | `*string` | Optional | Required only if `engine_type`==`remote_engine` |
| `Users` | [`[]models.SnmpUsmUser`](../../doc/models/snmp-usm-user.md) | Optional | SNMPv3 USM user definitions |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpUsm := models.SnmpUsm{
        EngineType:           models.ToPointer(models.SnmpUsmEngineTypeEnum_LOCALENGINE),
        RemoteEngineId:       models.ToPointer("00:00:00:0b:00:00:70:10:6f:08:b6:3f"),
        Users:                []models.SnmpUsmUser{
            models.SnmpUsmUser{
                AuthenticationPassword: models.ToPointer("authentication_password0"),
                AuthenticationType:     models.ToPointer(models.SnmpUsmUserAuthenticationTypeEnum_AUTHENTICATIONSHA384),
                EncryptionPassword:     models.ToPointer("encryption_password4"),
                EncryptionType:         models.ToPointer(models.SnmpUsmUserEncryptionTypeEnum_PRIVACY3DES),
                Name:                   models.ToPointer("name6"),
            },
            models.SnmpUsmUser{
                AuthenticationPassword: models.ToPointer("authentication_password0"),
                AuthenticationType:     models.ToPointer(models.SnmpUsmUserAuthenticationTypeEnum_AUTHENTICATIONSHA384),
                EncryptionPassword:     models.ToPointer("encryption_password4"),
                EncryptionType:         models.ToPointer(models.SnmpUsmUserEncryptionTypeEnum_PRIVACY3DES),
                Name:                   models.ToPointer("name6"),
            },
        },
    }

}
```

