
# Custom Header Signature



Documentation for accessing and setting credentials for csrfToken.

## Auth Credentials

| Name | Type | Description | Setter | Getter |
|  --- | --- | --- | --- | --- |
| X-CSRFToken | `string` | This protects the website against [Cross Site Request Forgery](https://en.wikipedia.org/wiki/Cross-site_request_forgery), all the POST / PUT / DELETE APIs needs to have CSRF token in the AJAX Request header when using Login/Password authentication (with or without MFA)<br><br>The CSRF Token is sent back by Mist in the Cookies from the Login Response API Call:<br>`cookies[csrftoken]`<br><br>The CSRF Token must be added in the HTTP Request Headers:<br><br>```<br>X-CSRFToken: vwvBuq9qkqaKh7lu8tNc0gkvBfEaLAmx<br>``` | `WithXCSRFToken` | `XCSRFToken` |



**Note:** Required auth credentials can be set using `WithCsrfTokenCredentials()` by providing a credentials instance with `NewCsrfTokenCredentials()` in the configuration initialization and accessed using the `CsrfTokenCredentials()` method in the configuration instance.

## Usage Example

### Client Initialization

You must provide credentials in the client as shown in the following code snippet.

```go
client := mistapi.NewClient(
    mistapi.CreateConfiguration(
        mistapi.WithCsrfTokenCredentials(
            mistapi.NewCsrfTokenCredentials("X-CSRFToken"),
        ),
    ),
)
```


