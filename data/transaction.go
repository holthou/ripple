package data

import "encoding/json"

type TxBase struct {
	TransactionType    TransactionType
	Flags              *TransactionFlag `json:",omitempty"`
	SourceTag          *uint32          `json:",omitempty"`
	Account            Account
	Sequence           uint32
	Fee                Value
	AccountTxnID       *Hash256        `json:",omitempty"`
	SigningPubKey      *PublicKey      `json:",omitempty"`
	TxnSignature       *VariableLength `json:",omitempty"`
	Signers            []SignerInfo    `json:",omitempty"`
	Memos              Memos           `json:",omitempty"`
	PreviousTxnID      *Hash256        `json:",omitempty"`
	LastLedgerSequence *uint32         `json:",omitempty"`
	Hash               Hash256         `json:"hash"`
	OperationLimit     *uint32         `json:",omitempty"`
	NetworkID          *uint32         `json:",omitempty"`
}

type SignerInfo struct {
	Signer Signer `json:",omitempty"`
}

type Signer struct {
	Account       Account
	TxnSignature  *VariableLength
	SigningPubKey *PublicKey
}

type Payment struct {
	TxBase
	Destination    Account
	Amount         Amount
	SendMax        *Amount  `json:",omitempty"`
	DeliverMin     *Amount  `json:",omitempty"`
	Paths          *PathSet `json:",omitempty"`
	DestinationTag *uint32  `json:",omitempty"`
	InvoiceID      *Hash256 `json:",omitempty"`
	TicketSequence *uint32  `json:",omitempty"`
}

type AccountSet struct {
	TxBase
	EmailHash      *Hash128        `json:",omitempty"`
	WalletLocator  *Hash256        `json:",omitempty"`
	WalletSize     *uint32         `json:",omitempty"`
	MessageKey     *VariableLength `json:",omitempty"`
	Domain         *VariableLength `json:",omitempty"`
	TransferRate   *uint32         `json:",omitempty"`
	TickSize       *uint8          `json:",omitempty"`
	SetFlag        *uint32         `json:",omitempty"`
	ClearFlag      *uint32         `json:",omitempty"`
	TicketSequence *uint32         `json:",omitempty"`
}

type AccountDelete struct {
	TxBase
	Destination    Account
	DestinationTag *uint32 `json:",omitempty"`
	TicketSequence *uint32 `json:",omitempty"`
}

type SetRegularKey struct {
	TxBase
	RegularKey     *RegularKey `json:",omitempty"`
	TicketSequence *uint32     `json:",omitempty"`
}

type OfferCreate struct {
	TxBase
	OfferSequence  *uint32 `json:",omitempty"`
	TakerPays      Amount
	TakerGets      Amount
	Expiration     *uint32 `json:",omitempty"`
	TicketSequence *uint32 `json:",omitempty"`
}

type OfferCancel struct {
	TxBase
	OfferSequence  uint32
	TicketSequence *uint32 `json:",omitempty"`
}

type TrustSet struct {
	TxBase
	LimitAmount    Amount
	QualityIn      *uint32 `json:",omitempty"`
	QualityOut     *uint32 `json:",omitempty"`
	TicketSequence *uint32 `json:",omitempty"`
}

type SetFee struct {
	TxBase
	BaseFee           Uint64Hex
	ReferenceFeeUnits uint32
	ReserveBase       uint32
	ReserveIncrement  uint32
}

type Amendment struct {
	TxBase
	Amendment Hash256
}

type EscrowCreate struct {
	TxBase
	Destination    Account
	Amount         Amount
	Digest         *Hash256 `json:",omitempty"`
	CancelAfter    *uint32  `json:",omitempty"`
	FinishAfter    *uint32  `json:",omitempty"`
	DestinationTag *uint32  `json:",omitempty"`
	TicketSequence *uint32  `json:",omitempty"`
}

type EscrowFinish struct {
	TxBase
	Owner          Account
	OfferSequence  uint32
	Method         *uint8   `json:",omitempty"`
	Digest         *Hash256 `json:",omitempty"`
	Proof          *Hash256 `json:",omitempty"`
	TicketSequence *uint32  `json:",omitempty"`
}

type EscrowCancel struct {
	TxBase
	Owner          Account
	OfferSequence  uint32
	TicketSequence *uint32 `json:",omitempty"`
}

type PaymentChannelCreate struct {
	TxBase
	Amount         Amount
	Destination    Account
	SettleDelay    uint32
	PublicKey      PublicKey
	CancelAfter    *uint32 `json:",omitempty"`
	DestinationTag *uint32 `json:",omitempty"`
	SourceTag      *uint32 `json:",omitempty"`
	TicketSequence *uint32 `json:",omitempty"`
}

