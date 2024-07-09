
# Tunnel Configs

## Structure

`TunnelConfigs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoProvision` | [`*models.TunnelConfigsAutoProvision`](../../doc/models/tunnel-configs-auto-provision.md) | Optional | - |
| `IkeLifetime` | `*int` | Optional | Only if:<br><br>* `provider`== `custom-ipsec` |
| `IkeMode` | [`*models.GatewayTemplateTunnelIkeModeEnum`](../../doc/models/gateway-template-tunnel-ike-mode-enum.md) | Optional | Only if:<br><br>* `provider`== `custom-ipsec`<br>**Default**: `"main"` |
| `IkeProposals` | [`[]models.GatewayTemplateTunnelIkeProposal`](../../doc/models/gateway-template-tunnel-ike-proposal.md) | Optional | if `provider`== `custom-ipsec` |
| `IpsecLifetime` | `*int` | Optional | if `provider`== `custom-ipsec` |
| `IpsecProposals` | [`[]models.GatewayTemplateTunnelIpsecProposal`](../../doc/models/gateway-template-tunnel-ipsec-proposal.md) | Optional | Only if:<br><br>* `provider`== `custom-ipsec` |
| `LocalId` | `*string` | Optional | Only if:<br><br>* `provider`== `zscaler-ipsec`<br>* `provider`==`jse-ipsec`<br>* `provider`== `custom-ipsec` |
| `Mode` | [`*models.GatewayTemplateTunnelModeEnum`](../../doc/models/gateway-template-tunnel-mode-enum.md) | Optional | **Default**: `"active-standby"` |
| `Primary` | [`*models.GatewayTemplateTunnelNode`](../../doc/models/gateway-template-tunnel-node.md) | Optional | - |
| `Probe` | [`*models.GatewayTemplateTunnelProbe`](../../doc/models/gateway-template-tunnel-probe.md) | Optional | Only if:<br><br>* `provider`== `custom-ipsec` |
| `Protocol` | [`*models.GatewayTemplateTunnelProtocolEnum`](../../doc/models/gateway-template-tunnel-protocol-enum.md) | Optional | Only if:<br><br>* `provider`== `custom-ipsec` |
| `Provider` | [`*models.TunnelProviderOptionsNameEnum`](../../doc/models/tunnel-provider-options-name-enum.md) | Optional | - |
| `Psk` | `*string` | Optional | Only if:<br><br>* `provider`== `zscaler-ipsec`<br>* `provider`==`jse-ipsec`<br>* `provider`== `custom-ipsec` |
| `Secondary` | [`*models.GatewayTemplateTunnelNode`](../../doc/models/gateway-template-tunnel-node.md) | Optional | - |
| `Version` | [`*models.GatewayTemplateTunnelVersionEnum`](../../doc/models/gateway-template-tunnel-version-enum.md) | Optional | Only if:<br><br>* `provider`== `custom-gre`<br>* `provider`== `custom-ipsec`<br>**Default**: `"2"` |

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
      "lng": 22.82
    },
    "primary": {
      "num_hosts": "num_hosts6",
      "wan_names": [
        "wan_names8"
      ]
    },
    "region": "auto",
    "secondary": {
      "num_hosts": "num_hosts8",
      "wan_names": [
        "wan_names0"
      ]
    }
  },
  "ike_lifetime": 236,
  "ike_proposals": [
    {
      "auth_algo": "sha2",
      "dh_group": "15",
      "enc_algo": "aes_gcm128"
    },
    {
      "auth_algo": "sha2",
      "dh_group": "15",
      "enc_algo": "aes_gcm128"
    }
  ],
  "ipsec_lifetime": 40
}
```

