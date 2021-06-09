package schema

import (
	"go/ast"
	"horm/dialect"
	"reflect"
)

//Field 映射数据库中的一列
type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}       //被映射的对象
	Name       string            //表明
	Fields     []*Field          //字段
	FieldNames []string          //字段名
	fieldMap   map[string]*Field //字段名与字段的映射关系
}

func (s *Schema) GetFiled(name string) *Field {
	return s.fieldMap[name]
}

func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, filed := range s.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(filed.Name).Interface())
	}
	return fieldValues
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) { //变量是非匿名切导出的
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("horm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
