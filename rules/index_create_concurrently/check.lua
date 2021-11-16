local function check(tree)
  local stmt
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на создание индекса
    if stmt.IndexStmt then
      if not(stmt.IndexStmt.concurrent) then
        return "Запрос: "..statement:query()..[[

Создание индекса - потенциально тяжелая операция, при выполнении запроса при создании индекса доступ к таблице
без ключевого слова `concurrently` будет заблокирован.

Решение:
Вместо: `CREATE INDEX "email_idx" ON "app_user" ("email")`.
Используйте: `CREATE INDEX CONCURRENTLY "email_idx" ON "app_user" ("email");`.

Документация: https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY

]]
      end
    end
  end
  return nil
end

return check
