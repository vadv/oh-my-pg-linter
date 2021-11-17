Удаление индекса - потенциально тяжелая операция, при выполнении запроса при создании индекса доступ к таблице
без ключевого слова `concurrently` будет заблокирован.

Решение:
Вместо: `DROP INDEX "email_idx";`.
Используйте: `DROP INDEX CONCURRENTLY "email_idx";`.

Документация: https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY
