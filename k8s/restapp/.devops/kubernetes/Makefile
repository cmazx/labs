.PHONY: init
init:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update
	helm install -f ./.helm/database/values.yaml postgresql bitnami/postgresql
	helm install -f ./.helm/restapp/values.yaml restapp ./.helm/restapp
port-forward:
	kubectl port-forward --namespace default svc/postgresql 5432:5432