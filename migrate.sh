#!/bin/bash

if [ $1 = "up" ]; then
	echo "Upgrade version $1"
	migrate -source file://./db/migrations/ -database 'mysql://root:123@tcp(localhost:3306)/tracking' up $2
elif [ $1 = "down" ]; then
	echo "Downgrade version $1"
	migrate -source file://./db/migrations/ -database 'mysql://root:123@tcp(localhost:3306)/tracking' down $2
fi
