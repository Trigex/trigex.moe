# trigex.moe Makefile
# ==============================================================================
# Build variables
# ==============================================================================
# Go command
GO_CMD := go

# Node command
NPM_CMD := npm

# Tools directory for local binaries.
# Use Make's built-in CURDIR variable to create an absolute path.
# This satisfies Go's requirement for an absolute GOBIN path.
TOOLS_DIR := $(CURDIR)/tools

# Templ command
# We define the full path to the locally installed templ binary.
TEMPL_CMD := $(TOOLS_DIR)/bin/templ

# sqlc command
SQLC_CMD := $(TOOLS_DIR)/bin/sqlc

# Tailwind/DaisyUI stylesheet build inputs/outputs
CSS_INPUT := assets/css/app.css
CSS_OUTPUT := assets/css/site.css
NPM_MARKER := node_modules/.npm-install-stamp

# Name of the output binary
BINARY_NAME := trigexmoe

# The main Go file that serves as the entry point
MAIN_GO := cmd/trigex.moe/main.go

# ==============================================================================
# Installation Variables (FreeBSD)
# ==============================================================================

# Check if the current user is root. If so, DOAS is empty. If not, it's "doas".
# This avoids running doas when already root.
DOAS := $(shell if [ `id -u` -eq 0 ]; then echo ""; else echo "doas"; fi)

# The name of the service and user
SERVICE_NAME := trigexmoe
SERVICE_USER := trigexmoe

# The name of the source rc.d file in the project directory
RCD_SOURCE_FILE := trigexmoe.rc

# Installation directories
INSTALL_DIR := /usr/local/bin
RCD_DIR := /usr/local/etc/rc.d
CONFIG_DIR := /usr/local/etc
SHARE_DIR := /usr/local/share/$(SERVICE_NAME)
DATA_DIR := $(SHARE_DIR)/data
DB_PATH := /var/db/$(SERVICE_NAME).sqlite

# Full paths for installation
INSTALL_PATH := $(INSTALL_DIR)/$(SERVICE_NAME)
SERVICE_FILE := $(RCD_DIR)/$(SERVICE_NAME)
CONFIG_PATH := $(CONFIG_DIR)/$(SERVICE_NAME).yaml
CONFIG_SOURCE_FILE := config.freebsd.yaml

# ==============================================================================
# Standard Targets
# ==============================================================================

# The default target, executed when you run make without arguments
.PHONY: all
all: build

# Build the application.
# This target depends on generated code and stylesheets.
.PHONY: build
build: generate sqlc css
	@echo "--> Building Go application..."
	$(GO_CMD) build -o $(BINARY_NAME) $(MAIN_GO)
	@echo "--> Build complete: $(BINARY_NAME)"

# Generate code using templ.
# This target has a file prerequisite: the templ binary itself.
.PHONY: generate
generate: $(TEMPL_CMD)
	@echo "--> Running 'templ generate'..."
	$(TEMPL_CMD) generate

# Generate sqlc Go code.
.PHONY: sqlc
sqlc: $(SQLC_CMD)
	@echo "--> Running 'sqlc generate'..."
	$(SQLC_CMD) generate

# Build the DaisyUI stylesheet.
.PHONY: css
css: $(CSS_OUTPUT)

package-lock.json: package.json
	@echo "--> Installing Node dependencies..."
	$(NPM_CMD) install

$(NPM_MARKER): package-lock.json
	@mkdir -p node_modules
	@touch $@

$(CSS_OUTPUT): $(CSS_INPUT) $(NPM_MARKER)
	@echo "--> Building DaisyUI stylesheet..."
	$(NPM_CMD) run build:css

# This is a file-based target. The commands here will only run if the file
# specified by $(TEMPL_CMD) does not exist.
$(TEMPL_CMD):
	@echo "--> templ command not found, installing locally to $(TOOLS_DIR)/bin..."
	@mkdir -p $(TOOLS_DIR)/bin
	@# By setting GOBIN to an absolute path, we satisfy `go install`'s requirement.
	GOBIN=$(TOOLS_DIR)/bin $(GO_CMD) install github.com/a-h/templ/cmd/templ@latest

$(SQLC_CMD):
	@echo "--> sqlc command not found, installing locally to $(TOOLS_DIR)/bin..."
	@mkdir -p $(TOOLS_DIR)/bin
	GOBIN=$(TOOLS_DIR)/bin $(GO_CMD) install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Run the compiled application locally.
