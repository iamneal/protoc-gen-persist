package service

import(
	"fmt"
	"database/sql"
	"reflect"
	"strings"
	_"github.com/lib/pq"
	//"github.com/golang/protobuf/ptypes/timestamp"
	//"github.com/jmoiron/sqlx/reflectx"
)


/*  goal of this project: take SMSGroupMstr struct,  => row in the database => smsgroupmstr struct */

func ParseStruct(st interface{}) {
	s := reflect.ValueOf(st).Elem()
	typeof := s.Type()
	fmt.Printf("typeof: %+v \n", typeof)
	s.Field(0).SetInt(77)
	for i:= 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeof.Field(i).Name, f.Type(), f.Interface())
	}
}

func SnakeMatcher(snake string) func(string) bool {
	splitSnake := strings.Split(snake, "_")
	for i, elm := range splitSnake {
		splitSnake[i] = strings.Title(elm)
	}
	camelCased := strings.Join(splitSnake, "")
	return func(camel string) bool {
		fmt.Printf("camelCased: %s  input: %s ", camelCased, camel)
		return (camel == camelCased)
	}
}

type ServiceImpl struct {
	db *sql.DB
}

func NewServiceImpl(conString string) *ServiceImpl {
	return &ServiceImpl{ }
}
func (s *ServiceImpl) Insert(query string, arguments []string, in interface{}, out interface{}) (error){
	argVals := make([]interface{}, len(arguments))
	ref := reflect.ValueOf(in).Elem()

	for i, arg := range arguments {
		argVals[i] = ref.FieldByNameFunc(SnakeMatcher(arg))
	}
	// can't use query row because *row has no columns field :(
	rows, err := s.db.Query(query, argVals...)

	if err != nil {
		return err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	for rows.Next() {
		if rows.Err() != nil {
			return err
		}
		refRes := reflect.ValueOf(out)
		if refRes.Kind() != reflect.Ptr {
			return fmt.Errorf("out object must be pointer to struct")
		}
		arrOutFields := make([]interface{}, len(cols))

		valRes := refRes.Elem()
		for i, col := range cols {
			// spread our out interface into an array of pointers to the original struct values
			arrOutFields[i] = valRes.FieldByNameFunc(SnakeMatcher(col)).Addr()
		}
		// scan into our out struct
		err = rows.Scan(arrOutFields...)
		return err
	}
	return nil
}

func (s *ServiceImpl) Select(query string, arguments[]string, st interface {}) (interface{}, error) {
	return nil, nil
}
