# MspExportInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**query** | **str** | The query to execute in order to generate the export. Use the Lucene Query Syntax to construct your query. You can use this to filter the entities included in your export. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_export_input import MspExportInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspExportInput from a JSON string
msp_export_input_instance = MspExportInput.from_json(json)
# print the JSON string representation of the object
print(MspExportInput.to_json())

# convert the object into a dict
msp_export_input_dict = msp_export_input_instance.to_dict()
# create an instance of MspExportInput from a dict
msp_export_input_form_dict = msp_export_input.from_dict(msp_export_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


