print("hello world lua")

function test()
    print("test func")
end

function test1(val)
    return val
end

function getValue(...)
    result  = ''
    print("result--->>>>"..result)

    local arg={...}
    for _,v in ipairs(arg) do
        print("val--->>>>"..v)
        result = result .. v
    end
    print("总共传入 " .. select("#",...) .. " 个数")
    local value = redis.call("Get",arg[1])
    print("当前值"..value)
    if( value - arg[2] >= 0 ) then
        local leftStock = redis.call("DecrBy" , arg[1],arg[2])
        print("剩余值为" .. leftStock );
        return leftStock
    else
        print("数量不够，无法扣减");
        return value - arg[2]
    end
    return -1
end

function getScriptValue(...)
    result  = ''
    print("result--->>>>"..result)

    local arg={...}
    for _,v in ipairs(arg) do
        print("val--->>>>"..v)
        result = result .. v
    end
    print("总共传入 " .. select("#",...) .. " 个数")
    return result
end