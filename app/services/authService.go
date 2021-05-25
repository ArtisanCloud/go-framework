package service

import (
	"crypto/rsa"
	"fmt"
	. "github.com/ArtisanCloud/go-framework/app/models"
	. "github.com/ArtisanCloud/go-framework/config"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

type AuthService struct {
	User    *User
}

var (
	StrPublicKeyPath  string
	StrPrivateKeyPath string
)

const InExpiredMonths = 3
const InExpiredDays = InExpiredMonths * 30
const InExpiredSecond = InExpiredDays * 24 * 60 * 60

/**
 ** 初始化构造函数
 */

// 模块初始化函数 import 包时被调用
func init() {

}

func SetupSSHKeyPath(ssh *SSHConfig) {
	StrPublicKeyPath = ssh.PublicKeyFile
	StrPrivateKeyPath = ssh.PrivateKeyFile
}

func NewAuthService() (r *AuthService) {
	r = &AuthService{
		User:    NewUser(),
	}

	return r
}


func (srv *AuthService) CreateToken(user *User) (string, bool) {

	signKey, err := ioutil.ReadFile(StrPrivateKeyPath)
	//signKey, err := ioutil.ReadFile(strPublicKeyPath)
	if err != nil {
		fmt.Println("Error reading private key %x", err)
		return "", false
	}

	var key *rsa.PrivateKey
	key, err = jwt.ParseRSAPrivateKeyFromPEM(signKey)
	//var key *rsa.PublicKey
	//key, err = jwt.ParseRSAPublicKeyFromPEM(signKey)

	//t := jwt.New(jwt.GetSigningMethod("RS256"))
	t := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["AccessToken"] = "bar"
	claims["UserUUID"] = user.UUID
	claims["exp"] = time.Now().Add(InExpiredDays).Unix()
	t.Claims = claims

	tokenString, err := t.SignedString(key)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("tokenString: " + tokenString)
	//var token *jwt.Token
	_, err = ParseTokenFromSignedTokenString(tokenString)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	return tokenString, true

}

func ParseTokenFromSignedTokenString(tokenString string) (*jwt.Token, error) {
	publicKey, err := ioutil.ReadFile(StrPublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("error reading public key file: %v\n", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, fmt.Errorf("error parsing RSA public key: %v\n", err)
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	return parsedToken, nil
}


func ParseAuthorization(authHeader string) (claims jwt.Claims, err error) {

	const BEARER_SCHEMA = "Bearer "
	tokenString := authHeader[len(BEARER_SCHEMA):]
	//fmt.Printf("tokenstring:%v\r\n", tokenString)

	token, err := jwt.Parse(tokenString, nil)
	if err != nil{
		log.Println("Problem with parsing", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok != true{
		log.Println("Problem with claims", ok)
	}

	return claims, nil
}

