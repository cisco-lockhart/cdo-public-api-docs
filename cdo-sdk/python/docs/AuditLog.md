# AuditLog

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the Audit Log. | 
**event_type** | **str** | The type of the Audit Log event. | [optional] 
**username** | **str** | The name/email of the of user the Audit Log refers to. | [optional] 
**event_description** | **str** | The description of the Audit Log event. | [optional] 
**event_time** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the Audit Log event was created. | [optional] 
**roles** | **List[str]** | The roles of the user who did the Audit log operation | [optional] 

## Example

```python
from cdo_sdk_python.models.audit_log import AuditLog

# TODO update the JSON string below
json = "{}"
# create an instance of AuditLog from a JSON string
audit_log_instance = AuditLog.from_json(json)
# print the JSON string representation of the object
print(AuditLog.to_json())

# convert the object into a dict
audit_log_dict = audit_log_instance.to_dict()
# create an instance of AuditLog from a dict
audit_log_form_dict = audit_log.from_dict(audit_log_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


