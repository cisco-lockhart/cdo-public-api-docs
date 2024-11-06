# MfaEventPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The Security Cloud Control API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[MfaEvent]**](MfaEvent.md) | The list of items retrieved. | [optional] 

## Example

```python
from cdo_sdk_python.models.mfa_event_page import MfaEventPage

# TODO update the JSON string below
json = "{}"
# create an instance of MfaEventPage from a JSON string
mfa_event_page_instance = MfaEventPage.from_json(json)
# print the JSON string representation of the object
print(MfaEventPage.to_json())

# convert the object into a dict
mfa_event_page_dict = mfa_event_page_instance.to_dict()
# create an instance of MfaEventPage from a dict
mfa_event_page_form_dict = mfa_event_page.from_dict(mfa_event_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


