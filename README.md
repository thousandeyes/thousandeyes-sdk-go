[![Go Reference](https://pkg.go.dev/badge/github.com/thousandeyes/thousandeyes-sdk-go/v3.svg)](https://pkg.go.dev/github.com/thousandeyes/thousandeyes-sdk-go/v3)
[![Release](https://img.shields.io/github/v/release/thousandeyes/thousandeyes-sdk-go?include_prereleases&sort=semver)](https://github.com/thousandeyes/thousandeyes-sdk-go/releases)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](./LICENSE)

# ThousandEyes Go SDK

## Description

`thousandeyes-sdk-go` provides generated Go clients for the
[ThousandEyes API v7](https://developer.cisco.com/docs/thousandeyes/v7/).
Each API domain is exposed as a separate package and uses the shared `client`
package for authentication, configuration, HTTP execution, and retries.

This project is maintained by the ThousandEyes team at Cisco. The v3 module
uses Go semantic import versioning through the `/v3` module path. Pin an exact
release when you need reproducible builds, and review the
[release notes](https://github.com/thousandeyes/thousandeyes-sdk-go/releases)
before upgrading.

## Installation and usage

### Requirements

The module's `go` directive is 1.23 and its suggested toolchain is Go 1.24.8,
as declared in [`go.mod`](./go.mod).

### Install the module

Add the v3 module to your project:

```sh
go get github.com/thousandeyes/thousandeyes-sdk-go/v3@latest
```

For reproducible builds, replace `latest` with a specific v3 release tag.
Import the shared client and the package for the API domain you need:

```go
import (
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
)
```

### API packages

Use the module installation above. Each published API domain includes generated
endpoint and model documentation:

- [Administrative](./administrative/README.md)
- [Agents](./agents/README.md)
- [Alerts](./alerts/README.md)
- [BGP monitors](./bgpmonitors/README.md)
- [Connectors](./connectors/README.md)
- [Credentials](./credentials/README.md)
- [Dashboards](./dashboards/README.md)
- [Emulation](./emulation/README.md)
- [Endpoint agents](./endpointagents/README.md)
- [Endpoint instant tests](./endpointinstanttests/README.md)
- [Endpoint labels](./endpointlabels/README.md)
- [Endpoint test results](./endpointtestresults/README.md)
- [Endpoint tests](./endpointtests/README.md)
- [Event detection](./eventdetection/README.md)
- [Instant tests](./instanttests/README.md)
- [Internet Insights](./internetinsights/README.md)
- [Snapshots](./snapshots/README.md)
- [Streaming](./streaming/README.md)
- [Tags](./tags/README.md)
- [Test results](./testresults/README.md)
- [Tests](./tests/README.md)
- [Usage](./usage/README.md)

### Authenticate and call an API

ThousandEyes API v7 uses bearer token authentication. Load tokens from an
environment variable or secret manager and never commit them to source control,
examples, fixtures, or logs.

The following example lists roles from the Administrative API:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
)

func main() {
	token := os.Getenv("TE_TOKEN")
	if token == "" {
		log.Fatal("TE_TOKEN is required")
	}

	configuration := client.NewConfiguration().WithAuthToken(token)
	apiClient := client.NewAPIClient(configuration)
	rolesAPI := (*administrative.RolesAPIService)(&apiClient.Common)

	request := rolesAPI.GetRoles()
	if aid := os.Getenv("TE_AID"); aid != "" {
		request = request.Aid(aid)
	}

	roles, response, err := request.Execute()
	if err != nil {
		if response != nil {
			log.Fatalf("list roles failed with %s: %v", response.Status, err)
		}
		log.Fatalf("list roles failed before receiving a response: %v", err)
	}

	fmt.Printf("received %d roles\n", len(roles.GetRoles()))
}
```

`client.NewConfiguration()` defaults to
`https://api.thousandeyes.com/v7`. Configuration methods can provide a token or
different API v7 endpoint, while fields on the configuration allow a custom
HTTP client, context, and user agent. Create and reuse a configured API client;
do not modify its configuration while requests are in flight.

`Debug` is disabled by default and should remain disabled in production. When
enabled, it dumps complete HTTP requests and responses, including payloads.
Common credential fields and authentication or cookie headers are redacted, but
URLs, other headers, and payloads can still contain sensitive API data. Treat
debug logs as sensitive and review them before persisting or sharing them.

Account-group context is optional and is set per request with `.Aid(aid)` on
operations that support it. The example reads it from `TE_AID` when present.

### Migrating from v2

Version 3 is a generated API v7 client and is not a drop-in replacement for the
hand-written v2 API v6 client. Expect to update imports and API calls:

| v2 | v3 |
| --- | --- |
| Module path ends in `/v2` | Module path ends in `/v3` |
| One root `thousandeyes` package | Shared `client` plus one package per API domain |
| Default endpoint is `https://api.thousandeyes.com/v6` | Default endpoint is `https://api.thousandeyes.com/v7` |
| `thousandeyes.NewClient(&thousandeyes.ClientOptions{AuthToken: token, AccountID: aid})` | `client.NewAPIClient(client.NewConfiguration().WithAuthToken(token))` |
| `AccountID` applies to all client calls | `.Aid(aid)` sets account context on each supported request |
| `sdk.GetRoles()` returns a model and error | `rolesAPI.GetRoles().Execute()` also returns the HTTP response |

Review the generated README and endpoint documentation for each API package.
API v7 model and method names may not have direct v2 equivalents.

### Pagination

Cursor-paginated request types expose `.Paginated()`. The returned pager fetches
one page at a time with `NextPage(ctx)` and reports whether another page may be
available through `HasMorePages()`. Use `.Cursor(value)` before `.Paginated()`
when you need to start from a known cursor.

When the response's item collection is known, the request also exposes
`.All(ctx)`. It returns a lazy, single-use iterator that follows `_links.next`,
stops fetching when the caller stops iterating, and yields an error if a page
cannot be retrieved:

```go
func listAllUserEvents(
	ctx context.Context,
	api *administrative.UserEventsAPIService,
) error {
	request := api.GetUserEvents().Window("1d")
	for event, err := range request.All(ctx) {
		if err != nil {
			return fmt.Errorf("list user events: %w", err)
		}
		fmt.Printf("event: %+v\n", event)
	}
	return nil
}
```

The pager rejects `_links.next` continuations with missing, empty, duplicate,
or previously seen cursor values instead of silently looping or restarting
pagination.

### Errors and retries

Generated `Execute` methods return a decoded model, an `*http.Response`, and an
`error`. Check the error before using the model. The response can be nil when a
request fails before the server responds or retry attempts are exhausted.

The shared client automatically retries recoverable connection failures, HTTP
429 responses, and most HTTP 5xx responses. It honors supported rate-limit and
`Retry-After` headers before falling back to exponential backoff. Because this
also applies to generated mutation operations, design non-idempotent workflows
to tolerate an ambiguous response and a retried request.

## Support

Use the
[ThousandEyes Community](https://community.cisco.com/t5/thousandeyes/bd-p/disc-thousandeyes)
for general best practices, help, tips, and examples. These resources may also
answer your question:

- [ThousandEyes Documentation](https://docs.thousandeyes.com/)
- [Internet and Cloud Intelligence Blog](https://www.thousandeyes.com/blog/)
- [Cisco ThousandEyes Blog](https://blogs.cisco.com/tag/cisco-thousandeyes?dtid=osscdc000283)
- [API developer support](https://developer.cisco.com/docs/thousandeyes/v7/developer-support/#developer-support)

For bug reports, feature requests, or questions about this SDK, contact
[ThousandEyes Support](https://docs.thousandeyes.com/product-documentation/getting-started/getting-support-from-thousandeyes#contacting-support).

## Roadmap and maintenance

This library is continuously updated alongside the
[ThousandEyes API v7](https://developer.cisco.com/docs/thousandeyes/v7/).

## Contributing

The API domain packages are generated from ThousandEyes OpenAPI definitions.
Generated files contain a `Code generated` notice and should not be edited as a
one-off fix. Changes to generated APIs must be made in the appropriate
specification or generator source and then regenerated. Contributions to shared
runtime code, documentation, tests, and workflows are welcome.

See [CONTRIBUTING.md](./CONTRIBUTING.md) for local development, formatting,
testing, and pull request guidance. Report security vulnerabilities privately
as described in [SECURITY.md](./SECURITY.md). All contributors must follow the
[Code of Conduct](./CODE_OF_CONDUCT.md).

## License

This project is licensed under the [Apache License 2.0](./LICENSE).

The ThousandEyes engineering team maintains this project and thanks William
Fleming, John Dyer, Joshua Blanchard, and all community contributors for their
work on the SDK.
