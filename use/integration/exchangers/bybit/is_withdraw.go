package bybit

import "context"

func (c *client) IsWithdraw(ctx context.Context, symbol string) bool {
	return true
}
