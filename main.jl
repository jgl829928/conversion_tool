
using CSV , DataFrames, DelimitedFiles ,ExcelReaders;


xls_dir="./xls"
csv_dir="./csv"
go_dir="./go"

allsheets=[]; #所有文件的所有表格
notReplace=[] #没有替换的坐标
ignoreCols = [
    "STR_CEHUA",
    "STR_TYPE",
    "STR_GOTO",
    "INT_SELECT",
    "STR_BUTTON",
    "INT_OPEN",
    "INT_MOVE",
    "INT_UI_SELECT",
    "STR_EFFECT_NAME_1",
    "STR_EFFECT_NAME_2",
    "STR_EFFECT_NAME_3",
    "STR_BUFF_NAME"
]
ignoreFiles = ["localization", "severslocalization","HeroIconRule","const"]
ignoreSheets = ["Sheet1"]

constTable="./xls/const.xlsx"



# 读取const列表
function constdata()
    f =  openxl(constTable);
    sheets=f.workbook[:sheet_names]();
    data = readxlsheet(constTable, 1)
    ptable = DataFrame(data);
    return ptable
end

# 判断类型
function valuetype(value)
    if(typeof(value)==String)
        return "string"
    elseif(float(value) % 1 == 0.0)
        return "int32"
    elseif(float(value) % 1 != 0.0)
        return "float32"
    end
end


# 忽略列数据
function delcols(data)
    nrows, ncols = size(data)
    temp=deepcopy(data);
    count=-1;
    for col in 1:ncols
        if(data[1,col] in ignoreCols)
            count=count+1;
            deletecols!(temp, col-count);
        end
    end
    return temp
end

# 替换变量
function replace_variate(data,file_info)
    nrows, ncols = size(data);
    intcols=[];
    coordinate=[] #坐标
    # 把含有int的列找出
    for col in 1:ncols
        if occursin("INT", data[1,col])
            push!(intcols,col)
        end
    end
    # 把需要替换的坐标保存
    for i in 1:length(intcols)
        for row in 2:nrows
            if(valuetype(data[row,intcols[i]])=="string")
               temp_item=(row,intcols[i])
               push!(coordinate,temp_item)
            end
        end
    end
    #替换数据
    constall=constdata()
    constrows=size(constall)[1]
    for i in 1:length(coordinate)
        row=coordinate[i][1]
        col=coordinate[i][2]
        isShowConst=false
        for constrow in 2:constrows
            if(data[row,col]==constall[constrow,1])
                data[row,col]=constall[constrow,3]
                isShowConst=true;
                break;
            end
        end
        if !isShowConst
            #错误提示  变量在const中找不到对应的值
            println("Table Check : error ceil find! file [$(file_info[1])] sheet [$(file_info[2])] pos( $row,$col ) ceilName $(data[1,col]) rightType [] ceilType [] ceilValue [$(data[row,col])]");
            push!(notReplace,(row,col))
            continue;
        end   
    end
    return data
end



#判断一列数据类型
function judge_coltype(data,file_info)
    nrows, ncols = size(data);
    for col in 1:ncols
        coltype=valuetype(data[2,col])
        for row in 2:nrows
            if((row,col) in notReplace) continue end;
            if(valuetype(data[row,col])!=coltype)
                #错误提示 对与第一行不同数据类型的进行报错
                println("Table Check : error ceil find! file [$(file_info[1])] sheet [$(file_info[2])] pos( $row,$col ) ceilName $(data[1,col]) rightType [] ceilType [] ceilValue [$(data[row,col])]");
            end
        end
    end
end


function to_csv(data,sheet)
    nrows, ncols = size(data)
    
    for col in 1:ncols
        oldname=names(data)[col];
        newname=Symbol(data[1,col]);
        namedata=Dict(oldname=>newname)
        rename!(data,namedata)
    end
    deleterows!(data, 1:1)

    csvfile="$csv_dir/$sheet.csv";
    read(`touch $csvfile`);
    CSV.write(csvfile,data);
    println("[xlstocsv] $sheet.csv has build")
end



