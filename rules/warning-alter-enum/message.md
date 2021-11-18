# Problem

`alter enum add value` will lead to problem with drivers that will not receive information about new enum value, because they use oid-cache.

[SQL-ALTERTYPE](https://www.postgresql.org/docs/11/sql-altertype.html)

# Solutions

Instead of:
```sql
alter type job_type add value 'NewType';
```

Use:
```sql
-- nolint:warning-alter-enum
alter type job_type add value 'NewType';
```

And restart of all connected apps is also required.
