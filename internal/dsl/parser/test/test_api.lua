local inspect = require("inspect")
local parser = require("parser")

local result, err = parser.parse([[
-- nolint:concurrently_if_not_exists
create index concurrently if not exists item_company_id_spu_idx
    on item using btree ("company_id", "spu");
ALTER TABLE "email" ADD CONSTRAINT "fk_user"
    FOREIGN KEY ("user_id") REFERENCES "user" ("id");
ALTER TABLE "email" ADD CONSTRAINT "fk_user"
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") NOT VALID;
]])
if err then error(err) end

local foundCreateIndex = false
for _, statement in pairs(result) do
  local stmt = statement:tree()
  if stmt.IndexStmt then
    foundCreateIndex = true
    if not(stmt.IndexStmt.accessMethod == "btree") then error (inspect(statement)) end
    if not(stmt.IndexStmt.idxname == "item_company_id_spu_idx") then error (inspect(statement)) end
    if not(stmt.IndexStmt.if_not_exists) then error (inspect(statement)) end
    if not(stmt.IndexStmt.concurrent) then error (inspect(statement)) end
    if not(statement:query() == [[-- nolint:concurrently_if_not_exists
create index concurrently if not exists item_company_id_spu_idx
    on item using btree ("company_id", "spu")]]) then
      error(statement:query())
    end
  end
end

if not foundCreateIndex then error("create index not found") end
