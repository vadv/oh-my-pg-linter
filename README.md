# Использование

Проверить миграции:

```shell
oh-my-pg-linter --rules=/etc/oh-my-pg-linter/rules check ./migrations/*.sql
```

Проверить тесты:

```shell
oh-my-pg-linter --rules=/etc/oh-my-pg-linter/rules test gin_fast_update
```

Запустить lua-файл:

```shell
oh-my-pg-linter run ./file.lua
```

# Как написать проверку

* Создать в директории rules директорию с названием правила.
* Создать файл check.lua который должен возвращать nil если проверка успешно завершена или строку если есть замечания.
* Создать файл test.lua который должен возвращать таблицу состоящую из таблиц { {sql = "text", passed = bool} }.