type PaymentChannelFund struct {
	TxBase
	Channel        Hash256
	Amount         Amount
	Expiration     *uint32 `json:",omitempty"`
	TicketSequence *uint32 `json:",omitempty"`
}

type PaymentChannelClaim struct {
	TxBase
	Channel        Hash256
	Balance        *Amount         `json:",omitempty"`
	Amount         *Amount         `json:",omitempty"`
	Signature      *VariableLength `json:",omitempty"`
	PublicKey      *PublicKey      `json:",omitempty"`
	TicketSequence *uint32         `json:",omitempty"`
}

// CheckCreate, CheckCash, CheckCancel enabled by amendment 157D2D480E006395B76F948E3E07A45A05FE10230D88A7993C71F97AE4B1F2D1

// https://ripple.com/build/transactions/#checkcreate
type CheckCreate struct {
	TxBase
	Destination    Account
	SendMax        Amount
	DestinationTag *uint32  `json:",omitempty"`
	Expiration     *uint32  `json:",omitempty"`
	InvoiceID      *Hash256 `json:",omitempty"`
	TicketSequence *uint32  `json:",omitempty"`
}

// https://ripple.com/build/transactions/#checkcash
// Must include one of Amount or DeliverMin
type CheckCash struct {
	TxBase
	CheckID        Hash256
	Amount         *Amount `json:",omitempty"`
	DeliverMin     *Amount `json:",omitempty"`
	TicketSequence *uint32 `json:",omitempty"`
}

// https://ripple.com/build/transactions/#checkcancel
type CheckCancel struct {
	TxBase
	CheckID        Hash256
	TicketSequence *uint32 `json:",omitempty"`
}

type TicketCreate struct {
	TxBase
	TicketCount    *uint32 `json:",omitempty"`
	TicketSequence *uint32 `json:",omitempty"`
}

type SignerListSet struct {
	TxBase
	SignerQuorum   uint32            `json:",omitempty"`
	SignerEntries  []SignerEntryInfo `json:",omitempty"`
	TicketSequence *uint32           `json:",omitempty"`
}

type UNLModify struct {
	TxBase
	UNLModifyDisabling uint8           `json:",omitempty"`
	UNLModifyValidator *VariableLength `json:",omitempty"`
}

type SetDepositPreAuth struct {
	TxBase
	Authorize      *Account `json:",omitempty"`
	Unauthorize    *Account `json:",omitempty"`
	TicketSequence *uint32  `json:",omitempty"`
}

type NFTokenMint struct {
	TxBase
	NFTokenTaxon   *uint32         `json:",omitempty"`
	TransferFee    *uint16         `json:",omitempty"`
	Issuer         *Account        `json:",omitempty"`
	URI            *VariableLength `json:",omitempty"`
	TicketSequence *uint32         `json:",omitempty"`
}

type NFTokenBurn struct {
	TxBase
	Owner          *Account `json:",omitempty"`
	TicketSequence *uint32  `json:",omitempty"`
}

type NFTokenCreateOffer struct {
	TxBase
	NFTokenID      *Hash256 `json:",omitempty"`
	Amount         *Amount  `json:",omitempty"`
	Destination    *Account `json:",omitempty"`
	Owner          *Account `json:",omitempty"`
	Expiration     *uint32  `json:",omitempty"`
	TicketSequence *uint32  `json:",omitempty"`
}

type NFTCancelOffer struct {
	TxBase
	NFTokenOffers  *Vector256 `json:",omitempty"`
	TicketSequence *uint32    `json:",omitempty"`
}

type NFTAcceptOffer struct {
	TxBase
	NFTokenBuyOffer  *Hash256 `json:",omitempty"`
	NFTokenSellOffer *Hash256 `json:",omitempty"`
	NFTokenBrokerFee *Amount  `json:",omitempty"`
	TicketSequence   *uint32  `json:",omitempty"`
}

type ImportTransaction struct {
	TxBase
	Blob *VariableLength `json:",omitempty"`
}

type InvokeTransaction struct {
	TxBase
	Destination Account
	Blob        *VariableLength `json:",omitempty"`
	InvoiceID   *Hash256        `json:",omitempty"`
}

// NicknameSet is deprecated (type 6).
type NicknameSet struct {
	TxBase
}

// Contract is deprecated (type 9).
type Contract struct {
	TxBase
}

// SpinalTap is deprecated (type 11).
type SpinalTap struct {
	TxBase
}

// Issue represents a currency+issuer pair (no amount value), used for ISSUE-typed fields.
type Issue struct {
	Currency Currency `json:"currency"`
	Issuer   *Account `json:"issuer,omitempty"`
}

