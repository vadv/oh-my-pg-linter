local function check(tree)
  local stmt
  local result = {}
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на drop индекса
    if stmt.DropStmt and (stmt.DropStmt.removeType == "OBJECT_INDEX") then
      if not(stmt.DropStmt.concurrent) then
        table.insert(result, statement)
      end
    end
  end
  return result
end

return check
