pip install setuptools
scripts/cli update-sdk-description

secret_arn=arn:aws:secretsmanager:us-west-2:107042026245:secret:jenkins-pypi-credentials-mD4NdK
echo -n "$(yellow Retrieving the secret value from AWS Secrets Manager)... "
# secret_value=$(aws secretsmanager get-secret-value --secret-id $secret_arn --query SecretString --output text)
secret_value=pypi-AgEIcHlwaS5vcmcCJDdlYzM4ZTVkLTExMTUtNGVhZi1iZWJjLWE4MzhjODU5MjVkZQACFlsxLFsiY2RvLXNkay1weXRob24iXV0AAixbMixbIjIzZWEyMmNiLWM1MDQtNGJjMC05NWFjLWNiOTA4MDdkNzZhYyJdXQAABiCbLPx9Y8wSqob6hGNZrr6E5O1HfuserkMBLZBNH5TGsg

package_name=$(grep -Eo 'NAME = "[^"]+"' cdo-sdk/python/setup.py | cut -d'"' -f2)
current_version=$(grep -Eo "VERSION = \"[0-9]+\.[0-9]+\.[0-9]+\"" cdo-sdk/python/setup.py | cut -d'"' -f2)

cd cdo-sdk/python
pip3 install wheel twine

echo -n "$(yellow Creating the Wheel and Source)... "
python3 setup.py sdist bdist_wheel

echo -n "$(yellow Publishing to PyPI)... "
# /home/jenkins/.local/bin/twine upload --username __token__ --password $secret_value dist/${package_name}-${current_version}.tar.gz dist/${package_name//-/_}-${current_version}-py3-none-any.whl
twine upload --username __token__ --password $secret_value dist/${package_name}-${current_version}.tar.gz dist/${package_name//-/_}-${current_version}-py3-none-any.whl

unset secret_value