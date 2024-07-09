
# Snmp Usmp User

## Structure

`SnmpUsmpUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthenticationPassword` | `*string` | Optional | Not required if `authentication_type`==`authentication_none`<br>include alphabetic, numeric, and special characters, but it cannot include control characters.<br>**Constraints**: *Minimum Length*: `7` |
| `AuthenticationType` | [`*models.SnmpUsmpUserAuthenticationTypeEnum`](../../doc/models/snmp-usmp-user-authentication-type-enum.md) | Optional | sha224, sha256, sha384, sha512 are supported in 21.1 and newer release |
| `EncryptionPassword` | `*string` | Optional | Not required if `encryption_type`==`privacy-none`<br>include alphabetic, numeric, and special characters, but it cannot include control characters<br>**Constraints**: *Minimum Length*: `8` |
| `EncryptionType` | [`*models.SnmpUsmpUserEncryptionTypeEnum`](../../doc/models/snmp-usmp-user-encryption-type-enum.md) | Optional | - |
| `Name` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "authentication_password": "authentication_password2",
  "authentication_type": "authentication_sha224",
  "encryption_password": "encryption_password6",
  "encryption_type": "privacy-aes128",
  "name": "name8"
}
```

