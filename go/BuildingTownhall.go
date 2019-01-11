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

type BuildingTownhallRow struct {
     INT_LEVEL   int32   `colname:"INT_LEVEL"`
 INT_COST_ITEM_TYPE_1   int32   `colname:"INT_COST_ITEM_TYPE_1"`
 INT_COST_ITEM_COUNT_1   int32   `colname:"INT_COST_ITEM_COUNT_1"`
 INT_COST_ITEM_TYPE_2   int32   `colname:"INT_COST_ITEM_TYPE_2"`
 INT_COST_ITEM_COUNT_2   int32   `colname:"INT_COST_ITEM_COUNT_2"`
 INT_COST_ITEM_TYPE_3   int32   `colname:"INT_COST_ITEM_TYPE_3"`
 INT_COST_ITEM_COUNT_3   int32   `colname:"INT_COST_ITEM_COUNT_3"`
 INT_COST_ITEM_TYPE_4   int32   `colname:"INT_COST_ITEM_TYPE_4"`
 INT_COST_ITEM_COUNT_4   int32   `colname:"INT_COST_ITEM_COUNT_4"`
 INT_COST_ITEM_TYPE_5   int32   `colname:"INT_COST_ITEM_TYPE_5"`
 INT_COST_ITEM_COUNT_5   int32   `colname:"INT_COST_ITEM_COUNT_5"`
 INT_TIME   int32   `colname:"INT_TIME"`
 INT_REQUIRE_BUILDING_1   int32   `colname:"INT_REQUIRE_BUILDING_1"`
 INT_REQUIRE_LEVEL_1   int32   `colname:"INT_REQUIRE_LEVEL_1"`
 INT_REQUIRE_BUILDING_2   int32   `colname:"INT_REQUIRE_BUILDING_2"`
 INT_REQUIRE_LEVEL_2   int32   `colname:"INT_REQUIRE_LEVEL_2"`
 INT_POWER   int32   `colname:"INT_POWER"`
 INT_SCOUT_ITEM_TYPE   int32   `colname:"INT_SCOUT_ITEM_TYPE"`
 INT_SCOUT_ITEM_COUNT   int32   `colname:"INT_SCOUT_ITEM_COUNT"`
 INT_ALLYWAREHOUSE_MAX   int32   `colname:"INT_ALLYWAREHOUSE_MAX"`
 INT_ALLYWAREHOUSE_MAX_PER_DAY   int32   `colname:"INT_ALLYWAREHOUSE_MAX_PER_DAY"`
 INT_PLUNDER_LIMIT   int32   `colname:"INT_PLUNDER_LIMIT"`
 INT_REWARD_TYPE_1   int32   `colname:"INT_REWARD_TYPE_1"`
 INT_REWARD_COUNT_1   int32   `colname:"INT_REWARD_COUNT_1"`

}

type BuildingTownhallTable struct {
	rows []BuildingTownhallRow
    each map[int]bool
	sync.RWMutex
}

var (
    BuildingTownhall = BuildingTownhallTable{}
)

