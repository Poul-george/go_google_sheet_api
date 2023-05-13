package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
)

func main() {
	// 認証情報の読み込み
	creds, err := getCredentials2()
	if err != nil {
		log.Fatalf("Failed to get credentials: %v", err)
	}

	// Google Sheets API クライアントの作成
	ctx := context.Background()
	client, err := sheets.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to create sheets client: %v", err)
	}

	// スプレッドシートの ID とシート名を指定
	spreadsheetId := ""
	sheetName := ""

	// スプレッドシートからrangeを指定しないで値を取得
	readRange := sheetName

	resp, err := client.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("wwFailed to read data from sheet: %v", err)
	}

	// 取得した値を出力
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		getResponseOutput(resp.Values)
	}

	rowData := []interface{}{"投稿したいデータ"}
	insertReq := &sheets.ValueRange{
		Values: [][]interface{}{rowData},
	}

	//データ登録
	_, err = client.Spreadsheets.Values.Append(spreadsheetId, sheetName+"!A1", insertReq).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		log.Fatalf("Failed to append row: %v", err)
	}

	resp, err = client.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("wwFailed to read data from sheet: %v", err)
	}

	// 取得した値を出力
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		getResponseOutput2(resp.Values)
	}
}

func getResponseOutput2(res [][]interface{}) {
	fmt.Println("Name, Major:")
	for _, row := range res {
		// スプレッドシートのカラムが 3 列であることを前提としている
		for i := 0; i < len(row); i++ {
			fmt.Printf("%v, ", row[i])
		}
		fmt.Println()
	}
}

// 認証情報の取得
func getCredentials2() (*google.Credentials, error) {
	credsFile := "credentials/go_sheet.json"
	b, err := ioutil.ReadFile(credsFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}

	// 認証情報を取得
	creds, err := google.CredentialsFromJSON(context.Background(), b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	return creds, nil
}
