package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	parser := NewParser()
	parser.SetContexter(newTestContexter())

	t.Run("1", func(t *testing.T) {
		parser.SetFormat("'abc'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "abc", tk)
	})
	t.Run("2", func(t *testing.T) {
		parser.SetFormat("'TSign '")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "TSign ", tk)
	})
	t.Run("3", func(t *testing.T) {
		parser.SetFormat("'TSign '+'A'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "TSign A", tk)
	})
	t.Run("4", func(t *testing.T) {
		parser.SetFormat("'TSign ' +'A'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "TSign A", tk)
	})
	t.Run("5", func(t *testing.T) {
		parser.SetFormat("'TSign ' + 'A'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "TSign A", tk)
	})
	t.Run("6", func(t *testing.T) {
		parser.SetFormat("'TSign '+ 'A'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "TSign A", tk)
	})
	t.Run("7", func(t *testing.T) {
		parser.SetFormat("'TSign '+ 'A'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "TSign A", tk)
	})
	t.Run("8", func(t *testing.T) {
		parser.SetFormat("'TSign ' + $SerialNumber + ':'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "TSign {SerialNumber}:", tk)
	})
	t.Run("9", func(t *testing.T) {
		parser.SetFormat("$ReqPath+'\\n'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "{ReqPath}\\n", tk)
	})
	t.Run("10", func(t *testing.T) {
		parser.SetFormat("$ReqPath+'\n'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "{ReqPath}\n", tk)
	})
	t.Run("11", func(t *testing.T) {
		parser.SetFormat("$ReqPath + '\n'")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "{ReqPath}\n", tk)
	})
	t.Run("12", func(t *testing.T) {
		parser.SetFormat("$ReqPath + '\n' + $RegBody")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "{ReqPath}\n{RegBody}", tk)
	})
	t.Run("13", func(t *testing.T) {
		parser.SetFormat("hmac_sha1()")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "{{hmac_sha1:[]}}", tk)
	})
	t.Run("14", func(t *testing.T) {
		parser.SetFormat("hmac_sha1('abc')")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "{{hmac_sha1:[abc]}}", tk)
	})
	t.Run("15", func(t *testing.T) {
		parser.SetFormat("hmac_sha1('abc', 'def')")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "{{hmac_sha1:[abc def]}}", tk)
	})
	t.Run("16", func(t *testing.T) {
		parser.SetFormat("hmac_sha1($RegPath + 'abc')")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "{{hmac_sha1:[{RegPath}abc]}}", tk)
	})
	t.Run("17", func(t *testing.T) {
		parser.SetFormat("urlsafe_base64(hmac_sha1($ReqPath + '\\n' + $ReqBody, $SecretKey))")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "{{urlsafe_base64:[{{hmac_sha1:[{ReqPath}\\n{ReqBody} {SecretKey}]}}]}}", tk)
	})
	t.Run("18", func(t *testing.T) {
		parser.SetFormat("'TSign ' + $SerialNumber + ':' + urlsafe_base64(hmac_sha1($ReqPath + '\\n' + $ReqBody, $SecretKey))")
		tk, err := parser.GetToken()
		assert.Nil(t, err)
		assert.Equal(t, "TSign {SerialNumber}:{{urlsafe_base64:[{{hmac_sha1:[{ReqPath}\\n{ReqBody} {SecretKey}]}}]}}", tk)
	})
}
