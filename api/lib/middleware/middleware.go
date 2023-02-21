package middleware

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
			c.Abort()
			return
		}

		parts := strings.Split(authToken, ";")
		if len(parts) != 3 {
			c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
			c.Abort()
			return
		}

		sessionId := parts[0]
		encodedSalt := parts[1]
		signature := parts[2]

		decodedSessionId, dsiErr := base64.StdEncoding.DecodeString(sessionId)
		if dsiErr != nil {
			c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
			c.Abort()
			return
		}

		decodedSignature, dsErr := base64.StdEncoding.DecodeString(signature)
		if dsErr != nil {
			c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
			c.Abort()
			return
		}

		parsedDecodedSessionId, pdsiErr := uuid.Parse(string(decodedSessionId[:]))
		if pdsiErr != nil {
			c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
			c.Abort()
			return
		}

		session, sessionErr := db.Client.Session.Get(db.Context, parsedDecodedSessionId)
		if sessionErr != nil {
			c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
			c.Abort()
			return
		}

		if session.Expiry.Before(time.Now()) {
			db.Client.Session.DeleteOne(session).Exec(db.Context)
			c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
			c.Abort()
			return
		}

		user, userErr := session.QueryUser().First(db.Context)
		if userErr != nil {
			c.JSON(401, exceptions.Builder("Authorization", exceptions.InvalidHeader, exceptions.Authtoken))
			c.Abort()
			return
		}

		publicKey := cryptography.ImportPublicKey(session.N, session.E)
		valid := cryptography.VerifySignature(publicKey, decodedSignature, user.Email+base64.StdEncoding.EncodeToString(user.StrengthenedMasterHash)+encodedSalt)

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
					dE = errors.New("Invalid email")
					dA = exceptions.Email
				}
				d = indexValue.String()
			default:
				d = indexValue.Interface()
			}

			if dE != nil {
				c.JSON(400, exceptions.Builder(jsonTagName, exceptions.ParsingParam, dA))
				c.Abort()
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

		bindingErr := c.Bind(input)
		if bindingErr != nil {
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
