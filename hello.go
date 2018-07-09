package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch event.Message.(type) {
			case *linebot.FlexMessage:

			case *linebot.TextMessage:
				var reply []linebot.Message
				// form
				form := "FlexForm{form ~ order\n______\nNama (sesuai line):Tolongin\nTempat Pemesanan atau Pengambilan Barang (patokan):KFC pasirkaliki\nAlamat Yang Dituju (patokan):Jl Dago Asri, Bandung\nNo Telepon (wajib diisi):081122334455\nCatatan:cepetan yah`Confirm{Yes|No}}"
				reply = append(reply, LineFlexForm(form))

				//Carousel
				carousel := "carousel{https://pbs.twimg.com/profile_images/994899259462795265/jiCT03Qf_400x400.jpg~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Alamat~1:1:1:2:1:1:Alamat~button;Lihat Menu~http://www.google.com~url|https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Lihat Menu~http://www.google.com~url|https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Lihat Menu~http://www.google.com~url|https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Lihat Menu~http://www.google.com~url|https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Lihat Menu~http://www.google.com~url|https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Lihat Menu~http://www.google.com~url|https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Lihat Menu~http://www.google.com~url|https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Lihat Menu~http://www.google.com~url|https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Lihat Menu~http://www.google.com~url|https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db~printer~printing;Alamat~1:1:1:2:1:1:Alamat~button;Rincian Toko~1:1:1:2:1:2:Rincian Toko~button;Lihat Menu~http://www.google.com~url}"
				reply = append(reply, LineFlexCarousel(carousel))

				//Confirm
				confirm := "Confirm{Are you sure?;Yes|No}"
				reply = append(reply, LineFlexConfirm(confirm))

				//Button
				button := "Button{Test;a|b|c|d}"
				reply = append(reply, LineFlexButton(button))

				if _, err = bot.ReplyMessage(
					event.ReplyToken,
					reply...,
				).Do(); err != nil {
					log.Print(err)
				}

			}
		}
	}
}
