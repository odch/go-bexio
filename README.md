go-bexio
========

A small Go client for the [bexio](https://docs.bexio.com/) API. Implements only
the subset needed by ODCH's applications (contacts, invoices: list/get/create,
issue, PDF).

Derived from [philhug/go-bexio](https://github.com/philhug/go-bexio); maintained
by ODCH under `github.com/odch/go-bexio`.

## Usage

```go
client := bexio.NewClient(httpClient) // httpClient carries the bearer token
client.Authorization = "<personal access token>" // or use an OAuth2 http.Client
contacts, err := client.Contacts.ListContacts()
```

Authentication is left to the caller: pass an OAuth2-configured `*http.Client`,
or set `client.Authorization` to a personal access token.

## License

MIT — see [LICENSE](LICENSE).
