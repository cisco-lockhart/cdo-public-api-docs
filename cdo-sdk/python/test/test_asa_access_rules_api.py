# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.1.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.api.asa_access_rules_api import ASAAccessRulesApi


class TestASAAccessRulesApi(unittest.TestCase):
    """ASAAccessRulesApi unit test stubs"""

    def setUp(self) -> None:
        self.api = ASAAccessRulesApi()

    def tearDown(self) -> None:
        pass

    def test_fetch_access_rule(self) -> None:
        """Test case for fetch_access_rule

        Get Access Rule
        """
        pass

    def test_list_access_rules(self) -> None:
        """Test case for list_access_rules

        Get Access Rules
        """
        pass


if __name__ == '__main__':
    unittest.main()
