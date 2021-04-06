package main

import (
    "log"
    "net/smtp"
    "encoding/json"
    "net/http"
    "context"
    "github.com/go-kit/kit/endpoint"
    httptransport "github.com/go-kit/kit/transport/http"
)

type MailService interface {
    Send(string) int
}

type mainService struct{}

func (mainService) Send(s string) int {
    from_email   := "from address"
    to_email     := "to address"
    subject_body := "subject" + "\n\n" + "body"
    status       := smtp.SendMail("server:port", nil, from_email, []string{to_email}, []byte(subject_body))
    if status != nil {
        log.Printf("Error from SMTP Server: %s", status)
        return 0
    }
    log.Print("Email Sent Successfully")
    return 1
}

// メール送信のリクエスト
type sendRequest struct {
  S string `json:"s"`
}

// メール送信のレスポンス
type sendResponse struct {
  V int `json:"v"`
}

// メール送信のエンドポイント
func makeSendEndpoint(svc MailService) endpoint.Endpoint {
  return func(_ context.Context, request interface{}) (interface{}, error) {
    req := request.(sendRequest)
    v := svc.Send(req.S)
    return sendResponse{v}, nil
  }
}

func decodeSendRequest(_ context.Context, r *http.Request) (interface{}, error) {
  var request sendRequest
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    return nil, err
  }
  return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
  return json.NewEncoder(w).Encode(response)
}

func main() {
    svc := mainService{}
    // 文字数をカウントする処理のハンドラー
    sendHandler := httptransport.NewServer(
      makeSendEndpoint(svc),
      decodeSendRequest,
      encodeResponse,
    )
    // ハンドラーセットする。
    http.Handle("/send", sendHandler)
    // 8081ポートでサーバーを起動する。
    log.Fatal(http.ListenAndServe(":8081", nil))
}

// 参考
// https://netcorecloud.com/tutorials/send-email-through-gmail-smtp-server-using-go/

// どうやらAUTHをサポートしていないらしい…
// 2021/04/05 13:31:11 Error from SMTP Server: smtp: server doesn't support AUTH