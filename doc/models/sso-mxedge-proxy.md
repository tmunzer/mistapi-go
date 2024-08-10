
# Sso Mxedge Proxy

if `idp_type`==`mxedge_proxy`, this requires `mist_nac` to be enabled on the mxcluster

## Structure

`SsoMxedgeProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.SsoMxedgeProxyAcctServer`](../../doc/models/sso-mxedge-proxy-acct-server.md) | Optional | - |
| `AuthServers` | [`[]models.SsoMxedgeProxyAuthServer`](../../doc/models/sso-mxedge-proxy-auth-server.md) | Optional | - |
| `MxclusterId` | `*uuid.UUID` | Optional | - |
| `OperatorName` | `*string` | Optional | Operator name as Radius attribute while proxying |
| `ProxyHosts` | `[]string` | Optional | public hostname/IPs |
| `Ssids` | `[]string` | Optional | SSIDs that support eduroam |

## Example (as JSON)

```json
{
  "mxcluster_id": "572586b7-f97b-a22b-526c-8b97a3f609c4",
  "proxy_hosts": [
    "mxedge1.corp.com",
    "63.1.3.5"
  ],
  "ssids": [
    "eduroam_test, eduroam_main"
  ],
  "acct_servers": [
    {
      "host": "host4",
      "port": 254,
      "secret": "secret0"
    },
    {
      "host": "host4",
      "port": 254,
      "secret": "secret0"
    },
    {
      "host": "host4",
      "port": 254,
      "secret": "secret0"
    }
  ],
  "auth_servers": [
    {
      "host": "host0",
      "port": 114,
      "secret": "secret4"
    },
    {
      "host": "host0",
      "port": 114,
      "secret": "secret4"
    },
    {
      "host": "host0",
      "port": 114,
      "secret": "secret4"
    }
  ],
  "operator_name": "operator_name2"
}
```

