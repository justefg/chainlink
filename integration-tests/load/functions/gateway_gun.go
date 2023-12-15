package loadfunctions

import (
	"crypto/ecdsa"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/tdh2/go/tdh2/tdh2easy"
	"github.com/smartcontractkit/wasp"

	"github.com/smartcontractkit/chainlink-testing-framework/networks"

	tc "github.com/smartcontractkit/chainlink/integration-tests/testconfig"
)

/* SingleFunctionCallGun is a gun that constantly requests randomness for one feed  */

type GatewaySecretsSetGun struct {
	Cfg                *tc.TestConfig
	Resty              *resty.Client
	SlotID             uint
	Method             string
	EthereumPrivateKey *ecdsa.PrivateKey
	ThresholdPublicKey *tdh2easy.PublicKey
	DONPublicKey       []byte
}

func NewGatewaySecretsSetGun(cfg *tc.TestConfig, method string, pKey *ecdsa.PrivateKey, tdh2PubKey *tdh2easy.PublicKey, donPubKey []byte) *GatewaySecretsSetGun {
	return &GatewaySecretsSetGun{
		Cfg:                cfg,
		Resty:              resty.New(),
		Method:             method,
		EthereumPrivateKey: pKey,
		ThresholdPublicKey: tdh2PubKey,
		DONPublicKey:       donPubKey,
	}
}

func callSecretsSet(m *GatewaySecretsSetGun) *wasp.CallResult {
	randNum := strconv.Itoa(rand.Intn(100000))
	randSlot := uint(rand.Intn(5))
	version := uint64(time.Now().UnixNano())
	expiration := int64(60 * 60 * 1000)
	secret := fmt.Sprintf("{\"ltsecret\": \"%s\"}", randNum)
	log.Debug().
		Uint("SlotID", randSlot).
		Str("MessageID", randNum).
		Uint64("Version", version).
		Int64("Expiration", expiration).
		Str("Secret", secret).
		Msg("Sending S4 envelope")
	secrets, err := EncryptS4Secrets(
		m.EthereumPrivateKey,
		m.ThresholdPublicKey,
		m.DONPublicKey,
		secret,
	)
	if err != nil {
		return &wasp.CallResult{Error: err.Error(), Failed: true}
	}
	network := networks.MustGetSelectedNetworkConfig(m.Cfg.Network)[0].Name
	if len(m.Cfg.Network.WalletKeys[network]) < 1 {
		panic(fmt.Sprintf("no wallet keys found for %s", network))
	}

	cfg := m.Cfg.Functions
	_, _, err = UploadS4Secrets(m.Resty, &S4SecretsCfg{
		GatewayURL:            *cfg.Common.GatewayURL,
		PrivateKey:            m.Cfg.Network.WalletKeys[network][0],
		MessageID:             randNum,
		Method:                "secrets_set",
		DonID:                 *cfg.Common.DONID,
		S4SetSlotID:           randSlot,
		S4SetVersion:          version,
		S4SetExpirationPeriod: expiration,
		S4SetPayload:          secrets,
	})
	if err != nil {
		return &wasp.CallResult{Error: err.Error(), Failed: true}
	}
	return &wasp.CallResult{}
}

func callSecretsList(m *GatewaySecretsSetGun) *wasp.CallResult {
	randNum := strconv.Itoa(rand.Intn(100000))
	randSlot := uint(rand.Intn(5))
	version := uint64(time.Now().UnixNano())
	expiration := int64(60 * 60 * 1000)
	network := networks.MustGetSelectedNetworkConfig(m.Cfg.Network)[0].Name
	if len(m.Cfg.Network.WalletKeys[network]) < 1 {
		panic(fmt.Sprintf("no wallet keys found for %s", network))
	}
	cfg := m.Cfg.Functions
	if err := ListS4Secrets(m.Resty, &S4SecretsCfg{
		GatewayURL:            *cfg.Common.GatewayURL,
		RecieverAddr:          *cfg.Common.Receiver,
		PrivateKey:            m.Cfg.Network.WalletKeys[network][0],
		MessageID:             randNum,
		Method:                m.Method,
		DonID:                 *cfg.Common.DONID,
		S4SetSlotID:           randSlot,
		S4SetVersion:          version,
		S4SetExpirationPeriod: expiration,
	}); err != nil {
		return &wasp.CallResult{Error: err.Error(), Failed: true}
	}
	return &wasp.CallResult{}
}

// Call implements example gun call, assertions on response bodies should be done here
func (m *GatewaySecretsSetGun) Call(_ *wasp.Generator) *wasp.CallResult {
	var res *wasp.CallResult
	switch m.Method {
	case "secrets_set":
		res = callSecretsSet(m)
	case "secrets_list":
		res = callSecretsList(m)
	default:
		panic("gateway gun must use either 'secrets_set' or 'list' methods")
	}
	return res
}
