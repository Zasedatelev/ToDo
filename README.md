# Тестовое задание для SkillsRock

## Необходимые условия

*   Установленный Go (версия 1.23.6 или выше)
*   Установленный Make
*   Настроенная база данных PostgreSQL (или другая, которую вы используете) и строка подключения.

## Установка и запуск

1.  **Клонируйте репозиторий:**

    ```bash
    git clone https://github.com/Zasedatelev/ToDo.git
    cd <имя_папки_репозитория>
    ```

2.  **Скачайте зависимости:**

    ```bash
    make init-deps
    ```

3.  **(Опционально) Настройте переменные окружения:**

    Вам может потребоваться настроить переменные окружения для подключения к базе данных через файл `local.env` или передается напрямую в командной строке.

## Запустите миграции создания базы данных:

    ```bash
    make migrateUp
    ```

## Команда для откатита миграции:

    ```bash
    make migrateDown
    ```

### Запуск приложения

Чтобы запустить приложение, выполните следующую команду:

```bash
make start
```