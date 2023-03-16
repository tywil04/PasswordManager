package middleware

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/exceptions"
	"PasswordManager/api/lib/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
			c.Abort()
			return
		}

		valid, user, session := helpers.ValidateSession(authToken)
		if valid {
			c.Set("authedUser", user)
			c.Set("authedSession", session)
			c.Next()
			return
		}

		c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
		c.Abort()
	}
}

func processParamsRecursive(c *gin.Context, inputValue reflect.Value) map[string]any {
	inputType := inputValue.Type()

	data := make(map[string]any, inputValue.NumField())

	for i := 0; i < inputValue.NumField(); i++ {
		if !c.IsAborted() {
			indexType := inputType.Field(i)
			indexValue := inputValue.Field(i)
			jsonTagName := indexType.Tag.Get("json")

			if indexType.Tag.Get("pmOptional") != "true" && indexValue.IsZero() {
				fmt.Println(indexValue.String())
				c.JSON(400, exceptions.Builder(jsonTagName, exceptions.MissingParam))
				c.Abort()
				return map[string]any{}
			} else if indexType.Tag.Get("pmOptional") == "true" && indexValue.IsZero() {
				continue
			}

			if indexValue.Kind() == reflect.Struct {
				data[jsonTagName] = processParamsRecursive(c, indexValue)
				continue
			}

			if (indexValue.Kind() == reflect.Slice || indexValue.Kind() == reflect.Array) && reflect.TypeOf(indexValue.Interface()).Elem().Kind() == reflect.Struct {
				tempData := make([]map[string]any, indexValue.Len())
				for j := 0; j < indexValue.Len(); j++ {
					tempData[j] = processParamsRecursive(c, indexValue.Index(j))
				}
				data[jsonTagName] = tempData
				continue
			}

			var d any
			var dE error
			var dA exceptions.Addition = exceptions.Undefined

			switch indexType.Tag.Get("pmParseType") {
			case "base64":
				d, dE = base64.StdEncoding.DecodeString(indexValue.String())
				dA = exceptions.Base64
			case "uuid":
				d, dE = uuid.Parse(indexValue.String())
				dA = exceptions.Uuid
			case "hexcolour":
				_, dE = strconv.ParseInt(indexValue.String(), 16, 64)
				dA = exceptions.HexColour
				d = indexValue.String()
			case "email":
				regex, _ := regexp.Compile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))`)
				valid := regex.MatchString(indexValue.String())
				if !valid {
					dE = errors.New("invalid email")
					dA = exceptions.Email
				}
				d = indexValue.String()
			case "emailCode":
				regex, _ := regexp.Compile(`^[a-zA-Z0-9]{4}\-[a-zA-Z0-9]{4}$`)
				valid := regex.MatchString(indexValue.String())
				if !valid {
					dE = errors.New("invalid email code")
					dA = exceptions.EmailCode
				}
				d = indexValue.String()
			case "totpCode":
				regex, _ := regexp.Compile(`^[0-9]{6}$`)
				valid := regex.MatchString(indexValue.String())
				if !valid {
					dE = errors.New("invalid totp code")
					dA = exceptions.TotpCode
				}
				d = indexValue.String()
			default:
				d = indexValue.Interface()
			}

			if dE != nil {
				c.JSON(400, exceptions.Builder(jsonTagName, exceptions.ParsingParam, dA))
				c.Abort()
				return map[string]any{}
			}

			data[jsonTagName] = d
		}
	}

	if !c.IsAborted() {
		return data
	}

	return map[string]any{}
}

func ProcessParams(structure any) gin.HandlerFunc {
	return func(c *gin.Context) {
		structureType := reflect.TypeOf(structure)
		input := reflect.New(structureType).Interface()

		bindingErr := c.ShouldBind(input)

		if bindingErr != nil && c.Request.ContentLength > 0 {
			fmt.Println(bindingErr)
			c.JSON(400, exceptions.Builder("body", exceptions.Invalid, exceptions.JsonOrXml))
			c.Abort()
			return
		}

		params := processParamsRecursive(c, reflect.ValueOf(input).Elem())
		if c.IsAborted() {
			return
		}

		c.Set("params", params)
		c.Next()
	}
}
