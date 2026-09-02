<seotitle>Cisco Secure Firewall, SCC Firewall Manager, Security Cloud Control Firewall Manager SDK, Security Cloud Control Firewall Manager CLI, Security Cloud Control Firewall Manager Terraform Provider, Security Cloud Control Firewall Manager Ansible Collection, Security Cloud Control Firewall Manager Automation, Claude Code, Codex</seotitle>

# Automation Tools and AI Agents

Alongside the REST API, we provide a set of tools to automate Security Cloud Control Firewall
Manager: a Python SDK, a command-line interface, an Ansible collection, a Terraform provider, and an
agent plugin for Claude Code and Codex.

All of these tools are built on top of the Security Cloud Control Firewall Manager API, and all of
them authenticate with the same API token. See
[Authentication](/docs/cisco-security-cloud-control-firewall-manager/authentication/) to create one.

Prefer to explore the API interactively before automating it? Start with our
[Postman Collections](/docs/cisco-security-cloud-control-firewall-manager/postman-collections/).

## Python SDK

We provide a Software Development Kit (SDK) in Python to help you interact with the Security Cloud
Control APIs.

The SDK can be downloaded from [PyPI](https://pypi.org/project/scc-firewall-manager-sdk/).

```shell
pip install scc-firewall-manager-sdk
```

You can find the SDK
[documentation here](https://scc-firewall-manager-sdk.readthedocs.io/en/latest/).

## CLI

The `sccfm-cli` command lets you drive Security Cloud Control Firewall Manager from your shell and
from scripts. Install it from [PyPI](https://pypi.org/project/cisco-sccfm-devkit/):

```shell
pipx install cisco-sccfm-devkit
```

Alternatively, install it with Homebrew:

```shell
brew tap CiscoDevNet/tap
brew trust --formula CiscoDevNet/tap/sccfm-cli  # required once on Homebrew 6.0+
brew install CiscoDevNet/tap/sccfm-cli
```

## Ansible Collection

The `cisco.sccfm` collection lets you manage Security Cloud Control Firewall Manager from Ansible
playbooks. Install it from
[Ansible Galaxy](https://galaxy.ansible.com/ui/repo/published/cisco/sccfm/):

```shell
ansible-galaxy collection install cisco.sccfm
```

## Terraform Provider

The Security Cloud Control Firewall Manager Terraform Provider can be used to onboard and manage
devices and other resources. Download it from the
[Terraform registry](https://registry.terraform.io/providers/CiscoDevNet/sccfm/latest/docs).

View examples of how to use the Terraform provider on
[GitHub](https://github.com/CiscoDevNet/terraform-provider-sccfm/tree/main/provider/examples).

## Claude Code and Codex Plugin

An `sccfm` agent plugin bundles guided installation, authentication setup, and operational
skills for the CLI and Ansible collection into Claude Code and Codex.

Claude Code:

```shell
/plugin marketplace add CiscoDevNet/sccfm-devkit
/plugin install sccfm@sccfm-devkit
```

Codex:

```shell
codex plugin marketplace add CiscoDevNet/sccfm-devkit
codex plugin add sccfm@sccfm-devkit
```

## Reference Documentation

Full CLI and Ansible reference documentation, along with source and examples, is available on
[GitHub](https://github.com/CiscoDevNet/sccfm-devkit) and the
[generated reference site](https://ciscodevnet.github.io/sccfm-devkit/).

## Reporting Issues

Please report bugs and feature requests for these tools on GitHub:

- [CLI, Ansible collection, and agent plugin](https://github.com/CiscoDevNet/sccfm-devkit/issues)
- [Terraform Provider](https://github.com/CiscoDevNet/terraform-provider-sccfm/issues)

For questions, best practices, and general feedback, use the
[Network Security forum](https://community.cisco.com/t5/network-security/bd-p/disc-network-security)
with the `Cisco Security Cloud Control` label. See
[Developer Support](/docs/cisco-security-cloud-control-firewall-manager/developer-support/) for
all support options.
