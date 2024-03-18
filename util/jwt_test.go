package util

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	_ "embed"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"testing"
)

//go:embed test-public.pem
var keyTest []byte

func TestRsa(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJkNzg3MGJjNi0yMDY3LTQ3MjAtYWNmNC04MjRhZTIzMWFiZDAifQ.eyJleHAiOjE3MTAxNTQ5MzEsImlhdCI6MTcxMDE0NzczMSwianRpIjoiMDkyZWNiNWEtZDk3OS00ZTYwLWEyMTMtYjk5NDFjYmU2OTZkIiwiaXNzIjoiaHR0cDovLzE5Mi4xNjguNC4yNjo4MDgwL2F1dGgvcmVhbG1zL2ZhYmFvIiwiYXVkIjpbInJlYWxtLW1hbmFnZW1lbnQiLCJhY2NvdW50Il0sInN1YiI6IjM0N2M5ZTllLTA3NmMtNDVlMy1iZTc0LWM0ODJmZmZjYzZlNSIsInR5cCI6IkJlYXJlciIsImF6cCI6InBrdWxhdyIsInNlc3Npb25fc3RhdGUiOiIwNjViMGUzNy0xOTQzLTQ1Y2UtYWIwNi00NmY0ZjNlOTVjZTMiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIioiXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIue7hOeuoeeQhuWRmCIsImxhd2Zpcm0iLCJmZWRlcmF0ZWQiLCJjYXJzaSJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJlYWxtLW1hbmFnZW1lbnQiOnsicm9sZXMiOlsidmlldy1yZWFsbSIsInZpZXctaWRlbnRpdHktcHJvdmlkZXJzIiwibWFuYWdlLWlkZW50aXR5LXByb3ZpZGVycyIsImltcGVyc29uYXRpb24iLCJyZWFsbS1hZG1pbiIsImNyZWF0ZS1jbGllbnQiLCJtYW5hZ2UtdXNlcnMiLCJxdWVyeS1yZWFsbXMiLCJ2aWV3LWF1dGhvcml6YXRpb24iLCJxdWVyeS1jbGllbnRzIiwicXVlcnktdXNlcnMiLCJtYW5hZ2UtZXZlbnRzIiwibWFuYWdlLXJlYWxtIiwidmlldy1ldmVudHMiLCJ2aWV3LXVzZXJzIiwidmlldy1jbGllbnRzIiwibWFuYWdlLWF1dGhvcml6YXRpb24iLCJtYW5hZ2UtY2xpZW50cyIsInF1ZXJ5LWdyb3VwcyJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsInZpZXctYXBwbGljYXRpb25zIiwidmlldy1jb25zZW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJtYW5hZ2UtY29uc2VudCIsImRlbGV0ZS1hY2NvdW50Iiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgYXV0aG9yaXphdGlvbiByb2xlcyBwa3VsYXctZXh0ZW5zaW9ucyBlbWFpbCB2NiBwcm9maWxlIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJsb2dpblR5cGUiOiJwYXNzd29yZCIsImdyb3VwVXNlck5hbWUiOiJ0ZXN0IiwiaXBBZGRyZXNzIjoiMTkyLjE2OC4wMDQuMTc2fjE5Mi4xNjguMDA0LjE3NiIsImV4dGVuc2lvbl9yb2xlcyI6eyJ3ZWl4aW4iOlsibGF3ZmlybSJdfSwibG9naW5UaW1lUGVyaW9kIjp7InBhc3N3b3JkIjowLCJvbmVDbGlja0xvZ2luIjoxMzEsIndlaXhpbiI6MjgyLCJjYXJzaSI6MjU0fSwicHJlZmVycmVkX3VzZXJuYW1lIjoidGVzdCIsImdyb3VwVXNlcklkIjoiMzQ3YzllOWUtMDc2Yy00NWUzLWJlNzQtYzQ4MmZmZmNjNmU1IiwicGhvbmVOdW1iZXIiOiIxMzUyMTk3Mjk5MSIsIm5pY2tuYW1lIjoidGVzdCIsImlzRmVkZXJhbFVzZXIiOjAsInBob25lTnVtYmVyVmVyaWZpZWQiOiJ0cnVlIiwiaXNHcm91cFVzZXIiOiIxIiwiZW1haWwiOiJ6emwxMjNAc2luYS5jb20ifQ.vnzS0NVQcBeryZPgu9s6X24gofkrGZjRcLQaxna2doc"
	publicKey, err1 := jwt.ParseRSAPublicKeyFromPEM(keyTest)
	if err1 != nil {
		fmt.Println("Error ParseRSAPublicKeyFromPEM:", err1)
	}
	parts := strings.Split(token, ".")

	method := jwt.GetSigningMethod("HS256")

	err := method.Verify(strings.Join(parts[0:2], "."), parts[2], publicKey)
	valid := err == nil
	fmt.Println(valid)

}

func validateRSASignature(publicKeyPEM, signature, data []byte) (bool, error) {
	// Parse the public key
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return false, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %s", err)
	}

	rsaPublicKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("not an RSA public key")
	}

	// Verify the signature
	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return false, fmt.Errorf("signature verification failed: %s", err)
	}

	return true, nil
}
