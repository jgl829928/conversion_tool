/*
    本文文件由脚本自动生成，禁止手动编辑
*/

package config

import (
    "ABEngine/log"
	"encoding/csv"
	"os"
	"reflect"
	"strconv"
	"sync"
    "path"
)

type BuildingTypeRow struct {
     INT_TYPE   int32   `colname:"INT_TYPE"`
 STR_NAME   string   `colname:"STR_NAME"`
 INT_BUILDING_REMOVABLE   int32   `colname:"INT_BUILDING_REMOVABLE"`
 INT_PRODUCATION_BUILDING   int32   `colname:"INT_PRODUCATION_BUILDING"`
 INT_AREA   int32   `colname:"INT_AREA"`
 STR_LIST_NAME   string   `colname:"STR_LIST_NAME"`

}

type BuildingTypeTable struct {
	rows []BuildingTypeRow
    each map[int]bool
	sync.RWMutex
}

var (
    BuildingType = BuildingTypeTable{}
)

func (this *BuildingTypeTable) Reload(csvfile string) (ok bool) {
	this.Lock()
	defer this.Unlock()

	records, ok := this.readfile(csvfile)

	if !ok {
		log.Erro("can't load csv %v", csvfile)
		return
	}

	this.rows = this.createRow(records)

    this.each = make(map[int]bool)
    for i := 0;i < len(this.rows);i++ {
        this.each[i] = true
    }

	return true
}

func (this *BuildingTypeTable) NumRows() int {
	this.RLock()
	defer this.RUnlock()

	return len(this.rows)
}

func (this *BuildingTypeTable) Row(i int) (row BuildingTypeRow, ok bool) {
	this.RLock()
	defer this.RUnlock()

	if i >= len(this.rows) {
		return
	}

	return this.rows[i], true
}

func (this *BuildingTypeTable) RowByINT_TYPE(INT_TYPE int32) (row BuildingTypeRow,rowIndex int, ok bool) {
	this.RLock()
	defer this.RUnlock()

	for i := 0; i < len(this.rows); i++ {
		if this.rows[i].INT_TYPE == INT_TYPE {
			return this.rows[i], i,true
		}
	}

	return
}

//--------------------------------------------------------- for GTable

func (this *BuildingTypeTable) RowType() BuildingTypeRow {
    return BuildingTypeRow{}
}

func (this *BuildingTypeTable) ByField(i int, field string) (row reflect.Value, ok bool) {
    this.RLock()
    defer this.RUnlock()

    if i >= len(this.rows) {
        return
    }

    fieldIndex, ok := this.fieldByName(field)

    if !ok {
        return
    }

    return reflect.ValueOf(this.rows[i]).Field(fieldIndex), true
}

func (this *BuildingTypeTable) GetRowListByFieldMap(intfieldMap map[string]int32,strfieldMap map[string]string,fltfieldMap map[string]float32,count int) (rows []int) {
	this.RLock()
	defer this.RUnlock()

    for i := range this.each {
        if count >= 0 && len(rows) >= count {
            break
        }

        if !this.matchIntField(i,intfieldMap) {
            continue
        }

        if !this.matchStrField(i,strfieldMap) {
            continue
        }

        if !this.matchFltField(i,fltfieldMap) {
            continue
        }

        rows = append(rows,i)
    }

	return
}

//---------------------------------------------------------  inner func

func (this *BuildingTypeTable) readfile(csvfile string) (records [][]string, ok bool) {
    if path.Ext(csvfile) != ".csv" {
        log.Erro("not csvfile %v",csvfile)
        return
    }

	fr, err := os.Open(csvfile)
	if err != nil {
        log.Erro("err %v",err)
		return
	}

	defer fr.Close()

	cr := csv.NewReader(fr)

	records, err = cr.ReadAll()

	if err != nil {
        log.Erro("err %v",err)
		return
	}

    return records,true
}

func (this *BuildingTypeTable) createRow(records [][]string) (rows []BuildingTypeRow) {
	if len(records) <= 0 {
		return
	}

    refType := reflect.TypeOf(BuildingTypeRow{})
	fieldTagMap := make(map[string]int)

	for i := 0; i < refType.NumField(); i++ {
		fieldTagMap[refType.Field(i).Tag.Get("colname")] = i
	}

	for rowIndex, row := range records {
		if rowIndex == 0 {
			continue
		}

		refElem := reflect.New(refType).Elem()
		for colIndex, colValue := range row {
			fieldIndex, ok := fieldTagMap[records[0][colIndex]]

			if !ok {
				continue
			}

			fieldValue := refElem.Field(fieldIndex)

			switch fieldValue.Kind() {
			case reflect.String:
				fieldValue.SetString(colValue)
			case reflect.Int32:
				numValue, _ := strconv.ParseInt(colValue, 0, 0)
				fieldValue.SetInt(numValue)
            case reflect.Float32:
				numValue, _ := strconv.ParseFloat(colValue, 32)
				fieldValue.SetFloat(numValue)
			}
		}

		rows = append(rows, refElem.Interface().(BuildingTypeRow))
	}

    log.Trace("load table BuildingType len(%v)",len(rows))
	return
}

