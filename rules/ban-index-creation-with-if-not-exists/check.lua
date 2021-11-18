local function check(tree)
  local stmt, result = nil, {}
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    if stmt.IndexStmt and stmt.IndexStmt.if_not_exists then
      table.insert(result, statement)
    end
  end
  return result
end

return check
