BUILD_DIR := build
SERVER_OUT := $(BUILD_DIR)/server

$(BUILD_DIR):
	@mkdir -p $(BUILD_DIR)

$(SERVER_OUT): server.c | $(BUILD_DIR)
	@gcc server.c -o $(SERVER_OUT)

run: $(SERVER_OUT)
	@./$(SERVER_OUT)

all: run