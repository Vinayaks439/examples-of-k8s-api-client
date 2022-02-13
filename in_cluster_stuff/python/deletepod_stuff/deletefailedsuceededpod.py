from kubernetes import client, config

def main(ns):
    config.load_incluster_config()

    v1 = client.CoreV1Api()
    try:
        list = kubectl.list_namespaced_pod(namespace=ns)
    except AttributeError  as e:
        print(f"Error {e}")
    else:
        for pod in list.items:
            phase = pod.status.phase
            if phase in ("Succeeded", "Failed"):
                not_running.append(pod.metadata.name)
        body = client.V1DeleteOptions()
        for i in not_running:
            try:
                print(f"deleting pod {i}")
                kubectl.delete_namespaced_pod(i, ns, body=body)
            except:
                print("Error maybe there are no pods to delete")
    finally:
        print(f"End of the function")

if __name__ == "__main__":
    current_namespace = open("/var/run/secrets/kubernetes.io/serviceaccount/namespace").read() # Unfortunately the k8s-client lib for python does not support currentnamespace when using in-cluster config https://github.com/kubernetes-client/python/issues/363
    main(current_namespace)