#转换为go文件 获取替代数据
function replace_value(data)
    rowType="";
    fieldByName="";
    typeByName="";
    intfieldValue="";
    fltfieldValue="";
    strfieldValue="";
    nrows, ncols = size(data);
    for col in 1:ncols
        coltype=valuetype(data[2,col])
        rowType="$rowType $(data[1,col])   $coltype   `colname:\"$(data[1,col])\"`\n";
        fieldByName ="$fieldByName case \"$(data[1,col])\" :\n return $(col-1),true\n";
        typeByName ="$typeByName case \"$(data[1,col])\" :\n return reflect.$coltype,true\n";
        if ( coltype=="int32")        
            intfieldValue = "$intfieldValue case $(col-1) :\n return this.rows[row].$(data[1,col]),true\n";
        elseif (coltype=="string")
            strfieldValue = "$strfieldValue case $(col-1) :\n return this.rows[row].$(data[1,col]),true\n";
        else
            fltfieldValue = "$fltfieldValue case $(col-1) :\n return this.rows[row].$(data[1,col]),true\n";
        end    
    end
    return (rowType="$(rowType)",fieldByName="$(fieldByName)",
            typeByName="$(typeByName)",intfieldValue="$(intfieldValue)"
            ,fltfieldValue="$(fltfieldValue)",strfieldValue="$(strfieldValue)");
end



function to_go(data,sheet)
    gofile="$go_dir/$sheet.go";
    read(`touch $gofile`);
    tempFile = read("golang_table.temp", String);
    tempFile=replace(tempFile,"%CSVNAME%"=>"$sheet");
    tempFile=replace(tempFile,"%FIRSTFIELD%"=>"$(data[1,1])");
    tempFile=replace(tempFile,"%FIRSTFIELDTYPE%"=>"$(valuetype("$(data[1,1])"))");
    tempFile=replace(tempFile,"%ROWTYPE%"=>"$(replace_value(data).rowType)");
    tempFile=replace(tempFile,"%FIELDBYNAME%"=>"$(replace_value(data).fieldByName)");
    tempFile=replace(tempFile,"%TYPEBYNAME%"=>"$(replace_value(data).typeByName)");
    tempFile=replace(tempFile,"%GetIntFieldValue%"=>"$(replace_value(data).intfieldValue)");
    tempFile=replace(tempFile,"%GetStrFieldValue%"=>"$(replace_value(data).strfieldValue)");
    tempFile=replace(tempFile,"%GetFltFieldValue%"=>"$(replace_value(data).fltfieldValue)");

    io = open(gofile, "w");
    write(io,tempFile);
    println("[xlstogo] $(sheet).go has build")
end


# 读取文件
function readsheet(file)
    file_dir="$xls_dir/$file";
    println("read file $file")
    f =  openxl(file_dir);
    sheets=f.workbook[:sheet_names]();
    for i in 1:size(sheets,1)
        # 忽略表格
        if(sheets[i] in ignoreSheets) 
            println("ignore $file sheet $(sheets[i])")
            continue 
        end;
        #错误提示  表格重复
        if(sheets[i] in allsheets)
            error("duplicated sheet $(sheets[i]) in file $file")
        else
            push!(allsheets,sheets[i])
        end

        println("read $file sheet $(sheets[i])")

        # 获取表格数据
        data = readxlsheet(file_dir, i)
        ptable = DataFrame(data);

        # 忽略列数据
        tableData=delcols(ptable);

        # 判断表格是否为空
        if(size(tableData)[1]==0)
            error("empty sheet $(sheets[i]) in file $file")
        end
        
        file_info=(file,sheets[i])

        # 替换变量,变量不存在进行报错
        tableData=replace_variate(tableData,file_info)

        # 判断数据类型，对数据类型不同的进行报错
        judge_coltype(tableData,file_info)

        # to_csv
        to_csv(tableData,sheets[i]);

        # to_go
        to_go(tableData,sheets[i]);
    end
    println("$file's sheets have converted successfully")
end


function main()
    run(`rm -rf ./csv ./go`)
    run(`mkdir csv go` )
    for (root, dirs, files) in walkdir(xls_dir)
        for file in files
            # 判断是否为xls或者xlsx文件
            if(split(file, ".")[2]!="xls"&&split(file, ".")[2]!="xlsx")
                continue;
            end
            # 忽略文件
            file_pre=split(file, ".xlsx")[1];
            if(file_pre in ignoreFiles) 
                println("ignore file $file")
                continue 
            end;
            # 读取文件
            readsheet(file);
        end
        println("sucess all done")
    end
    # run(`gofmt -w ./go/*.go`)
end


main()

