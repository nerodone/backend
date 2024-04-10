#!/bin/bash

choices=(
	"docker"
	"pg"
	"air"
	"test"
	"tmux"
)

run_docker() {
	docker build -t nero_backend --label nero_backend .
	docker run -d --env-file=.env -p 3000:3000 nero_backend:latest
}

run_air() {
	if [ -f ".air.toml" ] && command -v air >/dev/null; then
		air -c .air.toml
	else
		echo "air.toml not found or air not installed"
	fi
}

run_pgcli() {
	source .env
	pgcli "$XATA_PG"
}

run_tests() {
	go test -v ./...
}

run_tmux() {
	if tmux has-session -t nero >/dev/null; then
		tmux attach-session -t nero
		return 0
	fi

	tmux new-session -d -s nero
	tmux send-keys -t nero "nvim" C-m

	tmux new-window -S -n build
	tmux send-keys -t nero "air -c .air.toml" C-m
	tmux new-window -S -n test
	tmux send-keys -t nero "go test ./..." C-m

	tmux select-window -t nero:0
}

use_fzf() {
	if command -v fzf >/dev/null; then
		return 0
	else
		return 1
	fi
}

use_standard() {
	usage=" "
	for choice in "${choices[@]}"; do
		usage="$usage$choice "
	done
	echo "usage: ./dev.sh [$usage]"
}

choose() {
	if use_fzf; then
		choice=$(printf '%s\n' "${choices[@]}" | fzf)
		handle_choice "$choice"
	else
		use_standard
	fi
}

handle_choice() {
	case $1 in
	"docker")
		run_docker
		;;
	"pg")
		run_pgcli
		;;
	"air")
		run_air
		;;
	"test")
		run_tests
		;;
	"tmux")
		run_tmux
		;;
	*)
		use_standard
		;;
	esac
}

#          ╭──────────────────────────────────────────────────────────╮
#          │                           Main                           │
#          ╰──────────────────────────────────────────────────────────╯
if [ $# -eq 0 ]; then
	choose
else
	handle_choice "$1"
fi
