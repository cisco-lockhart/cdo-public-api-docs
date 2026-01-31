# ExportInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**query** | **str** | The query to execute in order to generate the export. Use the Lucene Query Syntax to construct your query. You can use this to filter the entities included in your export. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.export_input import ExportInput

# TODO update the JSON string below
json = "{}"
# create an instance of ExportInput from a JSON string
export_input_instance = ExportInput.from_json(json)
# print the JSON string representation of the object
print(ExportInput.to_json())

# convert the object into a dict
export_input_dict = export_input_instance.to_dict()
# create an instance of ExportInput from a dict
export_input_form_dict = export_input.from_dict(export_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


