# coding: utf-8

"""
    Cisco Security Cloud Control API

    Use the documentation to explore the endpoints Security Cloud Control has to offer

    The version of the OpenAPI document: 1.5.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.universal_ztna_settings import UniversalZtnaSettings

class TestUniversalZtnaSettings(unittest.TestCase):
    """UniversalZtnaSettings unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> UniversalZtnaSettings:
        """Test UniversalZtnaSettings
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `UniversalZtnaSettings`
        """
        model = UniversalZtnaSettings()
        if include_optional:
            return UniversalZtnaSettings(
                auto_deploy_enabled = False,
                domain_settings = [
                    cdo_sdk_python.models.domain_settings.DomainSettings(
                        fqdn = 'myftd.cisco.com', 
                        certificate = cdo_sdk_python.models.fmc_object_reference.FmcObjectReference(
                            uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                            link = 'https://us.manage.security.cisco.com/api/rest/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/object/networks/0276f012-3875-0ed3-0000-004294981364', ), 
                        interfaces = [
                            cdo_sdk_python.models.fmc_object_reference.FmcObjectReference(
                                uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                                link = 'https://us.manage.security.cisco.com/api/rest/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/object/networks/0276f012-3875-0ed3-0000-004294981364', )
                            ], )
                    ],
                source_nat_v4 = cdo_sdk_python.models.fmc_object_reference.FmcObjectReference(
                    uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                    link = 'https://us.manage.security.cisco.com/api/rest/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/object/networks/0276f012-3875-0ed3-0000-004294981364', ),
                source_nat_v6 = cdo_sdk_python.models.fmc_object_reference.FmcObjectReference(
                    uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                    link = 'https://us.manage.security.cisco.com/api/rest/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/object/networks/0276f012-3875-0ed3-0000-004294981364', )
            )
        else:
            return UniversalZtnaSettings(
        )
        """

    def testUniversalZtnaSettings(self):
        """Test UniversalZtnaSettings"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()