# Makefile for a Go project using templ

# ==============================================================================
# Project Variables
# ==============================================================================

# Go command
GO_CMD := go

# Go Path
# We get the GOPATH from the `go env` command to ensure we have the correct path.
GOPATH := $(shell go env GOPATH)

# Templ command
# We define the full path to the templ binary. This ensures we run the correct
# executable, even if GOPATH/bin is not in the system's PATH.
TEMPL_CMD := $(GOPATH)/bin/templ

# Name of the output binary
BINARY_NAME := trigexmoe

# The main Go file that serves as the entry point
MAIN_GO := cmd/main.go


# ==============================================================================
# Installation Variables (FreeBSD)
# ==============================================================================

# The name of the service and user
SERVICE_NAME := trigexmoe
SERVICE_USER := trigexmoe

# Installation directories
INSTALL_DIR := /usr/local/sbin
RCD_DIR := /usr/local/etc/rc.d

# Full paths for installation
INSTALL_PATH := $(INSTALL_DIR)/$(SERVICE_NAME)
SERVICE_FILE := $(RCD_DIR)/$(SERVICE_NAME)


# ==============================================================================
# FreeBSD rc.d Service File Content
# ==============================================================================

# We define the content of the rc.d script here. The double dollar signs ($$)
# are used to escape the single dollar signs for Make, so that they are
# correctly interpreted by the shell when the file is created.
define RCD_SCRIPT
#!/bin/sh

# PROVIDE: $(SERVICE_NAME)
# REQUIRE: LOGIN networking
# KEYWORD: shutdown

. /etc/rc.subr

name="$(SERVICE_NAME)"
rcvar="$${name}_enable"

pidfile="/var/run/$${name}.pid"
trigexmoe_user="$(SERVICE_USER)"
procname="$(INSTALL_PATH)"
logfile="/var/log/$${name}.log"

# Define custom start, stop, and status commands
start_cmd="$${name}_start"
stop_cmd="$${name}_stop"
status_cmd="$${name}_status"

# A placeholder command is needed for rc.subr to function correctly
command="/usr/bin/true"

trigexmoe_start()
{
	# Check if a PID file exists and if the process is actually running
	if [ -f "$${pidfile}" ] && kill -0 `cat $${pidfile}` 2>/dev/null; then
		echo "$${name} is already running."
		return 1
	fi

	echo "Starting $${name}."
	# Use su to run the process as the correct user, redirecting output.
	# The command is run in a subshell `()` to ensure the `&` backgrounds it correctly.
	su -m $${trigexmoe_user} -c "($${procname} > $${logfile} 2>&1 &)"

	# Give the process a moment to start up
	sleep 1

	# Find the PID of the new process and write it to the pidfile.
	# The pgrep pattern is anchored to match the exact process name.
	pgrep -u $${trigexmoe_user} -f "^$${procname}$$" > $${pidfile}
}

trigexmoe_stop()
{
	if [ ! -f "$${pidfile}" ] || ! kill -0 `cat $${pidfile}` 2>/dev/null; then
		echo "$${name} is not running."
		return 1
	fi

	echo "Stopping $${name}."
	# Send the TERM signal to the process ID found in the pidfile
	kill `cat $${pidfile}`
	# Remove the pidfile
	rm -f $${pidfile}
}

trigexmoe_status()
{
	if [ -f "$${pidfile}" ] && kill -0 `cat $${pidfile}` 2>/dev/null; then
		echo "$${name} is running as pid `cat $${pidfile}`."
	else
		echo "$${name} is not running."
	fi
}

load_rc_config $$name
: $${trigexmoe_enable:="NO"}

run_rc_command "$$1"
endef
# This makes the RCD_SCRIPT variable available to shell commands invoked by make.
export RCD_SCRIPT


# ==============================================================================
# Standard Targets
# ==============================================================================

# The default target, executed when you just run `make`.
.PHONY: all
all: build

# Build the application.
# This target depends on the 'generate' target.
.PHONY: build
build: generate
	@echo "--> Building Go application..."
	$(GO_CMD) build -o $(BINARY_NAME) $(MAIN_GO)
	@echo "--> Build complete: $(BINARY_NAME)"

# Generate code using templ.
# This target has a file prerequisite: the templ binary itself.
.PHONY: generate
generate: $(TEMPL_CMD)
	@echo "--> Running 'templ generate'..."
	$(TEMPL_CMD) generate

# This is a file-based target. The commands here will only run if the file
# specified by $(TEMPL_CMD) does not exist.
$(TEMPL_CMD):
	@echo "--> templ command not found, installing..."
	@# We ensure `go install` places the binary in the correct GOPATH/bin.
	GOBIN=$(GOPATH)/bin $(GO_CMD) install github.com/a-h/templ/cmd/templ@latest

# Run the compiled application locally.
.PHONY: run
run: build
	@echo "--> Running application locally..."
	./$(BINARY_NAME)

# Clean up build artifacts.
.PHONY: clean
clean:
	@echo "--> Cleaning up..."
	rm -f $(BINARY_NAME)


# ==============================================================================
# Deployment Targets (FreeBSD)
# ==============================================================================

# Install the application and service file.
# This target is intended for FreeBSD and requires root privileges.
.PHONY: install
install: build
	@echo "--> Installing for FreeBSD..."
	@if [ `uname -s` != "FreeBSD" ]; then \
		echo "Error: 'install' target is only for FreeBSD systems."; \
		exit 1; \
	fi
	@echo "--> Installing binary to $(INSTALL_PATH)..."
	sudo install -m 0755 $(BINARY_NAME) $(INSTALL_PATH)
	@echo "--> Installing rc.d service file to $(SERVICE_FILE)..."
	@# Create the rc.d script file using the variable defined above.
	sudo sh -c 'echo "$$RCD_SCRIPT" > $(SERVICE_FILE)'
	@# Make the rc.d script executable.
	sudo chmod 0755 $(SERVICE_FILE)
	@echo ""
	@echo "--> Installation complete."
	@echo "--> To enable the service, add trigexmoe_enable=\"YES\" to /etc/rc.conf"
	@echo "--> To start the service now, run: sudo service trigexmoe start"


# Uninstall the application and service file.
# This target is intended for FreeBSD and requires root privileges.
.PHONY: uninstall
uninstall:
	@echo "--> Uninstalling from FreeBSD..."
	@if [ `uname -s` != "FreeBSD" ]; then \
		echo "Error: 'uninstall' target is only for FreeBSD systems."; \
		exit 1; \
	fi
	@echo "--> Stopping service (if running)..."
	-sudo service $(SERVICE_NAME) stop
	@echo "--> Removing binary from $(INSTALL_PATH)..."
	sudo rm -f $(INSTALL_PATH)
	@echo "--> Removing rc.d service file from $(SERVICE_FILE)..."
	sudo rm -f $(SERVICE_FILE)
	@echo "--> Uninstallation complete."


# ==============================================================================
# Help
# ==============================================================================

# A helper target to show the available commands.
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  all         - (Default) Build the application."
	@echo "  build       - Generate templ files and compile the Go source code."
	@echo "  generate    - Run 'templ generate' (will install templ if missing)."
	@echo "  run         - Build and run the application locally."
	@echo "  clean       - Remove the compiled binary."
	@echo "  install     - (FreeBSD only) Install binary and rc.d service file."
	@echo "  uninstall   - (FreeBSD only) Remove binary and rc.d service file."

