BINARY_NAME = filemanager

SRC_FILES = main.go ui.go filemanager.go editor.go

.PHONY: all build run clean

all: build

# Сборка бинарного файла
build:
	@echo "Собираем проект..."
	@go build -o $(BINARY_NAME) $(SRC_FILES)
	@echo "Сборка завершена."

# Запуск программы
run: build
	@echo "Запускаем программу..."
	@./$(BINARY_NAME)

# Очистка скомпилированных файлов
clean:
	@echo "Чистим проект..."
	@rm -f $(BINARY_NAME)
	@echo "Очистка завершена."
