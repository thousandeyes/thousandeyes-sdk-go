# Contributing to the ThousandEyes Go SDK

Thank you for considering a contribution to the ThousandEyes Go SDK.

## Generated code

The API domain packages in this repository are generated from ThousandEyes
OpenAPI definitions. Generated files contain a `Code generated` notice.

Do not edit generated files as a one-off fix. Changes to generated APIs must be
made in the appropriate specification or generator source and then regenerated
so they remain reproducible. Pull requests that include generated output should
describe its source and the command or workflow used to produce it.

Contributions to shared runtime code, documentation, tests, and workflows are
welcome.

## Development setup

Install the Go version declared in `go.mod`, clone your fork, and download the
module dependencies:

```sh
go mod download
```

Tests must not require real ThousandEyes credentials or external network
access.

## Development workflow

1. Fork the repository.
2. Create a focused feature branch.
3. Make the change and add or update tests where applicable.
4. Format and validate the repository:

   ```sh
   gofmt -w path/to/changed_file.go
   go test -race ./...
   git diff --check
   ```

5. Push the branch and open a pull request.

Keep pull requests focused and explain what changed, why it is needed, and how
it was validated. Use concise, imperative commit subjects.

## Reporting bugs and suggesting features

For bugs, feature requests, or questions about the SDK, contact
[ThousandEyes Support](https://docs.thousandeyes.com/product-documentation/getting-started/getting-support-from-thousandeyes#contacting-support).

Do not report security vulnerabilities through a public GitHub issue. Follow
the private process in [SECURITY.md](./SECURITY.md).

All contributors must follow the [Code of Conduct](./CODE_OF_CONDUCT.md).
