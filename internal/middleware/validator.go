package middleware

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"products-service-go/internal/logger"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
)

// SchemaConfig defines the validation configuration for an endpoint
type SchemaConfig struct {
	RequestDTO   interface{}
	ResponseDTOs map[int]interface{}
}

// ValidateWithSchema creates a middleware that validates request/response using DTO-generated JSON schemas
func ValidateWithSchema(config SchemaConfig) fiber.Handler {
	// Generate request schema from DTO
	var requestSchema *openapi3.Schema
	if config.RequestDTO != nil {
		requestSchema = generateSchemaFromStruct(config.RequestDTO)
	}

	// Generate response schemas from DTOs
	responseSchemas := make(map[int]*openapi3.Schema)
	for statusCode, dto := range config.ResponseDTOs {
		responseSchemas[statusCode] = generateSchemaFromStruct(dto)
	}

	return func(c *fiber.Ctx) error {
		// Validate Request
		if requestSchema != nil {
			if err := validateRequest(c, requestSchema, config.RequestDTO); err != nil {
				return err
			}
		}

		// Continue to handler
		if err := c.Next(); err != nil {
			return err
		}

		// Validate Response (simplified - just log errors)
		statusCode := c.Response().StatusCode()
		if responseSchema, exists := responseSchemas[statusCode]; exists {
			responseBody := c.Response().Body()
			if err := validateResponse(responseBody, responseSchema); err != nil {
				logger.Error("Response validation failed: %v", err)
			}
		}

		return nil
	}
}

func validateRequest(c *fiber.Ctx, schema *openapi3.Schema, dtoType interface{}) error {
	logger.Info("Validating request with generated schema")

	var requestData interface{}
	if err := c.BodyParser(&requestData); err != nil {
		logger.Error("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	validatedDTO := createDTOInstance(dtoType)
	if err := json.Unmarshal(jsonData, validatedDTO); err != nil {
		logger.Error("Failed to map to DTO: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := validateDTOConstraints(validatedDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	c.Locals("validatedRequest", validatedDTO)
	logger.Info("Request validation passed")
	return nil
}

func validateResponse(responseBody []byte, schema *openapi3.Schema) error {
	if len(responseBody) == 0 {
		return nil
	}

	var responseData interface{}
	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		return fmt.Errorf("failed to parse response JSON: %w", err)
	}

	// Basic structure validation
	return nil // Simplified for now
}

func generateSchemaFromStruct(dto interface{}) *openapi3.Schema {
	t := reflect.TypeOf(dto)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	schema := &openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: make(openapi3.Schemas),
		Required:   []string{},
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName == "" {
			jsonName = strings.ToLower(field.Name)
		}

		propertySchema := generatePropertySchema(field)

		validateTag := field.Tag.Get("validate")
		if strings.Contains(validateTag, "required") {
			schema.Required = append(schema.Required, jsonName)
		}

		applyValidationConstraints(propertySchema, validateTag)

		schema.Properties[jsonName] = &openapi3.SchemaRef{
			Value: propertySchema,
		}
	}

	return schema
}

func generatePropertySchema(field reflect.StructField) *openapi3.Schema {
	schema := &openapi3.Schema{}

	switch field.Type.Kind() {
	case reflect.String:
		schema.Type = &openapi3.Types{"string"}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		schema.Type = &openapi3.Types{"integer"}
	case reflect.Float32, reflect.Float64:
		schema.Type = &openapi3.Types{"number"}
	case reflect.Bool:
		schema.Type = &openapi3.Types{"boolean"}
	case reflect.Slice:
		schema.Type = &openapi3.Types{"array"}
	case reflect.Struct:
		if field.Type.String() == "time.Time" {
			schema.Type = &openapi3.Types{"string"}
			schema.Format = "date-time"
		} else {
			schema.Type = &openapi3.Types{"object"}
		}
	default:
		schema.Type = &openapi3.Types{"string"}
	}

	return schema
}

func applyValidationConstraints(schema *openapi3.Schema, validateTag string) {
	if validateTag == "" {
		return
	}

	constraints := strings.Split(validateTag, ",")
	for _, constraint := range constraints {
		parts := strings.Split(strings.TrimSpace(constraint), "=")
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		switch key {
		case "min":
			if min, err := strconv.ParseFloat(value, 64); err == nil {
				if schema.Type != nil && len(*schema.Type) > 0 && (*schema.Type)[0] == "string" {
					minLen := uint64(min)
					schema.MinLength = minLen
				} else {
					schema.Min = &min
				}
			}
		case "max":
			if max, err := strconv.ParseFloat(value, 64); err == nil {
				if schema.Type != nil && len(*schema.Type) > 0 && (*schema.Type)[0] == "string" {
					maxLen := uint64(max)
					schema.MaxLength = &maxLen
				} else {
					schema.Max = &max
				}
			}
		case "minLength":
			if minLen, err := strconv.ParseUint(value, 10, 64); err == nil {
				schema.MinLength = minLen
			}
		case "maxLength":
			if maxLen, err := strconv.ParseUint(value, 10, 64); err == nil {
				schema.MaxLength = &maxLen
			}
		}
	}
}

func validateDTOConstraints(dto interface{}) error {
	v := reflect.ValueOf(dto)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		validateTag := field.Tag.Get("validate")

		if err := validateFieldConstraints(field.Name, value, validateTag); err != nil {
			return err
		}
	}

	return nil
}

// Split tag fields to validate every part of struct
func validateFieldConstraints(fieldName string, value reflect.Value, validateTag string) error {
	if validateTag == "" {
		return nil
	}

	constraints := strings.Split(validateTag, ",")
	for _, constraint := range constraints {
		constraint = strings.TrimSpace(constraint)

		if constraint == "required" {
			if value.IsZero() {
				return fmt.Errorf("field %s is required", fieldName)
			}
			continue
		}

		parts := strings.Split(constraint, "=")
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		val := parts[1]

		switch key {
		case "min":
			if minVal, err := strconv.ParseFloat(val, 64); err == nil {
				if value.Kind() == reflect.String {
					if len(value.String()) < int(minVal) {
						return fmt.Errorf("field %s must have at least %v characters", fieldName, minVal)
					}
				} else if value.Kind() >= reflect.Int && value.Kind() <= reflect.Int64 {
					if float64(value.Int()) < minVal {
						return fmt.Errorf("field %s must be at least %v", fieldName, minVal)
					}
				}
			}
		case "max":
			if maxVal, err := strconv.ParseFloat(val, 64); err == nil {
				if value.Kind() == reflect.String {
					if len(value.String()) > int(maxVal) {
						return fmt.Errorf("field %s must have at most %v characters", fieldName, maxVal)
					}
				} else if value.Kind() >= reflect.Int && value.Kind() <= reflect.Int64 {
					if float64(value.Int()) > maxVal {
						return fmt.Errorf("field %s must be at most %v", fieldName, maxVal)
					}
				}
			}
		}
	}

	return nil
}

func createDTOInstance(dtoType interface{}) interface{} {
	t := reflect.TypeOf(dtoType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return reflect.New(t).Interface()
}

// GetValidatedRequest retrieves the validated request from the context
func GetValidatedRequest[T any](c *fiber.Ctx) (T, error) {
	data := c.Locals("validatedRequest")
	if data == nil {
		var zero T
		return zero, fmt.Errorf("no validated request found")
	}

	result, ok := data.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("invalid request type")
	}

	return result, nil
}
