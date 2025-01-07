package pdf

import (
	"context"

	"github.com/chromedp/chromedp"
)

type RemoteChromeDevToolsContext struct {
	devtoolsURL string
	allocator   context.Context
	cancel      context.CancelFunc
}

func NewRemoteChromeDevToolsContext(ctx context.Context, devtoolsURL string) *RemoteChromeDevToolsContext {
	allocator, cancel := chromedp.NewRemoteAllocator(ctx, devtoolsURL)

	return &RemoteChromeDevToolsContext{
		devtoolsURL: devtoolsURL,
		allocator:   allocator,
		cancel:      cancel,
	}
}

func (r *RemoteChromeDevToolsContext) Cancel() {
	r.cancel()
}

func (r *RemoteChromeDevToolsContext) WithContext(cb func(context.Context)) {
	ctx, cancel := chromedp.NewContext(r.allocator)
	defer cancel()

	cb(ctx)
}
