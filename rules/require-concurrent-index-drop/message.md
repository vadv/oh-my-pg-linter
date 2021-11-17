Удаление индекса - потенциально тяжелая операция, при выполнении запроса при создании индекса доступ к таблице
без ключевого слова `concurrently` будет заблокирован.

Вместо:
```sql
DROP INDEX "email_idx";
```

Используйте:
```sql
DROP INDEX CONCURRENTLY "email_idx"
```

[SQL-CREATEINDEX-CONCURRENTLY](https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY)
