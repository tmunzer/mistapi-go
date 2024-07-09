
# Webhook Nac Events Event

## Structure

`WebhookNacEventsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | random mac |
| `AuthType` | `*string` | Optional | authentication type, e.g. "eap-tls", "eap-ttls", "eap-teap", "mab", "device-auth" |
| `Bssid` | `*string` | Optional | BSSID |
| `DryrunNacruleId` | `*uuid.UUID` | Optional | NAC Policy Dry Run Rule ID, if present and matched |
| `DryrunNacruleMatched` | `*bool` | Optional | True - if dryrun rule present and matched with priority, False - if not matched or not present |
| `IdpId` | `*uuid.UUID` | Optional | SSO ID, if present and used |
| `IdpRole` | `[]string` | Optional | IDP returned roles/groups for the user |
| `Mac` | `*string` | Optional | MAC address |
| `NacruleId` | `*uuid.UUID` | Optional | NAC Policy Rule ID, if matched |
| `NacruleMatched` | `*bool` | Optional | NAC Policy Rule Matched |
| `NasVendor` | `*string` | Optional | vendor of NAS device |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RandomMac` | `*bool` | Optional | AP MAC |
| `RespAttrs` | `[]string` | Optional | Radius attributes returned by NAC to NAS Devive |
| `SiteId` | `*uuid.UUID` | Optional | site id if assigned, null if not assigned |
| `Ssid` | `*string` | Optional | SSID |
| `Timestamp` | `*float64` | Optional | start time, in epoch |
| `Type` | `*string` | Optional | event type, e.g. NAC_CLIENT_PERMIT |
| `Username` | `*string` | Optional | Username presented by the client |
| `Vlan` | `*string` | Optional | Vlan ID |

## Example (as JSON)

```json
{
  "ap": "5c5b35513227",
  "auth_type": "802.1X",
  "bssid": "5c5b355fafcc",
  "dryrun_nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264012",
  "idp_id": "912ef72e-2239-4996-b81e-469e87a27cd6",
  "idp_role": [
    "itsuperusers",
    "vip"
  ],
  "mac": "ac3eb179e535",
  "nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62",
  "nas_vendor": "juniper-mist",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "resp_attrs": [
    "Tunnel-Type=VLAN",
    "Tunnel-Medium-Type=IEEE-802",
    "Tunnel-Private-Group-Id=750",
    "User-Name=anonymous"
  ],
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "##mist_nac",
  "timestamp": 1691512031.358188,
  "type": "NAC_CLIENT_PERMIT",
  "username": "user@deaflyz.net",
  "vlan": "750",
  "dryrun_nacrule_matched": false
}
```

