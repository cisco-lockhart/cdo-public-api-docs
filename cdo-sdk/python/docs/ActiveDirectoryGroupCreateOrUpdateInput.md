# ActiveDirectoryGroupCreateOrUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the Active Directory Group. SCC does not support special characters for this field. | 
**role** | **str** | This determines the role for all the users included in this Active Directory Group. | [optional] 
**group_identifier** | **str** | The unique identifier, represented as a UUID, of the Active Directory Group in your Identity Provider (IdP). | 
**issuer_url** | **str** | The Identity Provider (IdP) URL, which Cisco Defense Orchestrator will use to validate SAML assertions during the sign-in process. | 
**notes** | **str** | Any notes that are applicable to this Active Directory Group. | [optional] 

## Example

```python
from cdo_sdk_python.models.active_directory_group_create_or_update_input import ActiveDirectoryGroupCreateOrUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of ActiveDirectoryGroupCreateOrUpdateInput from a JSON string
active_directory_group_create_or_update_input_instance = ActiveDirectoryGroupCreateOrUpdateInput.from_json(json)
# print the JSON string representation of the object
print(ActiveDirectoryGroupCreateOrUpdateInput.to_json())

# convert the object into a dict
active_directory_group_create_or_update_input_dict = active_directory_group_create_or_update_input_instance.to_dict()
# create an instance of ActiveDirectoryGroupCreateOrUpdateInput from a dict
active_directory_group_create_or_update_input_form_dict = active_directory_group_create_or_update_input.from_dict(active_directory_group_create_or_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


