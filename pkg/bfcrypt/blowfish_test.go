package bfcrypt

import (
	"fmt"
	"testing"

	json "github.com/json-iterator/go"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt58(t *testing.T) {
	key := []byte("PPpn7ugdcTa4DTUdqSkxSeR7gTsv93Q576ug8QMdeLbVFPte")
	src := []byte("2342342342342342342342342342342")

	encodeStr, err := Encrypt58(src, key)
	assert.NoError(t, err)

	out, err := Decrypt58(encodeStr, key)
	assert.NoError(t, err)

	assert.Equal(t, src, out)
}

func TestEncrypt(t *testing.T) {
	key := []byte("PPpn7ugdcTa4DTUdqSkxSeR7gTsv93Q576ug8QMdeLbVFPte")
	src := []byte("2342342342342342342342342342342")

	encodeStr, err := Encrypt(src, key)
	assert.NoError(t, err)

	out, err := Decrypt(encodeStr, key)
	assert.NoError(t, err)

	assert.Equal(t, src, out)
}

func BenchmarkEncrypt(b *testing.B) {
	key := []byte("PPpn7ugdcTa4DTUdqSkxSeR7gTsv93Q576ug8QMdeLbVFPte")
	src := []byte("2342342342342342342342342342342")

	for i := 0; i < b.N; i++ {
		Encrypt(src, key)
	}
}

func ExampleEncrypt58() {
	// VksnConfig configuration for sn generator
	type VksnConfig struct {
		Name              string `json:"Name"`
		Version           string `json:"Version"`
		Debug             bool   `json:"Debug"`
		LogFullPath       string `json:"logFullPath"`
		LogFilenamePrefix string `json:"LogFilenamePrefix"`
		//
		Total int `json:"total"`
	}

	key := []byte("PPpn7ugdcTa4DTUdqSkxSeR7gTsv93Q576ug8QMdeLbVFPte")
	// price := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	set := make(map[string]string, 0)
	set["tsingson"] = "tsingson "

	cfg := &VksnConfig{
		Name:              "apkssssssssssssn",
		Debug:             true,
		LogFullPath:       "/Users/qinshen/go/bin/log/",
		LogFilenamePrefix: "apkserial-",
		Total:             5,
	}

	litter.Dump(cfg)

	fmt.Println("-------------")

	cfgByte, err := json.Marshal(cfg)
	if err != nil {
	}

	encodeStr, err := Encrypt58(cfgByte, key)
	// fmt.Println("encode=", encodeStr)

	litter.Dump(encodeStr)
	fmt.Println("-------------")

	price2, err := Decrypt58(encodeStr, key)

	cfg2 := &VksnConfig{}

	err = json.Unmarshal(price2, &cfg2)
	if err != nil {
	}
	fmt.Println("-------------")

	litter.Dump(cfg2)

	fmt.Println("-------------")
}
