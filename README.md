go言語でgoogle sheet apiを使用して、sheetの値取得、行データの追加

手順<br>
補足：go.modのmoduleは適宜自分のものに変更する<br>
1: google cloud consoleにいき、新しいプロジェクトを作成する<br><br>
2: メニューバーからAPIとサービスを選択しgoogle sheet apiを有効にする<br><br>
3: APIとサービスから認証情報を選択し、認証情報の作成を行う。その際サービスアカウントで作成する<br><br>
4: 作成したサービスアカウントの「キー」というタブから鍵を作成し、jsonで作成する<br><br>
5: 鍵を作成したらjsonがダウンロードされるので、そのjsonをgoプロジェクトにcredentialsというディレクトリを作成し、jsonの名前を変更し、credentials/go_sheet.jsonという形で置く。<br><br>
6: google spread sheetを作成し、左下の「シート1」となっているシートの名前を任意の名前に変更する。変更した名前をgo_google_sheet_test.goの28行めに入れる<br><br>
6: google spread sheetのurlからspreadsheetIdをとってきて、 go_google_sheet_test.goの27行めに入れる<br><br>
