
# Suppressed Alarm Applies

If `scope`==`site`. Object defines the scope (within the org e.g. whole org, and/or some site_groups, and/or some sites) for which the alarm service has to be suppressed for some `duration`

*This model accepts additional fields of type interface{}.*

## Structure

`SuppressedAlarmApplies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | UUID of the current org (if provided, the alarms will be suppressed at org level) |
| `SiteIds` | `[]uuid.UUID` | Optional | List of UUID of the sites within the org (if provided, the alarms will be suppressed for all the mentioned sites under the org) |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of UUID of the site groups within the org (if provided, the alarms will be suppressed for all the sites under those site groups) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "00001c94-0000-0000-0000-000000000000",
  "site_ids": [
    "000010c2-0000-0000-0000-000000000000",
    "000010c1-0000-0000-0000-000000000000",
    "000010c0-0000-0000-0000-000000000000"
  ],
  "sitegroup_ids": [
    "000006ec-0000-0000-0000-000000000000"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

