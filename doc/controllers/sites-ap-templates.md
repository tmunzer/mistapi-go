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
            "eth1,eth2": {
              "disabled": true,
              "dynamic_vlan": {
                "default_vlan_id": 999,
                "enabled": true
              },
              "port_vlan_id": 1,
              "vlan_id": 9,
              "vlan_ids": "1, 10, 50"
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

