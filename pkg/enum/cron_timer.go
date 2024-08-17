package enum

// ┌─────────────second 范围 (0 - 60)
// │ ┌───────────── min (0 - 59)
// │ │ ┌────────────── hour (0 - 23)
// │ │ │ ┌─────────────── day of month (1 - 31)
// │ │ │ │ ┌──────────────── month (1 - 12)
// │ │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
// │ │ │ │ │ │                  Saturday)
// │ │ │ │ │ │
// │ │ │ │ │ │
// * * * * * *
const (
	// CronEverySecond 每隔1秒
	CronEverySecond = "0/1 * * * * ?"
)
