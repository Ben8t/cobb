

dev:
	@docker container run --rm -it -v $(PWD):/tmp --entrypoint bash -w /tmp golang:latest