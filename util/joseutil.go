package util

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/go-jose/go-jose/v3"
	"time"
)

func Encrypt(data []byte) (encryptStr string, err error) {
	crypt, err := jose.NewEncrypter(jose.A128CBC_HS256, rcpt, nil)
	if err != nil {
		return "", err
	}
	obj, errEncrypt := crypt.Encrypt(data)
	if errEncrypt != nil {
		return "", errEncrypt
	}
	encryptStr, errorSerialize := obj.CompactSerialize()
	if errorSerialize != nil {
		return "", errorSerialize
	}
	return encryptStr, err

}

func Decrypt(str string) (decryptStr string, err error) {
	jwe, errParse := jose.ParseEncrypted(str)
	if errParse != nil {
		return "", errParse
	}
	decryptedKey, errDecrypt := jwe.Decrypt("password")
	if errDecrypt != nil {
		return "", errDecrypt
	}
	decryptStr = string(decryptedKey)
	return decryptStr, nil
}

var rcpt jose.Recipient

func init() {
	rcpt = jose.Recipient{
		Algorithm:  jose.PBES2_HS256_A128KW,
		Key:        "password",
		PBES2Count: 4096,
		PBES2Salt:  []byte("salt"),
	}
}

const (
	serverPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyOCNCy8x280vYwpVl26ToWPXkTFMdkSkLQ+UB6uCK5/1Ojw1+ZqKWhAZg4bNTG03QtG7xJOPJUH0S2Gf2CNbZjm2OEe1ZeY08JCFThZ4h5D6iuAOarmYjlSnzhdW8jpiSkm0Okci+xZkC1dWW2twjGgc0/E6o9JXYEggfrDWbllpGQOib+lx+mU51e0OMqu8TwSrLoa/d0a2dawPJnU3l1a6lyvhD7W/asUphMNeLKgmNf4wf92wElLlDkB2Qozfg1J/6l4JzBbgVLLljO0j7VKv7seyAnul55KmRd9tUjpdecf0FQPFZiNOTW+g4ivMNBQZvKAnsfgDboxy1PzW1QIDAQAB"
	payload         = `{"name":"pass1","value":"secret1","username":"user1"}`
)

func Rs256Enc() {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(serverPublicKey)
	if err != nil {
		panic(err)
	}

	pub, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		panic(err)
	}

	opts := new(jose.EncrypterOptions)
	opts.WithHeader("iat", time.Now().Unix())

	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.RSA_OAEP_256, Key: pub}, opts)
	if err != nil {
		panic(err)
	}

	jwe, err := encrypter.Encrypt([]byte(payload))
	if err != nil {
		panic(err)
	}

	jweCompact, err := jwe.CompactSerialize()
	if err != nil {
		panic(err)
	}

	fmt.Println(jweCompact)
}

