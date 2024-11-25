// backend-webUE/utils/operator.go
package utils

import (
	"backend-webUE/models"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func generateRandomMsisdn(length int) string {
	const digits = "0123456789"

	msisdn := make([]byte, length)
	for i := 0; i < length; i++ {
		msisdn[i] = digits[rand.Intn(len(digits))]
	}
	return string(msisdn)
}

type OperatorConfig struct {
	PlmnId            models.PlmnId
	Amf               string
	UeConfiguredNssai []models.Snssai
	UeDefaultNssai    []models.Snssai
	// Add other necessary fields from your config
}

type Operator struct {
	config *OperatorConfig
}

func NewOperator(cfg *OperatorConfig) *Operator {
	return &Operator{
		config: cfg,
	}
}

func (o *Operator) GenerateUe() *models.UeProfile {
	ue := &models.UeProfile{
		PlmnId:          o.config.PlmnId,
		Amf:             o.config.Amf,
		ConfiguredSlice: o.config.UeConfiguredNssai,
		DefaultSlice:    o.config.UeDefaultNssai,
		// Initialize other fields
	}

	// Generate random values for the UE profile
	ue.Supi = o.randSupi()
	ue.Key = o.randUeKey()
	ue.Op = o.randOp()
	ue.Imei = o.randImei()
	ue.Imeisv = o.randImeiSv()
	// Set other fields as needed

	return ue
}

func (o *Operator) randUeKey() string {
	return md5Hash(randSeq(16))
}

func (o *Operator) randOp() string {
	return md5Hash(randSeq(16))
}

func (o *Operator) randSupi() string {
	mcc := o.config.PlmnId.Mcc
	mnc := o.config.PlmnId.Mnc
	mcclen := len(mcc)
	mnclen := len(mnc)
	msisdnlen := 15 - mcclen - mnclen
	prefix := mcc + mnc
	msisdn := generateRandomMsisdn(msisdnlen)
	return "imsi-" + prefix + msisdn
}

func (o *Operator) randImei() string {
	// Generate a random IMEI
	return generateRandomMsisdn(15)
}

func (o *Operator) randImeiSv() string {
	// Generate a random IMEISV
	return generateRandomMsisdn(16)
}
