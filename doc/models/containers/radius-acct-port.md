
# Radius Acct Port

RADIUS Auth Port, value from 1 to 65535, default is 1813

## Class Name

`RadiusAcctPort`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.RadiusAcctPortContainer.FromNumber(int number) |
| `string` | models.RadiusAcctPortContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.RadiusAcctPortContainer.FromNumber(1)
```

## string

### Initialization Code

#### Example

```go
value := models.RadiusAcctPortContainer.FromString("String0")
```

