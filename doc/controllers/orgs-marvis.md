# Orgs Marvis

```go
orgsMarvis := client.OrgsMarvis()
```

## Class Name

`OrgsMarvis`


# Troubleshoot Org

Troubleshoot sites, devices, clients, and wired clients for maximum of last 7 days from current time. See search APIs for device information:

- [search Device](/#operations/searchOrgDevices)
- [search Wireless Client](/#operations/searchOrgWirelessClients)
- [search Wired Client](/#operations/searchOrgWiredClients)
- [search Wan Client](/#operations/searchOrgWanClients)

**NOTE**: requires Marvis subscription license

```go
TroubleshootOrg(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    siteId *uuid.UUID,
    start *string,
    end *string,
    mType *models.TroubleshootTypeEnum) (
    models.ApiResponse[models.ResponseTroubleshoot],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | **required** when troubleshooting device or a client |
| `siteId` | `*uuid.UUID` | Query, Optional | **required** when troubleshooting site |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `mType` | [`*models.TroubleshootTypeEnum`](../../doc/models/troubleshoot-type-enum.md) | Query, Optional | When troubleshooting site, type of network to troubleshoot |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseTroubleshoot](../../doc/models/response-troubleshoot.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsMarvis.TroubleshootOrg(ctx, orgId, nil, nil, nil, nil, nil)
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
  "end": 1655151856,
  "results": [
    {
      "category": "client",
      "reason": "slow association",
      "recommendation": "Ensure the IP helper-address is configured on the VLAN interface.",
      "text": "Clients of the AP had slow association 8% of the time on Bhavabhi and 5 GHz. ..."
    }
  ],
  "start": 1655065456
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

