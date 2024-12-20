# coding: utf-8

"""
    Cisco Security Cloud Control API

    Use the documentation to explore the endpoints Security Cloud Control has to offer

    The version of the OpenAPI document: 1.5.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, Field, StrictFloat, StrictInt, StrictStr, field_validator
from typing import Any, ClassVar, Dict, List, Optional, Union
from typing import Optional, Set
from typing_extensions import Self

class ChassisStatsHealthMetrics(BaseModel):
    """
    The chassis health metrics for the device.
    """ # noqa: E501
    fan1_rpm_avg: Optional[Union[StrictFloat, StrictInt]] = Field(default=None, description="The average speed of fan 1, if present, crucial for assessing the effectiveness of the system’s cooling mechanism under operational load.", alias="fan1RpmAvg")
    fan2_rpm_avg: Optional[Union[StrictFloat, StrictInt]] = Field(default=None, description="The average speed of fan 2, if present, crucial for assessing the effectiveness of the system’s cooling mechanism under operational load.", alias="fan2RpmAvg")
    fan3_rpm_avg: Optional[Union[StrictFloat, StrictInt]] = Field(default=None, description="The average speed of fan 3, if present, crucial for assessing the effectiveness of the system’s cooling mechanism under operational load.", alias="fan3RpmAvg")
    fan4_rpm_avg: Optional[Union[StrictFloat, StrictInt]] = Field(default=None, description="The average speed of fan 4, if present, crucial for assessing the effectiveness of the system’s cooling mechanism under operational load.", alias="fan4RpmAvg")
    psu1_output_status: Optional[StrictStr] = Field(default=None, alias="psu1OutputStatus")
    psu2_output_status: Optional[StrictStr] = Field(default=None, alias="psu2OutputStatus")
    psu1_input_status: Optional[StrictStr] = Field(default=None, alias="psu1InputStatus")
    psu2_input_status: Optional[StrictStr] = Field(default=None, alias="psu2InputStatus")
    psu1_fan_status: Optional[StrictStr] = Field(default=None, alias="psu1FanStatus")
    psu2_fan_status: Optional[StrictStr] = Field(default=None, alias="psu2FanStatus")
    __properties: ClassVar[List[str]] = ["fan1RpmAvg", "fan2RpmAvg", "fan3RpmAvg", "fan4RpmAvg", "psu1OutputStatus", "psu2OutputStatus", "psu1InputStatus", "psu2InputStatus", "psu1FanStatus", "psu2FanStatus"]

    @field_validator('psu1_output_status')
    def psu1_output_status_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['UP', 'DOWN']):
            raise ValueError("must be one of enum values ('UP', 'DOWN')")
        return value

    @field_validator('psu2_output_status')
    def psu2_output_status_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['UP', 'DOWN']):
            raise ValueError("must be one of enum values ('UP', 'DOWN')")
        return value

    @field_validator('psu1_input_status')
    def psu1_input_status_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['UP', 'DOWN']):
            raise ValueError("must be one of enum values ('UP', 'DOWN')")
        return value

    @field_validator('psu2_input_status')
    def psu2_input_status_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['UP', 'DOWN']):
            raise ValueError("must be one of enum values ('UP', 'DOWN')")
        return value

    @field_validator('psu1_fan_status')
    def psu1_fan_status_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['UP', 'DOWN']):
            raise ValueError("must be one of enum values ('UP', 'DOWN')")
        return value

    @field_validator('psu2_fan_status')
    def psu2_fan_status_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['UP', 'DOWN']):
            raise ValueError("must be one of enum values ('UP', 'DOWN')")
        return value

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of ChassisStatsHealthMetrics from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ChassisStatsHealthMetrics from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "fan1RpmAvg": obj.get("fan1RpmAvg"),
            "fan2RpmAvg": obj.get("fan2RpmAvg"),
            "fan3RpmAvg": obj.get("fan3RpmAvg"),
            "fan4RpmAvg": obj.get("fan4RpmAvg"),
            "psu1OutputStatus": obj.get("psu1OutputStatus"),
            "psu2OutputStatus": obj.get("psu2OutputStatus"),
            "psu1InputStatus": obj.get("psu1InputStatus"),
            "psu2InputStatus": obj.get("psu2InputStatus"),
            "psu1FanStatus": obj.get("psu1FanStatus"),
            "psu2FanStatus": obj.get("psu2FanStatus")
        })
        return _obj


