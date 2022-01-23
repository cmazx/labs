Для инициализации выполнить

```make init```

или 

	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update
	helm install -f ./.helm/database/values.yaml postgresql bitnami/postgresql
	helm install -f ./.helm/restapp/values.yaml restapp ./.helm/restapp

Коллекция postman: [restapp.postman_collection.json](restapp.postman_collection.json)