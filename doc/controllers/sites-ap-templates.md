# Sites AP Templates

```go
sitesAPTemplates := client.SitesAPTemplates()
```

## Class Name

`SitesAPTemplates`


# List Site Ap Templates Derived

Get the list of derived AP Templates for a site

```go
ListSiteApTemplatesDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[[]models.ApTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | Whether resolve the site variables |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ApTemplate](../../doc/models/ap-template.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := sitesAPTemplates.ListSiteApTemplatesDerived(ctx, siteId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
[
  {
    "ap_matching": {
      "enabled": true,
      "rules": [
        {
          "match_model": "string",
          "name": "string",
          "port_config": {
            "property1": {
              "additional_vlan_ids": [
                55,
                66
              ],
              "authentication_protocol": "pap",
              "disabled": true,
              "dynamic_vlan": {
                "default_vlan_id": 999,
                "enabled": true,
                "type": "string",
                "vlans": {
                  "1-10": null,
                  "user": null
                }
              },
              "enable_mac_auth": false,
              "forwarding": "all",
              "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
              "mxtunnel_name": "string",
              "port_auth": "none",
              "port_vlan_id": 1,
              "radius_config": {
                "acct_interim_interval": 0,
                "acct_servers": [
                  {
                    "host": "1.2.3.4",
                    "keywrap_enabled": true,
                    "keywrap_format": "hex",
                    "keywrap_kek": "1122334455",
                    "keywrap_mack": "1122334455",
                    "port": 1813,
                    "secret": "testing123"
                  }
                ],
                "auth_servers": [
                  {
                    "host": "1.2.3.4",
                    "keywrap_enabled": true,
                    "keywrap_format": "hex",
                    "keywrap_kek": "1122334455",
                    "keywrap_mack": "1122334455",
                    "port": 1812,
                    "secret": "testing123"
                  }
                ],
                "auth_servers_retries": 3,
                "auth_servers_timeout": 5,
                "coa_enabled": false,
                "coa_port": 3799,
                "network": "string",
                "source_ip": "string"
              },
              "radsec": {
                "enabled": true,
                "idle_timeout": 60,
                "mxcluster_ids": [
                  "572586b7-f97b-a22b-526c-8b97a3f609c4"
                ],
                "proxy_hosts": [
                  "mxedge1.local"
                ],
                "server_name": "radsec.abc.com",
                "servers": [
                  {
                    "host": "1.1.1.1",
                    "port": 1812
                  }
                ],
                "use_mxedge": true,
                "use_site_mxedge": false
              },
              "vlan_id": 9,
              "vlan_ids": [
                1,
                10,
                50
              ],
              "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
              "wxtunnel_remote_id": "wifiguest"
            },
            "property2": {
              "additional_vlan_ids": [
                55,
                66
              ],
              "authentication_protocol": "pap",
              "disabled": true,
              "dynamic_vlan": {
                "default_vlan_id": 999,
                "enabled": true,
                "type": "string",
                "vlans": {
                  "1-10": null,
                  "user": null
                }
              },
              "enable_mac_auth": false,
              "forwarding": "all",
              "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
              "mxtunnel_name": "string",
              "port_auth": "none",
              "port_vlan_id": 1,
              "radius_config": {
                "acct_interim_interval": 0,
                "acct_servers": [
                  {
                    "host": "1.2.3.4",
                    "keywrap_enabled": true,
                    "keywrap_format": "hex",
                    "keywrap_kek": "1122334455",
                    "keywrap_mack": "1122334455",
                    "port": 1813,
                    "secret": "testing123"
                  }
                ],
                "auth_servers": [
                  {
                    "host": "1.2.3.4",
                    "keywrap_enabled": true,
                    "keywrap_format": "hex",
                    "keywrap_kek": "1122334455",
                    "keywrap_mack": "1122334455",
                    "port": 1812,
                    "secret": "testing123"
                  }
                ],
                "auth_servers_retries": 3,
                "auth_servers_timeout": 5,
                "coa_enabled": false,
                "coa_port": 3799,
                "network": "string",
                "source_ip": "string"
              },
              "radsec": {
                "enabled": true,
                "idle_timeout": 60,
                "mxcluster_ids": [
                  "572586b7-f97b-a22b-526c-8b97a3f609c4"
                ],
                "proxy_hosts": [
                  "mxedge1.local"
                ],
                "server_name": "radsec.abc.com",
                "servers": [
                  {
                    "host": "1.1.1.1",
                    "port": 1812
                  }
                ],
                "use_mxedge": true,
                "use_site_mxedge": false
              },
              "vlan_id": 9,
              "vlan_ids": [
                1,
                10,
                50
              ],
              "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
              "wxtunnel_remote_id": "wifiguest"
            }
          }
        }
      ]
    },
    "created_time": 0,
    "for_site": true,
    "id": "497f6eca-6276-4993-bfeb-53cbbbba9f08",
    "modified_time": 0,
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
    "wifi": {
      "cisco_enabled": true,
      "disable_11k": false,
      "disable_radios_when_power_constrained": true,
      "enable_arp_spoof": true,
      "enable_shared_radio_scanning": false,
      "enabled": true,
      "locate_connected": false,
      "locate_unconnected": false,
      "mesh_allow_dfs": false,
      "mesh_enable_crm": true,
      "mesh_enabled": true,
      "proxy_arp": false
    }
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

