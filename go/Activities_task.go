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

type Activities_taskRow struct {
     INT_ID   int32   `colname:"INT_ID"`
 INT_ACTIVITIES_ID   int32   `colname:"INT_ACTIVITIES_ID"`
 INT_INT_PAGE   int32   `colname:"INT_INT_PAGE"`
 INT_TYPE   int32   `colname:"INT_TYPE"`
 STR_CONST_DES   string   `colname:"STR_CONST_DES"`
 STR_NAME   string   `colname:"STR_NAME"`
 STR_DESCRIPTION   string   `colname:"STR_DESCRIPTION"`
 INT_PARM_1   int32   `colname:"INT_PARM_1"`
 INT_PARM_2   int32   `colname:"INT_PARM_2"`
 INT_RANGE_2   int32   `colname:"INT_RANGE_2"`
 INT_END_VALUE   int32   `colname:"INT_END_VALUE"`
 INT_REWARD_MODE   int32   `colname:"INT_REWARD_MODE"`
 INT_MAX_RECEIVE_TIMES   int32   `colname:"INT_MAX_RECEIVE_TIMES"`
 INT_AUTO_AWARD_ITEM_TYPY_1   int32   `colname:"INT_AUTO_AWARD_ITEM_TYPY_1"`
 INT_AUTO_AWARD_ITEM_COUNT_1   int32   `colname:"INT_AUTO_AWARD_ITEM_COUNT_1"`
 INT_AUTO_AWARD_ITEM_TYPY_2   int32   `colname:"INT_AUTO_AWARD_ITEM_TYPY_2"`
 INT_AUTO_AWARD_ITEM_COUNT_2   int32   `colname:"INT_AUTO_AWARD_ITEM_COUNT_2"`
 INT_REWARD_1   int32   `colname:"INT_REWARD_1"`
 INT_COUNT_1   int32   `colname:"INT_COUNT_1"`
 INT_REWARD_2   int32   `colname:"INT_REWARD_2"`
 INT_COUNT_2   int32   `colname:"INT_COUNT_2"`
 INT_REWARD_3   int32   `colname:"INT_REWARD_3"`
 INT_COUNT_3   int32   `colname:"INT_COUNT_3"`
 INT_REWARD_4   int32   `colname:"INT_REWARD_4"`
 INT_COUNT_4   int32   `colname:"INT_COUNT_4"`
 INT_REWARD_5   int32   `colname:"INT_REWARD_5"`
 INT_COUNT_5   int32   `colname:"INT_COUNT_5"`
 INT_MIN_VALUE   int32   `colname:"INT_MIN_VALUE"`
 INT_MAX_VALUE   int32   `colname:"INT_MAX_VALUE"`
 INT_SORT_INDEX   int32   `colname:"INT_SORT_INDEX"`
 INT_VIP_LIMIT   int32   `colname:"INT_VIP_LIMIT"`
 INT_DISCOUNT_RATES   int32   `colname:"INT_DISCOUNT_RATES"`

}

type Activities_taskTable struct {
	rows []Activities_taskRow
    each map[int]bool
	sync.RWMutex
}

var (
    Activities_task = Activities_taskTable{}
)

