# DeploymentRunDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployer_username** | **str** | The name of the user performing the deploy. | [optional] 
**deployment_run_status** | **str** | The status of the deployment | [optional] 
**deployment_type** | **str** | The type of the deployment | [optional] 
**description** | **str** | A human-readable description of this deployment. | [optional] 
**device_deployment_statuses** | [**List[DeviceDeploymentStatusDto]**](DeviceDeploymentStatusDto.md) | The current status of each device that is being deployed to as part of this deployment. | [optional] 
**device_uids** | **List[str]** | The set of device UIDs that are part of this deployment. | [optional] 
**error_msg** | **str** | Error message if the deployment failed. | [optional] 
**failure_reason** | **str** | The reason the deployment failed, if applicable. | [optional] 
**ignore_warnings** | **bool** | Boolean indicating whether warnings identified by the deployment sub-system about this deployment were ignored. | [optional] 
**last_updated_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the deployment run was last updated. | [optional] 
**name** | **str** | The name of the deployment run. Deployment runs names are unique in a tenant in SCC Firewall Manager. | [optional] 
**submission_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the deployment run was triggered. | [optional] 
**transaction_uid** | **str** | The unique identifier of the CDO transaction that tracks this deployment. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the deployment run in SCC Firewall Manager. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.deployment_run_dto import DeploymentRunDto

# TODO update the JSON string below
json = "{}"
# create an instance of DeploymentRunDto from a JSON string
deployment_run_dto_instance = DeploymentRunDto.from_json(json)
# print the JSON string representation of the object
print(DeploymentRunDto.to_json())

# convert the object into a dict
deployment_run_dto_dict = deployment_run_dto_instance.to_dict()
# create an instance of DeploymentRunDto from a dict
deployment_run_dto_form_dict = deployment_run_dto.from_dict(deployment_run_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


