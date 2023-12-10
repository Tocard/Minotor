package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"minotor/data"
	"minotor/es"
	"minotor/thirdapp"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GetNodeStatus(c *gin.Context) {
}

// convertWeiFieldsToNumbers recursively converts string representations of "Wei" fields to actual numbers in a data structure.
func convertWeiFieldsToNumbers(data interface{}) {
	value := reflect.ValueOf(data)

	// If it's a pointer, dereference it
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// If it's a struct, iterate through its fields
	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fieldName := value.Type().Field(i).Name

			// Check if the field name ends with "Wei"
			if strings.HasSuffix(fieldName, "Wei") {
				// If it's a string, attempt conversion to float64
				if field.Kind() == reflect.String {
					floatValue, err := strconv.ParseFloat(field.Interface().(string), 64)
					if err == nil {
						// Update the value in place
						field.Set(reflect.ValueOf(floatValue))
					}
				}
			} else {
				// Recursively process non-"Wei" fields
				convertWeiFieldsToNumbers(field.Interface())
			}
		}
	}

	// If it's a slice or array, iterate through its elements
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)
			// Recursively process elements
			convertWeiFieldsToNumbers(element.Interface())
		}
	}

	// If it's a map, iterate through its values
	if value.Kind() == reflect.Map {
		for _, key := range value.MapKeys() {
			mapValue := value.MapIndex(key)
			// Recursively process values
			convertWeiFieldsToNumbers(mapValue.Interface())
		}
	}
}

func GetAllNodesStatus(c *gin.Context) {
	var OperatorsToBulk [][]byte
	Operators, code := thirdapp.HarvestAllOperatorsInfo()
	if code != 200 {
		c.String(code, string(Operators))
	}
	AllOperators := data.AllOperator{}
	err := json.Unmarshal(Operators, &AllOperators)
	if err != nil {
		c.String(500, err.Error())
	}

	clock := time.Now().Format(time.RFC3339)
	for _, Operator := range AllOperators.Data.Operator {
		Operator.Timestamp = clock
		convertWeiFieldsToNumbers(Operator)
		OperatorJson, _ := json.Marshal(Operator)
		OperatorsToBulk = append(OperatorsToBulk, OperatorJson)
	}

	es.BulkData("minotor-streamr-operator", OperatorsToBulk)
	c.String(201, fmt.Sprintf("%s", OperatorsToBulk))
}

func GetSlashingHistory(c *gin.Context) {
}
