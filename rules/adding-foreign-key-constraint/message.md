# Description

A foreign key constraint should be added with NOT VALID.
Adding a foreign key constraint requires a table scan and a SHARE ROW EXCLUSIVE lock on both tables, which blocks writes.
Adding the constraint as NOT VALID in one transaction and then using VALIDATE in another transaction will allow writes when adding the constraint.

# Problem

Adding a foreign key constraint requires a table scan and a SHARE ROW EXCLUSIVE lock on both tables, which blocks writes to each table.
This means no writes will be allowed to either table while the table you're altering is scanned to validate the constraint.

# Solution

To prevent blocking writes to tables, add the constraint as NOT VALID in one transaction, then VALIDATE CONSTRAINT in another.
While NOT VALID prevents row updates while running, it commits instantly if it can get a lock.
VALIDATE CONSTRAINT allows row updates while it scans the table.

# Adding constraint to existing table

Instead of:

```sql
ALTER TABLE "email" ADD CONSTRAINT "fk_user"
    FOREIGN KEY ("user_id") REFERENCES "user" ("id");
```

Use:

```sql
ALTER TABLE "email" ADD CONSTRAINT "fk_user"
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") NOT VALID;
ALTER TABLE "email" VALIDATE CONSTRAINT "fk_user";
```
