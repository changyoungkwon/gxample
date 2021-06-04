package cli

import (
	"fmt"
	"net/http"
	"time"

	"github.com/changyoungkwon/gxample/internal/cluster"
	"github.com/changyoungkwon/gxample/internal/config"
	"github.com/changyoungkwon/gxample/internal/logging"
	"github.com/changyoungkwon/gxample/internal/routes"
	"github.com/spf13/cobra"
)

var registerEureka = false

// ServeCmd represents the serve command
var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		if registerEureka {
			conn := cluster.GetConn()
			conn.SendHeartbeatForever()
		}
		svr := newServer()
		logging.Logger.Fatal(svr.ListenAndServe())
	},
}

func newServer() *http.Server {
	return &http.Server{
		Handler:      routes.Router(),
		Addr:         fmt.Sprintf("0.0.0.0:%d", config.Get().API.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func init() {
	ServeCmd.Flags().BoolVarP(&registerEureka, "eureka", "e", false, "register eureka")
}
