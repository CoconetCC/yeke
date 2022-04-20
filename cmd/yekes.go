package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var yekesCmd = &cobra.Command{
	Use:   "yekes",
	Short: "yeke的扫描模块",
	Long:  `从文件导入进行yeke扫描`,
	Run: func(cmd *cobra.Command, args []string) {
		color.RGBStyleFromString("58,191,219").Println(`
██    ██         ██   ██     by: CoconetCC   
░░██  ██         ░██  ██         
 ░░████    █████ ░██ ██    █████ 
  ░░██    ██░░░██░████    ██░░░██
   ░██   ░███████░██░██  ░███████
   ░██   ░██░░░░ ░██░░██ ░██░░░░ 
   ░██   ░░██████░██ ░░██░░██████
   ░░     ░░░░░░ ░░   ░░  ░░░░░░

`)
	},
}

var (
	yekes_scan      string
	yekes_output    string
	yekes_localfile string
	yekes_thread    int
	yekes_proxy     string
)

func init() {
	rootCmd.AddCommand(yekesCmd)
	yekesCmd.Flags().StringVarP(&yekes_scan, "scan", "s", "", "从本地payload.yaml读取poc进行扫描")
	yekesCmd.Flags().StringVarP(&yekes_localfile, "local", "l", "", "从本地文件读取资产，进行指纹识别，支持无协议，列如：192.168.1.1:9090 | http://192.168.1.1:9090")
	yekesCmd.Flags().StringVarP(&yekes_output, "output", "o", "", "输出所有结果，当前仅支持json和xlsx后缀的文件。")
	yekesCmd.Flags().IntVarP(&yekes_thread, "thread", "t", 100, "指纹识别线程大小。")
	yekesCmd.Flags().StringVarP(&proxy, "proxy", "p", "", "指定访问目标时的代理，支持http代理和socks5，例如：http://127.0.0.1:8080、socks5://127.0.0.1:8080")

}
