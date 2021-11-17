local function check(tree)
  local result = {}
  for _, statement in pairs(tree) do
    table.insert(result, statement)
  end
  return result
end

return check
