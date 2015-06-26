default: build

build:
	goxc -bc="darwin,linux,windows" -d "."
