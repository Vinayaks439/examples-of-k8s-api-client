from kubernetes import client, config


def loadconfig():
    config.load_incluster_config()

    v1 = client.CoreV1Api()
    return v1