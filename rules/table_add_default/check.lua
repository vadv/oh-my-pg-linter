local function check(tree)
  local stmt
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на alter table
    if stmt.AlterTableStmt then
      for _, cmd in pairs(stmt.AlterTableStmt.cmds) do
        if cmd.AlterTableCmd and cmd.AlterTableCmd and cmd.AlterTableCmd.def.ColumnDef.constraints then
          for _, constraint in pairs(cmd.AlterTableCmd.def.ColumnDef.constraints) do
            if constraint.Constraint and constraint.Constraint.contype == "CONSTR_NOTNULL" then
              return "Запрос: "..statement:query()..[[
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

]]
            end
          end
        end
      end
    end
  end
  return nil
end

return check
