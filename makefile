.DEFAULT_GOAL:=./bin/pomo
INSTALL_PATH=/usr/local/bin/pomo

./bin/pomo: cmd/pomo/main.go internal/cache.go internal/handlers.go
	@go build -o ./bin/pomo ./cmd/pomo

clean:
	@rm -rd ./bin
	@rm -rd ${HOME}/.cache/pomodoro

install: ./bin/pomo
	@install ./bin/pomo ${INSTALL_PATH}

uninstall: clean
	@rm ${INSTALL_PATH}
