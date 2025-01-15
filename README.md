# Welcome to the Cisco Security Cloud Control Public API Documentation

These docs are published to DevNet using PubHub.

See [https://pubhub.cisco.com/detail/5082] or contact siwarrie@cisco.com for details

## View your changes

Go to [https://devnetapps.cisco.com/docs/cisco-defense-orchestrator/].

# Contributions

- PRs with changes to this file will only accepted from contributors who work for Cisco, who use
  a Github account associated with their Cisco email address to create the PR.

## Adding a new microservice

- Edit `cloud-fw-mgr-api-docs.config.yaml` in the root of this repo, and add the new microservice to
  the `services` section. For example:

```yaml
  - name: "my-new-service"
    url: "https://staging.manage.security.cisco.com/api/platform/my-new-service/v3/api-docs.yaml"
    # Optional: if your endpoints are accessible on /api/platform/myapp/foos, but the public-api-gateway makes them available on /api/rest/v1/bars/foos, you should specify /v1/bars as the prefix. Start with a /, and no trailing slashes.
    prefixToAdd: "/v1/bars"
```
- Edit `api-changelog.md` to add a short description of what APIs your new microservice provides.
- Submit a PR (please see commit message guidelines).
- You will need approval from the CODEOWNERS. See `.github/CODEOWNERS` to see who to request
  approval from.
- Once your PR is approved, you will be (Cisco network only) be able to view the changes at
  [https://devnetapps.cisco.com/docs/cisco-security-cloud-control/](https://devnetapps.cisco.com/docs/cisco-security-cloud-control/).
- To get it published to [https://developer.cisco.com/docs/cisco-security-cloud-control](https://developer.cisco.com/docs/cisco-security-cloud-control),
  please work with the Cisco Devnet team or [Siddhu Warrier](mailto:siwarrie@cisco.com), as this is a manual process that will involve reviewing the APIs you have added to ensure
  that they meet Cisco style guidelines.

## Making API changes to an existing microservice

### My change is backwards-incompatible

To adhere to Cisco's policy around backwards compatibility, you **cannot** make backwards-incompatible changes to an existing microservice.
If you need to make a backwards-incompatible change, please work with [Jeremy Street](mailto:jerestre@cisco.com) and [Siddhu Warrier](mailto:siwarrie@cisco.com).

### My change is backwards-compatible

This can involve adding new APIs, or changing existing in APIs in a backwards-compatible manner.

- Edit `api-changelog.md` to add a short description of what APIs your new microservice provides.
- Submit a PR (please see commit message guidelines).
- You will need approval from the CODEOWNERS. See `.github/CODEOWNERS` to see who to request
  approval from.
- Once your PR is approved, you will be (Cisco network only) be able to view the changes at
  [https://devnetapps.cisco.com/docs/cisco-security-cloud-control/](https://devnetapps.cisco.com/docs/cisco-security-cloud-control/).
- To get it published to [https://developer.cisco.com/docs/cisco-security-cloud-control](https://developer.cisco.com/docs/cisco-security-cloud-control),
  please work with the Cisco Devnet team or [Siddhu Warrier](mailto:siwarrie@cisco.com), as this is a manual process that will involve reviewing the APIs you have added to ensure
  that they meet Cisco style guidelines.

## PR guidelines

All PR commits should follow
the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format.
This will help us automate the release process, and we will set the version of the CLI on the basis
of the commit made.

# The Golang CLI for API documentation generation

## Pre-requisites

- Install Golang: https://golang.org/doc/install
- Install cobra-cli: `go install github.com/spf13/cobra-cli@latest`
- Install pre-requisites:

```shell
go run main.go install-pre-reqs
```

## Development

### Add a new command

```shell
cobra-cli add <command-name>
```

### Testing

All tests are written using [Ginkgo](https://github.com/onsi/ginkgo)
and [Gomega](https://github.com/onsi/gomega).

To run your tests, use the following command:

```shell
cd /path/to/scripts/golang
ginkgo -r
```