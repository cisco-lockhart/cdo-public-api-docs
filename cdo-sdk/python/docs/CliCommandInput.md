# CliCommandInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | List of UIDs of the devices to execute the CLI script for. | 
**script** | **str** | The script executed to generate the CLI result. | 

## Example

```python
from cdo_sdk_python.models.cli_command_input import CliCommandInput

# TODO update the JSON string below
json = "{}"
# create an instance of CliCommandInput from a JSON string
cli_command_input_instance = CliCommandInput.from_json(json)
# print the JSON string representation of the object
print(CliCommandInput.to_json())

# convert the object into a dict
cli_command_input_dict = cli_command_input_instance.to_dict()
# create an instance of CliCommandInput from a dict
cli_command_input_form_dict = cli_command_input.from_dict(cli_command_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


