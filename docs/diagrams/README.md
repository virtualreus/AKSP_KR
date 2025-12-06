# Диаграммы архитектуры проекта

Все диаграммы созданы в формате Mermaid и могут быть визуализированы различными способами.

## Список диаграмм

1. **architecture-overview.md** - Схема основных компонентов системы
2. **microservices-interaction.md** - Схема взаимодействия между компонентами (sequence diagrams)
3. **api-contracts.md** - API и контракты между компонентами
4. **database-schema.md** - ER-диаграмма и модели данных
5. **deployment.md** - Схема физической инфраструктуры (Docker, CI/CD, Kubernetes)

## Способы визуализации

### 1. GitHub/GitLab (автоматически)

Диаграммы Mermaid автоматически отображаются в markdown файлах на GitHub и GitLab.

### 2. VS Code

Установите расширение "Markdown Preview Mermaid Support" для просмотра диаграмм в VS Code.

### 3. Онлайн редакторы

- [Mermaid Live Editor](https://mermaid.live/) - скопируйте код диаграммы и экспортируйте в PNG/SVG
- [Draw.io](https://app.diagrams.net/) - поддерживает импорт Mermaid

### 4. CLI инструменты

```bash
# Установка Mermaid CLI
npm install -g @mermaid-js/mermaid-cli

# Экспорт диаграммы в PNG
mmdc -i docs/diagrams/architecture-overview.md -o docs/diagrams/architecture-overview.png

# Экспорт всех диаграмм
for file in docs/diagrams/*.md; do
    mmdc -i "$file" -o "${file%.md}.png"
done
```

### 5. Для пояснительной записки

1. Откройте файл диаграммы в [Mermaid Live Editor](https://mermaid.live/)
2. Скопируйте код диаграммы (блок между `mermaid и `)
3. Вставьте в редактор
4. Экспортируйте в PNG (кнопка "Actions" → "Download PNG")
5. Вставьте изображение в Word документ

## Структура диаграмм

Каждая диаграмма содержит:

- Mermaid код для визуализации
- Текстовое описание компонентов
- Дополнительные детали и пояснения

## Примечания

- Все диаграммы соответствуют реальной архитектуре проекта
- Диаграммы можно редактировать напрямую в markdown файлах
- Для изменения диаграммы отредактируйте Mermaid код между тегами ```mermaid
