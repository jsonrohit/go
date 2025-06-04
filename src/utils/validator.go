package utils

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// ValidateStruct validates a struct based on validation tags
func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	// Handle pointer to struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if v.Kind() != reflect.Struct {
		return errors.New("input is not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("validate")

		if tag == "" {
			continue
		}

		fieldName := strings.ToLower(fieldType.Name)
		fieldValue := field.Interface()

		if err := validateField(fieldName, fieldValue, tag); err != nil {
			return err
		}
	}

	return nil
}

// validateField validates a single field based on validation rules
func validateField(fieldName string, value interface{}, tag string) error {
	rules := strings.Split(tag, ",")

	for _, rule := range rules {
		rule = strings.TrimSpace(rule)

		switch {
		case rule == "required":
			if isEmpty(value) {
				return errors.New(fieldName + " is required")
			}

		case strings.HasPrefix(rule, "min="):
			minStr := strings.TrimPrefix(rule, "min=")
			min, err := strconv.Atoi(minStr)
			if err != nil {
				continue
			}

			if err := validateMin(fieldName, value, min); err != nil {
				return err
			}

		case strings.HasPrefix(rule, "max="):
			maxStr := strings.TrimPrefix(rule, "max=")
			max, err := strconv.Atoi(maxStr)
			if err != nil {
				continue
			}

			if err := validateMax(fieldName, value, max); err != nil {
				return err
			}

		case rule == "email":
			if err := validateEmail(fieldName, value); err != nil {
				return err
			}
		}
	}

	return nil
}

// isEmpty checks if a value is empty
func isEmpty(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v) == ""
	case int, int8, int16, int32, int64:
		return v == 0
	case uint, uint8, uint16, uint32, uint64:
		return v == 0
	case float32, float64:
		return v == 0
	default:
		return value == nil
	}
}

// validateMin validates minimum value/length
func validateMin(fieldName string, value interface{}, min int) error {
	switch v := value.(type) {
	case string:
		if len(v) < min {
			return errors.New(fieldName + " must be at least " + strconv.Itoa(min) + " characters")
		}
	case int:
		if v < min {
			return errors.New(fieldName + " must be at least " + strconv.Itoa(min))
		}
	}
	return nil
}

// validateMax validates maximum value/length
func validateMax(fieldName string, value interface{}, max int) error {
	switch v := value.(type) {
	case string:
		if len(v) > max {
			return errors.New(fieldName + " must be at most " + strconv.Itoa(max) + " characters")
		}
	case int:
		if v > max {
			return errors.New(fieldName + " must be at most " + strconv.Itoa(max))
		}
	}
	return nil
}

// validateEmail validates email format
func validateEmail(fieldName string, value interface{}) error {
	if str, ok := value.(string); ok {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(str) {
			return errors.New(fieldName + " must be a valid email address")
		}
	}
	return nil
}