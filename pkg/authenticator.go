package sbb_api

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

const NotSoSecretKey = "c3eAd3eC3a7845dE98f73942b3d5f9c0"
const SbbCaCertificate = "MIIFtDCCA5ygAwIBAgIJAIbo7SRrMr9cMA0GCSqGSIb3DQEBCwUAMGcxCzAJBgNVBAYTAkNIMQ0wCwYDVQQIDARCZXJuMQ0wCwYDVQQHDARCZXJuMRQwEgYDVQQKDAtTQkIgQ0ZGIEZTUzELMAkGA1UECwwCSVQxFzAVBgNVBAMMDiouc2JibW9iaWxlLmNoMB4XDTE1MTAwNjA3NTAwMVoXDTM1MTAwMTA3NTAwMVowZzELMAkGA1UEBhMCQ0gxDTALBgNVBAgMBEJlcm4xDTALBgNVBAcMBEJlcm4xFDASBgNVBAoMC1NCQiBDRkYgRlNTMQswCQYDVQQLDAJJVDEXMBUGA1UEAwwOKi5zYmJtb2JpbGUuY2gwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCuPN4NAtvcJ/QekrEBhH1kCuP9ld84VH1mbGowUVN66gQIJnviZBOTsAZHGh1aZCnUXa0d0Nj0Srhp7s/+ITiPGOCpHBjjsQ3dkr6qGhy1kNMAv7B/vfAMiknMkVH0X7pKGJTTsF5HIpqJUhYKqX/zpOO1YKiSrMmmaAKcim0iKi/sg8nzQeVHcgx2E4hOMb4oAcOULVJUBmT29/BN678XzJ/eIBUYbJ0rYFRl70n06ADUdoaRRKCtSCjlayTUOhKgo+2unn2R0fl5PxxhZnLaHEIypveQx0mWyqgO++mjMUkLme436bEgLUNQs/oHqPE/PzXtsgltOKM+uyLhWXJNRDkpRhQ9TzrTCiEODIw1dhCnWJKScFbLcuXAokqfO1tUGgSd1dBb6ayTuRa5lb3zjTS3hKemG+knujxsmZfdKRjs61fO1v7H5ZyVRMochZugcaMF/9ThpIIVFwj9YIkzMwUNVnpEZeG6Tntcxp7otpATUGUoslczOvtrISNBNNG9uqq9bNB5BcNCvc3dc4VmRpH0iy77PMrF1zIde4z4pVeXc3Nidd1JfQoAYUy2AwnpRdnuL4QTlIJLCGBAltQbqkTZN6heWm1sJQzzy1BhN0NqPK6DNFvxhi6fR1tGuozNndtSHZRTIAusBxXROXB/f7FKqp/xjjm2OGRQrEB2twIDAQABo2MwYTAdBgNVHQ4EFgQUSwQViuGuvQceYK8EQGRELpt0urYwHwYDVR0jBBgwFoAUSwQViuGuvQceYK8EQGRELpt0urYwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAYYwDQYJKoZIhvcNAQELBQADggIBAFhTBtmilFC1gFeIDNniJiRJ3Pwz+9dFoMqiQkxnXsJPZkBXoanYDCHcFU/EfScz6dmIBfkljc/M5GiMJ9MfO+XHuMhYRk63ThxVz8WaQnRo2w5oxocaApV3yzYwpYH6+DEkW+Z0JQJUB9oSN+Xfv13Bgxmlc2WywIx56B1gyAWTx69uI3nNOfwFMxx9ngQcOYz2RYkmZDA1Rimn9QWASBX+32XJA1Kkher2gzHQoZ/29DUFy+mc5A68VoIegirtRJ115QQMAe77vLqvvVhsw+f7bNEk+Z0myBzUeEtS5hSdfnrbpSYq5aOxEuLKBjit2ZGatMbwOZ4EWc74QH/qSkxQQoYshGKBoN9wEp3EDRBOQuF3KiNgsEgymJHuAMucV8DuNxvzhtP30NeyuM6cirXM/97Pv1z0yrKFoLHq+ZLhSdX+yybMTw3Qqo5PYreBvOGsydg4Ocsz4NF2/zBa18NZAtpIk6uYo1LrJChuq2U8i60/vcmQMIFkj6/eHTGcqktaj29eKAwttKWYzJNDTHPskR4Yrsm8RRs2LCz27NNQIAQotSdFZWF7JMv/8n7+SHQDFKxRmbDJ4oMdIkq6D6/eva82j7qDSmoJXNA+zciRsY8cv5kGtsNEXLJuaTUHVjAg+dbvLvbAVrZqdAhxEK2gl1WbrUC1NnMXQvM38uZC"

func GetAuthorization(path string, date time.Time) string {
	hmacFunc := hmac.New(sha1.New, []byte(getKey()))
	hmacFunc.Write([]byte(fmt.Sprintf("%s%s", path, date.Format("2006-01-02"))))
	result := hmacFunc.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}

func getKey() string {
	certHash, err := getCertificateHash()
	if err != nil {
		return ""
	}
	return retrieveKey(certHash, NotSoSecretKey)
}

func retrieveKey(certHash string, key string) string {
	clearText := fmt.Sprintf("%s%s", certHash, key)
	h := sha256.New()
	h.Write([]byte(clearText))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func GetCACert() (*x509.Certificate, error) {
	certBin, err := base64.StdEncoding.DecodeString(SbbCaCertificate)
	if err != nil {
		return nil, err
	}
	return x509.ParseCertificate(certBin)
}

func getCertificateHash() (string, error) {
	certBin, err := base64.StdEncoding.DecodeString(SbbCaCertificate)
	if err != nil {
		return "", errors.New("unable to decode certificate from B64 string")
	}

	h := sha1.New()
	h.Write(certBin)
	hash := h.Sum(nil)

	return base64.StdEncoding.EncodeToString(hash), nil
}
