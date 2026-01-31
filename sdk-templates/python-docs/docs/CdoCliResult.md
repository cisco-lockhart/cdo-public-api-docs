# CdoCliResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device. | 
**error_msg** | **str** | The error message, if any. | [optional] 
**execution_uid** | **str** | The identifier, represented as a UUID, of the CLI execution that produced the CLI execution result. Note: A single executionUid may be associated with one or more execution results. | [optional] 
**result** | **str** | The result of the CLI execution. | [optional] 
**script** | **str** | The script executed to generate the CLI result. | [optional] 
**start_time** | **datetime** | The time (in UTC) at which the user run the CLI execution, represented using the RFC-3339 standard. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the CLI execution result. | 

## Example

```python
from scc_firewall_manager_sdk.models.cdo_cli_result import CdoCliResult

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


