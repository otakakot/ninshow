package model

type Cert struct {
	Kid string // Kid 鍵識別子
	Kty string // Kty RSAやEC等の暗号アルゴリズファミリー
	Use string // Use 公開鍵の用途
	Alg string // Alg 署名検証アルゴリズム
	N   string // N modulus 公開鍵を復元するための公開鍵の絶対値
	E   string // E exponent 公開鍵を復元するための指数値
}
