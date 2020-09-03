.PHONY: build clean publish

build: clean
	./scripts/build.sh

clean:
	rm -rf build
	rm -rf publish

publish: build
	./scripts/publish.sh