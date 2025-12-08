
# Site Setting Mxedge

Site Mist Edges form a cluster of RadSec Proxy servers

## Structure

`SiteSettingMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MistDas` | [`*models.MxedgeDas`](../../doc/models/mxedge-das.md) | Optional | Configure cloud-assisted dynamic authorization service on this cluster of mist edges |
| `MistNac` | [`*models.MxclusterNac`](../../doc/models/mxcluster-nac.md) | Optional | - |
| `MistNacedge` | [`*models.MistNacedge`](../../doc/models/mist-nacedge.md) | Optional | - |
| `Radsec` | [`*models.MxclusterRadsec`](../../doc/models/mxcluster-radsec.md) | Optional | MxEdge RadSec Configuration |

## Example (as JSON)

```json
{
  "mist_das": {
    "coa_servers": [
      {
        "disable_event_timestamp_check": false,
        "enabled": false,
        "host": "host8",
        "port": 28,
        "require_message_authenticator": false
      },
      {
        "disable_event_timestamp_check": false,
        "enabled": false,
        "host": "host8",
        "port": 28,
        "require_message_authenticator": false
      }
    ],
    "enabled": false
  },
  "mist_nac": {
    "acct_server_port": 70,
    "auth_server_port": 34,
    "client_ips": {
      "key0": {
        "require_message_authenticator": false,
        "secret": "secret4",
        "site_id": "0000197c-0000-0000-0000-000000000000",
        "vendor": "cisco-aironet"
      },
      "key1": {
        "require_message_authenticator": false,
        "secret": "secret4",
        "site_id": "0000197c-0000-0000-0000-000000000000",
        "vendor": "cisco-aironet"
      },
      "key2": {
        "require_message_authenticator": false,
        "secret": "secret4",
        "site_id": "0000197c-0000-0000-0000-000000000000",
        "vendor": "cisco-aironet"
      }
    },
    "enabled": false,
    "secret": "secret6"
  },
  "mist_nacedge": {
    "auth_ttl": 110,
    "default_dot1x_vlan": "default_dot1x_vlan4",
    "default_vlan": "default_vlan6",
    "enabled": false,
    "mxedge_hosts": [
      "mxedge_hosts7",
      "mxedge_hosts8",
      "mxedge_hosts9"
    ]
  },
  "radsec": {
    "acct_servers": [
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
      }
    ],
    "enabled": false,
    "match_ssid": false,
    "nas_ip_source": "tunnel"
  }
}
```

