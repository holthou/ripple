package data

import (
	"encoding/hex"
	"fmt"
	"github.com/test-go/testify/require"
	"strings"
	"testing"
)

func TestSignerListSet(t *testing.T) {
	siger := []string{
		"rNqui8PQUCiQ6Emyigsf3waEcKRQEvge4B",
		"rU19Gx4ftUEYv9d8pgkRviQ9bCAyh7FbKR",
		"rhzey2ok2z1Aho9U8xBRmdyH8B92bZzCbi",
		"rMGH87xMW8BPNDcjCBH5m7PhmKCZDHNA7E",
		"rwnz9Jd45MN6EkamrxFWqrvcxRyvaFiwrN",
	}
	sigerAccount := make([]*Account, 0, len(siger))
	for _, str := range siger {
		account, err := NewAccountFromAddress(str)
		require.NoError(t, err)
		fmt.Println(strings.ToUpper(hex.EncodeToString(account.Bytes())))
		sigerAccount = append(sigerAccount, account)
	}
	sender, err := NewAccountFromAddress("rfKsmLP6sTfVGDvga6rW6XbmSFUzc3G9f3")
	require.NoError(t, err)
	fmt.Printf("-- %s \n", strings.ToUpper(hex.EncodeToString(sender.Bytes())))
	fee, err := NewNativeValue(10)
	require.NoError(t, err)
	pk, err := hex.DecodeString("ED55E02DC3390812D45C6AD310C86D577209693583BBD2DD13717C092963FDBD4F")
	require.NoError(t, err)
	sg, err := hex.DecodeString("01B47532F6A64D1A34881D014B79DB7102EC8A4BDE728B68296F46A1097A0DDEAED9447F681B78997C09FFC9545EC9E6687F5901274CF4AD95161511C9961D08")
	require.NoError(t, err)
	targetStr := "12000C2403A065D320230000000368400000000000000A7321ED55E02DC3390812D45C6AD310C86D577209693583BBD2DD13717C092963FDBD4F744001B47532F6A64D1A34881D014B79DB7102EC8A4BDE728B68296F46A1097A0DDEAED9447F681B78997C09FFC9545EC9E6687F5901274CF4AD95161511C9961D088114453A8CB3100A157D6AE27E8F1F44D8FEB6C5B9D5F4EB130001811497D9B9AA807973A5202C14D04634538FA25C1960E1EB1300018114822516E2D524184354DEBCED6F4D9D06B662A0D0E1EB13000181142BCFC99ED62216759D23445037655B9275AB3698E1EB1300018114DE46C40D0303AFA60D81C8E046A4F3E8FE4F43FCE1EB130001811463AA7F225CF8F049165D283AD1B3A09418D9F65DE1F1"

	tx := &SignerListSet{
		SignerQuorum:  3,
		SignerEntries: make([]SignerEntryInfo, 0, 5),
	}
	tx.TransactionType = SIGNER_LIST_SET
	for _, a := range sigerAccount {
		one := uint16(1)
		tx.SignerEntries = append(tx.SignerEntries, SignerEntryInfo{SignerEntry{Account: a, SignerWeight: &one}})
	}
	tx.Account = *sender
	tx.Sequence = 60843475
	tx.Fee = *fee

	require.Equal(t, strings.ToUpper(hex.EncodeToString(pk)), "ED55E02DC3390812D45C6AD310C86D577209693583BBD2DD13717C092963FDBD4F")
	require.Equal(t, strings.ToUpper(hex.EncodeToString(sg)), "01B47532F6A64D1A34881D014B79DB7102EC8A4BDE728B68296F46A1097A0DDEAED9447F681B78997C09FFC9545EC9E6687F5901274CF4AD95161511C9961D08")
	tx.InitialiseForSigning()
	copy(tx.GetPublicKey().Bytes(), pk)
	*tx.GetSignature() = sg

	_, msg, err := Raw(tx)
	require.NoError(t, err)
	require.Equal(t, strings.ToUpper(hex.EncodeToString(msg)), targetStr)
}

