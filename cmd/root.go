package cmd

import (
	"os"

	"github.com/mobenaus/fc-pos-go-desafio-stress-test/internal/report"
	"github.com/mobenaus/fc-pos-go-desafio-stress-test/internal/stress"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc-pos-go-desafio-stress-test",
	Short: "Aplicação de stress-test para o curso de Go da FullCycle",
	Long:  `Essa aplicação irá executar requsições HTTP para uma URL determinada via parametro.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		ctx := cmd.Context()
		st := stress.NewStressTest(url, requests, concurrency)
		result, _ := st.Execute(ctx)
		report.DisplayReport(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("url", "", "URL para teste")
	rootCmd.MarkFlagRequired("url")
	rootCmd.Flags().Int("requests", 1000, "Numero de requests para disparar")
	rootCmd.Flags().Int("concurrency", 5, "Quantidade de processos concorrentes")
}
