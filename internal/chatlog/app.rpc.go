package chatlog

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/sjzar/chatlog/internal/chatlog/conf"
	"github.com/sjzar/chatlog/internal/wechat"
)

func (a *App) startRPC() {
	a.stopRPC()
	rpc.Register(a)
	rpc.HandleHTTP()

	l, err := net.Listen(a.ctx.RPCNetwork, a.ctx.RPCAddress)
	if err != nil {
		log.Fatal("[App::startRPC] listen error:", err)
	}
	a.rpcNetListener = l
	a.ctx.RCPRunning = true
	a.infoBar.UpdateRPC(
		true,
		l.Addr().Network(),
		l.Addr().String(),
	)
	http.Serve(l, nil)
	a.stopRPC()
}

func (a *App) stopRPC() {
	if a.rpcNetListener != nil {
		a.ctx.RCPRunning = false
		a.infoBar.UpdateRPC(
			false,
			a.rpcNetListener.Addr().Network(),
			a.rpcNetListener.Addr().String(),
		)
		a.rpcNetListener.Close()
		a.rpcNetListener = nil
	}
}

func (a *App) DecryptDBFiles(args *struct{}, reply *string) error {
	err := a.m.DecryptDBFiles()
	return err
}
func (a *App) StartService(args *struct{}, reply *string) error {
	err := a.m.StartService()
	return err
}
func (a *App) StopService(args *struct{}, reply *string) error {
	err := a.m.StopService()
	return err
}
func (a *App) StartAutoDecrypt(args *struct{}, reply *string) error {
	err := a.m.StartAutoDecrypt()
	return err
}
func (a *App) StopAutoDecrypt(args *struct{}, reply *string) error {
	err := a.m.StopAutoDecrypt()
	return err
}
func (a *App) RefreshSession(args *struct{}, reply *string) error {
	err := a.m.RefreshSession()
	return err
}

func (a *App) GetAccount(args *struct{}, reply *string) error {
	*reply = a.m.ctx.Account
	return nil
}
func (a *App) GetPlatform(args *struct{}, reply *string) error {
	*reply = a.m.ctx.Platform
	return nil
}
func (a *App) GetWXVersion(args *struct{}, reply *int) error {
	*reply = a.m.ctx.Version
	return nil
}
func (a *App) GetWXFullVersion(args *struct{}, reply *string) error {
	*reply = a.m.ctx.FullVersion
	return nil
}
func (a *App) GetDataDir(args *struct{}, reply *string) error {
	*reply = a.m.ctx.DataDir
	return nil
}
func (a *App) GetDataKey(args *struct{}, reply *string) error {
	err := a.m.GetDataKey()
	*reply = a.m.ctx.DataKey
	return err
}
func (a *App) GetDataUsage(args *struct{}, reply *string) error {
	*reply = a.m.ctx.DataUsage
	return nil
}
func (a *App) GetWorkDir(args *struct{}, reply *string) error {
	*reply = a.m.ctx.WorkDir
	return nil
}
func (a *App) GetWorkUsage(args *struct{}, reply *string) error {
	*reply = a.m.ctx.WorkUsage
	return nil
}
func (a *App) GetHTTPEnabled(args *struct{}, reply *bool) error {
	*reply = a.m.ctx.HTTPEnabled
	return nil
}
func (a *App) GetHTTPAddr(args *struct{}, reply *string) error {
	*reply = a.m.ctx.HTTPAddr
	return nil
}

func (a *App) GetAutoDecrypt(args *struct{}, reply *bool) error {
	*reply = a.m.ctx.AutoDecrypt
	return nil
}
func (a *App) GetLastSession(args *struct{}, reply *int64) error {
	*reply = a.m.ctx.LastSession.UnixMilli()
	return nil
}
func (a *App) GetWXPIP(args *struct{}, reply *int) error {
	*reply = a.m.ctx.PID
	return nil
}
func (a *App) GetWXExePath(args *struct{}, reply *string) error {
	*reply = a.m.ctx.ExePath
	return nil
}
func (a *App) GetWXStatus(args *struct{}, reply *string) error {
	*reply = a.m.ctx.Status
	return nil
}
func (a *App) GetCurrentWX(args *struct{}, reply *wechat.Account) error {
	*reply = *a.m.ctx.Current
	return nil
}

func (a *App) GetWXInstances(args *struct{}, reply *[]*wechat.Account) error {
	if len(a.m.ctx.WeChatInstances) == 0 {
		a.m.ctx.WeChatInstances = a.m.wechat.GetWeChatInstances()
		if len(a.m.ctx.WeChatInstances) >= 1 {
			a.m.ctx.SwitchCurrent(a.m.ctx.WeChatInstances[0])
		}
		if a.m.ctx.HTTPEnabled {
			// 启动HTTP服务
			if err := a.m.StartService(); err != nil {
				a.m.StopService()
			}
		}
	}
	*reply = a.m.ctx.WeChatInstances
	return nil
}

func (a *App) GetHistory(args *struct{}, reply *map[string]conf.ProcessConfig) error {
	*reply = a.m.ctx.History
	return nil
}

func (a *App) GetImgKey(args *struct{}, reply *string) error {
	*reply = a.m.ctx.ImgKey
	return nil
}
