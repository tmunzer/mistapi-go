
# Site Setting Mxedge

site mist edges form a cluster of radsecproxy servers

## Structure

`SiteSettingMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MistDas` | [`*models.MxedgeDas`](../../doc/models/mxedge-das.md) | Optional | configure cloud-assisted dynamic authorization service on this cluster of mist edges |
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
        "secret": "secret2"
      },
      {
        "disable_event_timestamp_check": false,
        "enabled": false,
        "host": "host8",
        "port": 28,
        "secret": "secret2"
      }
    ],
    "enabled": false
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
        "keywrap_format": "hex",
        "keywrap_kek": "keywrap_kek4",
        "keywrap_mack": "keywrap_mack6"
      }
    ],
    "enabled": false,
    "match_ssid": false,
    "proxy_hosts": [
      "proxy_hosts4"
    ]
  }
}
```