func (this *BuildingTypeTable) valueByField(row int,colname string) (interface{}, bool) {
	this.RLock()
	defer this.RUnlock()

    if row >= len(this.rows) {
        return nil, false
    }

    t, ok := this.typeByName(colname)

    if !ok {
        return nil, false
    }

    i, ok := this.fieldByName(colname)

    if !ok {
        return nil, false
    }

    switch t {
    case reflect.Int32:
        intData, ok := this.intFieldValue(i,row)

        if !ok {
            return nil, false
        }

        return intData, true
    case reflect.String:
        strData, ok := this.strFieldValue(i,row)

        if !ok {
            return nil, false
        }

        return strData, true
    case reflect.Float32:
        fltData, ok := this.fltFieldValue(i,row)

        if !ok {
            return nil, false
        }

        return fltData, true
    }

    return nil, false
}

func (this *BuildingTypeTable) fieldByName(colname string) (int,bool) {
    switch colname {
         case "INT_TYPE" :
 return 0,true
 case "STR_NAME" :
 return 1,true
 case "INT_BUILDING_REMOVABLE" :
 return 2,true
 case "INT_PRODUCATION_BUILDING" :
 return 3,true
 case "INT_AREA" :
 return 4,true
 case "STR_LIST_NAME" :
 return 5,true

        default:
            return 0,false
    }
}

func (this *BuildingTypeTable) typeByName(colname string) (reflect.Kind, bool) {
	switch colname {
	 case "INT_TYPE" :
 return reflect.int32,true
 case "STR_NAME" :
 return reflect.string,true
 case "INT_BUILDING_REMOVABLE" :
 return reflect.int32,true
 case "INT_PRODUCATION_BUILDING" :
 return reflect.int32,true
 case "INT_AREA" :
 return reflect.int32,true
 case "STR_LIST_NAME" :
 return reflect.string,true

    default:
		return reflect.Invalid, false
	}
}

func (this *BuildingTypeTable) intFieldValue(fieldIndex, row int) (data int32, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
     case 0 :
 return this.rows[row].INT_TYPE,true
 case 2 :
 return this.rows[row].INT_BUILDING_REMOVABLE,true
 case 3 :
 return this.rows[row].INT_PRODUCATION_BUILDING,true
 case 4 :
 return this.rows[row].INT_AREA,true

    default:
        return
    }
}

func (this *BuildingTypeTable) strFieldValue(fieldIndex, row int) (data string, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
     case 1 :
 return this.rows[row].STR_NAME,true
 case 5 :
 return this.rows[row].STR_LIST_NAME,true

    default:
        return
    }
}

func (this *BuildingTypeTable) fltFieldValue(fieldIndex, row int) (data float32, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
    
    default:
        return
    }
}

func (this *BuildingTypeTable) matchIntField(row int,intfieldMap map[string]int32) (match bool) {
    for colname,value := range intfieldMap {
        fieldIndex, ok := this.fieldByName(colname)

        if !ok {
            return
        }

        fieldValue, ok := this.intFieldValue(fieldIndex,row)

        if !ok {
            return
        }

        if fieldValue != value {
            return
        }
    }

    return true
}

func (this *BuildingTypeTable) matchStrField(row int,strfieldMap map[string]string) (match bool) {
     for colname,value := range strfieldMap {
        fieldIndex, ok := this.fieldByName(colname)

        if !ok {
            return
        }

        fieldValue, ok := this.strFieldValue(fieldIndex,row)

        if !ok {
            return
        }

        if fieldValue != value {
            return
        }
    }

    return true
}

func (this *BuildingTypeTable) matchFltField(row int,fltfieldMap map[string]float32) (match bool) {
    for colname,value := range fltfieldMap {
        fieldIndex, ok := this.fieldByName(colname)

        if !ok {
            return
        }

        fieldValue, ok := this.fltFieldValue(fieldIndex,row)

        if !ok {
            return
        }

        if fieldValue != value {
            return
        }
    }

    return true
}

