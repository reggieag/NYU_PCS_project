from setuptools import setup, find_packages

VERSION = '0.0.1'
DESCRIPTION = 'API library for Fuzzy Wuzzy PCS Project'

setup(
    name="api",
    version=VERSION,
    author="Zhengqi Xi",
    description=DESCRIPTION,
    packages=find_packages(),
    install_requires=[
        'attrs==20.3.0',
        'certifi==2020.12.5',
        'chardet==4.0.0',
        'idna==2.10',
        'isodate==0.6.0',
        'jsonschema==3.2.0',
        'oauthlib==3.1.0',
        'openapi-schema-validator==0.1.5',
        'openapi-spec-validator==0.3.0',
        'prance==0.20.2',
        'pyrsistent==0.17.3',
        'PyYAML==5.4.1',
        'requests==2.25.1',
        'requests-oauthlib==1.3.0',
        'semver==2.13.0',
        'six==1.15.0',
        'urllib3==1.26.4',
    ],
)