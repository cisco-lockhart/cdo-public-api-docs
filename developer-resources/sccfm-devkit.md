# SCCFM DevKit

The SCCFM DevKit provides a command-line interface and an Ansible collection for interacting with
Security Cloud Control Firewall Manager.

## CLI

Install the `sccfm-cli` command from [PyPI](https://pypi.org/project/cisco-sccfm-devkit/):

```shell
pipx install cisco-sccfm-devkit
```

## Ansible Collection

Install the `cisco.sccfm` collection from [Ansible Galaxy](https://galaxy.ansible.com/ui/repo/published/cisco/sccfm/):

```shell
ansible-galaxy collection install cisco.sccfm
```

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

## Documentation

Full CLI and Ansible reference documentation, along with source and examples, is available on
[GitHub](https://github.com/CiscoDevNet/sccfm-devkit) and the
[generated reference site](https://ciscodevnet.github.io/sccfm-devkit/).
