--[[
{
  CreateStmt = {
    oncommit = "ONCOMMIT_NOOP",
    relation = {
      inh = true,
      location = 14,
      relname = "app_user",
      relpersistence = "p"
    },
    tableElts = { {
        ColumnDef = {
          colname = "name",
          is_local = true,
          location = 31,
          typeName = {
            location = 38,
            names = { {
                String = {
                  str = "pg_catalog"
                }
              }, {
                String = {
                  str = "varchar"
                }
              } },
            typemod = -1,
            typmods = { {
                A_Const = {
                  location = 46,
                  val = {
                    Integer = {
                      ival = 100
                    }
                  }
                }
              } }
          }
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
    -- проверяем что statement на alter table
    if stmt.CreateStmt and stmt.CreateStmt.tableElts then
      for _, elts in pairs(stmt.CreateStmt.tableElts) do
        local is_pg_catalog, is_varchar = false, false
        if elts.ColumnDef and elts.ColumnDef.typeName and elts.ColumnDef.typeName.names then
          for _, name in pairs(elts.ColumnDef.typeName.names) do
            if name and name.String and name.String.str == "varchar" then
              is_varchar = true
            end
            if name and name.String and name.String.str == "pg_catalog" then
              is_pg_catalog = true
            end
          end
        end
        if is_varchar and is_pg_catalog then
          table.insert(result, statement)
        end
      end
    end
  end
  return result
end

return check
