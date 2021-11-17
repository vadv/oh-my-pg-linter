На версиях PostgreSQL ниже чем 11 версия, добавление поля с DEFAULT требует
полной перезаписи таблиц с `ACCESS EXCLUSIVE` локом.
https://www.postgresql.org/docs/10/sql-altertable.html#SQL-ALTERTABLE-NOTES
`ACCESS EXCLUSIVE` лок блокирует чтение/запись пока этот лок действует.

Решение:
Добавить колонку как null, потом установить default, заполнить ее, и удалить null.

Вместо:
ALTER TABLE "core_recipe" ADD COLUMN "foo" integer DEFAULT 10 NOT NULL;

Использовать:
ALTER TABLE "core_recipe" ADD COLUMN "foo" integer;
ALTER TABLE "core_recipe" ALTER COLUMN "foo" SET DEFAULT 10;
-- backfill column in batches
ALTER TABLE "core_recipe" ALTER COLUMN "foo" SET NOT NULL;
