# Blockchain DID Project

## Overview
This project implements a decentralized identity (DID) service using a simple blockchain. The service allows for the creation and management of DIDs, storing them in a blockchain structure. This README provides instructions for setting up the project for the first time, including requirements and deployment steps.

## Requirements

### Software Requirements
- **Go**: Version 1.21 or higher
- **Docker**: For containerization
- **Kubernetes**: A local cluster setup using `kind` (Kubernetes in Docker)
- **kubectl**: Command-line tool for interacting with Kubernetes clusters

### Go Dependencies
- `github.com/google/uuid`: For generating unique identifiers

## Setup Instructions

### Step 1: Clone the Repository
Clone the project repository to your local machine:

```bash
git clone https://github.com/yourusername/blockchain-did.git
cd blockchain-did
```

### Step 2: Install Go Dependencies
Navigate to the project directory and install the required Go dependencies:

```bash
go mod tidy
```

### Step 4: Set Up Kind Cluster
Create a `kind` cluster:

```bash
kind create cluster --name blockchain-cluster
```

### Step 5: Deploy to Kubernetes
Apply the Kubernetes configurations to set up the deployment and services. Make sure you have your YAML files ready (e.g., `bc-deploy.yaml`, `bc-service.yaml`).

```bash
kubectl apply -f kubernetes/
```

### Step 6: Access the Application
To access the application, you can use port forwarding:

```bash
kubectl port-forward service/blockchain-service 8080:8080 -n blockchain
```

Now, you can access the application at `http://localhost:8080`.

### Step 7: Access the Kubernetes Dashboard (Optional)
If you want to manage your cluster visually, you can set up the Kubernetes Dashboard:

1. Deploy the dashboard:

   ```bash
   kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.5.1/aio/deploy/recommended.yaml
   ```

2. Create a service account and cluster role binding:

   ```bash
   kubectl create serviceaccount dashboard-admin -n kubernetes-dashboard
   kubectl create clusterrolebinding dashboard-admin --clusterrole=cluster-admin --serviceaccount=kubernetes-dashboard:dashboard-admin
   ```

3. Get the access token:

   ```bash
   kubectl create token dashboard-admin -n kubernetes-dashboard
   ```

4. Start the proxy:

   ```bash
   kubectl proxy
   ```

5. Access the dashboard at `http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/`.

### Step 8: Clean Up
When you're done, you can delete the `kind` cluster:

```bash
kind delete cluster --name blockchain-cluster
```

## Conclusion
You have now set up the Blockchain DID project. You can create DIDs and manage them using the blockchain. If you have any questions or need further assistance, feel free to reach out!
