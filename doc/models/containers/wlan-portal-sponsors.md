
# Wlan Portal Sponsors

Object of allowed sponsors email with name. Required if `sponsor_enabled` is `true` and `sponsor_email_domains` is empty. Property key is the sponsor email, Property value is the sponsor name. List of email allowed for backward compatibility

## Class Name

`WlanPortalSponsors`

## Cases

| Type | Factory Method |
|  --- | --- |
| `[]string` | models.WlanPortalSponsorsContainer.FromArrayOfString([]string arrayOfString) |
| `map[string]string` | models.WlanPortalSponsorsContainer.FromMapOfString(map[string]string mapOfString) |

## []string

### Initialization Code

#### Example

```go
value := models.WlanPortalSponsorsContainer.FromArrayOfString([]string{
    "String1",
})
```

## map[string]string

### Initialization Code

#### Example

```go
value := models.WlanPortalSponsorsContainer.FromMapOfString(map[string]string{
    "sponsor1@company.com": "FirstName1 LastName1",
    "sponsor2@company.com": "FirstName2 LastName2",
})
```

