#!/bin/bash

# Скрипт для экспорта PlantUML диаграмм в PNG
# Требуется: plantuml (установить через brew install plantuml или java -jar plantuml.jar)

echo "Экспорт PlantUML диаграмм в PNG..."

# Проверка наличия plantuml
if ! command -v plantuml &> /dev/null; then
    echo "Ошибка: plantuml не установлен"
    echo "Установите: brew install plantuml"
    echo "Или скачайте: http://plantuml.com/download"
    exit 1
fi

# Переход в директорию скрипта
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# Экспорт всех .puml файлов
for file in *.puml; do
    if [ -f "$file" ]; then
        filename=$(basename "$file" .puml)
        echo "Экспорт: $filename.puml -> $filename.png"
        plantuml -tpng "$file" -o .
    fi
done

echo "Готово! Все PlantUML диаграммы экспортированы в PNG."

