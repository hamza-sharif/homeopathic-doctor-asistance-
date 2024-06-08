
export PROJECT = homeopathic-doctor-asistance-
DEV_IMAGE := ${PROJECT}_dev
IMPORT_PATH := github.com/hamza-sharif/${PROJECT}

# all non-windows environments
ROOT := $(shell pwd)
DOCKER_RUN := docker run --rm \
	-v ${ROOT}:/${PROJECT}/src/${IMPORT_PATH} \
	${DEV_IMAGE}
# Default target
all: build

# Build the project
build: gen-code
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)
	@echo "Build completed."

# Run tests
test:
	@echo "Running tests..."
	@go test $(PKG)
	@echo "Tests completed."

# Clean the build directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR) $(SWAGGER_GEN_DIR)
	@echo "Cleanup completed."

# Install the application
install:
	@echo "Installing $(APP_NAME)..."
	@go install $(SRC_DIR)
	@echo "Installation completed."

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

# Generate Swagger code
gen-code:
	@echo "Generating Swagger code..."
	${DOCKRUN} bash ./scripts/swagger.sh
	@echo "Swagger code generation completed."

# Display help
help:
	@echo "Usage:"
	@echo "  make         - Build the project"
	@echo "  make build   - Build the project"
	@echo "  make test    - Run tests"
	@echo "  make clean   - Clean the build directory"
	@echo "  make install - Install the application"
	@echo "  make run     - Build and run the application"
	@echo "  make gen     - Generate Swagger code"
	@echo "  make help    - Display this help message"
