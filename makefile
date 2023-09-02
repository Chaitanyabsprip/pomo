.DEFAULT_GOAL:=./bin/pomo

./bin/pomo: main.go cache.go handlers.go
	@go build -o ./bin/pomo

start: ./bin/pomo
	@./bin/pomo start

show: ./bin/pomo
	@./bin/pomo

stop: ./bin/pomo
	@./bin/pomo stop

clean:
	@rm ./bin/pomo
	@rm "${HOME}/.cache/pomodoro/*timer" 2>/dev/null

install: ./bin/pomo
	@install ./bin/pomo /usr/local/bin/pomo

uninstall: clean
	@rm ${HOME}/.config/bin/pomo
