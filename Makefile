export RULES_DIR=./rules

lint: lua-lint go-lint

lua-lint:
	luacheck . --codes --globals os

go-lint:
	golangci-lint run --deadline=5m -v

test-rules: build
	./bin/oh-my-pg-linter -r $(RULES_DIR) test-all

build:
	go build -o ./bin/oh-my-pg-linter ./cmd

test: lua-lint test-rules
	go test ./... -v -race

install: build
	mkdir -p $(DESTDIR)/usr/bin/ $(DESTDIR)/etc/oh-my-pg-linter/rules
	install -c -m 755 ./bin/oh-my-pg-linter $(DESTDIR)/usr/bin/
	@cp -av ./rules/* $(DESTDIR)/etc/oh-my-pg-linter
