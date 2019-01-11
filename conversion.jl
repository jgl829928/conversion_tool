using CSV , DataFrames, DelimitedFiles ,ExcelReaders;


xls_dir="./xls"
csv_dir="./csv"
go_dir="./go"


allsheets=[];
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
ignoreFiles = ["localization", "severslocalization","HeroIconRule"]
ignoreSheets = ["Sheet1"]


function valuetype(str)
    if occursin("INT",str)
        return "int32";
    elseif occursin("STR",str)'
        return "string";
    else 
        return "float32";
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

# 获取替代数据
function replace_value(data)
    rowType="";
    fieldByName="";
    typeByName="";
    intfieldValue="";
    fltfieldValue="";
    strfieldValue="";
    nrows, ncols = size(data);
    for col in 1:ncols
        coltype=valuetype("$(data[1,col])")
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


function readsheet(file)
    println("read file $file")
    f =  openxl(file);
    sheets=f.workbook[:sheet_names]();
    for i in 1:size(sheets,1)
        # 忽略表格
        if(sheets[i] in ignoreSheets) 
            println("ignore $file sheet $(sheets[i])")
            continue 
        end;
        # 表格重复
        if(sheets[i] in allsheets)
            error("duplicated sheet $(sheets[i]) in file $file")
        else
            push!(allsheets,sheets[i])
        end

        println("read $file sheet $(sheets[i])")
        # 获取表格数据
        
        data = readxlsheet(file, i)

        ptable = DataFrame(data);
        tableData=delcols(ptable);
        if(size(tableData)[1]==0)
            error("empty sheet $(sheets[i]) in file $file")
        end

        # to_csv
        csvfile="$csv_dir/$(sheets[i]).csv";
        read(`touch $csvfile`);
        CSV.write(csvfile,tableData);
        println("[xlstocsv] $(sheets[i]).csv has build")

        # to_go
        gofile="$go_dir/$(sheets[i]).go";
        read(`touch $gofile`);
        tempFile = read("golang_table.temp", String);
        tempFile=replace(tempFile,"%CSVNAME%"=>"$(sheets[i])");
        tempFile=replace(tempFile,"%FIRSTFIELD%"=>"$(tableData[1,1])");
        tempFile=replace(tempFile,"%FIRSTFIELDTYPE%"=>"$(valuetype("$(tableData[1,1])"))");
        tempFile=replace(tempFile,"%ROWTYPE%"=>"$(replace_value(tableData).rowType)");
        tempFile=replace(tempFile,"%FIELDBYNAME%"=>"$(replace_value(tableData).fieldByName)");
        tempFile=replace(tempFile,"%TYPEBYNAME%"=>"$(replace_value(tableData).typeByName)");
        tempFile=replace(tempFile,"%GetIntFieldValue%"=>"$(replace_value(tableData).intfieldValue)");
        tempFile=replace(tempFile,"%GetStrFieldValue%"=>"$(replace_value(tableData).strfieldValue)");
        tempFile=replace(tempFile,"%GetFltFieldValue%"=>"$(replace_value(tableData).fltfieldValue)");

        io = open(gofile, "w");
        write(io,tempFile);
        println("[xlstocsv] $(sheets[i]).go has build")
    end
    println("$file's sheets have converted successfully")
end


try
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
            file_dir="$xls_dir/$file";
            readsheet(file_dir);
        end
        println("sucess all done")
    end
catch y
    println(y)
end




















