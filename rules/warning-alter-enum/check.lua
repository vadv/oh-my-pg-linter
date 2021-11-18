--[[
{
  AlterEnumStmt = {
    newVal = "AnsibleLocalRunJob",
    newValIsAfter = true,
    typeName = { {
        String = {
          str = "job_type"
        }
      } }
  }
}
--]]
local function check(tree)
  local stmt
  local result = {}
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    if stmt.AlterEnumStmt and stmt.AlterEnumStmt.newVal then
      table.insert(result, statement)
    end
  end
  return result
end

return check
