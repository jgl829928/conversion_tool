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

type BuildingCountRow struct {
     INT_TOWNHALL_LEVEL   int32   `colname:"INT_TOWNHALL_LEVEL"`
 INT_SMELTER_COUNT   int32   `colname:"INT_SMELTER_COUNT"`
 INT_OIL_FIELD_COUNT   int32   `colname:"INT_OIL_FIELD_COUNT"`
 INT_REPAIR_SHOP_COUNT   int32   `colname:"INT_REPAIR_SHOP_COUNT"`
 INT_POWER_PLANT_COUNT   int32   `colname:"INT_POWER_PLANT_COUNT"`
 INT_RUBBER_PLANT_COUNT   int32   `colname:"INT_RUBBER_PLANT_COUNT"`
 INT_AMMO_PLANT_COUNT   int32   `colname:"INT_AMMO_PLANT_COUNT"`
 INT_BUNKER_COUNT   int32   `colname:"INT_BUNKER_COUNT"`

}

type BuildingCountTable struct {
	rows []BuildingCountRow
    each map[int]bool
	sync.RWMutex
}

var (
    BuildingCount = BuildingCountTable{}
)

func (this *BuildingCountTable) Reload(csvfile string) (ok bool) {
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

func (this *BuildingCountTable) NumRows() int {
	this.RLock()
	defer this.RUnlock()

	return len(this.rows)
}

func (this *BuildingCountTable) Row(i int) (row BuildingCountRow, ok bool) {
	this.RLock()
	defer this.RUnlock()

	if i >= len(this.rows) {
		return
	}

	return this.rows[i], true
}

func (this *BuildingCountTable) RowByINT_TOWNHALL_LEVEL(INT_TOWNHALL_LEVEL int32) (row BuildingCountRow,rowIndex int, ok bool) {
	this.RLock()
	defer this.RUnlock()

	for i := 0; i < len(this.rows); i++ {
		if this.rows[i].INT_TOWNHALL_LEVEL == INT_TOWNHALL_LEVEL {
			return this.rows[i], i,true
		}
	}

	return
}

//--------------------------------------------------------- for GTable

func (this *BuildingCountTable) RowType() BuildingCountRow {
    return BuildingCountRow{}
}

func (this *BuildingCountTable) ByField(i int, field string) (row reflect.Value, ok bool) {
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

func (this *BuildingCountTable) GetRowListByFieldMap(intfieldMap map[string]int32,strfieldMap map[string]string,fltfieldMap map[string]float32,count int) (rows []int) {
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

func (this *BuildingCountTable) readfile(csvfile string) (records [][]string, ok bool) {
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

func (this *BuildingCountTable) createRow(records [][]string) (rows []BuildingCountRow) {
	if len(records) <= 0 {
		return
	}

    refType := reflect.TypeOf(BuildingCountRow{})
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

		rows = append(rows, refElem.Interface().(BuildingCountRow))
	}

    log.Trace("load table BuildingCount len(%v)",len(rows))
	return
}

func (this *BuildingCountTable) valueByField(row int,colname string) (interface{}, bool) {
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

func (this *BuildingCountTable) fieldByName(colname string) (int,bool) {
    switch colname {
         case "INT_TOWNHALL_LEVEL" :
 return 0,true
 case "INT_SMELTER_COUNT" :
 return 1,true
 case "INT_OIL_FIELD_COUNT" :
 return 2,true
 case "INT_REPAIR_SHOP_COUNT" :
 return 3,true
 case "INT_POWER_PLANT_COUNT" :
 return 4,true
 case "INT_RUBBER_PLANT_COUNT" :
 return 5,true
 case "INT_AMMO_PLANT_COUNT" :
 return 6,true
 case "INT_BUNKER_COUNT" :
 return 7,true

        default:
            return 0,false
    }
}

func (this *BuildingCountTable) typeByName(colname string) (reflect.Kind, bool) {
	switch colname {
	 case "INT_TOWNHALL_LEVEL" :
 return reflect.int32,true
 case "INT_SMELTER_COUNT" :
 return reflect.int32,true
 case "INT_OIL_FIELD_COUNT" :
 return reflect.int32,true
 case "INT_REPAIR_SHOP_COUNT" :
 return reflect.int32,true
 case "INT_POWER_PLANT_COUNT" :
 return reflect.int32,true
 case "INT_RUBBER_PLANT_COUNT" :
 return reflect.int32,true
 case "INT_AMMO_PLANT_COUNT" :
 return reflect.int32,true
 case "INT_BUNKER_COUNT" :
 return reflect.int32,true

    default:
		return reflect.Invalid, false
	}
}

func (this *BuildingCountTable) intFieldValue(fieldIndex, row int) (data int32, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
     case 0 :
 return this.rows[row].INT_TOWNHALL_LEVEL,true
 case 1 :
 return this.rows[row].INT_SMELTER_COUNT,true
 case 2 :
 return this.rows[row].INT_OIL_FIELD_COUNT,true
 case 3 :
 return this.rows[row].INT_REPAIR_SHOP_COUNT,true
 case 4 :
 return this.rows[row].INT_POWER_PLANT_COUNT,true
 case 5 :
 return this.rows[row].INT_RUBBER_PLANT_COUNT,true
 case 6 :
 return this.rows[row].INT_AMMO_PLANT_COUNT,true
 case 7 :
 return this.rows[row].INT_BUNKER_COUNT,true

    default:
        return
    }
}

func (this *BuildingCountTable) strFieldValue(fieldIndex, row int) (data string, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
    
    default:
        return
    }
}

func (this *BuildingCountTable) fltFieldValue(fieldIndex, row int) (data float32, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
    
    default:
        return
    }
}

func (this *BuildingCountTable) matchIntField(row int,intfieldMap map[string]int32) (match bool) {
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

func (this *BuildingCountTable) matchStrField(row int,strfieldMap map[string]string) (match bool) {
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

func (this *BuildingCountTable) matchFltField(row int,fltfieldMap map[string]float32) (match bool) {
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

