local function parseTree(content)
	local parser = require("parser")
	local result, err = parser.parse(content)
	if err then error(err) end
	return result
end

return parseTree
