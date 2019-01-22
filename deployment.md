## Deployment
### Initial deployment to Google Kubernetes Engine
Execute the following from the Google Cloud Platform console.
For more details, see https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app.

1) Clone the repo
```
git clone https://github.com/MichaelReiter/snapshot-api-gateway.git
```

2) Change directory
```
cd snapshot-api-gateway
```

3) Set the PROJECT_ID environment variable
```
export PROJECT_ID="$(gcloud config get-value project -q)"
```

4) Build the Docker container
```
docker build -t gcr.io/${PROJECT_ID}/snapshot-api-gateway .
```

5) Authorize Google Cloud (one time sign in)
```
gcloud auth configure-docker
```

6) Upload the Docker image to the Google Container Registry
```
docker push gcr.io/${PROJECT_ID}/snapshot-api-gateway
```

7) Create a container cluster
```
gcloud container clusters create snapshot --num-nodes=3
```

8) List container instances
```
gcloud compute instances list
```

9) Deploy the application to Kubernetes Engine
```
kubectl run hello-web --image=gcr.io/${PROJECT_ID}/snapshot-api-gateway --port 8000
```

10) Expose the application to the internet
```
kubectl expose deployment snapshot --type=LoadBalancer --port 80 --target-port 8000
```

11) View service details
```
kubectl get service
```

### Scaling the application
- Update replicas
```
kubectl scale deployment snapshot --replicas=3
```

- See Kubernetes pod details
```
kubectl get pods
```

### Deploy a new app version
1) Build an updated image
```
docker build -t gcr.io/${PROJECT_ID}/snapshot-api-gateway:v2 .
```

2) Push the image to the Google Container Registry
```
docker push gcr.io/${PROJECT_ID}/snapshot-api-gateway:v2
```

3) Apply a rolling update to the existing deployment
```
kubectl set image deployment/snapshot 
```

### Take down the deployment
1) Delete the service
```
kubectl delete service hello-web
```

2) Delete the container cluster
```
gcloud container clusters delete snapshot
```
