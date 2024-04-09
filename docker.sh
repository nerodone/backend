#!/bin/bash

docker build -t nero_backend . &&
	docker run -d --env-file=.env -p 3000:3000 nero_backend