const (
	clientPrivateKey = "MIIEowIBAAKCAQEAyOCNCy8x280vYwpVl26ToWPXkTFMdkSkLQ+UB6uCK5/1Ojw1+ZqKWhAZg4bNTG03QtG7xJOPJUH0S2Gf2CNbZjm2OEe1ZeY08JCFThZ4h5D6iuAOarmYjlSnzhdW8jpiSkm0Okci+xZkC1dWW2twjGgc0/E6o9JXYEggfrDWbllpGQOib+lx+mU51e0OMqu8TwSrLoa/d0a2dawPJnU3l1a6lyvhD7W/asUphMNeLKgmNf4wf92wElLlDkB2Qozfg1J/6l4JzBbgVLLljO0j7VKv7seyAnul55KmRd9tUjpdecf0FQPFZiNOTW+g4ivMNBQZvKAnsfgDboxy1PzW1QIDAQABAoIBABtNAX5SpLkV0unn/qf9HE1j72lJRPYR8Co7osfEFIXvUHngswlAaqtA4ubQTFilw0vui1F1KJBMvXH9S6cpbwOetlhBzrKRrgI/8g4h2bG7D4IOX4c7wDMf6EqHO4biTneAOuFsx6FOcrxj21XDbWF3yOs5LtoS/VMVmmz9hbcra5hLq21IO6wAbWxKIoBxlqwsz9NaO0RF4Du9JHiuhfUE3F4ndwer5plgzILs/eY2ZMmI9dm78lE7c5dO3c6orhtaQZX7SB8ZclZzolGUlQIshrMavsmcgJdPm/hYmsdMHYsfZUCgnh3TRgXNdPm6AG7+zeYKmDEs5r/Xtp7BSsUCgYEA//pLDNFsB2EOkjWEhztGaVqy5s5zl0islaUB1O1/NLI4wKEfeEqX3uzHWj49FeLuKVql/JgboeAancfjTTh7B/J7RnXORHrrqhrPJbzIwXScE+1ISFi6pqx2HFPQJksTlMEn31nPrFxk2k1juvPhFF9NhU6C4A0eTEOCgaWOVbMCgYEAyOUHhDQDPOSyGQxnEUDp7wb1a7tifoka0Vsovb+WBnjac5+Kouj1ktp/qqqz5mrivjs0WeGrsXFBpzh/VJ//8V/pkDCtmfbKFPmnQVDZYZ3GP5WA2lkY2o/r9S5fFr7BPsFkqWGew7/eA/4bZWwuFkaOZLHuFGeSLiLkRw0d7VcCgYEAuhEf+KNJ2VtNki3gtcP6fJ53KJqtHyyuduBIopQ4QP/DjRIRpAjkuCB5EDnpgT1BTUIdIa0XeOVDs6kWqo0BsVcrEJ6VXuXS5AU0ygOeEjFHOpziS+RjWMBH4nNx/EPaaei1qP7JnEpU7PIQKlcYJwzXdq0JKP+gJGN2O/MxsU0CgYBc2x/CR5hHiR2Y3la1bt+yD/FYPmCkRBMVOaF3MxrRGwM99jxKWItJuZzQ8d39XuI3M35bHSgS+Fp0RHT2VI4kr0Rx1U4ooB1/3HUmh4wyxo5fzWalhhEq1OMnjKt/A+SmcRVqBSkxKXohNk7LbllPCoW+nU+CNu0mThJdlP6EOQKBgCSwCwlCDNMRyiFMYNTeXdAzO98sxvrpSFYV/8sx/jlZIvrC0QCcgahdxXBhOrII1S+xc6V/jOqBZot5OkCknLjIQiKuwLKYi4e81bzzpx0M2VjF53u/tk/jGijaaJxvun6AeK8cRjxXm0RlO58vaY99AiJvuASHetkQ5i1UOgw0"
	jwe              = "eyJhbGciOiJSU0EtT0FFUC0yNTYiLCJlbmMiOiJBMjU2R0NNIiwiaWF0IjoxNjk5NjAyNjczfQ.X0AGND8pRD34OYVKowepzy_8bBRcZ2C4aqe1nWZwq1kRhjeCJ2XJL2EiF7GS0c194hXHGj1ZmmthG8giz8a7xnb43npdVH5NfCuNrLztbwedagnxUUohecMN2WVj4D_gujDHhZFDjNhem1uAyP3mbIAummaiacIMGwRISENP1-co56hc6NRGmh67nZNPdZC75K7f8v7pjUZBnook1bNQGy8gjuKSHCREDxGNStCml_mp2CshOiO8YBsPs9ckHQ5ndxwcFcotUnjDYL2QzBuCUkj8bCKOJWCinX_64Ycg29Jx2qzsD9nfK71MUj0oDiVqdkAtyI1QBrve6OrYULqGlw.X9e4fJRXgjhf2OZx.ETadp98mne6Q0FmWQnziiIoqGPp4dA4kmvM2PrDL5lWGNCk-FxQYAScJp38t9QoRgutcMkk.qquSaa5jb07tLPupzQ6c9w\n"
)

func Rs256Dec() {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(clientPrivateKey)

	if err != nil {
		panic(err)
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	encryptedJwe, err := jose.ParseEncrypted(jwe)
	if err != nil {
		panic(err)
	}

	decrypted, err := encryptedJwe.Decrypt(privateKey)

	fmt.Println("response: ", string(decrypted))
}
