
# Snmp Usmp User

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpUsmpUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthenticationPassword` | `*string` | Optional | Not required if `authentication_type`==`authentication_none`. Include alphabetic, numeric, and special characters, but it cannot include control characters.<br>**Constraints**: *Minimum Length*: `7` |
| `AuthenticationType` | [`*models.SnmpUsmpUserAuthenticationTypeEnum`](../../doc/models/snmp-usmp-user-authentication-type-enum.md) | Optional | sha224, sha256, sha384, sha512 are supported in 21.1 and newer release. enum: `authentication_md5`, `authentication_none`, `authentication_sha`, `authentication_sha224`, `authentication_sha256`, `authentication_sha384`, `authentication_sha512` |
| `EncryptionPassword` | `*string` | Optional | Not required if `encryption_type`==`privacy-none`. Include alphabetic, numeric, and special characters, but it cannot include control characters<br>**Constraints**: *Minimum Length*: `8` |
| `EncryptionType` | [`*models.SnmpUsmpUserEncryptionTypeEnum`](../../doc/models/snmp-usmp-user-encryption-type-enum.md) | Optional | enum: `privacy-3des`, `privacy-aes128`, `privacy-des`, `privacy-none` |
| `Name` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "authentication_password": "authentication_password8",
  "authentication_type": "authentication_none",
  "encryption_password": "encryption_password2",
  "encryption_type": "privacy-des",
  "name": "name4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

