#!/bin/bash
set -e

PID_FILE="/opt/homebrew/var/postgresql@17/postmaster.pid"

if [ -f "$PID_FILE" ]; then
	OLD_PID=$(head -1 "$PID_FILE")

	if ! ps -p $OLD_PID >/dev/null 2 >&1; then
		echo "Removing stale PID file..."
		rm -f "$PID_FILE"
	fi
fi

chown -R postgres:postgres /opt/homebrew/var/postgresql@17
chmod 700 /opt/homebrew/var/postgresql@17

exec docker-entrypoint.sh postgres
