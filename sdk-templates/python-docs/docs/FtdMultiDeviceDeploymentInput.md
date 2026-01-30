# FtdMultiDeviceDeploymentInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployment_notes** | **str** | Specify notes, if any, for the deployment. | [optional] 
**description** | **str** | Specify a human-readable description for the deployment. | [optional] 
**device_uids** | **List[str]** | The set of unique identifiers, represented as UUIDs, of the devices to upgrade in Security Cloud Control. All of the devices in the list have to be compatible with the upgrade package. | 
**ignore_warnings** | **bool** | Specify whether to ignore warnings generated during the pre-validation of the deployment job and proceed with the deployment regardless. **Warning**: Do not set this to &#x60;true&#x60; unless you know what you are doing. | [optional] [default to False]

## Example

```python
from scc_firewall_manager_sdk.models.ftd_multi_device_deployment_input import FtdMultiDeviceDeploymentInput

# TODO update the JSON string below
json = "{}"
# create an instance of FtdMultiDeviceDeploymentInput from a JSON string
ftd_multi_device_deployment_input_instance = FtdMultiDeviceDeploymentInput.from_json(json)
# print the JSON string representation of the object
print(FtdMultiDeviceDeploymentInput.to_json())

# convert the object into a dict
ftd_multi_device_deployment_input_dict = ftd_multi_device_deployment_input_instance.to_dict()
# create an instance of FtdMultiDeviceDeploymentInput from a dict
ftd_multi_device_deployment_input_form_dict = ftd_multi_device_deployment_input.from_dict(ftd_multi_device_deployment_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


