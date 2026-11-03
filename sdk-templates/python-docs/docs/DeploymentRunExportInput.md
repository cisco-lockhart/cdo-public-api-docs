# DeploymentRunExportInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**query** | **str** | The query to execute in order to generate the export. Use the Lucene Query Syntax to construct your query. You can use this to filter the deployment runs included in your export. Filter on the searchable fields of a deployment run, as with the list endpoint. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.deployment_run_export_input import DeploymentRunExportInput

# TODO update the JSON string below
json = "{}"
# create an instance of DeploymentRunExportInput from a JSON string
deployment_run_export_input_instance = DeploymentRunExportInput.from_json(json)
# print the JSON string representation of the object
print(DeploymentRunExportInput.to_json())

# convert the object into a dict
deployment_run_export_input_dict = deployment_run_export_input_instance.to_dict()
# create an instance of DeploymentRunExportInput from a dict
deployment_run_export_input_form_dict = deployment_run_export_input.from_dict(deployment_run_export_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


