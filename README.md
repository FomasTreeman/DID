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

### Step 3: Test with docker-compose

** COMING SOON **

## Conclusion
You have now set up the Blockchain DID project. You can create DIDs and manage them using the blockchain. If you have any questions or need further assistance, feel free to reach out!
