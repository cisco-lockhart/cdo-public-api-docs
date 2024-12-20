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

from pydantic import BaseModel, ConfigDict, Field, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from cdo_sdk_python.models.single_content import SingleContent
from typing import Optional, Set
from typing_extensions import Self

class GroupContent(BaseModel):
    """
    The contents of an object group
    """ # noqa: E501
    literals: Optional[List[SingleContent]] = Field(default=None, description="List of content literals")
    referenced_object_uids: Optional[List[StrictStr]] = Field(default=None, description="Set of UIDs of the group's referenced objects", alias="referencedObjectUids")
    __properties: ClassVar[List[str]] = ["literals", "referencedObjectUids"]

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
        """Create an instance of GroupContent from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in literals (list)
        _items = []
        if self.literals:
            for _item in self.literals:
                if _item:
                    _items.append(_item.to_dict())
            _dict['literals'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GroupContent from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "literals": [SingleContent.from_dict(_item) for _item in obj["literals"]] if obj.get("literals") is not None else None,
            "referencedObjectUids": obj.get("referencedObjectUids")
        })
        return _obj


