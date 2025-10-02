
# Sso Mxedge Proxy

If `idp_type`==`mxedge_proxy`, this requires `mist_nac` to be enabled on the mxcluster

*This model accepts additional fields of type interface{}.*

## Structure

`SsoMxedgeProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.SsoMxedgeProxyAcctServer`](../../doc/models/sso-mxedge-proxy-acct-server.md) | Optional | - |
| `AuthServers` | [`[]models.SsoMxedgeProxyAuthServer`](../../doc/models/sso-mxedge-proxy-auth-server.md) | Optional | - |
| `MxclusterId` | `*uuid.UUID` | Optional | - |
| `OperatorName` | `*string` | Optional | Operator name as Radius attribute while proxying |
| `ProxyHosts` | `[]string` | Optional | Public hostname/IPs |
| `Ssids` | `[]string` | Optional | SSIDs that support eduroam |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
      "secret": "secret0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "host": "host4",
      "port": 254,
      "secret": "secret0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "host": "host4",
      "port": 254,
      "secret": "secret0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "auth_servers": [
    {
      "host": "host0",
      "port": 114,
      "require_message_authenticator": false,
      "retry": 126,
      "secret": "secret4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "host": "host0",
      "port": 114,
      "require_message_authenticator": false,
      "retry": 126,
      "secret": "secret4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "host": "host0",
      "port": 114,
      "require_message_authenticator": false,
      "retry": 126,
      "secret": "secret4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "operator_name": "operator_name2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

