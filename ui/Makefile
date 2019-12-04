deps:
	@if [ ! -d "node_modules" ]; then yarn install; fi

build: deps
	./node_modules/brunch/bin/brunch b

serve: deps
	./node_modules/brunch/bin/brunch w --server
