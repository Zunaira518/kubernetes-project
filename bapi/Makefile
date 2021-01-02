build:
	docker build -t bapi .
run:
	-docker stop kubernetesProject
	-docker rm kubernetesProject
	docker run -d -p 8080:8080 --name kubernetesProject kubernetesProject