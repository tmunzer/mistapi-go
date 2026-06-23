
# Tunnel Config Ipsec Proposal

IPsec proposal settings for custom IPsec tunnels

## Structure

`TunnelConfigIpsecProposal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | [`*models.TunnelConfigAuthAlgoEnum`](../../doc/models/tunnel-config-auth-algo-enum.md) | Optional | enum: `md5`, `sha1`, `sha2` |
| `DhGroup` | [`*models.TunnelConfigDhGroupEnum`](../../doc/models/tunnel-config-dh-group-enum.md) | Optional | Only if `provider`==`custom-ipsec`. Diffie-Hellman group for IPsec phase 2. enum: `1`, `14`, `15`, `16`, `19`, `2`, `20`, `21`, `24`, `5`. `14` is the default 2048-bit group; `19`, `20`, `21`, and `24` are ECP groups<br><br>**Default**: `"14"` |
| `EncAlgo` | [`models.Optional[models.TunnelConfigEncAlgoEnum]`](../../doc/models/tunnel-config-enc-algo-enum.md) | Optional | enum: `3des`, `aes128`, `aes256`, `aes_gcm128`, `aes_gcm256`<br><br>**Default**: `"aes256"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigIpsecProposal := models.TunnelConfigIpsecProposal{
        AuthAlgo:             models.ToPointer(models.TunnelConfigAuthAlgoEnum_SHA2),
        DhGroup:              models.ToPointer(models.TunnelConfigDhGroupEnum_ENUM14),
        EncAlgo:              models.NewOptional(models.ToPointer(models.TunnelConfigEncAlgoEnum_AES256)),
    }

}
```

