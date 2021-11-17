# Problem

Ensure all index deletions use the `CONCURRENTLY` option.
This rule ignores indexes added to tables deleted in the same transaction.
During a normal index delete updates are blocked. `CONCURRENTLY` avoids the issue of blocking.

[SQL-CREATEINDEX-CONCURRENTLY](https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY)

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
