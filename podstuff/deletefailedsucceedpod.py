import base64
from google.auth import compute_engine
import google.auth.transport.requests
import os
from google.cloud.container_v1 import ClusterManagerClient
from kubernetes import client

PROJECT_ID = os.environ["PROJECT_ID"]
REGION     = os.environ["REGION"]
CLUSTER_ID = os.environ["CLUSTER_ID"]
scopes = ['https://www.googleapis.com/auth/cloud-platform']
creds,proj = google.auth.default(scopes=scopes)
req = google.auth.transport.requests.Request()
creds.refresh(req)

cmc = ClusterManagerClient(credentials=creds)
gke_cluster = cmc.get_cluster(name=f'projects/{PROJECT_ID}/locations/{REGION}/clusters/{CLUSTER_ID}')

cert = base64.b64decode(gke_cluster.master_auth.cluster_ca_certificate)
cert_filename = 'cluster_ca_cert'
cert_file = open(cert_filename, 'wb')
cert_file.write(cert)
cert_file.close()

cluster_config = client.Configuration()
cluster_config.host = f"https://{gke_cluster.endpoint}:443"
cluster_config.ssl_ca_cert = cert_filename
cluster_config.api_key = {"authorization": "Bearer " + creds.token}
client.Configuration.set_default(cluster_config)
not_running = []
def main(ns):
    kubectl = client.CoreV1Api()
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
    main("default")
    os.remove("cluster_ca_cert")
