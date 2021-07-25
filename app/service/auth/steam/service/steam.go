package service

import (
	pb "backend/app/service/auth/steam/api/grpc/v1"
	accountSrv "backend/app/service/user/account/api/grpc/v1"
	"backend/pkg/ecode"
	"context"
	"github.com/yohcop/openid-go"
	"google.golang.org/grpc/codes"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

const (
	callbackURL = "/auth/steam/callback"
)

// NoOpDiscoveryCache implements the DiscoveryCache interface and doesn't cache anything.
// For a simple website, I'm not sure you need a cache.
type NoOpDiscoveryCache struct{}

// Put is a no op.
func (n *NoOpDiscoveryCache) Put(id string, info openid.DiscoveredInfo) {}

// Get always returns nil.
func (n *NoOpDiscoveryCache) Get(id string) openid.DiscoveredInfo {
	return nil
}

var nonceStore = openid.NewSimpleNonceStore()
var discoveryCache = &NoOpDiscoveryCache{}

func (s *Service) SteamLogin(ctx context.Context, req *pb.SteamLoginReq) (resp *pb.SteamLoginResp, err error) {
	resp = &pb.SteamLoginResp{}

	resp.Redirect, err = openid.RedirectURL(
		s.config.SteamOpenID.Endpoint,
		s.config.SteamOpenID.Host+callbackURL,
		s.config.SteamOpenID.Host+"/")
	return
}

func (s *Service) SteamCallback(ctx context.Context, req *pb.SteamCallbackReq) (resp *pb.SteamCallbackResp, err error) {
	resp = &pb.SteamCallbackResp{}

	if req.Openid == nil {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid openid filed")
		return
	}

	payload := s.config.JWT.PayloadFormContext(ctx)
	fullURL, _ := url.Parse(req.Openid.ReturnTo)

	// OpenID verify
	value := fullURL.Query()
	value.Set("openid.ns", req.Openid.Ns)
	value.Set("openid.mode", req.Openid.Mode)
	value.Set("openid.op_endpoint", req.Openid.OpEndpoint)
	value.Set("openid.claimed_id", req.Openid.ClaimedId)
	value.Set("openid.identity", req.Openid.Identity)
	value.Set("openid.return_to", req.Openid.ReturnTo)
	value.Set("openid.response_nonce", req.Openid.ResponseNonce)
	value.Set("openid.assoc_handle", req.Openid.AssocHandle)
	value.Set("openid.signed", req.Openid.Signed)
	value.Set("openid.sig", req.Openid.Sig)
	fullURL.RawQuery = value.Encode()

	client := &http.Client{}
	if len(s.config.SteamOpenID.HttpProxy) > 0 {
		proxy, _ := url.Parse("http://" + s.config.SteamOpenID.HttpProxy)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}
	openID := openid.NewOpenID(client)
	id, err := openID.Verify(fullURL.String(), discoveryCache, nonceStore)
	if err != nil {
		return
	}

	// Get SteamID
	r := regexp.MustCompile(`https?://steamcommunity.com/openid/id/(\d+)`).FindStringSubmatch(id)
	if len(r) < 2 {
		err = ecode.Errorf(codes.Unknown, "failed to get SteamID")
		return
	}
	steamID, _ := strconv.ParseInt(r[1], 10, 64)
	payload.Set("steam_id", r[1])

	// Get UID
	uid, err := s.account.UID(ctx, &accountSrv.UIDReq{SteamId: steamID})
	if err != nil {
		return
	}
	payload.Set("uid", uid.Uid)

	// Add to connect token
	resp.JwtString, err = s.config.JWT.NewTokenString(*payload)
	return
}
