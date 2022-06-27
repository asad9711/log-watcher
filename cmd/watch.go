/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "utility to read/watch a remote file",
	Long: `utility to read/watch a remote file`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("watch cmd called - args - ", args)
		fileName, _ := cmd.Flags().GetString("fileName")
		log.Println("fileName - ", fileName)
		numberOfLines, _ := cmd.Flags().GetString("numberOfLines")

		// make http request to server
		requestBody, _ := json.Marshal(map[string]string{
			"numberOfLines": numberOfLines,
			"fileName":      fileName,
		})
		requestBuffer := bytes.NewBuffer(requestBody)
		resp, err := http.Post("http://localhost:8090/read", "application/text", requestBuffer)
		if err != nil {
			log.Println("error while making post request - ", err.Error())
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Println("resp is - ", sb)

	},
}

func init() {
	rootCmd.AddCommand(watchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	watchCmd.Flags().StringP("fileName", "f", "", "name of file to read from")
	watchCmd.Flags().StringP("numberOfLines", "n", "", "number of lines to read")
	watchCmd.MarkFlagRequired("fileName")

}