// 验证签名格式
func TestMultiSign(t *testing.T) {
	// FE153AE47F7E90755C7E851267C1CBA2B91D7BC574DC8F94736F3047BD2CD6B9
	type SignInfo struct {
		account string
		pk      string
		sg      string
	}
	signList := []SignInfo{
		{"rwnz9Jd45MN6EkamrxFWqrvcxRyvaFiwrN",
			"ED0D0C4028D416D565C45ABD675E56D1114DB79BD7B788608AE8F38B93636B118B",
			"E949A301D12BE2825D300BEED564EA95F4A8D1D341E33FDAD96C486B23DC712B6FCBAFC04899F2CD1B2EF857808ABA57DCA24640297860429ABF2E21D0FC0F07",
		},
		{"rU19Gx4ftUEYv9d8pgkRviQ9bCAyh7FbKR",
			"ED07CDB4A6FD3D67F16526AC1EF4DA1E9C28C780BF6771BE85FEC94F19F464CDCC",
			"E1E4F06FEDF78F8A9224AA0B3C04FFCA1EB59006B3F9799BBA0FF3AD5CEDE09EEC5D7D6173DE93D0A5820E7F066250B22C3E9A7EB4C66CFCF4D74F212BBD2109",
		},
		{"rNqui8PQUCiQ6Emyigsf3waEcKRQEvge4B",
			"ED588CB6EEB8FC7D430C7F4623202244DCE9571FCAF237368AC151DF4B233BD195",
			"44F96ECA6D4BDBEA2527855FBF07C202D9FA9BF5D1BCFD4A82454F9B1A2E1C81D0F387EF097BDC585FE6FADC28EB3265F06FCB9054263BAFFFC77961C4645D05",
		},
	}
	signerArray := make([]SignerInfo, 0, len(signList))
	for _, info := range signList {
		a, err := NewAccountFromAddress(info.account)
		require.NoError(t, err)
		pk, err := hex.DecodeString(info.pk)
		require.NoError(t, err)
		sg, err := hex.DecodeString(info.sg)
		require.NoError(t, err)
		signer := Signer{
			Account:       *a,
			SigningPubKey: new(PublicKey),
			TxnSignature:  new(VariableLength),
		}
		copy(signer.SigningPubKey.Bytes(), pk)
		*signer.TxnSignature = sg
		signerArray = append(signerArray, SignerInfo{signer})
	}
	sender, err := NewAccountFromAddress("rfKsmLP6sTfVGDvga6rW6XbmSFUzc3G9f3")
	require.NoError(t, err)
	destination, err := NewAccountFromAddress("rPTm7fS3QKgtYqSgD3RjkJ31EYB3a8bjL2")
	require.NoError(t, err)
	value, err := NewNativeValue(199711494)
	require.NoError(t, err)
	fmt.Printf("--f %s \n", strings.ToUpper(hex.EncodeToString(sender.Bytes())))
	fmt.Printf("--t %s \n", strings.ToUpper(hex.EncodeToString(destination.Bytes())))
	fee, err := NewNativeValue(4000)
	require.NoError(t, err)
	targetStr := "12000022800000002403ACFDBB2E90D48169201B049F386761400000000BE75B06684000000000000FA073008114453A8CB3100A157D6AE27E8F1F44D8FEB6C5B9D58314F662B3A88B96E1BAAA262DD242048BD0554A5033F3E0107321ED0D0C4028D416D565C45ABD675E56D1114DB79BD7B788608AE8F38B93636B118B7440E949A301D12BE2825D300BEED564EA95F4A8D1D341E33FDAD96C486B23DC712B6FCBAFC04899F2CD1B2EF857808ABA57DCA24640297860429ABF2E21D0FC0F07811463AA7F225CF8F049165D283AD1B3A09418D9F65DE1E0107321ED07CDB4A6FD3D67F16526AC1EF4DA1E9C28C780BF6771BE85FEC94F19F464CDCC7440E1E4F06FEDF78F8A9224AA0B3C04FFCA1EB59006B3F9799BBA0FF3AD5CEDE09EEC5D7D6173DE93D0A5820E7F066250B22C3E9A7EB4C66CFCF4D74F212BBD21098114822516E2D524184354DEBCED6F4D9D06B662A0D0E1E0107321ED588CB6EEB8FC7D430C7F4623202244DCE9571FCAF237368AC151DF4B233BD195744044F96ECA6D4BDBEA2527855FBF07C202D9FA9BF5D1BCFD4A82454F9B1A2E1C81D0F387EF097BDC585FE6FADC28EB3265F06FCB9054263BAFFFC77961C4645D05811497D9B9AA807973A5202C14D04634538FA25C1960E1F1"

	//构造payment交易
	payment := &Payment{
		Destination: *destination,
		Amount: Amount{
			Value: value,
		},
	}

	payment.TransactionType = PAYMENT
	last := uint32(77543527)
	payment.LastLedgerSequence = &last
	payment.Account = *sender
	payment.Sequence = 61668795
	payment.Fee = *fee
	tagNum := uint32(2429845865)
	payment.DestinationTag = &tagNum
	payment.Flags = new(TransactionFlag)
	*payment.Flags = *payment.Flags | TxCanonicalSignature
	if payment.GetBase().SigningPubKey == nil {
		payment.GetBase().SigningPubKey = new(PublicKey)
	}
	payment.GetBase().Signers = signerArray

	_, msg, err := Raw(payment)
	require.NoError(t, err)
	require.Equal(t, strings.ToUpper(hex.EncodeToString(msg)), targetStr)
}

// 验证签名前的数据是否正确
func TestMultiSign2(t *testing.T) {
	sender, err := NewAccountFromAddress("r9LqNeG6qHxjeUocjvVki2XR35weJ9mZgQ")
	require.NoError(t, err)
	destination, err := NewAccountFromAddress("rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh")
	require.NoError(t, err)
	siger, err := NewAccountFromAddress("rJZdUusLDtY9NEsGea7ijqhVrXv98rYBYN")
	require.NoError(t, err)
	value, err := NewNativeValue(1000)
	require.NoError(t, err)
	fee, err := NewNativeValue(10)
	require.NoError(t, err)
	targetStr := "534D5400120000228000000024000000016140000000000003E868400000000000000A730081145B812C9D57731E27A2DA8B1830195F88EF32A3B68314B5F762798A53D543A014CAF8B297CFF8F2F937E8C0A5ABEF242802EFED4B041E8F2D4A8CC86AE3D1"

	//构造payment交易
	payment := &Payment{
		Destination: *destination,
		Amount: Amount{
			Value: value,
		},
	}
	payment.TransactionType = PAYMENT
	payment.Account = *sender
	payment.Sequence = 1
	payment.Fee = *fee

	payment.Flags = new(TransactionFlag)
	*payment.Flags = *payment.Flags | TxCanonicalSignature
	if payment.GetBase().SigningPubKey == nil {
		payment.GetBase().SigningPubKey = new(PublicKey)
	}

	_, msg, err := SigningHash(payment)
	require.NoError(t, err)

	sigDigest := append(HP_TRANSACTION_MultiSig.Bytes(), msg...)
	sigDigest = append(sigDigest, siger.Bytes()...)

	require.Equal(t, strings.ToUpper(hex.EncodeToString(sigDigest)), targetStr)
}
