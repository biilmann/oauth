package oauth

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// Test example from https://dev.twitter.com/oauth/overview/creating-signatures
func TestSignature(t *testing.T) {
	method := "POST"
	body := strings.NewReader("status=Hello%20Ladies%20%2b%20Gentlemen%2c%20a%20signed%20OAuth%20request%21")
	uri := "https://api.twitter.com/1/statuses/update.json?include_entities=true"

	in, _ := http.NewRequest(method, uri, body)
	in.Header.Set("Authorization", `OAuth `+
		`oauth_consumer_key="xvz1evFS4wEEPTGEFPHBog",`+
		`oauth_nonce="kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg",`+
		`oauth_signature_method="HMAC-SHA1",`+
		`oauth_timestamp="1318622958",`+
		`oauth_token="370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb",`+
		`oauth_version="1.0"`)

	out := "tnnArxj06cWHq44gCs1OSKk%2FjLY%3D"

	consumer := &Consumer{
		"xvz1evFS4wEEPTGEFPHBog",
		"kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw",
	}
	token := &Token{
		"370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb",
		"LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE",
	}

	if a := consumer.Signature(in, token); a != out {
		t.Errorf("Signature() should be \n%v\n not %v", out, a)
	}
}

// Test example from RFC 5849
func TestBaseUri(t *testing.T) {
	in, _ := url.Parse("http://example.com/r%20v/X?id=123")
	out := "http://example.com/r%20v/X"

	if a := baseUri(in); a != out {
		t.Errorf("baseUri(%v) should be %v not %v", in, out, a)
	}
}

// Test example from https://dev.twitter.com/oauth/overview/percent-encoding-parameters
func TestEncode(t *testing.T) {
	tests := map[string]string{
		`Ladies + Gentlemen`: `Ladies%20%2B%20Gentlemen`,
		`An encoded string!`: `An%20encoded%20string%21`,
		`Dogs, Cats & Mice`:  `Dogs%2C%20Cats%20%26%20Mice`,
		`☃`:                  `%E2%98%83`,
	}

	for in, out := range tests {
		if a := encode(in); a != out {
			t.Errorf("encode(%v) should be %v not %v", in, out, a)
		}
	}
}

// Test example from RFC 5849
func TestRequestParameters(t *testing.T) {
	method := "POST"
	body := strings.NewReader("c2&a3=2+q")
	uri := "https://example.com/request?b5=%3D%253D&a3=a&c%40=&a2=r%20b"

	in, _ := http.NewRequest(method, uri, body)
	in.Header.Set("Authorization", `OAuth realm="Example",`+
		`oauth_consumer_key="9djdj82h48djs9d2",`+
		`oauth_token="kkk9d7dh3k39sjv7",`+
		`oauth_signature_method="HMAC-SHA1",`+
		`oauth_timestamp="137131201",`+
		`oauth_nonce="7d8f3e4a",`+
		`oauth_signature="djosJKDKJSD8743243%2Fjdk33klY%3D"`)

	out := `a2=r%20b&a3=2%20q&a3=a&b5=%3D%253D&c%40=&c2=&oauth_consumer_key=` +
		`9djdj82h48djs9d2&oauth_nonce=7d8f3e4a&oauth_signature_method=HMAC-S` +
		`HA1&oauth_timestamp=137131201&oauth_token=kkk9d7dh3k39sjv7`

	if a := requestParameters(in); a != out {
		t.Errorf("requestParameters() should be \n%v\n not \n%v\n", out, a)
	}
}
