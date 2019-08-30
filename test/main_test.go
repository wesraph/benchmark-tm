package test

import (
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/wesraph/benchmark-tm/abci/did/v1"
	"github.com/wesraph/benchmark-tm/test/common"
	"github.com/wesraph/benchmark-tm/test/data"
	"github.com/wesraph/benchmark-tm/test/utils"
)

func TestRegisterMasterNodeEcdsa(t *testing.T) {
	privKey, err := utils.GetPrivateKeyFromStringEcdsa(data.MasterNodePrivEcdsa)
	if err != nil {
		panic(err)
	}

	x509EncodedPub, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		panic(err)
	}
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	var param did.RegisterMasterNodeParam
	param.NodeID = data.MasterNodeID
	param.PublicKey = string(pemEncodedPub)
	param.MasterPublicKey = string(pemEncodedPub)
	param.NodeName = ""

	err = common.RegisterMasterNode(data.MasterNodeID, data.MasterNodePrivEcdsa, param, common.EcdsaPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PASS")
}

func TestSetTxEcdsa(t *testing.T) {
	var param did.SetTxParam
	param.From = `6abface6-ad51-4ec6-bcf6-17e6042f7eee-AAAA`
	param.To = `efc19d99-df9f-4dc4-a4bc-b54496ac878d-AAAA`
	param.Price = 100.0
	param.Amount = 0.00000001
	err := common.SetTx(data.MasterNodeID, data.MasterNodePrivEcdsa, param, common.EcdsaPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PASS")
}

//func TestRegisterMasterNode(t *testing.T) {
//privKey, err := utils.GetPrivateKeyFromString(data.MasterNodePrivRSA)
//if err != nil {
//panic(err)
//}
//publicKeyBytes, err := utils.GeneratePublicKey(&privKey.PublicKey)
//if err != nil {
//log.Fatal(err.Error())
//}
//var param did.RegisterMasterNodeParam
//param.NodeID = data.MasterNodeID
//param.PublicKey = string(publicKeyBytes)
//param.MasterPublicKey = string(publicKeyBytes)
//param.NodeName = ""
//common.RegisterMasterNode(t, data.MasterNodeID, data.MasterNodePrivRSA, param, common.RSAPrivateKey)
//}

//func TestSetTx(t *testing.T) {
//var param did.SetTxParam
//param.From = `6abface6-ad51-4ec6-bcf6-17e6042f7eee-AAAA`
//param.To = `efc19d99-df9f-4dc4-a4bc-b54496ac878d-AAAA`
//param.Price = 100.0
//param.Amount = 0.00000001
//common.SetTx(t, data.MasterNodeID, data.MasterNodePrivRSA, param, common.RSAPrivateKey)
//}

//func TestSetValidator(t *testing.T) {
//var param did.SetValidatorParam
//param.PublicKey = `kRKM3mkPlogAhWLARAoE9nG+i+fFbZLQDMZoS1O50So=`
//param.Power = 100
//common.SetValidator(t, data.MasterNodeID, data.MasterNodePrivRSA, param, common.RSAPrivateKey)
//}
