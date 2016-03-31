package providers

import (
	"github.com/gosuri/uitable"
	"github.com/shirou/gopsutil/host"
)

func UserStats() string {
	userStats, err := host.Users()
	if err != nil {
		panic(err)
	}
	userTable := uitable.New()
	userTable.AddRow("USERS")
	userTable.AddRow("#", "User", "Terminal", "Host", "Started")
	for i, user := range userStats {
		userTable.AddRow(i, user.User, user.Terminal, user.Host, user.Started)
	}

	return userTable.String()
}
