
# Issued Client Certificate

## Structure

`IssuedClientCertificate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CertProvider` | `*string` | Optional | - |
| `CommonName` | `*string` | Optional | - |
| `CreatedTime` | `*string` | Optional | When the certificate has been created |
| `DeviceId` | `*uuid.UUID` | Optional | - |
| `ExpireTime` | `*string` | Optional | When the certificate will expire |
| `SerialNumber` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "cert_provider": "byod",
  "common_name": "john@corp.com",
  "created_time": "2025-08-18 10:10:30.949165+00:00",
  "device_id": "00000000-0000-0000-1000-d8695a0f9e61",
  "expire_time": "2026-08-18 10:06:00+00:00",
  "serial_number": "91984382552102771A2B3C4E5F224719956718003374658"
}
```

