package jwt

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stdeemene/go-travel2/config/env"
	"github.com/stdeemene/go-travel2/pkg/user/model"
	"github.com/stdeemene/go-travel2/utils/response"
	"net/http"
	"strings"
	"time"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	contextTokenClaims = contextKey("tokenClaims")
	contextEmailKey    = contextKey("email")
)

type TokenClaims struct {
	UserId        string `json:"id"`
	UserEmail     string `json:"email"`
	UserFirstname string `json:"firstname"`
	UserLastname  string `json:"lastname"`
	UserRole      string `json:"role"`
	UserIsActive  bool   `json:"isActive"`
	jwt.StandardClaims
}

type TokenDetails struct {
	UserId              string `json:"id"`
	UserRole            string `json:"role"`
	UserFirstname       string `json:"firstname"`
	UserLastName        string `json:"lastname"`
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	UserIsActive        bool   `json:"isActive"`
	AccessUUID          string `json:"access_token_uuid"`
	RefreshUUID         string `json:"refresh_token_uuid"`
	AccessTokenExpires  int64  `json:"access_token_exp"`
	RefreshTokenExpires int64  `json:"refresh_token_exp"`
}

func GenerateJwtToken(user *model.User) (*TokenDetails, error) {
	accessKey := env.GetEnvWithKey("JWT_ACCESS_KEY")
	refreshKey := env.GetEnvWithKey("JWT_REFRESH_KEY")
	issuer := env.GetEnvWithKey("JWT_ISSUER")
	td := &TokenDetails{}
	td.AccessTokenExpires = time.Now().Add(time.Hour * 24).Unix()
	td.AccessUUID = uuid.New().String()

	td.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.New().String()
	var accessSigningKey = []byte(accessKey)
	var refreshSigningKey = []byte(refreshKey)

	claims := &TokenClaims{
		UserId:        user.ID.Hex(),
		UserEmail:     user.Email,
		UserFirstname: user.Firstname,
		UserLastname:  user.Lastname,
		UserIsActive:  user.IsActive,
		UserRole:      user.Role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: td.AccessTokenExpires,
			Subject:   user.Email,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(accessSigningKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := token.SignedString(refreshSigningKey)
	if err != nil {
		return nil, err
	}

	td.AccessToken = accessToken
	td.RefreshToken = refreshToken
	td.UserRole = user.Role
	td.UserId = user.ID.Hex()
	td.UserFirstname = user.Firstname
	td.UserLastName = user.Lastname
	td.UserIsActive = user.IsActive
	return td, nil
}

func ProtectApi(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessKey := env.GetEnvWithKey("JWT_ACCESS_KEY")

		if strings.Contains(r.URL.Path, "/auth/") || strings.Contains(r.URL.Path, "/swagger/") {
			h.ServeHTTP(w, r)
		} else {
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader == "" {
				response.BaseResponse(w, http.StatusUnauthorized, "An authorization header is required")
				return
			}

			tokenString := strings.Split(authorizationHeader, " ")
			if len(tokenString) != 2 {
				response.BaseResponse(w, http.StatusUnauthorized, "Please pass the authorization header as <Bearer APIKEY>")
				return
			}
			tknStr := tokenString[1]

			tc := TokenClaims{}
			token, err := jwt.ParseWithClaims(tknStr, &tc, func(token *jwt.Token) (interface{}, error) {
				return []byte(accessKey), nil
			})
			if err != nil || !token.Valid {
				response.BaseResponse(w, http.StatusUnauthorized, "Invalid Access Token")
			} else {
				NewContext(r.Context(), &tc)
				ctx := context.WithValue(r.Context(), contextEmailKey, tc.UserEmail)
				h.ServeHTTP(w, r.WithContext(ctx))

				// NewContext(r.Context(), &tc)
				// h.ServeHTTP(w, r)
			}
		}

	})

}

// 	// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context, tokenClaims *TokenClaims) context.Context {
	return context.WithValue(ctx, contextTokenClaims, tokenClaims)
}

// 	// FromContext returns the User value stored øin ctx, if any.
func FromContext(ctx context.Context) (*TokenClaims, bool) {
	tc, ok := ctx.Value(contextTokenClaims).(*TokenClaims)
	return tc, ok
}

// AuthToken gets the auth token from the context.
func ContextEmail(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(contextEmailKey).(string)
	return email, ok
}

func Refresh(w http.ResponseWriter, r *http.Request) (*TokenDetails, error) {

	// // params := mux.Vars(r)
	// // tknStr := params["token"]
	// tknStr := new(models.RefreshReq)
	// err := json.NewDecoder(r.Body).Decode(&tknStr)
	// if err != nil {
	// 	middleware.BaseResponse(w, http.StatusBadRequest, err.Error())
	// }

	// fmt.Println("tknStr", tknStr.Token)
	td := &TokenDetails{}
	td.AccessTokenExpires = time.Now().Add(time.Hour * 24).Unix()
	td.AccessUUID = uuid.New().String()

	td.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.New().String()

	// tc := TokenClaims{}

	// tclaims, bool := FromContext(r.Context())
	// fmt.Println("claims", tclaims)
	// fmt.Println("bool", bool)

	// tkn, err := jwt.ParseWithClaims(tknStr.Token, &tc, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(config.Jwt.AccessKey), nil
	// })
	// fmt.Println("tc", tc)
	// fmt.Println("tkn", tkn)
	// if !tkn.Valid {
	// 	return nil, utility.NewError("Invalid Access Token")
	// }
	// if err != nil {
	// 	if err == jwt.ErrSignatureInvalid {
	// 		return nil, utility.NewError("Invalid Access Token")
	// 	}
	// 	return nil, utility.NewError("Bad Request")
	// }

	// (END) The code uptil this point is the same as the first part of the `Welcome` route
	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// var accessSigningKey = []byte(config.Jwt.AccessKey)
	// var refreshSigningKey = []byte(config.Jwt.RefreshKey)
	// Now, create a new token for the current use, with a renewed expiration time
	// expirationTime := time.Now().Add(5 * time.Minute)
	// claims.ExpiresAt = expirationTime.Unix()
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString(jwtKey)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// claims := &TokenClaims{
	// 	UserEmail:     td.UserLastName,
	// 	UserFirstname: user.Firstname,
	// 	UserLastname:  user.Lastname,
	// 	UserIsActive:  user.IsActive,
	// 	UserRole:      user.Role,
	// 	StandardClaims: jwt.StandardClaims{
	// 		Issuer:    config.Jwt.Issuer,
	// 		ExpiresAt: td.AccessTokenExpires,
	// 		Subject:   user.Email,
	// 	},
	// }
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// accessToken, err := token.SignedString(accessSigningKey)
	// if err != nil {
	// 	return nil, err
	// }

	// refreshToken, err := token.SignedString(refreshSigningKey)
	// if err != nil {
	// 	return nil, err
	// }
	// td.AccessToken = accessToken
	// td.RefreshToken = refreshToken
	// td.UserRole = user.Role
	// td.UserFirstname = user.Firstname
	// td.UserLastName = user.Lastname
	// td.UserIsActive = user.IsActive
	return td, nil
}
