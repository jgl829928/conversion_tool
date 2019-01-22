# 判断数据类型
function judgeType(value)
    if(typeof(value)==String)
        return "string"
    elseif(float(value) % 1 == 0.0)
        return "int32"
    elseif(float(value) % 1 != 0.0)
        return "float32"
    end
end

println(judgeType("12"))