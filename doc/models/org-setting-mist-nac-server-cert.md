
# Org Setting Mist Nac Server Cert

radius server cert to be presented in EAP TLS

## Structure

`OrgSettingMistNacServerCert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cert` | `*string` | Optional | - |
| `Key` | `*string` | Optional | - |
| `Password` | `*string` | Optional | private key password (optional) |

## Example (as JSON)

```json
{
  "cert": "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
  "key": "-----BEGIN PRI...",
  "password": "password6"
}
```

