
# Gateway Template Tunnel Ipsec Proposal

## Structure

`GatewayTemplateTunnelIpsecProposal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | [`*models.TunnelConfigsAuthAlgoEnum`](../../doc/models/tunnel-configs-auth-algo-enum.md) | Optional | - |
| `DhGroup` | [`*models.TunnelConfigsDhGroupEnum`](../../doc/models/tunnel-configs-dh-group-enum.md) | Optional | Only if `provider`== `custom-ipsec`. Possible values:<br><br>* 1<br>* 2 (1024-bit)<br>* 5<br>* 14 (default, 2048-bit)<br>* 15 (3072-bit)<br>* 16 (4096-bit)<br>* 19 (256-bit ECP)<br>* 20 (384-bit ECP)<br>* 21 (521-bit ECP)<br>* 24 (2048-bit ECP)<br>**Default**: `"14"` |
| `EncAlgo` | [`models.Optional[models.TunnelConfigsEncAlgoEnum]`](../../doc/models/tunnel-configs-enc-algo-enum.md) | Optional | **Default**: `"aes256"` |

## Example (as JSON)

```json
{
  "dh_group": "14",
  "enc_algo": "aes256",
  "auth_algo": "sha1"
}
```

