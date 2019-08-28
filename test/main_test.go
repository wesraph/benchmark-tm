package test

import (
	"log"
	"testing"

	"github.com/ndidplatform/smart-contract/abci/did/v1"
	"github.com/ndidplatform/smart-contract/test/common"
	"github.com/ndidplatform/smart-contract/test/data"
	"github.com/ndidplatform/smart-contract/test/utils"
)

func TestRegisterMasterNode(t *testing.T) {
	privKey := utils.GetPrivateKeyFromString(data.MasterNodePrivK)
	publicKeyBytes, err := utils.GeneratePublicKey(&privKey.PublicKey)
	if err != nil {
		log.Fatal(err.Error())
	}
	var param did.RegisterMasterNodeParam
	param.NodeID = data.MasterNodeID
	param.PublicKey = string(publicKeyBytes)
	param.MasterPublicKey = string(publicKeyBytes)
	param.NodeName = ""
	common.RegisterMasterNode(t, data.MasterNodeID, data.MasterNodePrivK, param)
}

func TestSetTx(t *testing.T) {
	var param did.SetTxParam
	param.From = `6abface6-ad51-4ec6-bcf6-17e6042f7eee-AAAA`
	param.To = `efc19d99-df9f-4dc4-a4bc-b54496ac878d-AAAA`
	param.Price = 100.0
	param.Amount = 0.00000001
	common.SetTx(t, data.MasterNodeID, data.MasterNodePrivK, param)
}

func TestSetValidator(t *testing.T) {
	var param did.SetValidatorParam
	param.PublicKey = `kRKM3mkPlogAhWLARAoE9nG+i+fFbZLQDMZoS1O50So=`
	param.Power = 100
	common.SetValidator(t, data.MasterNodeID, data.MasterNodePrivK, param)
}
