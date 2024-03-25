build:
	podman build --tag docker.io/vishnu930417/minikube:latest .
push:
	podman push docker.io/vishnu930417/minikube:latest
	for img in $(podman images --filter dangling=true | awk 'NR>1{print $3}'); do