
# Site Setting Mxedge

site mist edges form a cluster of radsecproxy servers

## Structure

`SiteSettingMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MistDas` | [`*models.MxedgeDas`](../../doc/models/mxedge-das.md) | Optional | configure cloud-assisted dynamic authorization service on this cluster of mist edges |
| `MistNac` | [`*models.MxclusterNac`](../../doc/models/mxcluster-nac.md) | Optional | - |
| `MistNacedge` | [`*models.MistNacedge`](../../doc/models/mist-nacedge.md) | Optional | - |
| `Radsec` | [`*models.MxclusterRadsec`](../../doc/models/mxcluster-radsec.md) | Optional | MxEdge Radsec Configuration |

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
        "vendor": "cisco-meraki"
      },
      "key1": {
        "require_message_authenticator": false,
        "secret": "secret4",
        "site_id": "0000197c-0000-0000-0000-000000000000",
        "vendor": "cisco-meraki"
      },
      "key2": {
        "require_message_authenticator": false,
        "secret": "secret4",
        "site_id": "0000197c-0000-0000-0000-000000000000",
        "vendor": "cisco-meraki"
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
        "keywrap_enabled": false,
        "keywrap_format": "ascii",
        "keywrap_kek": "keywrap_kek4",
        "keywrap_mack": "keywrap_mack6"
      }
    ],
    "enabled": false,
    "match_ssid": false,
    "nas_ip_source": "tunnel"
  }
}
```

