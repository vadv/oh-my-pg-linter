# filepath

## Usage

```lua
local filepath = require("filepath")

-- filepath.ext()
local result = filepath.ext("/var/tmp/file.name")
if not(result == ".name") then error("ext") end

-- filepath.basename()
local result = filepath.basename("/var/tmp/file.name")
if not(result == "file.name") then error("basename") end

-- filepath.dir()
local result = filepath.dir("/var/tmp/file.name")
if not(result == "/var/tmp") then error("dir") end

-- filepath.join()
local result = filepath.join("/var", "tmp", "file.name")
if not(result == "/var/tmp/file.name") then error("join") end

-- filepath.glob()
local result = filepath.glob("/var/*/*.name")
if not(result[1] == "/var/tmp/file.name") then error("glob") end

-- filepath.work_dir()
local wd, err = filepath.work_dir()

-- filepath.script_dir()
local sd, err = filepath.script_dir()
```

