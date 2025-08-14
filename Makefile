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
MAIN_GO := main.go


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

# We define the content of the rc.d script in a single variable with \n for
# newlines. This is compatible with both GNU and BSD Make.
# The double dollar signs ($$) escape the dollar sign for the shell.
RCD_SCRIPT_CONTENT := '#!/bin/sh\n\n# PROVIDE: $(SERVICE_NAME)\n# REQUIRE: LOGIN networking\n# KEYWORD: shutdown\n\n. /etc/rc.subr\n\nname="$(SERVICE_NAME)"\nrcvar="$${name}_enable"\n\npidfile="/var/run/$${name}.pid"\ntrigexmoe_user="$(SERVICE_USER)"\nprocname="$(INSTALL_PATH)"\nlogfile="/var/log/$${name}.log"\n\n# Define custom start, stop, and status commands\nstart_cmd="$${name}_start"\nstop_cmd="$${name}_stop"\nstatus_cmd="$${name}_status"\n\n# A placeholder command is needed for rc.subr to function correctly\ncommand="/usr/bin/true"\n\ntrigexmoe_start()\n{\n\t# Check if a PID file exists and if the process is actually running\n\tif [ -f "$${pidfile}" ] && kill -0 `cat $${pidfile}` 2>/dev/null; then\n\t\techo "$${name} is already running."\n\t\treturn 1\n\tfi\n\n\techo "Starting $${name}."\n\t# Use su to run the process as the correct user, redirecting output.\n\t# The command is run in a subshell `()` to ensure the `&` backgrounds it correctly.\n\tsu -m $${trigexmoe_user} -c "($${procname} > $${logfile} 2>&1 &)"\n\n\t# Give the process a moment to start up\n\tsleep 1\n\n\t# Find the PID of the new process and write it to the pidfile.\n\t# The pgrep pattern is anchored to match the exact process name.\n\tpgrep -u $${trigexmoe_user} -f "^$${procname}$$" > $${pidfile}\n}\n\ntrigexmoe_stop()\n{\n\tif [ ! -f "$${pidfile}" ] || ! kill -0 `cat $${pidfile}` 2>/dev/null; then\n\t\techo "$${name} is not running."\n\t\treturn 1\n\tfi\n\n\techo "Stopping $${name}."\n\t# Send the TERM signal to the process ID found in the pidfile\n\tkill `cat $${pidfile}`\n\t# Remove the pidfile\n\trm -f $${pidfile}\n}\n\ntrigexmoe_status()\n{\n\tif [ -f "$${pidfile}" ] && kill -0 `cat $${pidfile}` 2>/dev/null; then\n\t\techo "$${name} is running as pid `cat $${pidfile}`."\n\telse\n\t\techo "$${name} is not running."\n\tfi\n}\n\nload_rc_config $$name\n: $${trigexmoe_enable:="NO"}\n\nrun_rc_command "$$1"\n'

# This makes the RCD_SCRIPT_CONTENT variable available to shell commands invoked by make.
export RCD_SCRIPT_CONTENT


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
	@# Use printf with %b to interpret the newline characters in the variable.
	sudo sh -c 'printf -- "%b" "$$RCD_SCRIPT_CONTENT" > $(SERVICE_FILE)'
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
