package main
import (
	"fmt"
	// "encoding/json"
	// "net/http"
	"os"
	"github.com/joho/godotenv"
	"time"
	"strconv"
	"crypto/sha256"
	"crypto/hmac"
	"net/http"
)

type Headers struct{
	Authorization string
	sign string
	t string
	nonce string
}

type Paramaters struct{

}

func main(){
	token, secret := loadEnv()
	t, sign, nonce := sign(token, secret)
	fmt.Println(t,sign,nonce)
	url := "https://api.switch-bot.com/v1.1/devices"
	req, _ := http.NewRequest("GET",url, nil)
	req,Header.Set()
	resp, err := http.Get("")
}

func sign(token string, secret string) (string, string, string){
	t := strconv.FormatInt(time.Now().Unix()*1000, 10)
	var nonce = ""
	string_to_sign := token + t + nonce
	byte_sign, _ := getBinaryBySHA256WithKey(string_to_sign, secret)
	sign := string(byte_sign)
	return t, sign, nonce
}




func getBinaryBySHA256(s string) []byte {
    r := sha256.Sum256([]byte(s))
    return r[:]
}
func getBinaryBySHA256WithKey(msg, key string) ([]byte, error) {
    mac := hmac.New(sha256.New, getBinaryBySHA256(key))
    _, err := mac.Write([]byte(msg))
    return mac.Sum(nil), err
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