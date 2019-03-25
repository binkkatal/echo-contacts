all: setup build run

setup:
	echo "Setting up the system....."
	. ./setup/setup.sh

build:
	echo "Building the project....."
	. ./setup/build.sh

run:
	./bin/echo-contacts