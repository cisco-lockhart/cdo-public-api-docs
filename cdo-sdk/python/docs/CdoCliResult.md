# CdoCliResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the CLI execution result. | 
**result** | **str** | The result of the CLI execution. | [optional] 
**error_msg** | **str** | The error message, if any. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device. | 
**start_time** | **datetime** | The time (in UTC) at which the user run the CLI execution, represented using the RFC-3339 standard. | [optional] 
**script** | **str** | The script executed to generate the CLI result. | [optional] 

## Example

```python
from cdo_sdk_python.models.cdo_cli_result import CdoCliResult

# TODO update the JSON string below
json = "{}"
# create an instance of CdoCliResult from a JSON string
cdo_cli_result_instance = CdoCliResult.from_json(json)
# print the JSON string representation of the object
print(CdoCliResult.to_json())

# convert the object into a dict
cdo_cli_result_dict = cdo_cli_result_instance.to_dict()
# create an instance of CdoCliResult from a dict
cdo_cli_result_form_dict = cdo_cli_result.from_dict(cdo_cli_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


