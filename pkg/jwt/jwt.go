package jwt

import (
	"errors"
	"github.com/kataras/jwt"
	"time"
)

type Attribute string
type AttributeAction int

const (
	ARRIBUTE_ACTIONS_NO_PERMISSIONS AttributeAction = iota // 0
	ATTRIBUTE_ACTIONS_VIEW                                 // 1
	ATTRIBUTE_ACTIONS_EDIT                                 // 2
	ATTRIBUTE_ACTIONS_DELETE                               // 3
)

type JwtType int

const (
	JWT_TYPE_UNSPECIFIED JwtType = iota
	JWT_TYPE_ACCESS_TOKEN
	JWT_TYPE_RFRESH_TOKEN
)

type UserClaims struct {
	jwt.Claims
	OrgId      string                           `json:"orgId,required"`
	Attributes map[Attribute]map[Attribute]bool `json:"attributes"`
}

type RefreshClaims struct {
	jwt.Claims
	OrgId   string
	Version int
}

type JWT struct {
	Secret          string
	Algo            jwt.Alg
	Issuer          string
	AccessLiveness  time.Duration
	RefreshLiveness time.Duration
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
		Algo:   jwt.RS256,
	}
}

func (j JWT) signA(claims jwt.Claims, custClaims UserClaims) ([]byte, error) {
	token, err := jwt.Sign(j.Algo, j.Secret, claims, custClaims, jwt.MaxAge(j.AccessLiveness))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (j JWT) signR(claims jwt.Claims, custClaims RefreshClaims) ([]byte, error) {
	token, err := jwt.Sign(j.Algo, j.Secret, claims, custClaims, jwt.MaxAge(j.RefreshLiveness))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (j JWT) Issue(typ JwtType, userId string, orgId string, attribs map[Attribute]map[Attribute]bool, version int) ([]byte, error) {
	now := time.Now()

	if typ == JWT_TYPE_UNSPECIFIED {
		return nil, errors.New("unspecified jwt type")
	}

	if typ == JWT_TYPE_ACCESS_TOKEN {
		standardClaims := jwt.Claims{
			Subject:  userId,
			Expiry:   now.Add(j.AccessLiveness).Unix(),
			IssuedAt: now.Unix(),
			Issuer:   j.Issuer,
		}

		userClaims := UserClaims{
			OrgId:      orgId,
			Attributes: attribs,
		}

		sign, err := j.signA(standardClaims, userClaims)
		if err != nil {
			return nil, err
		}

		return sign, nil
	}

	if typ == JWT_TYPE_RFRESH_TOKEN {
		standardClaims := jwt.Claims{
			Subject:  userId,
			Expiry:   now.Add(j.AccessLiveness).Unix(),
			IssuedAt: now.Unix(),
			Issuer:   j.Issuer,
		}

		refreshClaims := RefreshClaims{
			OrgId:   orgId,
			Version: version,
		}

		sign, err := j.signR(standardClaims, refreshClaims)
		if err != nil {
			return nil, err
		}

		return sign, nil
	}

	return nil, errors.New("no jwt type selected")
}

func (j JWT) Verify(token string) (*UserClaims, error) {
	verifiedToken, err := jwt.Verify(j.Algo, j.Secret, []byte(token))
	if err != nil {
		return nil, err
	}

	var claims UserClaims
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
