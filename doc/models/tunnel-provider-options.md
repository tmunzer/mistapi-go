
# Tunnel Provider Options

## Structure

`TunnelProviderOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Jse` | [`*models.TunnelProviderOptionsJse`](../../doc/models/tunnel-provider-options-jse.md) | Optional | For jse-ipsec, this allows provisioning of adequate resource on JSE. Make sure adequate licenses are added |
| `Prisma` | [`*models.TunnelProviderOptionsPrisma`](../../doc/models/tunnel-provider-options-prisma.md) | Optional | - |
| `Zscaler` | [`*models.TunnelProviderOptionsZscaler`](../../doc/models/tunnel-provider-options-zscaler.md) | Optional | For zscaler-ipsec and zscaler-gre |

## Example (as JSON)

```json
{
  "jse": {
    "num_users": 186,
    "org_name": "org_name6"
  },
  "prisma": {
    "service_account_name": "service_account_name6"
  },
  "zscaler": {
    "aup_block_internet_until_accepted": false,
    "aup_enabled": false,
    "aup_force_ssl_inspection": false,
    "aup_timeout_in_days": 104,
    "auth_required": false
  }
}
```

