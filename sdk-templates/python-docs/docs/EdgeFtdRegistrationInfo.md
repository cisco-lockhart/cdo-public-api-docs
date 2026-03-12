# EdgeFtdRegistrationInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**hostname** | **str** | The FQDN or public IP of the FTD being registered to the cdFMC. This hostname must be accessible from the public internet, as the cdFMC will use it to initiate the connection to the FTD. This is only present when the FTD is registered using the cdFMC-triggered CLI flow (as opposed to the default flow, where the FTD connects to the cdFMC). | [optional] 
**nat_id** | **str** | The NAT ID used during cdFMC-triggered registration. This is only present when the FTD is registered using the cdFMC-triggered CLI flow, where the cdFMC initiates the connection to the FTD (as opposed to the default flow, where the FTD connects to the cdFMC). | [optional] 
**reg_key** | **str** | The registration key used to authenticate the cdFMC-triggered registration. This is only present when the FTD is registered using the cdFMC-triggered CLI flow, where the cdFMC initiates the connection to the FTD (as opposed to the default flow, where the FTD connects to the cdFMC). | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.edge_ftd_registration_info import EdgeFtdRegistrationInfo

# TODO update the JSON string below
json = "{}"
# create an instance of EdgeFtdRegistrationInfo from a JSON string
edge_ftd_registration_info_instance = EdgeFtdRegistrationInfo.from_json(json)
# print the JSON string representation of the object
print(EdgeFtdRegistrationInfo.to_json())

# convert the object into a dict
edge_ftd_registration_info_dict = edge_ftd_registration_info_instance.to_dict()
# create an instance of EdgeFtdRegistrationInfo from a dict
edge_ftd_registration_info_form_dict = edge_ftd_registration_info.from_dict(edge_ftd_registration_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


