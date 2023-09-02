.DEFAULT_GOAL:=./bin/pomo
INSTALL_PATH=/usr/local/bin/pomo

./bin/pomo: main.go cache.go handlers.go
	@go build -o ./bin/pomo

clean:
	@rm -rd ./bin
	@rm -rd ${HOME}/.cache/pomodoro

install: ./bin/pomo
	@install ./bin/pomo ${INSTALL_PATH}

uninstall: clean
	@rm ${INSTALL_PATH}