func (this *Activities_taskTable) Reload(csvfile string) (ok bool) {
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

func (this *Activities_taskTable) NumRows() int {
	this.RLock()
	defer this.RUnlock()

	return len(this.rows)
}

func (this *Activities_taskTable) Row(i int) (row Activities_taskRow, ok bool) {
	this.RLock()
	defer this.RUnlock()

	if i >= len(this.rows) {
		return
	}

	return this.rows[i], true
}

func (this *Activities_taskTable) RowByINT_ID(INT_ID int32) (row Activities_taskRow,rowIndex int, ok bool) {
	this.RLock()
	defer this.RUnlock()

	for i := 0; i < len(this.rows); i++ {
		if this.rows[i].INT_ID == INT_ID {
			return this.rows[i], i,true
		}
	}

	return
}

//--------------------------------------------------------- for GTable

func (this *Activities_taskTable) RowType() Activities_taskRow {
    return Activities_taskRow{}
}

func (this *Activities_taskTable) ByField(i int, field string) (row reflect.Value, ok bool) {
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

func (this *Activities_taskTable) GetRowListByFieldMap(intfieldMap map[string]int32,strfieldMap map[string]string,fltfieldMap map[string]float32,count int) (rows []int) {
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

func (this *Activities_taskTable) readfile(csvfile string) (records [][]string, ok bool) {
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

func (this *Activities_taskTable) createRow(records [][]string) (rows []Activities_taskRow) {
	if len(records) <= 0 {
		return
	}

    refType := reflect.TypeOf(Activities_taskRow{})
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

		rows = append(rows, refElem.Interface().(Activities_taskRow))
	}

    log.Trace("load table Activities_task len(%v)",len(rows))
	return
}

func (this *Activities_taskTable) valueByField(row int,colname string) (interface{}, bool) {
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

func (this *Activities_taskTable) fieldByName(colname string) (int,bool) {
    switch colname {
         case "INT_ID" :
 return 0,true
 case "INT_ACTIVITIES_ID" :
 return 1,true
 case "INT_INT_PAGE" :
 return 2,true
 case "INT_TYPE" :
 return 3,true
 case "STR_CONST_DES" :
 return 4,true
 case "STR_NAME" :
 return 5,true
 case "STR_DESCRIPTION" :
 return 6,true
 case "INT_PARM_1" :
 return 7,true
 case "INT_PARM_2" :
 return 8,true
 case "INT_RANGE_2" :
 return 9,true
 case "INT_END_VALUE" :
 return 10,true
 case "INT_REWARD_MODE" :
 return 11,true
 case "INT_MAX_RECEIVE_TIMES" :
 return 12,true
 case "INT_AUTO_AWARD_ITEM_TYPY_1" :
 return 13,true
 case "INT_AUTO_AWARD_ITEM_COUNT_1" :
 return 14,true
 case "INT_AUTO_AWARD_ITEM_TYPY_2" :
 return 15,true
 case "INT_AUTO_AWARD_ITEM_COUNT_2" :
 return 16,true
 case "INT_REWARD_1" :
 return 17,true
 case "INT_COUNT_1" :
 return 18,true
 case "INT_REWARD_2" :
 return 19,true
 case "INT_COUNT_2" :
 return 20,true
 case "INT_REWARD_3" :
 return 21,true
 case "INT_COUNT_3" :
 return 22,true
 case "INT_REWARD_4" :
 return 23,true
 case "INT_COUNT_4" :
 return 24,true
 case "INT_REWARD_5" :
 return 25,true
 case "INT_COUNT_5" :
 return 26,true
 case "INT_MIN_VALUE" :
 return 27,true
 case "INT_MAX_VALUE" :
 return 28,true
 case "INT_SORT_INDEX" :
 return 29,true
 case "INT_VIP_LIMIT" :
 return 30,true
 case "INT_DISCOUNT_RATES" :
 return 31,true

        default:
            return 0,false
    }
}

func (this *Activities_taskTable) typeByName(colname string) (reflect.Kind, bool) {
	switch colname {
	 case "INT_ID" :
 return reflect.int32,true
 case "INT_ACTIVITIES_ID" :
 return reflect.int32,true
 case "INT_INT_PAGE" :
 return reflect.int32,true
 case "INT_TYPE" :
 return reflect.int32,true
 case "STR_CONST_DES" :
 return reflect.string,true
 case "STR_NAME" :
 return reflect.string,true
 case "STR_DESCRIPTION" :
 return reflect.string,true
 case "INT_PARM_1" :
 return reflect.int32,true
 case "INT_PARM_2" :
 return reflect.int32,true
 case "INT_RANGE_2" :
 return reflect.int32,true
 case "INT_END_VALUE" :
 return reflect.int32,true
 case "INT_REWARD_MODE" :
 return reflect.int32,true
 case "INT_MAX_RECEIVE_TIMES" :
 return reflect.int32,true
 case "INT_AUTO_AWARD_ITEM_TYPY_1" :
 return reflect.int32,true
 case "INT_AUTO_AWARD_ITEM_COUNT_1" :
 return reflect.int32,true
 case "INT_AUTO_AWARD_ITEM_TYPY_2" :
 return reflect.int32,true
 case "INT_AUTO_AWARD_ITEM_COUNT_2" :
 return reflect.int32,true
 case "INT_REWARD_1" :
 return reflect.int32,true
 case "INT_COUNT_1" :
 return reflect.int32,true
 case "INT_REWARD_2" :
 return reflect.int32,true
 case "INT_COUNT_2" :
 return reflect.int32,true
 case "INT_REWARD_3" :
 return reflect.int32,true
 case "INT_COUNT_3" :
 return reflect.int32,true
 case "INT_REWARD_4" :
 return reflect.int32,true
 case "INT_COUNT_4" :
 return reflect.int32,true
 case "INT_REWARD_5" :
 return reflect.int32,true
 case "INT_COUNT_5" :
 return reflect.int32,true
 case "INT_MIN_VALUE" :
 return reflect.int32,true
 case "INT_MAX_VALUE" :
 return reflect.int32,true
 case "INT_SORT_INDEX" :
 return reflect.int32,true
 case "INT_VIP_LIMIT" :
 return reflect.int32,true
 case "INT_DISCOUNT_RATES" :
 return reflect.int32,true

    default:
		return reflect.Invalid, false
	}
}

func (this *Activities_taskTable) intFieldValue(fieldIndex, row int) (data int32, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
     case 0 :
 return this.rows[row].INT_ID,true
 case 1 :
 return this.rows[row].INT_ACTIVITIES_ID,true
 case 2 :
 return this.rows[row].INT_INT_PAGE,true
 case 3 :
 return this.rows[row].INT_TYPE,true
 case 7 :
 return this.rows[row].INT_PARM_1,true
 case 8 :
 return this.rows[row].INT_PARM_2,true
 case 9 :
 return this.rows[row].INT_RANGE_2,true
 case 10 :
 return this.rows[row].INT_END_VALUE,true
 case 11 :
 return this.rows[row].INT_REWARD_MODE,true
 case 12 :
 return this.rows[row].INT_MAX_RECEIVE_TIMES,true
 case 13 :
 return this.rows[row].INT_AUTO_AWARD_ITEM_TYPY_1,true
 case 14 :
 return this.rows[row].INT_AUTO_AWARD_ITEM_COUNT_1,true
 case 15 :
 return this.rows[row].INT_AUTO_AWARD_ITEM_TYPY_2,true
 case 16 :
 return this.rows[row].INT_AUTO_AWARD_ITEM_COUNT_2,true
 case 17 :
 return this.rows[row].INT_REWARD_1,true
 case 18 :
 return this.rows[row].INT_COUNT_1,true
 case 19 :
 return this.rows[row].INT_REWARD_2,true
 case 20 :
 return this.rows[row].INT_COUNT_2,true
 case 21 :
 return this.rows[row].INT_REWARD_3,true
 case 22 :
 return this.rows[row].INT_COUNT_3,true
 case 23 :
 return this.rows[row].INT_REWARD_4,true
 case 24 :
 return this.rows[row].INT_COUNT_4,true
 case 25 :
 return this.rows[row].INT_REWARD_5,true
 case 26 :
 return this.rows[row].INT_COUNT_5,true
 case 27 :
 return this.rows[row].INT_MIN_VALUE,true
 case 28 :
 return this.rows[row].INT_MAX_VALUE,true
 case 29 :
 return this.rows[row].INT_SORT_INDEX,true
 case 30 :
 return this.rows[row].INT_VIP_LIMIT,true
 case 31 :
 return this.rows[row].INT_DISCOUNT_RATES,true

    default:
        return
    }
}

func (this *Activities_taskTable) strFieldValue(fieldIndex, row int) (data string, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
     case 4 :
 return this.rows[row].STR_CONST_DES,true
 case 5 :
 return this.rows[row].STR_NAME,true
 case 6 :
 return this.rows[row].STR_DESCRIPTION,true

    default:
        return
    }
}

func (this *Activities_taskTable) fltFieldValue(fieldIndex, row int) (data float32, ok bool) {
    if len(this.rows) <= row {
        return
    }

    switch fieldIndex {
    
    default:
        return
    }
}

func (this *Activities_taskTable) matchIntField(row int,intfieldMap map[string]int32) (match bool) {
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

func (this *Activities_taskTable) matchStrField(row int,strfieldMap map[string]string) (match bool) {
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

func (this *Activities_taskTable) matchFltField(row int,fltfieldMap map[string]float32) (match bool) {
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

