// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/additionalfield"
	"PasswordManager/ent/emailchallenge"
	"PasswordManager/ent/password"
	"PasswordManager/ent/schema"
	"PasswordManager/ent/session"
	"PasswordManager/ent/totpcredential"
	"PasswordManager/ent/url"
	"PasswordManager/ent/user"
	"PasswordManager/ent/webauthnchallenge"
	"PasswordManager/ent/webauthncredential"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	additionalfieldFields := schema.AdditionalField{}.Fields()
	_ = additionalfieldFields
	// additionalfieldDescKey is the schema descriptor for key field.
	additionalfieldDescKey := additionalfieldFields[1].Descriptor()
	// additionalfield.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	additionalfield.KeyValidator = additionalfieldDescKey.Validators[0].(func([]byte) error)
	// additionalfieldDescKeyIv is the schema descriptor for keyIv field.
	additionalfieldDescKeyIv := additionalfieldFields[2].Descriptor()
	// additionalfield.KeyIvValidator is a validator for the "keyIv" field. It is called by the builders before save.
	additionalfield.KeyIvValidator = additionalfieldDescKeyIv.Validators[0].(func([]byte) error)
	// additionalfieldDescValue is the schema descriptor for value field.
	additionalfieldDescValue := additionalfieldFields[3].Descriptor()
	// additionalfield.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	additionalfield.ValueValidator = additionalfieldDescValue.Validators[0].(func([]byte) error)
	// additionalfieldDescValueIv is the schema descriptor for valueIv field.
	additionalfieldDescValueIv := additionalfieldFields[4].Descriptor()
	// additionalfield.ValueIvValidator is a validator for the "valueIv" field. It is called by the builders before save.
	additionalfield.ValueIvValidator = additionalfieldDescValueIv.Validators[0].(func([]byte) error)
	// additionalfieldDescID is the schema descriptor for id field.
	additionalfieldDescID := additionalfieldFields[0].Descriptor()
	// additionalfield.DefaultID holds the default value on creation for the id field.
	additionalfield.DefaultID = additionalfieldDescID.Default.(func() uuid.UUID)
	emailchallengeFields := schema.EmailChallenge{}.Fields()
	_ = emailchallengeFields
	// emailchallengeDescCode is the schema descriptor for code field.
	emailchallengeDescCode := emailchallengeFields[1].Descriptor()
	// emailchallenge.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	emailchallenge.CodeValidator = emailchallengeDescCode.Validators[0].(func(string) error)
	// emailchallengeDescExpiry is the schema descriptor for expiry field.
	emailchallengeDescExpiry := emailchallengeFields[2].Descriptor()
	// emailchallenge.DefaultExpiry holds the default value on creation for the expiry field.
	emailchallenge.DefaultExpiry = emailchallengeDescExpiry.Default.(func() time.Time)
	// emailchallengeDescID is the schema descriptor for id field.
	emailchallengeDescID := emailchallengeFields[0].Descriptor()
	// emailchallenge.DefaultID holds the default value on creation for the id field.
	emailchallenge.DefaultID = emailchallengeDescID.Default.(func() uuid.UUID)
	passwordFields := schema.Password{}.Fields()
	_ = passwordFields
	// passwordDescName is the schema descriptor for name field.
	passwordDescName := passwordFields[1].Descriptor()
	// password.NameValidator is a validator for the "name" field. It is called by the builders before save.
	password.NameValidator = passwordDescName.Validators[0].(func([]byte) error)
	// passwordDescNameIv is the schema descriptor for nameIv field.
	passwordDescNameIv := passwordFields[2].Descriptor()
	// password.NameIvValidator is a validator for the "nameIv" field. It is called by the builders before save.
	password.NameIvValidator = passwordDescNameIv.Validators[0].(func([]byte) error)
	// passwordDescUsername is the schema descriptor for username field.
	passwordDescUsername := passwordFields[3].Descriptor()
	// password.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	password.UsernameValidator = passwordDescUsername.Validators[0].(func([]byte) error)
	// passwordDescUsernameIv is the schema descriptor for usernameIv field.
	passwordDescUsernameIv := passwordFields[4].Descriptor()
	// password.UsernameIvValidator is a validator for the "usernameIv" field. It is called by the builders before save.
	password.UsernameIvValidator = passwordDescUsernameIv.Validators[0].(func([]byte) error)
	// passwordDescPassword is the schema descriptor for password field.
	passwordDescPassword := passwordFields[5].Descriptor()
	// password.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	password.PasswordValidator = passwordDescPassword.Validators[0].(func([]byte) error)
	// passwordDescPasswordIv is the schema descriptor for passwordIv field.
	passwordDescPasswordIv := passwordFields[6].Descriptor()
	// password.PasswordIvValidator is a validator for the "passwordIv" field. It is called by the builders before save.
	password.PasswordIvValidator = passwordDescPasswordIv.Validators[0].(func([]byte) error)
	// passwordDescID is the schema descriptor for id field.
	passwordDescID := passwordFields[0].Descriptor()
	// password.DefaultID holds the default value on creation for the id field.
	password.DefaultID = passwordDescID.Default.(func() uuid.UUID)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescN is the schema descriptor for n field.
	sessionDescN := sessionFields[1].Descriptor()
	// session.NValidator is a validator for the "n" field. It is called by the builders before save.
	session.NValidator = sessionDescN.Validators[0].(func([]byte) error)
	// sessionDescID is the schema descriptor for id field.
	sessionDescID := sessionFields[0].Descriptor()
	// session.DefaultID holds the default value on creation for the id field.
	session.DefaultID = sessionDescID.Default.(func() uuid.UUID)
	totpcredentialFields := schema.TotpCredential{}.Fields()
	_ = totpcredentialFields
	// totpcredentialDescCreatedAt is the schema descriptor for createdAt field.
	totpcredentialDescCreatedAt := totpcredentialFields[1].Descriptor()
	// totpcredential.DefaultCreatedAt holds the default value on creation for the createdAt field.
	totpcredential.DefaultCreatedAt = totpcredentialDescCreatedAt.Default.(func() time.Time)
	// totpcredentialDescSecret is the schema descriptor for secret field.
	totpcredentialDescSecret := totpcredentialFields[2].Descriptor()
	// totpcredential.SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	totpcredential.SecretValidator = totpcredentialDescSecret.Validators[0].(func(string) error)
	// totpcredentialDescValidated is the schema descriptor for validated field.
	totpcredentialDescValidated := totpcredentialFields[3].Descriptor()
	// totpcredential.DefaultValidated holds the default value on creation for the validated field.
	totpcredential.DefaultValidated = totpcredentialDescValidated.Default.(bool)
	// totpcredentialDescID is the schema descriptor for id field.
	totpcredentialDescID := totpcredentialFields[0].Descriptor()
	// totpcredential.DefaultID holds the default value on creation for the id field.
	totpcredential.DefaultID = totpcredentialDescID.Default.(func() uuid.UUID)
	urlFields := schema.Url{}.Fields()
	_ = urlFields
	// urlDescURL is the schema descriptor for url field.
	urlDescURL := urlFields[1].Descriptor()
	// url.URLValidator is a validator for the "url" field. It is called by the builders before save.
	url.URLValidator = urlDescURL.Validators[0].(func([]byte) error)
	// urlDescUrlIv is the schema descriptor for urlIv field.
	urlDescUrlIv := urlFields[2].Descriptor()
	// url.UrlIvValidator is a validator for the "urlIv" field. It is called by the builders before save.
	url.UrlIvValidator = urlDescUrlIv.Validators[0].(func([]byte) error)
	// urlDescID is the schema descriptor for id field.
	urlDescID := urlFields[0].Descriptor()
	// url.DefaultID holds the default value on creation for the id field.
	url.DefaultID = urlDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescStrengthenedMasterHash is the schema descriptor for strengthenedMasterHash field.
	userDescStrengthenedMasterHash := userFields[2].Descriptor()
	// user.StrengthenedMasterHashValidator is a validator for the "strengthenedMasterHash" field. It is called by the builders before save.
	user.StrengthenedMasterHashValidator = userDescStrengthenedMasterHash.Validators[0].(func([]byte) error)
	// userDescStrengthenedMasterHashSalt is the schema descriptor for strengthenedMasterHashSalt field.
	userDescStrengthenedMasterHashSalt := userFields[3].Descriptor()
	// user.StrengthenedMasterHashSaltValidator is a validator for the "strengthenedMasterHashSalt" field. It is called by the builders before save.
	user.StrengthenedMasterHashSaltValidator = userDescStrengthenedMasterHashSalt.Validators[0].(func([]byte) error)
	// userDescProtectedDatabaseKey is the schema descriptor for protectedDatabaseKey field.
	userDescProtectedDatabaseKey := userFields[4].Descriptor()
	// user.ProtectedDatabaseKeyValidator is a validator for the "protectedDatabaseKey" field. It is called by the builders before save.
	user.ProtectedDatabaseKeyValidator = userDescProtectedDatabaseKey.Validators[0].(func([]byte) error)
	// userDescProtectedDatabaseKeyIv is the schema descriptor for protectedDatabaseKeyIv field.
	userDescProtectedDatabaseKeyIv := userFields[5].Descriptor()
	// user.ProtectedDatabaseKeyIvValidator is a validator for the "protectedDatabaseKeyIv" field. It is called by the builders before save.
	user.ProtectedDatabaseKeyIvValidator = userDescProtectedDatabaseKeyIv.Validators[0].(func([]byte) error)
	// userDescVerified is the schema descriptor for verified field.
	userDescVerified := userFields[7].Descriptor()
	// user.DefaultVerified holds the default value on creation for the verified field.
	user.DefaultVerified = userDescVerified.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	webauthnchallengeFields := schema.WebAuthnChallenge{}.Fields()
	_ = webauthnchallengeFields
	// webauthnchallengeDescID is the schema descriptor for id field.
	webauthnchallengeDescID := webauthnchallengeFields[0].Descriptor()
	// webauthnchallenge.DefaultID holds the default value on creation for the id field.
	webauthnchallenge.DefaultID = webauthnchallengeDescID.Default.(func() uuid.UUID)
	webauthncredentialFields := schema.WebAuthnCredential{}.Fields()
	_ = webauthncredentialFields
	// webauthncredentialDescCreatedAt is the schema descriptor for createdAt field.
	webauthncredentialDescCreatedAt := webauthncredentialFields[2].Descriptor()
	// webauthncredential.DefaultCreatedAt holds the default value on creation for the createdAt field.
	webauthncredential.DefaultCreatedAt = webauthncredentialDescCreatedAt.Default.(func() time.Time)
	// webauthncredentialDescID is the schema descriptor for id field.
	webauthncredentialDescID := webauthncredentialFields[0].Descriptor()
	// webauthncredential.DefaultID holds the default value on creation for the id field.
	webauthncredential.DefaultID = webauthncredentialDescID.Default.(func() uuid.UUID)
}
