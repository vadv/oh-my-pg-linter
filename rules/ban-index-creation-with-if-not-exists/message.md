# Problem

Using `IF NOT EXISTS` with `CONCURRENTLY` is dangerous.
This is a problem of non-idempotency of the migration, because restarting of the migration can lead to `invalid` index state.

Instead of:

```sql
CREATE INDEX CONCURRENTLY IF NOT EXISTS dist_id_temp_idx ON distributors (dist_id);
```

Use:

```sql
DO
$$
    BEGIN
        IF EXISTS (
              SELECT indexrelid::oid::regclass FROM pg_index WHERE NOT indisvalid AND indexrelid::oid::regclass::TEXT = 'dist_id_temp_idx'
            )
        THEN
            DROP INDEX dist_id_temp_idx;
        END IF;
        IF NOT EXISTS (
              SELECT indexrelid::oid::regclass FROM pg_index WHERE indexrelid::oid::regclass::TEXT = 'dist_id_temp_idx'
            )
        THEN
            CREATE INDEX CONCURRENTLY dist_id_temp_idx ON distributors (dist_id);
        END IF;
    END
$$;
```
