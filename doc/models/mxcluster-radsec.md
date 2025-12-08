
# Mxcluster Radsec

MxEdge RadSec Configuration

## Structure

`MxclusterRadsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.MxclusterRadsecAcctServer`](../../doc/models/mxcluster-radsec-acct-server.md) | Optional | List of RADIUS accounting servers, optional, order matters where the first one is treated as primary<br><br>**Constraints**: *Unique Items Required* |
| `AuthServers` | [`[]models.MxclusterRadsecAuthServer`](../../doc/models/mxcluster-radsec-auth-server.md) | Optional | List of RADIUS authentication servers, order matters where the first one is treated as primary<br><br>**Constraints**: *Unique Items Required* |
| `Enabled` | `*bool` | Optional | Whether to enable service on Mist Edge i.e. RADIUS proxy over TLS |
| `MatchSsid` | `*bool` | Optional | Whether to match ssid in request message to select from a subset of RADIUS servers |
| `NasIpSource` | [`*models.MxclusterRadsecNasIpSourceEnum`](../../doc/models/mxcluster-radsec-nas-ip-source-enum.md) | Optional | SSpecify NAS-IP-ADDRESS, NAS-IPv6-ADDRESS to use with auth_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`<br><br>**Default**: `"any"` |
| `ProxyHosts` | `[]string` | Optional | Hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to `tunterm_hosts` |
| `ServerSelection` | [`*models.MxclusterRadsecServerSelectionEnum`](../../doc/models/mxcluster-radsec-server-selection-enum.md) | Optional | When ordered, Mist Edge will prefer and go back to the first radius server if possible. enum: `ordered`, `unordered`<br><br>**Default**: `"ordered"` |
| `SrcIpSource` | [`*models.MxclusterRadsecSrcIpSourceEnum`](../../doc/models/mxcluster-radsec-src-ip-source-enum.md) | Optional | Specify IP address to connect to auth_servers and acct_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`<br><br>**Default**: `"any"` |

## Example (as JSON)

```json
{
  "nas_ip_source": "any",
  "server_selection": "ordered",
  "src_ip_source": "any",
  "acct_servers": [
    {
      "host": "host4",
      "port": 254,
      "secret": "secret0",
      "ssids": [
        "ssids5",
        "ssids6"
      ]
    },
    {
      "host": "host4",
      "port": 254,
      "secret": "secret0",
      "ssids": [
        "ssids5",
        "ssids6"
      ]
    },
    {
      "host": "host4",
      "port": 254,
      "secret": "secret0",
      "ssids": [
        "ssids5",
        "ssids6"
      ]
    }
  ],
  "auth_servers": [
    {
      "host": "host0",
      "inband_status_check": false,
      "inband_status_interval": 160,
      "keywrap_enabled": false,
      "keywrap_format": "ascii"
    },
    {
      "host": "host0",
      "inband_status_check": false,
      "inband_status_interval": 160,
      "keywrap_enabled": false,
      "keywrap_format": "ascii"
    },
    {
      "host": "host0",
      "inband_status_check": false,
      "inband_status_interval": 160,
      "keywrap_enabled": false,
      "keywrap_format": "ascii"
    }
  ],
  "enabled": false,
  "match_ssid": false
}
```

