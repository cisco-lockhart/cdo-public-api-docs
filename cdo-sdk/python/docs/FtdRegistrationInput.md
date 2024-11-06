# FtdRegistrationInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ftd_uid** | **str** | The unique identifier, represented as a UUID, of the FTD device in Security Cloud Control for which registration should be triggered. | 

## Example

```python
from cdo_sdk_python.models.ftd_registration_input import FtdRegistrationInput

# TODO update the JSON string below
json = "{}"
# create an instance of FtdRegistrationInput from a JSON string
ftd_registration_input_instance = FtdRegistrationInput.from_json(json)
# print the JSON string representation of the object
print(FtdRegistrationInput.to_json())

# convert the object into a dict
ftd_registration_input_dict = ftd_registration_input_instance.to_dict()
# create an instance of FtdRegistrationInput from a dict
ftd_registration_input_form_dict = ftd_registration_input.from_dict(ftd_registration_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


