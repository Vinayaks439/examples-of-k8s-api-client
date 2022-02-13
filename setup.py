from setuptools import find_packages, setup
setup(
    name='examples-of-k8s-client',
    packages=find_packages(include=['inclusterconfigmod']),
    version='0.1.0',
    description='Examples of K8s Client By Vinayak',
    author='Me',
    setup_requires=['pytest-runner'],
    tests_require=['pytest==4.4.1'],
    url="https://github.com/Vinayaks439/examples-of-k8s-api-client",
    license='MIT',
    install_requires=["google-auth","google-cloud-container","kubernetes","requests"],
)