# Getting Started

## Requirements

- Python 3.8+

## Installation

```bash
pip install scc-firewall-manager-sdk
```

## Authentication

The SDK uses Bearer token (JWT) authentication. You can obtain a token from the Security Cloud Control console.

```python
from scc_firewall_manager_sdk import ApiClient, Configuration
import os

configuration = Configuration(
    host="https://api.us.security.cisco.com/firewall",
    access_token=os.environ["BEARER_TOKEN"]
)
```

### Regional Endpoints

| Region | Host                                          |
|--------|-----------------------------------------------|
| US     | `https://api.us.security.cisco.com/firewall`  |
| EU     | `https://api.eu.security.cisco.com/firewall`  |
| APJ    | `https://api.apj.security.cisco.com/firewall` |
| AU     | `https://api.au.security.cisco.com/firewall`  |
| IN     | `https://api.in.security.cisco.com/firewall`  |
| UAE    | `https://api.uae.security.cisco.com/firewall` |

## Example Usage

### List Devices

```python
from scc_firewall_manager_sdk import Configuration, ApiClient, InventoryApi
import os

configuration = Configuration(
    host="https://api.us.security.cisco.com/firewall", # Replace `us` with your region
    access_token=os.environ["BEARER_TOKEN"]
)

with ApiClient(configuration) as api_client:
    inventory_api = InventoryApi(api_client)
    devices = inventory_api.get_devices()
    for device in devices.items:
        print(f"Device: {device.name} - {device.device_type}")
```