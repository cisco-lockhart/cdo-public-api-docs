# DeploymentRunPatchInput

Fields to update on a deployment run. Only `deploymentNotes` and `labels` can be updated; a request containing any other field is rejected. Fields omitted from the request are left unchanged. Set a field to `null` (or, for `deploymentNotes`, an empty string, and for `labels`, an object with no labels) to clear it.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployment_notes** | **str** | The notes for the deployment run. Omit to leave unchanged; set to &#x60;null&#x60; or an empty string to clear. | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.deployment_run_patch_input import DeploymentRunPatchInput

# TODO update the JSON string below
json = "{}"
# create an instance of DeploymentRunPatchInput from a JSON string
deployment_run_patch_input_instance = DeploymentRunPatchInput.from_json(json)
# print the JSON string representation of the object
print(DeploymentRunPatchInput.to_json())

# convert the object into a dict
deployment_run_patch_input_dict = deployment_run_patch_input_instance.to_dict()
# create an instance of DeploymentRunPatchInput from a dict
deployment_run_patch_input_form_dict = deployment_run_patch_input.from_dict(deployment_run_patch_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


