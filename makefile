CACHEFILE=${HOME}/.cache/pomotimer

.DEFAULT_GOAL:=./bin/pomo

./bin/pomo: main.go cache.go handler.go
	@go build -o ./bin/pomo

start: ./bin/pomo
	@./bin/pomo start

show: ./bin/pomo
	@./bin/pomo

stop: ./bin/pomo ${CACHEFILE}
	@./bin/pomo stop

clean:
	@rm ./bin/pomo
	@if [ -f "${CACHEFILE}" ]; then rm "${CACHEFILE}"; fi

install: ./bin/pomo
	@install ./bin/pomo /usr/local/bin/pomo

uninstall: clean
	@rm ${HOME}/.config/bin/pomo
