
# Sponsor Link Validity Duration

Optional if `sponsor_enabled`==`true`. How long to remain valid sponsored guest request approve/deny link received in email, in minutes. Value is between 5 and 60.

## Class Name

`SponsorLinkValidityDuration`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.SponsorLinkValidityDurationContainer.FromNumber(int number) |
| `string` | models.SponsorLinkValidityDurationContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.SponsorLinkValidityDurationContainer.FromNumber(60)
```

## string

### Initialization Code

#### Example

```go
value := models.SponsorLinkValidityDurationContainer.FromString("String0")
```

