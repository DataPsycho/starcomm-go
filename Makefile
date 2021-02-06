.PHONY: post serve buildd served killd removebin removed

app_name = starcomm-go:latest

post:
	curl -X POST -H "Content-Type: application/json" -d '{"from": "Saru", "message": "Saru to Discovery. Over!"}' http://0.0.0.0:8080/discovery

serve:
	clear
	rm -rf ./$(app_name) & \
	go build -o $(app_name) . & \
	./$(app_name)

buildd:
	docker build -t $(app_name) .

served:
	sudo docker run --detach -p 8080:8080 $(app_name)


killd:
	@echo 'Killing container...'
	docker ps | grep $(app_name) | awk '{print $$1}' | xargs -r docker stop

removebin:
	rm -rf $(app_name)
	clear

removed:
	docker ps -a | grep $(app_name) | awk '{print $$1}' | xargs docker rm


