local function check(tree)
  local stmt
  local result = {}
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на создание индекса
    if stmt.IndexStmt then
      if stmt.IndexStmt.accessMethod == "gin" then
        local fastupdate = true
        if stmt.IndexStmt.options then
          for _, param in pairs(stmt.IndexStmt.options) do
            if param.DefElem and param.DefElem.defname == "fastupdate" and param.DefElem.arg
                and param.DefElem.arg.String and param.DefElem.arg.String.str == "false" then
              fastupdate = false
            end
          end
        end
        if fastupdate then
          table.insert(result, statement)
        end
      end
    end
  end
  return result
end

return check
