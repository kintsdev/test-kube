# test-kube

This is a simple Go application designed to run in a Kubernetes cluster. The application prints the following environment variables to the console:

- Pod Name
- Namespace
- Node Name
- Host IP
- Pod IP

## Prerequisites

- Docker (optional)
- Kubernetes cluster
- kubectl configured to interact with your cluster
- Ingress controller installed (e.g., NGINX Ingress Controller)

## Usage

```bash
kubectl apply -f deployment.yaml
```

## License

[MIT](LICENSE)