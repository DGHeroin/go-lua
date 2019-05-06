local message = 'helloworld123'

glua.set(string.len(message), message)

local result = glua.get()
print('get>> ', result)
local count = 0
while true do
    count = count + 1
    glua.get()
    if count > 10*1000*1000 then break end
end