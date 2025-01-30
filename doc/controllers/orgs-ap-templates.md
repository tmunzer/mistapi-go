# Orgs AP Templates

```go
orgsAPTemplates := client.OrgsAPTemplates()
```

## Class Name

`OrgsAPTemplates`

## Methods

* [Create Org Aptemplate](../../doc/controllers/orgs-ap-templates.md#create-org-aptemplate)
* [Delete Org Aptemplate](../../doc/controllers/orgs-ap-templates.md#delete-org-aptemplate)
* [Get Org Aptemplate](../../doc/controllers/orgs-ap-templates.md#get-org-aptemplate)
* [List Org Aptemplates](../../doc/controllers/orgs-ap-templates.md#list-org-aptemplates)
* [Update Org Aptemplate](../../doc/controllers/orgs-ap-templates.md#update-org-aptemplate)


# Create Org Aptemplate

Create Org AP Template

```go
CreateOrgAptemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.ApTemplate) (
    models.ApiResponse[models.ApTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.ApTemplate`](../../doc/models/ap-template.md) | Body, Optional | - |

## Response Type

[`models.ApTemplate`](../../doc/models/ap-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.ApTemplate{
    ApMatching:           models.ApTemplateMatching{
        Enabled:              models.ToPointer(true),
        Rules:                []models.ApTemplateMatchingRule{
            models.ApTemplateMatchingRule{
                MatchModel:           models.ToPointer("string"),
                Name:                 models.ToPointer("string"),
                PortConfig:           map[string]models.ApPortConfig{
                    "property1": models.ApPortConfig{
                        Disabled:             models.ToPointer(true),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(999),
                            Enabled:              models.ToPointer(true),
                            Type:                 models.ToPointer("string"),
                            Vlans:                map[string]string{
                                "1-10": "vlans1",
                                "user": "vlans2",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_ALL),
                        MxTunnelId:           models.ToPointer(uuid.MustParse("08cd7499-5841-51c8-e663-fb16b6f3b45e")),
                        MxtunnelName:         models.ToPointer("string"),
                        PortAuth:             models.ToPointer(models.ApPortConfigPortAuthEnum_NONE),
                        PortVlanId:           models.ToPointer(1),
                        RadiusConfig:         models.ToPointer(models.RadiusConfig{
                            AcctInterimInterval:  models.ToPointer(0),
                            AcctServers:          []models.RadiusAcctServer{
                                models.RadiusAcctServer{
                                    Host:                 "1.2.3.4",
                                    KeywrapEnabled:       models.ToPointer(true),
                                    KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                                    KeywrapKek:           models.ToPointer("1122334455"),
                                    KeywrapMack:          models.ToPointer("1122334455"),
                                    Port:                 models.ToPointer(1813),
                                    Secret:               "testing123",
                                },
                            },
                            AuthServers:          []models.RadiusAuthServer{
                                models.RadiusAuthServer{
                                    Host:                        "1.2.3.4",
                                    KeywrapEnabled:              models.ToPointer(true),
                                    KeywrapFormat:               models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                                    KeywrapKek:                  models.ToPointer("1122334455"),
                                    KeywrapMack:                 models.ToPointer("1122334455"),
                                    Port:                        models.ToPointer(1812),
                                    Secret:                      "testing123",
                                },
                            },
                            AuthServersRetries:   models.ToPointer(3),
                            AuthServersTimeout:   models.ToPointer(5),
                            CoaEnabled:           models.ToPointer(false),
                            CoaPort:              models.ToPointer(3799),
                            Network:              models.ToPointer("string"),
                            SourceIp:             models.ToPointer("string"),
                        }),
                        Radsec:               models.ToPointer(models.Radsec{
                            Enabled:              models.ToPointer(true),
                            IdleTimeout:          models.ToPointer(60),
                            MxclusterIds:         []uuid.UUID{
                                uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4"),
                            },
                            ProxyHosts:           []string{
                                "mxedge1.local",
                            },
                            ServerName:           models.ToPointer("radsec.abc.com"),
                            Servers:              []models.RadsecServer{
                                models.RadsecServer{
                                    Host:                 models.ToPointer("1.1.1.1"),
                                    Port:                 models.ToPointer(1812),
                                },
                            },
                            UseMxedge:            models.ToPointer(true),
                            UseSiteMxedge:        models.ToPointer(false),
                        }),
                        VlanId:               models.ToPointer(9),
                        VlandIds:             []int{
                            1,
                            10,
                            50,
                        },
                        WxtunnelId:           models.ToPointer(uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")),
                        WxtunnelRemoteId:     models.ToPointer("wifiguest"),
                        AdditionalProperties: map[string]interface{}{
                            "additional_vlan_ids": interface{}("5566"),
                            "authentication_protocol": interface{}("pap"),
                        },
                    },
                    "property2": models.ApPortConfig{
                        Disabled:             models.ToPointer(true),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(999),
                            Enabled:              models.ToPointer(true),
                            Type:                 models.ToPointer("string"),
                            Vlans:                map[string]string{
                                "1-10": "vlans1",
                                "user": "vlans2",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_ALL),
                        MxTunnelId:           models.ToPointer(uuid.MustParse("08cd7499-5841-51c8-e663-fb16b6f3b45e")),
                        MxtunnelName:         models.ToPointer("string"),
                        PortAuth:             models.ToPointer(models.ApPortConfigPortAuthEnum_NONE),
                        PortVlanId:           models.ToPointer(1),
                        RadiusConfig:         models.ToPointer(models.RadiusConfig{
                            AcctInterimInterval:  models.ToPointer(0),
                            AcctServers:          []models.RadiusAcctServer{
                                models.RadiusAcctServer{
                                    Host:                 "1.2.3.4",
                                    KeywrapEnabled:       models.ToPointer(true),
                                    KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                                    KeywrapKek:           models.ToPointer("1122334455"),
                                    KeywrapMack:          models.ToPointer("1122334455"),
                                    Port:                 models.ToPointer(1813),
                                    Secret:               "testing123",
                                },
                            },
                            AuthServers:          []models.RadiusAuthServer{
                                models.RadiusAuthServer{
                                    Host:                        "1.2.3.4",
                                    KeywrapEnabled:              models.ToPointer(true),
                                    KeywrapFormat:               models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                                    KeywrapKek:                  models.ToPointer("1122334455"),
                                    KeywrapMack:                 models.ToPointer("1122334455"),
                                    Port:                        models.ToPointer(1812),
                                    Secret:                      "testing123",
                                },
                            },
                            AuthServersRetries:   models.ToPointer(3),
                            AuthServersTimeout:   models.ToPointer(5),
                            CoaEnabled:           models.ToPointer(false),
                            CoaPort:              models.ToPointer(3799),
                            Network:              models.ToPointer("string"),
                            SourceIp:             models.ToPointer("string"),
                        }),
                        Radsec:               models.ToPointer(models.Radsec{
                            Enabled:              models.ToPointer(true),
                            IdleTimeout:          models.ToPointer(60),
                            MxclusterIds:         []uuid.UUID{
                                uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4"),
                            },
                            ProxyHosts:           []string{
                                "mxedge1.local",
                            },
                            ServerName:           models.ToPointer("radsec.abc.com"),
                            Servers:              []models.RadsecServer{
                                models.RadsecServer{
                                    Host:                 models.ToPointer("1.1.1.1"),
                                    Port:                 models.ToPointer(1812),
                                },
                            },
                            UseMxedge:            models.ToPointer(true),
                            UseSiteMxedge:        models.ToPointer(false),
                        }),
                        VlanId:               models.ToPointer(9),
                        VlandIds:             []int{
                            1,
                            10,
                            50,
                        },
                        WxtunnelId:           models.ToPointer(uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")),
                        WxtunnelRemoteId:     models.ToPointer("wifiguest"),
                        AdditionalProperties: map[string]interface{}{
                            "additional_vlan_ids": interface{}("5566"),
                            "authentication_protocol": interface{}("pap"),
                        },
                    },
                },
            },
        },
    },
    Wifi:                 models.ToPointer(models.ApTemplateWifi{
        CiscoEnabled:                      models.ToPointer(true),
        Disable11k:                        models.ToPointer(false),
        DisableRadiosWhenPowerConstrained: models.ToPointer(true),
        EnableArpSpoof:                    models.ToPointer(true),
        EnableSharedRadioScanning:         models.ToPointer(false),
        Enabled:                           models.ToPointer(true),
        LocateConnected:                   models.ToPointer(false),
        LocateUnconnected:                 models.ToPointer(false),
        MeshAllowDfs:                      models.ToPointer(false),
        MeshEnableCrm:                     models.ToPointer(true),
        MeshEnabled:                       models.ToPointer(true),
        ProxyArp:                          models.ToPointer(false),
    }),
}

apiResponse, err := orgsAPTemplates.CreateOrgAptemplate(ctx, orgId, &body)
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
            "vland_ids": [
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
            "vland_ids": [
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
  "id": "497f6eca-6276-4993-bfeb-53cbbbba8f08",
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
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Org Aptemplate

Delete existing AP Template

```go
DeleteOrgAptemplate(
    ctx context.Context,
    orgId uuid.UUID,
    aptemplateId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `aptemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

aptemplateId := uuid.MustParse("00001752-0000-0000-0000-000000000000")

resp, err := orgsAPTemplates.DeleteOrgAptemplate(ctx, orgId, aptemplateId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org Aptemplate

Get AP Template

```go
GetOrgAptemplate(
    ctx context.Context,
    orgId uuid.UUID,
    aptemplateId uuid.UUID) (
    models.ApiResponse[models.ApTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `aptemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ApTemplate`](../../doc/models/ap-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

aptemplateId := uuid.MustParse("00001752-0000-0000-0000-000000000000")

apiResponse, err := orgsAPTemplates.GetOrgAptemplate(ctx, orgId, aptemplateId)
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
            "vland_ids": [
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
            "vland_ids": [
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
  "id": "497f6eca-6276-4993-bfeb-53cbbbba8f08",
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
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Aptemplates

Get List of Org AP Templates

```go
ListOrgAptemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.ApTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.ApTemplate`](../../doc/models/ap-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsAPTemplates.ListOrgAptemplates(ctx, orgId, &limit, &page)
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
              "vland_ids": [
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
              "vland_ids": [
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


# Update Org Aptemplate

Update AP Template

```go
UpdateOrgAptemplate(
    ctx context.Context,
    orgId uuid.UUID,
    aptemplateId uuid.UUID,
    body *models.ApTemplate) (
    models.ApiResponse[models.ApTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `aptemplateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.ApTemplate`](../../doc/models/ap-template.md) | Body, Optional | - |

## Response Type

[`models.ApTemplate`](../../doc/models/ap-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

aptemplateId := uuid.MustParse("00001752-0000-0000-0000-000000000000")

body := models.ApTemplate{
    ApMatching:           models.ApTemplateMatching{
        Enabled:              models.ToPointer(true),
        Rules:                []models.ApTemplateMatchingRule{
            models.ApTemplateMatchingRule{
                MatchModel:           models.ToPointer("string"),
                Name:                 models.ToPointer("string"),
                PortConfig:           map[string]models.ApPortConfig{
                    "property1": models.ApPortConfig{
                        Disabled:             models.ToPointer(true),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(999),
                            Enabled:              models.ToPointer(true),
                            Type:                 models.ToPointer("string"),
                            Vlans:                map[string]string{
                                "1-10": "vlans1",
                                "user": "vlans2",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_ALL),
                        MxTunnelId:           models.ToPointer(uuid.MustParse("08cd7499-5841-51c8-e663-fb16b6f3b45e")),
                        MxtunnelName:         models.ToPointer("string"),
                        PortAuth:             models.ToPointer(models.ApPortConfigPortAuthEnum_NONE),
                        PortVlanId:           models.ToPointer(1),
                        RadiusConfig:         models.ToPointer(models.RadiusConfig{
                            AcctInterimInterval:  models.ToPointer(0),
                            AcctServers:          []models.RadiusAcctServer{
                                models.RadiusAcctServer{
                                    Host:                 "1.2.3.4",
                                    KeywrapEnabled:       models.ToPointer(true),
                                    KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                                    KeywrapKek:           models.ToPointer("1122334455"),
                                    KeywrapMack:          models.ToPointer("1122334455"),
                                    Port:                 models.ToPointer(1813),
                                    Secret:               "testing123",
                                },
                            },
                            AuthServers:          []models.RadiusAuthServer{
                                models.RadiusAuthServer{
                                    Host:                        "1.2.3.4",
                                    KeywrapEnabled:              models.ToPointer(true),
                                    KeywrapFormat:               models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                                    KeywrapKek:                  models.ToPointer("1122334455"),
                                    KeywrapMack:                 models.ToPointer("1122334455"),
                                    Port:                        models.ToPointer(1812),
                                    Secret:                      "testing123",
                                },
                            },
                            AuthServersRetries:   models.ToPointer(3),
                            AuthServersTimeout:   models.ToPointer(5),
                            CoaEnabled:           models.ToPointer(false),
                            CoaPort:              models.ToPointer(3799),
                            Network:              models.ToPointer("string"),
                            SourceIp:             models.ToPointer("string"),
                        }),
                        Radsec:               models.ToPointer(models.Radsec{
                            Enabled:              models.ToPointer(true),
                            IdleTimeout:          models.ToPointer(60),
                            MxclusterIds:         []uuid.UUID{
                                uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4"),
                            },
                            ProxyHosts:           []string{
                                "mxedge1.local",
                            },
                            ServerName:           models.ToPointer("radsec.abc.com"),
                            Servers:              []models.RadsecServer{
                                models.RadsecServer{
                                    Host:                 models.ToPointer("1.1.1.1"),
                                    Port:                 models.ToPointer(1812),
                                },
                            },
                            UseMxedge:            models.ToPointer(true),
                            UseSiteMxedge:        models.ToPointer(false),
                        }),
                        VlanId:               models.ToPointer(9),
                        VlandIds:             []int{
                            1,
                            10,
                            50,
                        },
                        WxtunnelId:           models.ToPointer(uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")),
                        WxtunnelRemoteId:     models.ToPointer("wifiguest"),
                        AdditionalProperties: map[string]interface{}{
                            "additional_vlan_ids": interface{}("5566"),
                            "authentication_protocol": interface{}("pap"),
                        },
                    },
                    "property2": models.ApPortConfig{
                        Disabled:             models.ToPointer(true),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(999),
                            Enabled:              models.ToPointer(true),
                            Type:                 models.ToPointer("string"),
                            Vlans:                map[string]string{
                                "1-10": "vlans1",
                                "user": "vlans2",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_ALL),
                        MxTunnelId:           models.ToPointer(uuid.MustParse("08cd7499-5841-51c8-e663-fb16b6f3b45e")),
                        MxtunnelName:         models.ToPointer("string"),
                        PortAuth:             models.ToPointer(models.ApPortConfigPortAuthEnum_NONE),
                        PortVlanId:           models.ToPointer(1),
                        RadiusConfig:         models.ToPointer(models.RadiusConfig{
                            AcctInterimInterval:  models.ToPointer(0),
                            AcctServers:          []models.RadiusAcctServer{
                                models.RadiusAcctServer{
                                    Host:                 "1.2.3.4",
                                    KeywrapEnabled:       models.ToPointer(true),
                                    KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                                    KeywrapKek:           models.ToPointer("1122334455"),
                                    KeywrapMack:          models.ToPointer("1122334455"),
                                    Port:                 models.ToPointer(1813),
                                    Secret:               "testing123",
                                },
                            },
                            AuthServers:          []models.RadiusAuthServer{
                                models.RadiusAuthServer{
                                    Host:                        "1.2.3.4",
                                    KeywrapEnabled:              models.ToPointer(true),
                                    KeywrapFormat:               models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                                    KeywrapKek:                  models.ToPointer("1122334455"),
                                    KeywrapMack:                 models.ToPointer("1122334455"),
                                    Port:                        models.ToPointer(1812),
                                    Secret:                      "testing123",
                                },
                            },
                            AuthServersRetries:   models.ToPointer(3),
                            AuthServersTimeout:   models.ToPointer(5),
                            CoaEnabled:           models.ToPointer(false),
                            CoaPort:              models.ToPointer(3799),
                            Network:              models.ToPointer("string"),
                            SourceIp:             models.ToPointer("string"),
                        }),
                        Radsec:               models.ToPointer(models.Radsec{
                            Enabled:              models.ToPointer(true),
                            IdleTimeout:          models.ToPointer(60),
                            MxclusterIds:         []uuid.UUID{
                                uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4"),
                            },
                            ProxyHosts:           []string{
                                "mxedge1.local",
                            },
                            ServerName:           models.ToPointer("radsec.abc.com"),
                            Servers:              []models.RadsecServer{
                                models.RadsecServer{
                                    Host:                 models.ToPointer("1.1.1.1"),
                                    Port:                 models.ToPointer(1812),
                                },
                            },
                            UseMxedge:            models.ToPointer(true),
                            UseSiteMxedge:        models.ToPointer(false),
                        }),
                        VlanId:               models.ToPointer(9),
                        VlandIds:             []int{
                            1,
                            10,
                            50,
                        },
                        WxtunnelId:           models.ToPointer(uuid.MustParse("7dae216d-7c98-a51b-e068-dd7d477b7216")),
                        WxtunnelRemoteId:     models.ToPointer("wifiguest"),
                        AdditionalProperties: map[string]interface{}{
                            "additional_vlan_ids": interface{}("5566"),
                            "authentication_protocol": interface{}("pap"),
                        },
                    },
                },
            },
        },
    },
    Wifi:                 models.ToPointer(models.ApTemplateWifi{
        CiscoEnabled:                      models.ToPointer(true),
        Disable11k:                        models.ToPointer(false),
        DisableRadiosWhenPowerConstrained: models.ToPointer(true),
        EnableArpSpoof:                    models.ToPointer(true),
        EnableSharedRadioScanning:         models.ToPointer(false),
        Enabled:                           models.ToPointer(true),
        LocateConnected:                   models.ToPointer(false),
        LocateUnconnected:                 models.ToPointer(false),
        MeshAllowDfs:                      models.ToPointer(false),
        MeshEnableCrm:                     models.ToPointer(true),
        MeshEnabled:                       models.ToPointer(true),
        ProxyArp:                          models.ToPointer(false),
    }),
}

apiResponse, err := orgsAPTemplates.UpdateOrgAptemplate(ctx, orgId, aptemplateId, &body)
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
            "vland_ids": [
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
            "vland_ids": [
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
  "id": "497f6eca-6276-4993-bfeb-53cbbbba8f08",
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
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

