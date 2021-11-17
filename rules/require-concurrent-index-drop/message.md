# Problem

Удаление индекса - потенциально тяжелая операция, при выполнении запроса при создании индекса доступ к таблице
без ключевого слова `concurrently` будет заблокирован.

# Solutions

Instead of:
```sql
DROP INDEX "email_idx";
```

Use:
```sql
DROP INDEX CONCURRENTLY "email_idx"
```

[SQL-CREATEINDEX-CONCURRENTLY](https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY)
