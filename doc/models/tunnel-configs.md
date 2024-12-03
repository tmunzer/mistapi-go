
# Tunnel Configs

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelConfigs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoProvision` | [`*models.TunnelConfigsAutoProvision`](../../doc/models/tunnel-configs-auto-provision.md) | Optional | - |
| `IkeLifetime` | `*int` | Optional | Only if `provider`== `custom-ipsec` |
| `IkeMode` | [`*models.GatewayTemplateTunnelIkeModeEnum`](../../doc/models/gateway-template-tunnel-ike-mode-enum.md) | Optional | Only if `provider`== `custom-ipsec`. enum: `aggressive`, `main`<br>**Default**: `"main"` |
| `IkeProposals` | [`[]models.GatewayTemplateTunnelIkeProposal`](../../doc/models/gateway-template-tunnel-ike-proposal.md) | Optional | if `provider`== `custom-ipsec` |
| `IpsecLifetime` | `*int` | Optional | if `provider`== `custom-ipsec` |
| `IpsecProposals` | [`[]models.GatewayTemplateTunnelIpsecProposal`](../../doc/models/gateway-template-tunnel-ipsec-proposal.md) | Optional | Only if  `provider`== `custom-ipsec` |
| `LocalId` | `*string` | Optional | Only if:<br><br>* `provider`== `zscaler-ipsec`<br>* `provider`==`jse-ipsec`<br>* `provider`== `custom-ipsec` |
| `Mode` | [`*models.GatewayTemplateTunnelModeEnum`](../../doc/models/gateway-template-tunnel-mode-enum.md) | Optional | enum: `active-active`, `active-standby`<br>**Default**: `"active-standby"` |
| `Networks` | `[]string` | Optional | networks reachable via this tunnel |
| `Primary` | [`*models.GatewayTemplateTunnelNode`](../../doc/models/gateway-template-tunnel-node.md) | Optional | - |
| `Probe` | [`*models.GatewayTemplateTunnelProbe`](../../doc/models/gateway-template-tunnel-probe.md) | Optional | Only if `provider`== `custom-ipsec` |
| `Protocol` | [`*models.GatewayTemplateTunnelProtocolEnum`](../../doc/models/gateway-template-tunnel-protocol-enum.md) | Optional | Only if `provider`== `custom-ipsec`. enum: `gre`, `ipsec` |
| `Provider` | [`*models.TunnelProviderOptionsNameEnum`](../../doc/models/tunnel-provider-options-name-enum.md) | Optional | enum: `custom-ipsec`, `customer-gre`, `jse-ipsec`, `zscaler-gre`, `zscaler-ipsec` |
| `Psk` | `*string` | Optional | Only if:<br><br>* `provider`== `zscaler-ipsec`<br>* `provider`==`jse-ipsec`<br>* `provider`== `custom-ipsec` |
| `Secondary` | [`*models.GatewayTemplateTunnelNode`](../../doc/models/gateway-template-tunnel-node.md) | Optional | - |
| `Version` | [`*models.GatewayTemplateTunnelVersionEnum`](../../doc/models/gateway-template-tunnel-version-enum.md) | Optional | Only if `provider`== `custom-gre` or `provider`== `custom-ipsec`. enum: `1`, `2`<br>**Default**: `"2"` |
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
      "num_hosts": "num_hosts6",
      "wan_names": [
        "wan_names8"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "secondary": {
      "num_hosts": "num_hosts8",
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
  "ike_lifetime": 136,
  "ike_proposals": [
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
  "ipsec_lifetime": 196,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

