DIR = $(shell pwd)
BASENAME = $(shell basename ${DIR})
BACKEND = ${DIR}/backend

run:
	go run main.go -c config.toml -logtostderr=true
build:
	@echo ${DIR} 
	@echo ${BASENAME} 
	@echo ${BACKEND} 
	@cd ${BACKEND} && go build
