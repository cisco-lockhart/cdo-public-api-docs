# DeviceDeploymentStatusDto

The current status of each device that is being deployed to as part of this deployment.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**completion_status** | [**CompletionStatusDto**](CompletionStatusDto.md) |  | [optional] 
**error_details** | [**List[ErrorDetails]**](ErrorDetails.md) | Detailed error information reported for this device during deployment. | [optional] 
**error_msg** | **str** | A summary error message if the deployment failed for this device. | [optional] 
**name** | **str** | The name of the device in SCC Firewall Manager. | [optional] 
**status** | **str** | The overall status of the deployment on this device | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device in SCC Firewall Manager. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_deployment_status_dto import DeviceDeploymentStatusDto

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceDeploymentStatusDto from a JSON string
device_deployment_status_dto_instance = DeviceDeploymentStatusDto.from_json(json)
# print the JSON string representation of the object
print(DeviceDeploymentStatusDto.to_json())

# convert the object into a dict
device_deployment_status_dto_dict = device_deployment_status_dto_instance.to_dict()
# create an instance of DeviceDeploymentStatusDto from a dict
device_deployment_status_dto_form_dict = device_deployment_status_dto.from_dict(device_deployment_status_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


