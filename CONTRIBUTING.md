# How to Contribute

Thank you for your interest in contributing to this repository. If you are not a Cisco employee, we 
unfortunately cannot accept pull request submissions from you. However, we welcome you to open issues.

If you are a Cisco employee who is working on a microservice that will become part of this repository, read on.

# Adding a new microservice

- Edit `cloud-fw-mgr-api-docs.config.yaml` in the root of this repo, and add the new microservice to
  the `services` section. For example:

```yaml
  - name: "my-new-service"
    url: "https://staging.manage.security.cisco.com/api/platform/my-new-service/v3/api-docs.yaml"
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

# Making API changes to an existing microservice

## My change is backwards-incompatible

To adhere to Cisco's policy around backwards compatibility, you **cannot** make backwards-incompatible changes to an existing microservice.
If you need to make a backwards-incompatible change, please work with [Jeremy Street](mailto:jerestre@cisco.com) and [Siddhu Warrier](mailto:siwarrie@cisco.com).

## My change is backwards-compatible

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