
# Snmp Usm User

SNMPv3 USM user definition

## Structure

`SnmpUsmUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthenticationPassword` | `*string` | Optional | Not required if `authentication_type`==`authentication-none`. Include alphabetic, numeric, and special characters, but it cannot include control characters.<br><br>**Constraints**: *Minimum Length*: `7` |
| `AuthenticationType` | [`*models.SnmpUsmUserAuthenticationTypeEnum`](../../doc/models/snmp-usm-user-authentication-type-enum.md) | Optional | sha224, sha256, sha384, sha512 are supported in 21.1 and newer release. enum: `authentication-md5`, `authentication-none`, `authentication-sha`, `authentication-sha224`, `authentication-sha256`, `authentication-sha384`, `authentication-sha512` |
| `EncryptionPassword` | `*string` | Optional | Not required if `encryption_type`==`privacy-none`. Include alphabetic, numeric, and special characters, but it cannot include control characters<br><br>**Constraints**: *Minimum Length*: `8` |
| `EncryptionType` | [`*models.SnmpUsmUserEncryptionTypeEnum`](../../doc/models/snmp-usm-user-encryption-type-enum.md) | Optional | enum: `privacy-3des`, `privacy-aes128`, `privacy-des`, `privacy-none` |
| `Name` | `*string` | Optional | Username for the SNMPv3 USM user |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpUsmUser := models.SnmpUsmUser{
        AuthenticationPassword: models.ToPointer("authentication_password4"),
        AuthenticationType:     models.ToPointer(models.SnmpUsmUserAuthenticationTypeEnum_AUTHENTICATIONSHA512),
        EncryptionPassword:     models.ToPointer("encryption_password8"),
        EncryptionType:         models.ToPointer(models.SnmpUsmUserEncryptionTypeEnum_PRIVACYDES),
        Name:                   models.ToPointer("name0"),
    }

}
```

