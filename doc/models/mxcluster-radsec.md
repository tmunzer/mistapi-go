
# Mxcluster Radsec

MxEdge Radsec Configuration

## Structure

`MxclusterRadsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.MxclusterRadsecAcctServer`](../../doc/models/mxcluster-radsec-acct-server.md) | Optional | list of RADIUS accounting servers, optional, order matters where the first one is treated as primary<br>**Constraints**: *Unique Items Required* |
| `AuthServers` | [`[]models.MxclusterRadsecAuthServer`](../../doc/models/mxcluster-radsec-auth-server.md) | Optional | list of RADIUS authentication servers, order matters where the first one is treated as primary<br>**Constraints**: *Unique Items Required* |
| `Enabled` | `*bool` | Optional | whether to enable service on Mist Edge i.e. RADIUS proxy over TLS |
| `MatchSsid` | `*bool` | Optional | whether to match ssid in request message to select from a subset of RADIUS servers |
| `ProxyHosts` | `[]string` | Optional | hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to `tunterm_hosts` |
| `ServerSelection` | [`*models.MxclusterRadsecServerSelectionEnum`](../../doc/models/mxcluster-radsec-server-selection-enum.md) | Optional | When ordered, Mist Edge will prefer and go back to the first radius server if possible. enum: `ordered`, `unordered`<br>**Default**: `"ordered"` |
| `Source` | [`*models.MxclusterRadsecSourceEnum`](../../doc/models/mxcluster-radsec-source-enum.md) | Optional | Specify source address to use when connecting to RADIUS servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`<br>**Default**: `"any"` |

## Example (as JSON)

```json
{
  "server_selection": "ordered",
  "source": "any",
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
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek4",
      "keywrap_mack": "keywrap_mack6"
    },
    {
      "host": "host0",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek4",
      "keywrap_mack": "keywrap_mack6"
    },
    {
      "host": "host0",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek4",
      "keywrap_mack": "keywrap_mack6"
    }
  ],
  "enabled": false,
  "match_ssid": false,
  "proxy_hosts": [
    "proxy_hosts8",
    "proxy_hosts9",
    "proxy_hosts0"
  ]
}
```

