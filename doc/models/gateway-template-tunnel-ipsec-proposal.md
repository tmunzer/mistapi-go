
# Gateway Template Tunnel Ipsec Proposal

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayTemplateTunnelIpsecProposal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | [`*models.TunnelConfigsAuthAlgoEnum`](../../doc/models/tunnel-configs-auth-algo-enum.md) | Optional | enum: `md5`, `sha1`, `sha2` |
| `DhGroup` | [`*models.TunnelConfigsDhGroupEnum`](../../doc/models/tunnel-configs-dh-group-enum.md) | Optional | Only if `provider`== `custom-ipsec`. enum:<br><br>* 1<br>* 2 (1024-bit)<br>* 5<br>* 14 (default, 2048-bit)<br>* 15 (3072-bit)<br>* 16 (4096-bit)<br>* 19 (256-bit ECP)<br>* 20 (384-bit ECP)<br>* 21 (521-bit ECP)<br>* 24 (2048-bit ECP)<br>**Default**: `"14"` |
| `EncAlgo` | [`models.Optional[models.TunnelConfigsEncAlgoEnum]`](../../doc/models/tunnel-configs-enc-algo-enum.md) | Optional | enum: `3des`, `aes128`, `aes256`, `aes_gcm128`, `aes_gcm256`<br>**Default**: `"aes256"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dh_group": "14",
  "enc_algo": "aes256",
  "auth_algo": "sha2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

