
# Tunnel Config Ike Proposal

IKE proposal settings for custom IPsec tunnels

## Structure

`TunnelConfigIkeProposal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAlgo` | [`*models.TunnelConfigAuthAlgoEnum`](../../doc/models/tunnel-config-auth-algo-enum.md) | Optional | enum: `md5`, `sha1`, `sha2` |
| `DhGroup` | [`*models.TunnelConfigIkeDhGroupEnum`](../../doc/models/tunnel-config-ike-dh-group-enum.md) | Optional | Diffie-Hellman group for IKE phase 1. enum: `1`, `14`, `15`, `16`, `19`, `2`, `20`, `21`, `24`, `5`. `14` is the default 2048-bit group; `19`, `20`, `21`, and `24` are ECP groups<br><br>**Default**: `"14"` |
| `EncAlgo` | [`models.Optional[models.TunnelConfigEncAlgoEnum]`](../../doc/models/tunnel-config-enc-algo-enum.md) | Optional | enum: `3des`, `aes128`, `aes256`, `aes_gcm128`, `aes_gcm256`<br><br>**Default**: `"aes256"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigIkeProposal := models.TunnelConfigIkeProposal{
        AuthAlgo:             models.ToPointer(models.TunnelConfigAuthAlgoEnum_MD5),
        DhGroup:              models.ToPointer(models.TunnelConfigIkeDhGroupEnum_ENUM14),
        EncAlgo:              models.NewOptional(models.ToPointer(models.TunnelConfigEncAlgoEnum_AES256)),
    }

}
```

