
# Response Certificate

If the current Org CA certificate is set to expire within 30 days, a pending certificate will be returned along with the expected auto-renewal timestamp.

## Structure

`ResponseCertificate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cert` | `string` | Required | - |
| `PendingCert` | `*string` | Optional | - |
| `PendingCertExpiry` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "cert": "cert4",
  "pending_cert": "pending_cert8",
  "pending_cert_expiry": 96
}
```

