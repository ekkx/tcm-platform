package main

import "github.com/ekkx/tcmrsv"

// OfficialClient は公式予約サイトとの通信を抽象化するインターフェース。
// テスト時に mock に差し替えることでスケジューラーのロジックを検証できる。
type OfficialClient interface {
	Login(params *tcmrsv.LoginParams) error
	Reserve(params *tcmrsv.ReserveParams) error
}

// OfficialClientFactory は公式サイトクライアントの生成関数。
// 本番では tcmrsv.New を、テストでは mock を返す。
type OfficialClientFactory func() OfficialClient

// tcmrsvClient は tcmrsv.Client を OfficialClient インターフェースに適合させるアダプタ。
type tcmrsvClient struct {
	inner *tcmrsv.Client
}

func newTCMRSVClientFactory() OfficialClientFactory {
	return func() OfficialClient {
		return &tcmrsvClient{inner: tcmrsv.New()}
	}
}

func (c *tcmrsvClient) Login(params *tcmrsv.LoginParams) error {
	return c.inner.Login(params)
}

func (c *tcmrsvClient) Reserve(params *tcmrsv.ReserveParams) error {
	return c.inner.Reserve(params)
}