.PHONY: run
run: build
	@echo "--> Running application locally..."
	./$(BINARY_NAME)

# Clean up build artifacts and local tools.
.PHONY: clean
clean:
	@echo "--> Cleaning up..."
	rm -f $(BINARY_NAME)
	rm -rf $(TOOLS_DIR)

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
	@if [ ! -f "$(RCD_SOURCE_FILE)" ]; then \
		echo "Error: Service file '$(RCD_SOURCE_FILE)' not found."; \
		echo "Please create it in the project directory."; \
		exit 1; \
	fi
	@if [ ! -f "$(CONFIG_SOURCE_FILE)" ]; then \
		echo "Error: Config template '$(CONFIG_SOURCE_FILE)' not found."; \
		exit 1; \
	fi
	@echo "--> Checking for service user '$(SERVICE_USER)'..."
	@if ! id -u $(SERVICE_USER) >/dev/null 2>&1; then \
		echo "--> Service user not found. Creating user '$(SERVICE_USER)'..."; \
		$(DOAS) pw useradd $(SERVICE_USER) -s /usr/sbin/nologin -d /nonexistent -c "Service user for $(SERVICE_NAME)" -w no; \
	else \
		echo "--> Service user already exists."; \
	fi
	@echo "--> Installing binary to $(INSTALL_PATH)..."
	$(DOAS) install -m 0755 $(BINARY_NAME) $(INSTALL_PATH)
	@echo "--> Installing rc.d service file to $(SERVICE_FILE)..."
	$(DOAS) install -m 0755 $(RCD_SOURCE_FILE) $(SERVICE_FILE)
	@echo "--> Ensuring FreeBSD data directory exists at $(DATA_DIR)..."
	$(DOAS) install -d -m 0755 -o $(SERVICE_USER) -g $(SERVICE_USER) $(DATA_DIR)
	@echo "--> Ensuring database file exists at $(DB_PATH)..."
	$(DOAS) touch $(DB_PATH)
	$(DOAS) chown $(SERVICE_USER):$(SERVICE_USER) $(DB_PATH)
	@echo "--> Installing config template to $(CONFIG_PATH) (if missing)..."
	@if [ ! -f "$(CONFIG_PATH)" ]; then \
		$(DOAS) install -m 0640 -o $(SERVICE_USER) -g $(SERVICE_USER) $(CONFIG_SOURCE_FILE) $(CONFIG_PATH); \
		echo "--> Installed default config file. Update admin credentials before starting service."; \
	else \
		echo "--> Config already exists, leaving it unchanged."; \
	fi
	@echo ""
	@echo "--> Installation complete."
	@echo "--> To enable the service, add trigexmoe_enable=\"YES\" to /etc/rc.conf"
	@echo "--> To start the service now, run: $(DOAS) service trigexmoe start"

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
	-$(DOAS) service $(SERVICE_NAME) stop
	@echo "--> Removing binary from $(INSTALL_PATH)..."
	$(DOAS) rm -f $(INSTALL_PATH)
	@echo "--> Removing rc.d service file from $(SERVICE_FILE)..."
	$(DOAS) rm -f $(SERVICE_FILE)
	@echo ""
	@echo "--> Uninstallation complete."
	@echo "--> NOTE: The service user '$(SERVICE_USER)' was not removed."
	@echo "--> To remove the user, run: make uninstall-user"

# Removes the service user.
.PHONY: uninstall-user
uninstall-user:
	@echo "--> Removing service user '$(SERVICE_USER)' from FreeBSD..."
	@if [ `uname -s` != "FreeBSD" ]; then \
		echo "Error: 'uninstall-user' target is only for FreeBSD systems."; \
		exit 1; \
	fi
	$(DOAS) pw userdel $(SERVICE_USER)

# ==============================================================================
# Help
# ==============================================================================

# A helper target to show the available commands.
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  all              - (Default) Build the application."
	@echo "  build            - Generate templ files and compile the Go source code."
	@echo "  generate         - Run 'templ generate' (will install templ if missing)."
	@echo "  run              - Build and run the application locally."
	@echo "  clean            - Remove the compiled binary and local tools."
	@echo "  install          - (FreeBSD only) Create user, install binary and rc.d service file."
	@echo "  uninstall        - (FreeBSD only) Remove binary and rc.d service file."
	@echo "  uninstall-user   - (FreeBSD only) Remove the service user account."
