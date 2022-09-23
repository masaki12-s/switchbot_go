package main
import (
	"fmt"
	// "encoding/json"
	// "net/http"
	"os"
	"github.com/joho/godotenv"
)
func main(){
	token, secret := loadEnv()
	fmt.Println(token,secret)
}

func sign(string, string) (string, string, string){
	return "a","b","c"
}

func loadEnv() (string, string){
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")
	
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	
	// .envの SAMPLE_MESSAGEを取得して、messageに代入します。
	token := os.Getenv("token")
	secret := os.Getenv("secret")
	return token, secret
}