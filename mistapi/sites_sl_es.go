package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// SitesSLEs represents a controller struct.
type SitesSLEs struct {
    baseController
}

// NewSitesSLEs creates a new instance of SitesSLEs.
// It takes a baseController as a parameter and returns a pointer to the SitesSLEs.
func NewSitesSLEs(baseController baseController) *SitesSLEs {
    sitesSLEs := SitesSLEs{baseController: baseController}
    return &sitesSLEs
}

// GetSiteSleClassifierDetails takes context, siteId, scope, scopeId, metric, classifier, start, end, duration as parameters and
// returns an models.ApiResponse with models.SleClassifierSummary data and
// an error if there was an issue with the request or response.
// Get SLE classifier details
func (s *SitesSLEs) GetSiteSleClassifierDetails(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SleSummaryScopeEnum,
    scopeId string,
    metric string,
    classifier string,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SleClassifierSummary],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/classifier/%v/summary", siteId, scope, scopeId, metric, classifier),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result models.SleClassifierSummary
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleClassifierSummary](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleMetricClassifiers takes context, siteId, scope, scopeId, metric as parameters and
// returns an models.ApiResponse with []string data and
// an error if there was an issue with the request or response.
// Get the list of classifiers for a specific metric
func (s *SitesSLEs) GetSiteSleMetricClassifiers(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleMetricClassifiersScopeParametersEnum,
    scopeId string,
    metric string) (
    models.ApiResponse[[]string],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/classifiers", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    
    var result []string
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]string](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleHistogram takes context, siteId, scope, scopeId, metric, start, end, duration as parameters and
// returns an models.ApiResponse with models.SleHistogram data and
// an error if there was an issue with the request or response.
// Get the histogram for the SLE metric
func (s *SitesSLEs) GetSiteSleHistogram(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleHistogramScopeParametersEnum,
    scopeId string,
    metric string,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SleHistogram],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/histogram", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result models.SleHistogram
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleHistogram](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleImpactSummary takes context, siteId, scope, scopeId, metric, start, end, duration, fields, classifier as parameters and
// returns an models.ApiResponse with models.SleImpactSummary data and
// an error if there was an issue with the request or response.
// Get impact summary counts optionally filtered by classifier and failure type
// 
// * Wireless SLE Fields: `wlan`, `device_type`, `device_os` ,`band`, `ap`, `server`, `mxedge`
// * Wired SLE Fields: `switch`, `client`, `vlan`, `interface`, `chassis`
// * WAN SLE Fields: `gateway`, `client`, `interface`, `chassis`, `peer_path`, `gateway_zones`
func (s *SitesSLEs) GetSiteSleImpactSummary(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactSummaryScopeParametersEnum,
    scopeId string,
    metric string,
    start *int,
    end *int,
    duration *string,
    fields *models.SiteSleImpactSummaryFieldsParameterEnum,
    classifier *string) (
    models.ApiResponse[models.SleImpactSummary],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/impact-summary", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if fields != nil {
        req.QueryParam("fields", *fields)
    }
    if classifier != nil {
        req.QueryParam("classifier", *classifier)
    }
    
    var result models.SleImpactSummary
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleImpactSummary](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleImpactedApplications takes context, siteId, scope, scopeId, metric, start, end, duration, classifier as parameters and
// returns an models.ApiResponse with models.SleImpactedApplications data and
// an error if there was an issue with the request or response.
// For WAN SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type
func (s *SitesSLEs) GetSiteSleImpactedApplications(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleScopeEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedApplications],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/impacted-applications", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if classifier != nil {
        req.QueryParam("classifier", *classifier)
    }
    
    var result models.SleImpactedApplications
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleImpactedApplications](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleImpactedAps takes context, siteId, scope, scopeId, metric, start, end, duration, classifier as parameters and
// returns an models.ApiResponse with models.SleImpactedAps data and
// an error if there was an issue with the request or response.
// For Wireless SLEs. Get list of impacted APs optionally filtered by classifier and failure type
func (s *SitesSLEs) GetSiteSleImpactedAps(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedApsScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedAps],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/impacted-aps", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if classifier != nil {
        req.QueryParam("classifier", *classifier)
    }
    
    var result models.SleImpactedAps
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleImpactedAps](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleImpactedChassis takes context, siteId, scope, scopeId, metric, start, end, duration, classifier as parameters and
// returns an models.ApiResponse with models.SleImpactedChassis data and
// an error if there was an issue with the request or response.
// For Wired and WAN SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type
func (s *SitesSLEs) GetSiteSleImpactedChassis(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedChassisScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedChassis],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/impacted-chassis", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if classifier != nil {
        req.QueryParam("classifier", *classifier)
    }
    
    var result models.SleImpactedChassis
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleImpactedChassis](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleImpactedWiredClients takes context, siteId, scope, scopeId, metric, start, end, duration, classifier as parameters and
// returns an models.ApiResponse with models.SleImpactedClients data and
// an error if there was an issue with the request or response.
// For Wired SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type
func (s *SitesSLEs) GetSiteSleImpactedWiredClients(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedClientsScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedClients],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/impacted-clients", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if classifier != nil {
        req.QueryParam("classifier", *classifier)
    }
    
    var result models.SleImpactedClients
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleImpactedClients](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleImpactedGateways takes context, siteId, scope, scopeId, metric, start, end, duration, classifier as parameters and
// returns an models.ApiResponse with models.SleImpactedGateways data and
// an error if there was an issue with the request or response.
// For WAN SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type
func (s *SitesSLEs) GetSiteSleImpactedGateways(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedGatewaysScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedGateways],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/impacted-gateways", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if classifier != nil {
        req.QueryParam("classifier", *classifier)
    }
    
    var result models.SleImpactedGateways
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleImpactedGateways](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleImpactedInterfaces takes context, siteId, scope, scopeId, metric, start, end, duration, classifier as parameters and
// returns an models.ApiResponse with models.SleImpactedInterfaces data and
// an error if there was an issue with the request or response.
// For Wired and WAN SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type
func (s *SitesSLEs) GetSiteSleImpactedInterfaces(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedInterfacesScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedInterfaces],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/impacted-interfaces", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if classifier != nil {
        req.QueryParam("classifier", *classifier)
    }
    
    var result models.SleImpactedInterfaces
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleImpactedInterfaces](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleImpactedSwitches takes context, siteId, scope, scopeId, metric, start, end, duration, classifier as parameters and
// returns an models.ApiResponse with models.SleImpactedSwitches data and
// an error if there was an issue with the request or response.
// For Wired SLEs. Get list of impacted switches optionally filtered by classifier and failure type
func (s *SitesSLEs) GetSiteSleImpactedSwitches(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedSwitchesScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedSwitches],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/impacted-switches", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if classifier != nil {
        req.QueryParam("classifier", *classifier)
    }
    
    var result models.SleImpactedSwitches
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleImpactedSwitches](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleImpactedWirelessClients takes context, siteId, scope, scopeId, metric, start, end, duration, classifier as parameters and
// returns an models.ApiResponse with models.SleImpactedUsers data and
// an error if there was an issue with the request or response.
// For Wireless SLEs. Get list of impacted wireless users optionally filtered by classifier and failure type
func (s *SitesSLEs) GetSiteSleImpactedWirelessClients(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedUsersScopeParameterEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedUsers],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/impacted_users", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if classifier != nil {
        req.QueryParam("classifier", *classifier)
    }
    
    var result models.SleImpactedUsers
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleImpactedUsers](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleSummary takes context, siteId, scope, scopeId, metric, start, end, duration as parameters and
// returns an models.ApiResponse with models.SleSummary data and
// an error if there was an issue with the request or response.
// Get the summary for the SLE metric
func (s *SitesSLEs) GetSiteSleSummary(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleMetricSummaryScopeParametersEnum,
    scopeId string,
    metric string,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SleSummary],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/summary", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result models.SleSummary
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleSummary](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSleThreshold takes context, siteId, scope, scopeId, metric as parameters and
// returns an models.ApiResponse with models.SleThreshold data and
// an error if there was an issue with the request or response.
// Get the SLE threshold
func (s *SitesSLEs) GetSiteSleThreshold(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleThresholdScopeParameterEnum,
    scopeId string,
    metric string) (
    models.ApiResponse[models.SleThreshold],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/threshold", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    
    var result models.SleThreshold
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleThreshold](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReplaceSiteSleThreshold takes context, siteId, scope, scopeId, metric, body as parameters and
// returns an models.ApiResponse with models.SleThreshold data and
// an error if there was an issue with the request or response.
// Replace the SLE threshold
func (s *SitesSLEs) ReplaceSiteSleThreshold(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleThresholdScopeParameterEnum,
    scopeId string,
    metric string,
    body *models.SleThreshold) (
    models.ApiResponse[models.SleThreshold],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/threshold", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SleThreshold
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleThreshold](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteSleThreshold takes context, siteId, scope, scopeId, metric, body as parameters and
// returns an models.ApiResponse with models.SleThreshold data and
// an error if there was an issue with the request or response.
// Update the SLE threshold
func (s *SitesSLEs) UpdateSiteSleThreshold(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleThresholdScopeParameterEnum,
    scopeId string,
    metric string,
    body *models.SleThreshold) (
    models.ApiResponse[models.SleThreshold],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metric/%v/threshold", siteId, scope, scopeId, metric),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SleThreshold
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SleThreshold](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSlesMetrics takes context, siteId, scope, scopeId as parameters and
// returns an models.ApiResponse with models.SiteSleMetrics data and
// an error if there was an issue with the request or response.
// Get the list of metrics for the given scope
func (s *SitesSLEs) GetSiteSlesMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleMetricsScopeParametersEnum,
    scopeId string) (
    models.ApiResponse[models.SiteSleMetrics],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/sle/%v/%v/metrics", siteId, scope, scopeId),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    
    var result models.SiteSleMetrics
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteSleMetrics](decoder)
    return models.NewApiResponse(result, resp), err
}
