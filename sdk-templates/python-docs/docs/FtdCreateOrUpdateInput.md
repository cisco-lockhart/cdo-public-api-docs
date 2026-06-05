# FtdCreateOrUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_type** | **str** | Specify the type of the FTD. The only supported type of FTD is CDFMC_MANAGED_FTD. | 
**fmc_access_policy_uid** | **str** | Specify the unique identifier, represented as a UUID, of the FMC access policy to apply to this device. | 
**ftd_hostname** | **str** | The FQDN or public IP of the FTD being registered to the cdFMC. This hostname must be accessible from the public internet, as the cdFMC will use it to initiate the connection to the FTD. Required when useCdFmcTriggeredRegistration is true; must not be specified otherwise. | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 
**licenses** | **List[str]** | Specify a set of licenses to apply to the device. | 
**name** | **str** | Specify a human-readable name for the device. | 
**nat_id** | **str** | The NAT ID used during cdFMC-triggered registration, where the cdFMC initiates the connection to the FTD (as opposed to the default flow, where the FTD connects to the cdFMC). Required when useCdFmcTriggeredRegistration is true; must not be specified otherwise. | [optional] 
**performance_tier** | **str** | Specify the performance tier of the FTDv (required only if isVirtual is set to true). | [optional] 
**reg_key** | **str** | The registration key used to authenticate the cdFMC-triggered registration, where the cdFMC initiates the connection to the FTD (as opposed to the default flow, where the FTD connects to the cdFMC). Required when useCdFmcTriggeredRegistration is true; must not be specified otherwise. | [optional] 
**use_cd_fmc_triggered_registration** | **bool** | If your FTD is publicly accessible, then you can choose to use the cdFMC triggered registration flow. If this field is set to true, the ftdHostname, natId, and regKey fields must be specified. | [optional] 
**virtual** | **bool** | Indicate whether the FTD is a virtual or a physical device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ftd_create_or_update_input import FtdCreateOrUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of FtdCreateOrUpdateInput from a JSON string
ftd_create_or_update_input_instance = FtdCreateOrUpdateInput.from_json(json)
# print the JSON string representation of the object
print(FtdCreateOrUpdateInput.to_json())

# convert the object into a dict
ftd_create_or_update_input_dict = ftd_create_or_update_input_instance.to_dict()
# create an instance of FtdCreateOrUpdateInput from a dict
ftd_create_or_update_input_form_dict = ftd_create_or_update_input.from_dict(ftd_create_or_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


