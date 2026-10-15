# RemediationReportDto

Details of one terminal NAT remediation report.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**completed_at** | **datetime** | Time the remediation reached its final status, in UTC (RFC 3339). | [optional] 
**remediation_job_uid** | **str** | Identifier of the remediation job, useful for support and troubleshooting. | [optional] 
**report_url** | **str** | Download URL for the remediation PDF. Expires 30 minutes after generation; retrieve the report again to obtain a fresh URL. Null when no PDF is recorded for this attempt. | [optional] 
**started_at** | **datetime** | Time the remediation started, in UTC (RFC 3339). Omitted if unavailable. | [optional] 
**status** | **str** | Final status of the remediation attempt. | [optional] 
**triggered_by** | **str** | Username of the user who initiated the remediation. | [optional] 
**version** | **int** | Positive remediation-attempt version within the insight. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.remediation_report_dto import RemediationReportDto

# TODO update the JSON string below
json = "{}"
# create an instance of RemediationReportDto from a JSON string
remediation_report_dto_instance = RemediationReportDto.from_json(json)
# print the JSON string representation of the object
print(RemediationReportDto.to_json())

# convert the object into a dict
remediation_report_dto_dict = remediation_report_dto_instance.to_dict()
# create an instance of RemediationReportDto from a dict
remediation_report_dto_form_dict = remediation_report_dto.from_dict(remediation_report_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


