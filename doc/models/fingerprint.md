
# Fingerprint

## Structure

`Fingerprint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Family` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Mfg` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Os` | `*string` | Optional | - |
| `OsType` | `*string` | Optional | - |
| `RandomMac` | `*bool` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "family": "family8",
  "mac": "mac0",
  "mfg": "mfg4",
  "model": "model4"
}
```

