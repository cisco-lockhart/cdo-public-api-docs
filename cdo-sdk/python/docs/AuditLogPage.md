# AuditLogPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[AuditLog]**](AuditLog.md) | The list of items retrieved. | [optional] 

## Example

```python
from cdo_sdk_python.models.audit_log_page import AuditLogPage

# TODO update the JSON string below
json = "{}"
# create an instance of AuditLogPage from a JSON string
audit_log_page_instance = AuditLogPage.from_json(json)
# print the JSON string representation of the object
print(AuditLogPage.to_json())

# convert the object into a dict
audit_log_page_dict = audit_log_page_instance.to_dict()
# create an instance of AuditLogPage from a dict
audit_log_page_form_dict = audit_log_page.from_dict(audit_log_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


