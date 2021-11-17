--[[
{
  AlterTableStmt = {
    cmds = { {
        AlterTableCmd = {
          behavior = "DROP_RESTRICT",
          def = {
            Constraint = {
              conname = "positive_balance",
              contype = "CONSTR_CHECK",
              initially_valid = true,
              location = 21,
              raw_expr = {
                A_Expr = {
                  kind = "AEXPR_OP",
                  lexpr = {
                    ColumnRef = {
                      fields = { {
                          String = {
                            str = "balance"
                          }
                        } },
                      location = 58
                    }
                  },
                  location = 68,
                  name = { {
                      String = {
                        str = ">="
                      }
                    } },
                  rexpr = {
                    A_Const = {
                      location = 71,
                      val = {
                        Integer = {
                          ival = 0
                        }
                      }
                    }
                  }
                }
              }
            }
          },
          subtype = "AT_AddConstraint"
        }
      } },
    relation = {
      inh = true,
      location = 13,
      relname = "a",
      relpersistence = "p"
    },
    relkind = "OBJECT_TABLE"
  }
}
--]]
local function check(tree)
  local stmt
  local result = {}
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на alter table
    if stmt.AlterTableStmt and stmt.AlterTableStmt.cmds then
      for _, cmd in pairs(stmt.AlterTableStmt.cmds) do
        if cmd.AlterTableCmd and cmd.AlterTableCmd.def and cmd.AlterTableCmd.def.Constraint then
          if cmd.AlterTableCmd.def.Constraint.contype == "CONSTR_CHECK"
              and cmd.AlterTableCmd.def.Constraint.initially_valid then
            table.insert(result, statement)
          end
        end
      end
    end
  end
  return result
end

return check
