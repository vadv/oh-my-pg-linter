Создание индекса - потенциально тяжелая операция, при выполнении запроса при создании индекса доступ к таблице
без ключевого слова `concurrently` будет заблокирован.

Решение:
Вместо: `CREATE INDEX "email_idx" ON "app_user" ("email")`.
Используйте: `CREATE INDEX CONCURRENTLY "email_idx" ON "app_user" ("email");`.

Документация: https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY
