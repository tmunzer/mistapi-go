
# Getting Started with Mist API

## Introduction

> Version: **2412.1.16**
> 
> Date: **January 6, 2024**

---


### Additional Documentation

* [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html)
* [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html)
* [Mist Product Updates](https://www.juniper.net/documentation/us/en/software/mist/product-updates/)

---


### Helpful Resources

* [API Sandbox and Exercises](https://api-class.mist.com/)
* [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace)
* [API Demo Apps](https://apps.mist-lab.fr/)
* [Juniper Blog](https://blogs.juniper.net/)

---


### Requirements

The SDK requires **Go version 1.18 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## Installation

The following section explains how to use the mistapi library in a new project.

### 1. Add SDK as a Dependency to the Application

- Add the following lines to your application's `go.mod` file:

```go
replace mistapi => ".\\mist-api-go_generic_lib" // local path to the SDK

require mistapi v0.0.0
```

- Resolve the dependencies in the updated `go.mod` file, using the `go get` command.

## Test the SDK

`Go Testing` is used as the testing framework. To run the unit tests of the SDK, navigate to the root directory of the SDK and run the following command in the terminal:

```bash
$ go test
```

## Initialize the API Client

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `environment` | `Environment` | The API environment. <br> **Default: `Environment.MIST_GLOBAL_01`** |
| `httpConfiguration` | [`HttpConfiguration`](doc/http-configuration.md) | Configurable http client options like timeout and retries. |
| `loggerConfiguration` | [`LoggerConfiguration`](doc/logger-configuration.md) | Represents the logger configurations for API calls |
| `apiTokenCredentials` | [`ApiTokenCredentials`](doc/auth/custom-header-signature.md) | The Credentials Setter for Custom Header Signature |
| `basicAuthCredentials` | [`BasicAuthCredentials`](doc/auth/basic-authentication.md) | The Credentials Setter for Basic Authentication |
| `csrfTokenCredentials` | [`CsrfTokenCredentials`](doc/auth/custom-header-signature-1.md) | The Credentials Setter for Custom Header Signature |

The API client can be initialized as follows:

```go
client := mistapi.NewClient(
    mistapi.CreateConfiguration(
        mistapi.WithHttpConfiguration(
            mistapi.CreateHttpConfiguration(
                mistapi.WithTimeout(0),
            ),
        ),
        mistapi.WithEnvironment(mistapi.MIST_GLOBAL_01),
        mistapi.WithApiTokenCredentials(
            mistapi.NewApiTokenCredentials("Authorization"),
        ),
        mistapi.WithBasicAuthCredentials(
            mistapi.NewBasicAuthCredentials(
                "Username",
                "Password",
            ),
        ),
        mistapi.WithCsrfTokenCredentials(
            mistapi.NewCsrfTokenCredentials("X-CSRFToken"),
        ),
        mistapi.WithLoggerConfiguration(
            mistapi.WithLevel("info"),
            mistapi.WithRequestConfiguration(
                mistapi.WithRequestBody(true),
            ),
            mistapi.WithResponseConfiguration(
                mistapi.WithResponseHeaders(true),
            ),
        ),
    ),
)
```

## Environments

The SDK can be configured to use a different environment for making API calls. Available environments are:

### Fields

| Name | Description |
|  --- | --- |
| Mist Global 01 | **Default** |
| Mist Global 02 | - |
| Mist Global 03 | - |
| Mist Global 04 | - |
| Mist EMEA 01 | - |
| Mist EMEA 02 | - |
| Mist EMEA 03 | - |
| Mist APAC 01 | - |

## Authorization

This API uses the following authentication schemes.

* [`apiToken (Custom Header Signature)`](doc/auth/custom-header-signature.md)
* [`basicAuth (Basic Authentication)`](doc/auth/basic-authentication.md)
* [`csrfToken (Custom Header Signature)`](doc/auth/custom-header-signature-1.md)

## List of APIs

* [Admins](doc/controllers/admins.md)
* [Admins Login](doc/controllers/admins-login.md)
* [Admins Login-O Auth 2](doc/controllers/admins-login-o-auth-2.md)
* [Admins Logout](doc/controllers/admins-logout.md)
* [Admins Lookup](doc/controllers/admins-lookup.md)
* [Admins Recover Password](doc/controllers/admins-recover-password.md)
* [Constants Definitions](doc/controllers/constants-definitions.md)
* [Constants Events](doc/controllers/constants-events.md)
* [Constants Models](doc/controllers/constants-models.md)
* [Installer](doc/controllers/installer.md)
* [MS Ps](doc/controllers/ms-ps.md)
* [MS Ps Admins](doc/controllers/ms-ps-admins.md)
* [MS Ps Inventory](doc/controllers/ms-ps-inventory.md)
* [MS Ps Licenses](doc/controllers/ms-ps-licenses.md)
* [MS Ps Logo](doc/controllers/ms-ps-logo.md)
* [MS Ps Logs](doc/controllers/ms-ps-logs.md)
* [MS Ps Marvis](doc/controllers/ms-ps-marvis.md)
* [MS Ps Org Groups](doc/controllers/ms-ps-org-groups.md)
* [MS Ps Orgs](doc/controllers/ms-ps-orgs.md)
* [MS Ps SL Es](doc/controllers/ms-ps-sl-es.md)
* [MS Ps SSO](doc/controllers/ms-ps-sso.md)
* [MS Ps SSO Roles](doc/controllers/ms-ps-sso-roles.md)
* [MS Ps Tickets](doc/controllers/ms-ps-tickets.md)
* [Orgs](doc/controllers/orgs.md)
* [Orgs Admins](doc/controllers/orgs-admins.md)
* [Orgs Alarms](doc/controllers/orgs-alarms.md)
* [Orgs Alarm Templates](doc/controllers/orgs-alarm-templates.md)
* [Orgs Antivirus Profiles](doc/controllers/orgs-antivirus-profiles.md)
* [Orgs API Tokens](doc/controllers/orgs-api-tokens.md)
* [Orgs AP Templates](doc/controllers/orgs-ap-templates.md)
* [Orgs Asset Filters](doc/controllers/orgs-asset-filters.md)
* [Orgs Assets](doc/controllers/orgs-assets.md)
* [Orgs Cert](doc/controllers/orgs-cert.md)
* [Orgs Clients-Marvis](doc/controllers/orgs-clients-marvis.md)
* [Orgs Clients-NAC](doc/controllers/orgs-clients-nac.md)
* [Orgs Clients-SDK](doc/controllers/orgs-clients-sdk.md)
* [Orgs Clients-Wan](doc/controllers/orgs-clients-wan.md)
* [Orgs Clients-Wired](doc/controllers/orgs-clients-wired.md)
* [Orgs Clients-Wireless](doc/controllers/orgs-clients-wireless.md)
* [Orgs Cradlepoint](doc/controllers/orgs-cradlepoint.md)
* [Orgs CRL](doc/controllers/orgs-crl.md)
* [Orgs Device Profiles](doc/controllers/orgs-device-profiles.md)
* [Orgs Devices](doc/controllers/orgs-devices.md)
* [Orgs Devices-Others](doc/controllers/orgs-devices-others.md)
* [Orgs Devices-SSR](doc/controllers/orgs-devices-ssr.md)
* [Orgs EVPN Topologies](doc/controllers/orgs-evpn-topologies.md)
* [Orgs Gateway Templates](doc/controllers/orgs-gateway-templates.md)
* [Orgs Guests](doc/controllers/orgs-guests.md)
* [Orgs IDP Profiles](doc/controllers/orgs-idp-profiles.md)
* [Orgs Inventory](doc/controllers/orgs-inventory.md)
* [Orgs JSE](doc/controllers/orgs-jse.md)
* [Orgs JSI](doc/controllers/orgs-jsi.md)
* [Orgs Licenses](doc/controllers/orgs-licenses.md)
* [Orgs Linked Applications](doc/controllers/orgs-linked-applications.md)
* [Orgs Logs](doc/controllers/orgs-logs.md)
* [Orgs Maps](doc/controllers/orgs-maps.md)
* [Orgs Marvis](doc/controllers/orgs-marvis.md)
* [Orgs Mx Clusters](doc/controllers/orgs-mx-clusters.md)
* [Orgs Mx Edges](doc/controllers/orgs-mx-edges.md)
* [Orgs Mx Tunnels](doc/controllers/orgs-mx-tunnels.md)
* [Orgs NACCRL](doc/controllers/orgs-naccrl.md)
* [Orgs NACIDP](doc/controllers/orgs-nacidp.md)
* [Orgs NAC Portals](doc/controllers/orgs-nac-portals.md)
* [Orgs NAC Rules](doc/controllers/orgs-nac-rules.md)
* [Orgs NAC Tags](doc/controllers/orgs-nac-tags.md)
* [Orgs Networks](doc/controllers/orgs-networks.md)
* [Orgs Network Templates](doc/controllers/orgs-network-templates.md)
* [Orgs Premium Analytics](doc/controllers/orgs-premium-analytics.md)
* [Orgs Psk Portals](doc/controllers/orgs-psk-portals.md)
* [Orgs Psks](doc/controllers/orgs-psks.md)
* [Orgs RF Templates](doc/controllers/orgs-rf-templates.md)
* [Orgs SCEP](doc/controllers/orgs-scep.md)
* [Orgs SDK Invites](doc/controllers/orgs-sdk-invites.md)
* [Orgs SDK Templates](doc/controllers/orgs-sdk-templates.md)
* [Orgs Sec Intel Profiles](doc/controllers/orgs-sec-intel-profiles.md)
* [Orgs Sec Policies](doc/controllers/orgs-sec-policies.md)
* [Orgs Service Policies](doc/controllers/orgs-service-policies.md)
* [Orgs Services](doc/controllers/orgs-services.md)
* [Orgs Setting](doc/controllers/orgs-setting.md)
* [Orgs Setting Zscaler](doc/controllers/orgs-setting-zscaler.md)
* [Orgs Sitegroups](doc/controllers/orgs-sitegroups.md)
* [Orgs Sites](doc/controllers/orgs-sites.md)
* [Orgs Site Templates](doc/controllers/orgs-site-templates.md)
* [Orgs SL Es](doc/controllers/orgs-sl-es.md)
* [Orgs SSO](doc/controllers/orgs-sso.md)
* [Orgs SSO Roles](doc/controllers/orgs-sso-roles.md)
* [Orgs Stats](doc/controllers/orgs-stats.md)
* [Orgs Stats-Assets](doc/controllers/orgs-stats-assets.md)
* [Orgs Stats-BGP Peers](doc/controllers/orgs-stats-bgp-peers.md)
* [Orgs Stats-Devices](doc/controllers/orgs-stats-devices.md)
* [Orgs Stats-Mx Edges](doc/controllers/orgs-stats-mx-edges.md)
* [Orgs Stats-Other Devices](doc/controllers/orgs-stats-other-devices.md)
* [Orgs Stats-Ports](doc/controllers/orgs-stats-ports.md)
* [Orgs Stats-Sites](doc/controllers/orgs-stats-sites.md)
* [Orgs Stats-Tunnels](doc/controllers/orgs-stats-tunnels.md)
* [Orgs Stats-VPN Peers](doc/controllers/orgs-stats-vpn-peers.md)
* [Orgs Subscriptions](doc/controllers/orgs-subscriptions.md)
* [Orgs Tickets](doc/controllers/orgs-tickets.md)
* [Orgs User MA Cs](doc/controllers/orgs-user-ma-cs.md)
* [Orgs Vars](doc/controllers/orgs-vars.md)
* [Orgs VP Ns](doc/controllers/orgs-vp-ns.md)
* [Orgs Webhooks](doc/controllers/orgs-webhooks.md)
* [Orgs Wlans](doc/controllers/orgs-wlans.md)
* [Orgs WLAN Templates](doc/controllers/orgs-wlan-templates.md)
* [Orgs Wx Rules](doc/controllers/orgs-wx-rules.md)
* [Orgs Wx Tags](doc/controllers/orgs-wx-tags.md)
* [Orgs Wx Tunnels](doc/controllers/orgs-wx-tunnels.md)
* [Samples Webhooks](doc/controllers/samples-webhooks.md)
* [Self Account](doc/controllers/self-account.md)
* [Self Alarms](doc/controllers/self-alarms.md)
* [Self API Token](doc/controllers/self-api-token.md)
* [Self Audit Logs](doc/controllers/self-audit-logs.md)
* [Self MFA](doc/controllers/self-mfa.md)
* [Self O Auth 2](doc/controllers/self-o-auth-2.md)
* [Sites](doc/controllers/sites.md)
* [Sites Alarms](doc/controllers/sites-alarms.md)
* [Sites Anomaly](doc/controllers/sites-anomaly.md)
* [Sites Applications](doc/controllers/sites-applications.md)
* [Sites AP Templates](doc/controllers/sites-ap-templates.md)
* [Sites Asset Filters](doc/controllers/sites-asset-filters.md)
* [Sites Assets](doc/controllers/sites-assets.md)
* [Sites Beacons](doc/controllers/sites-beacons.md)
* [Sites Clients-NAC](doc/controllers/sites-clients-nac.md)
* [Sites Clients-Wan](doc/controllers/sites-clients-wan.md)
* [Sites Clients-Wired](doc/controllers/sites-clients-wired.md)
* [Sites Clients-Wireless](doc/controllers/sites-clients-wireless.md)
* [Sites Device Profiles](doc/controllers/sites-device-profiles.md)
* [Sites Devices](doc/controllers/sites-devices.md)
* [Sites Devices-Others](doc/controllers/sites-devices-others.md)
* [Sites Devices-WAN Cluster](doc/controllers/sites-devices-wan-cluster.md)
* [Sites Devices-Wired](doc/controllers/sites-devices-wired.md)
* [Sites Devices-Wired-Virtual Chassis](doc/controllers/sites-devices-wired-virtual-chassis.md)
* [Sites Devices-Wireless](doc/controllers/sites-devices-wireless.md)
* [Sites Events](doc/controllers/sites-events.md)
* [Sites EVPN Topologies](doc/controllers/sites-evpn-topologies.md)
* [Sites Gateway Templates](doc/controllers/sites-gateway-templates.md)
* [Sites Guests](doc/controllers/sites-guests.md)
* [Sites Insights](doc/controllers/sites-insights.md)
* [Sites JSE](doc/controllers/sites-jse.md)
* [Sites Licenses](doc/controllers/sites-licenses.md)
* [Sites Location](doc/controllers/sites-location.md)
* [Sites Maps](doc/controllers/sites-maps.md)
* [Sites Maps-Auto-Placement](doc/controllers/sites-maps-auto-placement.md)
* [Sites Maps-Auto-Zone](doc/controllers/sites-maps-auto-zone.md)
* [Sites Mx Edges](doc/controllers/sites-mx-edges.md)
* [Sites Networks](doc/controllers/sites-networks.md)
* [Sites Network Templates](doc/controllers/sites-network-templates.md)
* [Sites Psks](doc/controllers/sites-psks.md)
* [Sites Rfdiags](doc/controllers/sites-rfdiags.md)
* [Sites RF Templates](doc/controllers/sites-rf-templates.md)
* [Sites Rogues](doc/controllers/sites-rogues.md)
* [Sites RRM](doc/controllers/sites-rrm.md)
* [Sites RSSI Zones](doc/controllers/sites-rssi-zones.md)
* [Sites Sec Intel Profiles](doc/controllers/sites-sec-intel-profiles.md)
* [Sites Service Policies](doc/controllers/sites-service-policies.md)
* [Sites Services](doc/controllers/sites-services.md)
* [Sites Setting](doc/controllers/sites-setting.md)
* [Sites Site Templates](doc/controllers/sites-site-templates.md)
* [Sites Skyatp](doc/controllers/sites-skyatp.md)
* [Sites SL Es](doc/controllers/sites-sl-es.md)
* [Sites Stats](doc/controllers/sites-stats.md)
* [Sites Stats-Apps](doc/controllers/sites-stats-apps.md)
* [Sites Stats-Assets](doc/controllers/sites-stats-assets.md)
* [Sites Stats-Beacons](doc/controllers/sites-stats-beacons.md)
* [Sites Stats-BGP Peers](doc/controllers/sites-stats-bgp-peers.md)
* [Sites Stats-Calls](doc/controllers/sites-stats-calls.md)
* [Sites Stats-Clients SDK](doc/controllers/sites-stats-clients-sdk.md)
* [Sites Stats-Clients Wireless](doc/controllers/sites-stats-clients-wireless.md)
* [Sites Stats-Devices](doc/controllers/sites-stats-devices.md)
* [Sites Stats-Discovered Switches](doc/controllers/sites-stats-discovered-switches.md)
* [Sites Stats-Mx Edges](doc/controllers/sites-stats-mx-edges.md)
* [Sites Stats-Ports](doc/controllers/sites-stats-ports.md)
* [Sites Stats-Wx Rules](doc/controllers/sites-stats-wx-rules.md)
* [Sites Stats-Zones](doc/controllers/sites-stats-zones.md)
* [Sites Synthetic Tests](doc/controllers/sites-synthetic-tests.md)
* [Sites UI Settings](doc/controllers/sites-ui-settings.md)
* [Sitesv Beacons](doc/controllers/sitesv-beacons.md)
* [Sites VP Ns](doc/controllers/sites-vp-ns.md)
* [Sites WAN Usages](doc/controllers/sites-wan-usages.md)
* [Sites Webhooks](doc/controllers/sites-webhooks.md)
* [Sites Wlans](doc/controllers/sites-wlans.md)
* [Sites Wx Rules](doc/controllers/sites-wx-rules.md)
* [Sites Wx Tags](doc/controllers/sites-wx-tags.md)
* [Sites Wx Tunnels](doc/controllers/sites-wx-tunnels.md)
* [Sites Zones](doc/controllers/sites-zones.md)
* [Utilities Common](doc/controllers/utilities-common.md)
* [Utilities LAN](doc/controllers/utilities-lan.md)
* [Utilities Location](doc/controllers/utilities-location.md)
* [Utilities Mx Edge](doc/controllers/utilities-mx-edge.md)
* [Utilities PCA Ps](doc/controllers/utilities-pca-ps.md)
* [Utilities Upgrade](doc/controllers/utilities-upgrade.md)
* [Utilities WAN](doc/controllers/utilities-wan.md)
* [Utilities Wi-Fi](doc/controllers/utilities-wi-fi.md)

## Classes Documentation

* [HttpConfiguration](doc/http-configuration.md)
* [RetryConfiguration](doc/retry-configuration.md)
* [LoggerConfiguration](doc/logger-configuration.md)
* [RequestLoggerConfiguration](doc/request-logger-configuration.md)
* [ResponseLoggerConfiguration](doc/response-logger-configuration.md)