// SetHook supporting types.

type HookParam struct {
	HookParameterName  *VariableLength `json:",omitempty"`
	HookParameterValue *VariableLength `json:",omitempty"`
}

type HookParamInfo struct {
	HookParameter HookParam `json:",omitempty"`
}

type HookGrant struct {
	HookHash  *Hash256 `json:",omitempty"`
	Authorize *Account `json:",omitempty"`
}

type HookGrantInfo struct {
	HookGrant HookGrant `json:",omitempty"`
}

type HookEntry struct {
	CreateCode     *VariableLength `json:",omitempty"`
	HookHash       *Hash256        `json:",omitempty"`
	HookOn         *Hash256        `json:",omitempty"` // UINT256 bitmask
	HookApiVersion *uint16         `json:",omitempty"`
	HookNamespace  *Hash256        `json:",omitempty"`
	HookParameters []HookParamInfo `json:",omitempty"`
	HookGrants     []HookGrantInfo `json:",omitempty"`
	Flags          *uint32         `json:",omitempty"`
}

type HookEntryInfo struct {
	Hook HookEntry `json:",omitempty"`
}

type SetHook struct {
	TxBase
	Hooks []HookEntryInfo `json:",omitempty"`
}

// URIToken transactions.

type URITokenMint struct {
	TxBase
	URI         *VariableLength `json:",omitempty"`
	Digest      *Hash256        `json:",omitempty"`
	Amount      *Amount         `json:",omitempty"`
	Destination *Account        `json:",omitempty"`
}

type URITokenBurn struct {
	TxBase
	URITokenID Hash256
}

type URITokenBuy struct {
	TxBase
	URITokenID Hash256
	Amount     Amount
}

type URITokenCreateSellOffer struct {
	TxBase
	URITokenID  Hash256
	Amount      Amount
	Destination *Account `json:",omitempty"`
}

type URITokenCancelSellOffer struct {
	TxBase
	URITokenID Hash256
}

// GenesisMint transaction supporting types.

type GenesisMintEntry struct {
	Destination Account
	Amount      Amount
}

type GenesisMintInfo struct {
	GenesisMint GenesisMintEntry `json:",omitempty"`
}

type GenesisMintTx struct {
	TxBase
	GenesisMints []GenesisMintInfo `json:",omitempty"`
}

// ClaimReward resets accumulators and claims a reward from a hook.
type ClaimReward struct {
	TxBase
	Issuer        *Account `json:",omitempty"`
	ClaimCurrency *Issue   `json:",omitempty"`
}

// EmitFailure is a pseudo-transaction recording a hook emit failure.
type EmitFailure struct {
	TxBase
	LedgerSequence  uint32
	TransactionHash Hash256
}

// UNLReport is a pseudo-transaction updating the negative UNL report.
type UNLReport struct {
	TxBase
	LedgerSequence  uint32
	ActiveValidator json.RawMessage `json:",omitempty"`
	ImportVLKey     json.RawMessage `json:",omitempty"`
}

func (t *TxBase) GetBase() *TxBase                    { return t }
func (t *TxBase) GetType() string                     { return txNames[t.TransactionType] }
func (t *TxBase) GetTransactionType() TransactionType { return t.TransactionType }
func (t *TxBase) Prefix() HashPrefix                  { return HP_TRANSACTION_ID }
func (t *TxBase) GetPublicKey() *PublicKey            { return t.SigningPubKey }
func (t *TxBase) GetSignature() *VariableLength       { return t.TxnSignature }
func (t *TxBase) SigningPrefix() HashPrefix           { return HP_TRANSACTION_SIGN }
func (t *TxBase) PathSet() PathSet                    { return PathSet(nil) }
func (t *TxBase) GetHash() *Hash256                   { return &t.Hash }

func (t *TxBase) Compare(other *TxBase) int {
	switch {
	case t.Account.Equals(other.Account):
		switch {
		case t.Sequence == other.Sequence:
			return t.GetHash().Compare(*other.GetHash())
		case t.Sequence < other.Sequence:
			return -1
		default:
			return 1
		}
	case t.Account.Less(other.Account):
		return -1
	default:
		return 1
	}
}

func (t *TxBase) InitialiseForSigning() {
	if t.SigningPubKey == nil {
		t.SigningPubKey = new(PublicKey)
	}
	if t.TxnSignature == nil {
		t.TxnSignature = new(VariableLength)
	}
}

func (o *OfferCreate) Ratio() *Value {
	return o.TakerPays.Ratio(o.TakerGets)
}

func (p *Payment) PathSet() PathSet {
	if p.Paths == nil {
		return PathSet(nil)
	}
	return *p.Paths
}
