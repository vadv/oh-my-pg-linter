# Использование

Проверить миграции:

```shell
oh-my-pg-linter check ./migrations/*.sql
```

Добавить директории с дополнительными проверками (переопределение - кто последний тот и папа):

```shell
oh-my-pg-linter check --rules=/path/to/additional/dir/with/rules:/and/another/dir/rules ./migrations/*.sql
```

Проверить тесты:

```shell
oh-my-pg-linter --rules=/etc/oh-my-pg-linter/rules test ban-gin-fast-update
```

```shell
oh-my-pg-linter --rules=/etc/oh-my-pg-linter/rules test-all
```

Запустить lua-файл (тестирование):

```shell
oh-my-pg-linter run ./file.lua
```

# Установка

```shell
go install github.com/vadv/oh-my-pg-linter@latest
```

# Как написать проверку

* Создать в директории rules директорию с названием правила.
* Создать файл check.lua который должен возвращать таблицу с замечаниями по запросам.
* Создать файл messages.md с описанием проблемы.
* Создать файл test.lua который должен возвращать таблицу состоящую из таблиц `{ {sql = "text", passed = bool} }`.

# Как добавить nolint

```sql
-- nolint:require-concurrent-index-creation,ban-gin-fast-update
create index on inventory using gin(groups) with (fastupdate = false);
```
