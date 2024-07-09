
# Suppressed Alarm Applies

if `scope`==`site`
Object defines the scope (within the org e.g. whole org, and/or some site_groups, and/or some sites) for which the alarm service has to be suppressed for some `duration`

## Structure

`SuppressedAlarmApplies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | UUID of the current org (if provided, the alarms will be suppressed at org level) |
| `SiteIds` | `[]uuid.UUID` | Optional | List of UUID of the sites within the org (if provided, the alarms will be suppressed for all the mentioned sites under the org) |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of UUID of the site groups within the org (if provided, the alarms will be suppressed for all the sites under those site groups) |

## Example (as JSON)

```json
{
  "org_id": "00001516-0000-0000-0000-000000000000",
  "site_ids": [
    "00001dcc-0000-0000-0000-000000000000",
    "00001dcd-0000-0000-0000-000000000000"
  ],
  "sitegroup_ids": [
    "0000267e-0000-0000-0000-000000000000",
    "0000267f-0000-0000-0000-000000000000"
  ]
}
```
