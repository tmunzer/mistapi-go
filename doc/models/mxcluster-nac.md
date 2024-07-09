
# Mxcluster Nac

## Structure

`MxclusterNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServerPort` | `*int` | Optional | **Default**: `1813` |
| `AuthServerPort` | `*int` | Optional | **Default**: `1812` |
| `ClientIps` | [`map[string]models.MxclusterNacClientIp`](../../doc/models/mxcluster-nac-client-ip.md) | Optional | Property key is the RADIUS Client IP/Subnet. |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Secret` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "acct_server_port": 1813,
  "auth_server_port": 1812,
  "enabled": false,
  "secret": "testing123",
  "client_ips": {
    "key0": {
      "secret": "secret4",
      "site_id": "0000197c-0000-0000-0000-000000000000",
      "vendor": "paloalto"
    },
    "key1": {
      "secret": "secret4",
      "site_id": "0000197c-0000-0000-0000-000000000000",
      "vendor": "paloalto"
    },
    "key2": {
      "secret": "secret4",
      "site_id": "0000197c-0000-0000-0000-000000000000",
      "vendor": "paloalto"
    }
  }
}
```

