
# Tunnel Config

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoProvision` | [`*models.TunnelConfigAutoProvision`](../../doc/models/tunnel-config-auto-provision.md) | Optional | - |
| `IkeLifetime` | `*int` | Optional | Only if `provider`==`custom-ipsec` |
| `IkeMode` | [`*models.TunnelConfigIkeModeEnum`](../../doc/models/tunnel-config-ike-mode-enum.md) | Optional | Only if `provider`==`custom-ipsec`. enum: `aggressive`, `main`<br>**Default**: `"main"` |
| `IkeProposals` | [`[]models.TunnelConfigIkeProposal`](../../doc/models/tunnel-config-ike-proposal.md) | Optional | if `provider`==`custom-ipsec` |
| `IpsecLifetime` | `*int` | Optional | if `provider`==`custom-ipsec` |
| `IpsecProposals` | [`[]models.TunnelConfigIpsecProposal`](../../doc/models/tunnel-config-ipsec-proposal.md) | Optional | Only if  `provider`==`custom-ipsec` |
| `LocalId` | `*string` | Optional | Required if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `Mode` | [`*models.TunnelConfigTunnelModeEnum`](../../doc/models/tunnel-config-tunnel-mode-enum.md) | Optional | Required if `provider`==`zscaler-gre`, `provider`==`jse-ipsec`. enum: `active-active`, `active-standby`<br>**Default**: `"active-standby"` |
| `Networks` | `[]string` | Optional | if `provider`==`custom-ipsec`, networks reachable via this tunnel |
| `Primary` | [`*models.TunnelConfigNode`](../../doc/models/tunnel-config-node.md) | Optional | Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `Probe` | [`*models.TunnelConfigProbe`](../../doc/models/tunnel-config-probe.md) | Optional | Only if `provider`==`custom-ipsec` |
| `Protocol` | [`*models.TunnelConfigProtocolEnum`](../../doc/models/tunnel-config-protocol-enum.md) | Optional | Only if `provider`==`custom-ipsec`. enum: `gre`, `ipsec` |
| `Provider` | [`*models.TunnelConfigProviderEnum`](../../doc/models/tunnel-config-provider-enum.md) | Optional | Only if `auto_provision.enabled`==`false`. enum: `custom-ipsec`, `customer-gre`, `jse-ipsec`, `zscaler-gre`, `zscaler-ipsec` |
| `Psk` | `*string` | Optional | Required if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `Secondary` | [`*models.TunnelConfigNode`](../../doc/models/tunnel-config-node.md) | Optional | Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `Version` | [`*models.TunnelConfigVersionEnum`](../../doc/models/tunnel-config-version-enum.md) | Optional | Only if `provider`==`custom-gre` or `provider`==`custom-ipsec`. enum: `1`, `2`<br>**Default**: `"2"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ike_mode": "main",
  "mode": "active-standby",
  "version": "2",
  "auto_provision": {
    "enable": false,
    "latlng": {
      "lat": 144.64,
      "lng": 22.82,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "primary": {
      "probe_ips": [
        "probe_ips7",
        "probe_ips8",
        "probe_ips9"
      ],
      "wan_names": [
        "wan_names8"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "provider": "jse-ipsec",
    "region": "region6",
    "secondary": {
      "probe_ips": [
        "probe_ips9",
        "probe_ips0",
        "probe_ips1"
      ],
      "wan_names": [
        "wan_names0"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "ike_lifetime": 118,
  "ike_proposals": [
    {
      "auth_algo": "sha1",
      "dh_group": "19",
      "enc_algo": "aes_gcm256",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "auth_algo": "sha1",
      "dh_group": "19",
      "enc_algo": "aes_gcm256",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "ipsec_lifetime": 178,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

