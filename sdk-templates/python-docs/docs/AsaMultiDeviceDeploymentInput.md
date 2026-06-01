# AsaMultiDeviceDeploymentInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployment_notes** | **str** | Specify notes, if any, for the deployment. | [optional] 
**description** | **str** | Specify a human-readable description for the deployment. | [optional] 
**device_uids** | **List[str]** | The set of unique identifiers, represented as UUIDs, of the ASA devices to deploy in Security Cloud Control. | 

## Example

```python
from scc_firewall_manager_sdk.models.asa_multi_device_deployment_input import AsaMultiDeviceDeploymentInput

# TODO update the JSON string below
json = "{}"
# create an instance of AsaMultiDeviceDeploymentInput from a JSON string
asa_multi_device_deployment_input_instance = AsaMultiDeviceDeploymentInput.from_json(json)
# print the JSON string representation of the object
print(AsaMultiDeviceDeploymentInput.to_json())

# convert the object into a dict
asa_multi_device_deployment_input_dict = asa_multi_device_deployment_input_instance.to_dict()
# create an instance of AsaMultiDeviceDeploymentInput from a dict
asa_multi_device_deployment_input_form_dict = asa_multi_device_deployment_input.from_dict(asa_multi_device_deployment_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


