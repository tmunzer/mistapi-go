
# Response Org Site Sle

Organization site SLE response for Wi-Fi, wired, or WAN

## Class Name

`ResponseOrgSiteSle`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.OrgSiteSleWifi`](../../../doc/models/org-site-sle-wifi.md) | models.ResponseOrgSiteSleContainer.FromOrgSiteSleWifi(models.OrgSiteSleWifi orgSiteSleWifi) |
| [`models.OrgSiteWiredWifi`](../../../doc/models/org-site-wired-wifi.md) | models.ResponseOrgSiteSleContainer.FromOrgSiteWiredWifi(models.OrgSiteWiredWifi orgSiteWiredWifi) |
| [`models.OrgSiteWanWifi`](../../../doc/models/org-site-wan-wifi.md) | models.ResponseOrgSiteSleContainer.FromOrgSiteWanWifi(models.OrgSiteWanWifi orgSiteWanWifi) |

## models.OrgSiteSleWifi

### Initialization Code

#### Example

```go
value := models.ResponseOrgSiteSleContainer.FromOrgSiteSleWifi(models.OrgSiteSleWifi{
    End:                  float64(108.82),
    Interval:             40,
    Limit:                216,
    Page:                 154,
    Results:              []models.OrgSiteSleWifiResult{
        models.OrgSiteSleWifiResult{
            ApAvailability:       float64(204.4),
            SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        },
    },
    Start:                float64(64.88),
    Total:                122,
})
```

## models.OrgSiteWiredWifi

### Initialization Code

#### Example

```go
value := models.ResponseOrgSiteSleContainer.FromOrgSiteWiredWifi(models.OrgSiteWiredWifi{
    End:                  float64(152.44),
    Interval:             50,
    Limit:                226,
    Page:                 112,
    Results:              []models.OrgSiteSleWiredResult{
        models.OrgSiteSleWiredResult{
            SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
            SwitchHealth:         float64(218.74),
        },
    },
    Start:                float64(108.5),
    Total:                132,
})
```

## models.OrgSiteWanWifi

### Initialization Code

#### Example

```go
value := models.ResponseOrgSiteSleContainer.FromOrgSiteWanWifi(models.OrgSiteWanWifi{
    End:                  float64(34.8),
    Interval:             62,
    Limit:                18,
    Page:                 132,
    Results:              []models.OrgSiteSleWanResult{
        models.OrgSiteSleWanResult{
            GatewayHealth:        float64(241.74),
            SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        },
    },
    Start:                float64(246.86),
    Total:                112,
})
```

