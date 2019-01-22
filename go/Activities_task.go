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
     6.0000999e7   int32   `colname:"6.0000999e7"`
 100000.0   int32   `colname:"100000.0"`
 0.0   int32   `colname:"0.0"`
 1033.0   int32   `colname:"1033.0"`
 首次充值次数，12121212   string   `colname:"首次充值次数，12121212"`
 首次充值   string   `colname:"首次充值"`
 首次充值   string   `colname:"首次充值"`
 0.0   int32   `colname:"0.0"`
 0.0   int32   `colname:"0.0"`
 0.0   int32   `colname:"0.0"`
 1.0   int32   `colname:"1.0"`
 2002.0   int32   `colname:"2002.0"`
 1.0   int32   `colname:"1.0"`
 0.0   int32   `colname:"0.0"`
 0.0   int32   `colname:"0.0"`
 0.0   int32   `colname:"0.0"`
 0.0   int32   `colname:"0.0"`
 270059.0   int32   `colname:"270059.0"`
 1.0   int32   `colname:"1.0"`
 190002.0   int32   `colname:"190002.0"`
 2.0   int32   `colname:"2.0"`
 200058.0   int32   `colname:"200058.0"`
 1.0   int32   `colname:"1.0"`
 10009.0   int32   `colname:"10009.0"`
 50.0   int32   `colname:"50.0"`
 300001.0   int32   `colname:"300001.0"`
 2.0   int32   `colname:"2.0"`
 0.0   int32   `colname:"0.0"`
 0.0   int32   `colname:"0.0"`
 1.0   int32   `colname:"1.0"`
 0.0   int32   `colname:"0.0"`
 0.0   int32   `colname:"0.0"`

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

func (this *Activities_taskTable) RowBy6.0000999e7(6.0000999e7 string) (row Activities_taskRow,rowIndex int, ok bool) {
	this.RLock()
	defer this.RUnlock()

	for i := 0; i < len(this.rows); i++ {
		if this.rows[i].6.0000999e7 == 6.0000999e7 {
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
         case "6.0000999e7" :
 return 0,true
 case "100000.0" :
 return 1,true
 case "0.0" :
 return 2,true
 case "1033.0" :
 return 3,true
 case "首次充值次数，12121212" :
 return 4,true
 case "首次充值" :
 return 5,true
 case "首次充值" :
 return 6,true
 case "0.0" :
 return 7,true
 case "0.0" :
 return 8,true
 case "0.0" :
 return 9,true
 case "1.0" :
 return 10,true
 case "2002.0" :
 return 11,true
 case "1.0" :
 return 12,true
 case "0.0" :
 return 13,true
 case "0.0" :
 return 14,true
 case "0.0" :
 return 15,true
 case "0.0" :
 return 16,true
 case "270059.0" :
 return 17,true
 case "1.0" :
 return 18,true
 case "190002.0" :
 return 19,true
 case "2.0" :
 return 20,true
 case "200058.0" :
 return 21,true
 case "1.0" :
 return 22,true
 case "10009.0" :
 return 23,true
 case "50.0" :
 return 24,true
 case "300001.0" :
 return 25,true
 case "2.0" :
 return 26,true
 case "0.0" :
 return 27,true
 case "0.0" :
 return 28,true
 case "1.0" :
 return 29,true
 case "0.0" :
 return 30,true
 case "0.0" :
 return 31,true

        default:
            return 0,false
    }
}

func (this *Activities_taskTable) typeByName(colname string) (reflect.Kind, bool) {
	switch colname {
	 case "6.0000999e7" :
 return reflect.int32,true
 case "100000.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "1033.0" :
 return reflect.int32,true
 case "首次充值次数，12121212" :
 return reflect.string,true
 case "首次充值" :
 return reflect.string,true
 case "首次充值" :
 return reflect.string,true
 case "0.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "1.0" :
 return reflect.int32,true
 case "2002.0" :
 return reflect.int32,true
 case "1.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "270059.0" :
 return reflect.int32,true
 case "1.0" :
 return reflect.int32,true
 case "190002.0" :
 return reflect.int32,true
 case "2.0" :
 return reflect.int32,true
 case "200058.0" :
 return reflect.int32,true
 case "1.0" :
 return reflect.int32,true
 case "10009.0" :
 return reflect.int32,true
 case "50.0" :
 return reflect.int32,true
 case "300001.0" :
 return reflect.int32,true
 case "2.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "1.0" :
 return reflect.int32,true
 case "0.0" :
 return reflect.int32,true
 case "0.0" :
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
 return this.rows[row].6.0000999e7,true
 case 1 :
 return this.rows[row].100000.0,true
 case 2 :
 return this.rows[row].0.0,true
 case 3 :
 return this.rows[row].1033.0,true
 case 7 :
 return this.rows[row].0.0,true
 case 8 :
 return this.rows[row].0.0,true
 case 9 :
 return this.rows[row].0.0,true
 case 10 :
 return this.rows[row].1.0,true
 case 11 :
 return this.rows[row].2002.0,true
 case 12 :
 return this.rows[row].1.0,true
 case 13 :
 return this.rows[row].0.0,true
 case 14 :
 return this.rows[row].0.0,true
 case 15 :
 return this.rows[row].0.0,true
 case 16 :
 return this.rows[row].0.0,true
 case 17 :
 return this.rows[row].270059.0,true
 case 18 :
 return this.rows[row].1.0,true
 case 19 :
 return this.rows[row].190002.0,true
 case 20 :
 return this.rows[row].2.0,true
 case 21 :
 return this.rows[row].200058.0,true
 case 22 :
 return this.rows[row].1.0,true
 case 23 :
 return this.rows[row].10009.0,true
 case 24 :
 return this.rows[row].50.0,true
 case 25 :
 return this.rows[row].300001.0,true
 case 26 :
 return this.rows[row].2.0,true
 case 27 :
 return this.rows[row].0.0,true
 case 28 :
 return this.rows[row].0.0,true
 case 29 :
 return this.rows[row].1.0,true
 case 30 :
 return this.rows[row].0.0,true
 case 31 :
 return this.rows[row].0.0,true

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
 return this.rows[row].首次充值次数，12121212,true
 case 5 :
 return this.rows[row].首次充值,true
 case 6 :
 return this.rows[row].首次充值,true

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

