# CliMacroExecuteInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | List of UIDs of the devices to execute the CLI macro for. | 
**macro_uid** | **str** | The unique identifier, represented as a UUID, of the CLI macro. | [optional] 
**parameters** | **Dict[str, str]** | Parameters provided for the CLI macro execution as key-value pairs. | [optional] 

## Example

```python
from cdo_sdk_python.models.cli_macro_execute_input import CliMacroExecuteInput

# TODO update the JSON string below
json = "{}"
# create an instance of CliMacroExecuteInput from a JSON string
cli_macro_execute_input_instance = CliMacroExecuteInput.from_json(json)
# print the JSON string representation of the object
print(CliMacroExecuteInput.to_json())

# convert the object into a dict
cli_macro_execute_input_dict = cli_macro_execute_input_instance.to_dict()
# create an instance of CliMacroExecuteInput from a dict
cli_macro_execute_input_form_dict = cli_macro_execute_input.from_dict(cli_macro_execute_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


