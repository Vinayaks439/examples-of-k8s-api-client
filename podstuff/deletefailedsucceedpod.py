from google.auth import compute_engine
import google.auth.transport.requests
import os
from google.cloud.container_v1 import ClusterManagerClient
from kubernetes import client

PROJECT_ID = os.environ["PROJECT_ID"]
REGION     = os.environ["REGION"]
CLUSTER_ID = os.environ["CLUSTER_ID"]
creds = compute_engine.Credentials()
req = google.auth.transport.requests.Request()
creds.refresh(req)

cmc = ClusterManagerClient(credentials=creds)
gke_cluster = cmc.get_cluster(name=f'projects/{PROJECT_ID}/locations/{REGION}/clusters/{CLUSTER_ID}')

cluster_config = client.Configuration()
cluster_config.host = f"https://{gke_cluster.endpoint}:443"
cluster_config.verify_ssl = False
cluster_config.api_key = {"authorization": "Bearer " + creds.token}
client.Configuration.set_default(cluster_config)
not_running = []
def main(ns):
    kubectl = client.CoreV1Api()
    try:
        list = kubectl.list_namespaced_pod(namespace=ns)
    except AttributeError  as e:
        print(f"Error {e}")
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

if __name__ == "__main__":
    main("default")