func (this *BuildingTownhallTable) Reload(csvfile string) (ok bool) {
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

func (this *BuildingTownhallTable) NumRows() int {
	this.RLock()
	defer this.RUnlock()

	return len(this.rows)
}

func (this *BuildingTownhallTable) Row(i int) (row BuildingTownhallRow, ok bool) {
	this.RLock()
	defer this.RUnlock()

	if i >= len(this.rows) {
		return
	}

	return this.rows[i], true
}

func (this *BuildingTownhallTable) RowByINT_LEVEL(INT_LEVEL int32) (row BuildingTownhallRow,rowIndex int, ok bool) {
	this.RLock()
	defer this.RUnlock()

	for i := 0; i < len(this.rows); i++ {
		if this.rows[i].INT_LEVEL == INT_LEVEL {
			return this.rows[i], i,true
		}
	}

	return
}

//--------------------------------------------------------- for GTable

func (this *BuildingTownhallTable) RowType() BuildingTownhallRow {
    return BuildingTownhallRow{}
}

func (this *BuildingTownhallTable) ByField(i int, field string) (row reflect.Value, ok bool) {
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

func (this *BuildingTownhallTable) GetRowListByFieldMap(intfieldMap map[string]int32,strfieldMap map[string]string,fltfieldMap map[string]float32,count int) (rows []int) {
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

func (this *BuildingTownhallTable) readfile(csvfile string) (records [][]string, ok bool) {
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

func (this *BuildingTownhallTable) createRow(records [][]string) (rows []BuildingTownhallRow) {
	if len(records) <= 0 {
		return
	}

    refType := reflect.TypeOf(BuildingTownhallRow{})
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

		rows = append(rows, refElem.Interface().(BuildingTownhallRow))
	}

    log.Trace("load table BuildingTownhall len(%v)",len(rows))
	return
}

func (this *BuildingTownhallTable) valueByField(row int,colname string) (interface{}, bool) {
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

func (this *BuildingTownhallTable) fieldByName(colname string) (int,bool) {
    switch colname {
         case "INT_LEVEL" :
 return 0,true
 case "INT_COST_ITEM_TYPE_1" :
 return 1,true
 case "INT_COST_ITEM_COUNT_1" :
 return 2,true
 case "INT_COST_ITEM_TYPE_2" :
 return 3,true
 case "INT_COST_ITEM_COUNT_2" :
 return 4,true
 case "INT_COST_ITEM_TYPE_3" :
 return 5,true
 case "INT_COST_ITEM_COUNT_3" :
 return 6,true
 case "INT_COST_ITEM_TYPE_4" :
 return 7,true
 case "INT_COST_ITEM_COUNT_4" :
 return 8,true
 case "INT_COST_ITEM_TYPE_5" :
 return 9,true
 case "INT_COST_ITEM_COUNT_5" :
 return 10,true
 case "INT_TIME" :
 return 11,true
 case "INT_REQUIRE_BUILDING_1" :
 return 12,true
 case "INT_REQUIRE_LEVEL_1" :
 return 13,true
 case "INT_REQUIRE_BUILDING_2" :
 return 14,true
 case "INT_REQUIRE_LEVEL_2" :
 return 15,true
 case "INT_POWER" :
 return 16,true
 case "INT_SCOUT_ITEM_TYPE" :
 return 17,true
 case "INT_SCOUT_ITEM_COUNT" :
 return 18,true
 case "INT_ALLYWAREHOUSE_MAX" :
 return 19,true
 case "INT_ALLYWAREHOUSE_MAX_PER_DAY" :
 return 20,true
 case "INT_PLUNDER_LIMIT" :
 return 21,true
 case "INT_REWARD_TYPE_1" :
 return 22,true
 case "INT_REWARD_COUNT_1" :
 return 23,true

        default:
            return 0,false
    }
}

func (this *BuildingTownhallTable) typeByName(colname string) (reflect.Kind, bool) {
	switch colname {
	 case "INT_LEVEL" :
 return reflect.int32,true
 case "INT_COST_ITEM_TYPE_1" :
 return reflect.int32,true
 case "INT_COST_ITEM_COUNT_1" :
 return reflect.int32,true
 case "INT_COST_ITEM_TYPE_2" :
 return reflect.int32,true
 case "INT_COST_ITEM_COUNT_2" :
 return reflect.int32,true
 case "INT_COST_ITEM_TYPE_3" :
 return reflect.int32,true
 case "INT_COST_ITEM_COUNT_3" :
 return reflect.int32,true
 case "INT_COST_ITEM_TYPE_4" :
 return reflect.int32,true
 case "INT_COST_ITEM_COUNT_4" :
 return reflect.int32,true
 case "INT_COST_ITEM_TYPE_5" :
 return reflect.int32,true
 case "INT_COST_ITEM_COUNT_5" :
 return reflect.int32,true
 case "INT_TIME" :
 return reflect.int32,true
 case "INT_REQUIRE_BUILDING_1" :
 return reflect.int32,true
 case "INT_REQUIRE_LEVEL_1" :
 return reflect.int32,true
 case "INT_REQUIRE_BUILDING_2" :
 return reflect.int32,true
 case "INT_REQUIRE_LEVEL_2" :
 return reflect.int32,true
 case "INT_POWER" :
 return reflect.int32,true
 case "INT_SCOUT_ITEM_TYPE" :
 return reflect.int32,true
 case "INT_SCOUT_ITEM_COUNT" :
 return reflect.int32,true
 case "INT_ALLYWAREHOUSE_MAX" :
 return reflect.int32,true
 case "INT_ALLYWAREHOUSE_MAX_PER_DAY" :
 return reflect.int32,true
 case "INT_PLUNDER_LIMIT" :
 return reflect.int32,true
 case "INT_REWARD_TYPE_1" :
 return reflect.int32,true
 case "INT_REWARD_COUNT_1" :
 return reflect.int32,true

    default:
		return reflect.Invalid, false
	}
}

func (this *BuildingTownhallTable) intFieldValue(fieldIndex, row int) (data int32, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
     case 0 :
 return this.rows[row].INT_LEVEL,true
 case 1 :
 return this.rows[row].INT_COST_ITEM_TYPE_1,true
 case 2 :
 return this.rows[row].INT_COST_ITEM_COUNT_1,true
 case 3 :
 return this.rows[row].INT_COST_ITEM_TYPE_2,true
 case 4 :
 return this.rows[row].INT_COST_ITEM_COUNT_2,true
 case 5 :
 return this.rows[row].INT_COST_ITEM_TYPE_3,true
 case 6 :
 return this.rows[row].INT_COST_ITEM_COUNT_3,true
 case 7 :
 return this.rows[row].INT_COST_ITEM_TYPE_4,true
 case 8 :
 return this.rows[row].INT_COST_ITEM_COUNT_4,true
 case 9 :
 return this.rows[row].INT_COST_ITEM_TYPE_5,true
 case 10 :
 return this.rows[row].INT_COST_ITEM_COUNT_5,true
 case 11 :
 return this.rows[row].INT_TIME,true
 case 12 :
 return this.rows[row].INT_REQUIRE_BUILDING_1,true
 case 13 :
 return this.rows[row].INT_REQUIRE_LEVEL_1,true
 case 14 :
 return this.rows[row].INT_REQUIRE_BUILDING_2,true
 case 15 :
 return this.rows[row].INT_REQUIRE_LEVEL_2,true
 case 16 :
 return this.rows[row].INT_POWER,true
 case 17 :
 return this.rows[row].INT_SCOUT_ITEM_TYPE,true
 case 18 :
 return this.rows[row].INT_SCOUT_ITEM_COUNT,true
 case 19 :
 return this.rows[row].INT_ALLYWAREHOUSE_MAX,true
 case 20 :
 return this.rows[row].INT_ALLYWAREHOUSE_MAX_PER_DAY,true
 case 21 :
 return this.rows[row].INT_PLUNDER_LIMIT,true
 case 22 :
 return this.rows[row].INT_REWARD_TYPE_1,true
 case 23 :
 return this.rows[row].INT_REWARD_COUNT_1,true

    default:
        return
    }
}

func (this *BuildingTownhallTable) strFieldValue(fieldIndex, row int) (data string, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
    
    default:
        return
    }
}

func (this *BuildingTownhallTable) fltFieldValue(fieldIndex, row int) (data float32, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
    
    default:
        return
    }
}

func (this *BuildingTownhallTable) matchIntField(row int,intfieldMap map[string]int32) (match bool) {
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

func (this *BuildingTownhallTable) matchStrField(row int,strfieldMap map[string]string) (match bool) {
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

func (this *BuildingTownhallTable) matchFltField(row int,fltfieldMap map[string]float32) (match bool) {
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
