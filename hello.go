package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
			case *linebot.TextMessage:
				newFlex := "Flex{Header{}|Hero{https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db}|Body{Horizontal{Alamat~FlexAction{1:1:1:1:1:1:1:Rincian Toko}};Vertical{FlexButton{1:1:1:1:1:1:2:Menu}~FlexButton{1:1:1:1:1:1:3:Tolongin Beliin}}}|Footer{}}"
				newFlex = strings.Replace(strings.TrimSuffix(newFlex, "}"), "Flex{", "", -1)
				flex := strings.Split(newFlex, "|")
				fmt.Println("flex component", flex)
				var flexBubbleContainer *linebot.BubbleContainer
				// flexBubbleContainer.Type = linebot.FlexContainerTypeBubble
				var lineFlexHero *linebot.ImageComponent
				for _, flexComponent := range flex {
					// var lineFlexHero *linebot.ImageComponent
					// var lineFlexBody *linebot.BoxComponent
					if strings.Contains(flexComponent, "Header{") {
						//TODO
					} else if strings.Contains(flexComponent, "Hero{") {
						//TODO
						flexHero := strings.Replace(strings.TrimSuffix(flexComponent, "}"), "Hero{", "", -1)
						lineFlexHero = &linebot.ImageComponent{
							Type: linebot.FlexComponentTypeImage,
							URL:  flexHero,
						}
						// flexBubbleContainer.Hero = lineFlexHero
					} else if strings.Contains(flexComponent, "Body{") {
						//TODO
						// var flexBodyContent []linebot.FlexComponent
						// flexBody := strings.Split(strings.Replace(strings.TrimSuffix(flexComponent, "}"), "Body{", "", -1), ";")
						// for _, flexBodyComponent := range flexBody {
						// 	// var flexBodyContentBox linebot.BoxComponent
						// 	if strings.Contains(flexBodyComponent, "Horizontal{") {
						// 		var lineFlexBodyHorizontal *linebot.BoxComponent
						// 		var lineFlexBodyHorizontalComponent []linebot.FlexComponent
						// 		flexBodyHorizontal := strings.Split(strings.Replace(strings.TrimSuffix(flexBodyComponent, ""), "Horizontal{", "", -1), "~")
						// 		for _, flexBodyHorizontalComponent := range flexBodyHorizontal {
						// 			if strings.Contains(flexBodyHorizontalComponent, "FlexAction{") {
						// 				flexActionLabel := strings.Replace(strings.TrimSuffix(flexBodyHorizontalComponent, "}"), "FlexAction{", "", -1)
						// 				var flexAction string
						// 				if strings.Contains(flexAction, ":") {
						// 					flexActionElements := strings.Split(flexActionLabel, ":")
						// 					flexAction = flexActionElements[len(flexActionElements)-1]
						// 				} else {
						// 					flexAction = flexActionLabel
						// 				}
						// 				lineFlexAction := linebot.NewMessageTemplateAction(flexActionLabel, flexAction)
						// 				lineFlexBodyButton := &linebot.ButtonComponent{
						// 					Type:   linebot.FlexComponentTypeButton,
						// 					Action: lineFlexAction,
						// 				}
						// 				lineFlexBodyHorizontalComponent = append(lineFlexBodyHorizontalComponent, lineFlexBodyButton)
						// 			} else {
						// 				lineFlexText := &linebot.TextComponent{
						// 					Type: linebot.FlexComponentTypeText,
						// 					Text: flexBodyHorizontalComponent,
						// 				}
						// 				lineFlexBodyHorizontalComponent = append(lineFlexBodyHorizontalComponent, lineFlexText)
						// 			}
						// 		}
						// 		lineFlexBodyHorizontal = &linebot.BoxComponent{
						// 			Type:     linebot.FlexComponentTypeBox,
						// 			Layout:   "horizontal",
						// 			Contents: lineFlexBodyHorizontalComponent,
						// 		}
						// 		flexBodyContent = append(flexBodyContent, lineFlexBodyHorizontal)
						// 	} else if strings.Contains(flexBodyComponent, "Vertical{") {
						// 		var lineFlexBodyVertical *linebot.BoxComponent
						// 		var lineFlexBodyVerticalComponent []linebot.FlexComponent
						// 		flexBodyVertical := strings.Split(strings.Replace(strings.TrimSuffix(flexBodyComponent, ""), "Vertical{", "", -1), "~")
						// 		for _, flexBodyVerticalComponent := range flexBodyVertical {
						// 			if strings.Contains(flexBodyVerticalComponent, "FlexAction{") {
						// 				flexActionLabel := strings.Replace(strings.TrimSuffix(flexBodyVerticalComponent, "}"), "FlexAction{", "", -1)
						// 				var flexAction string
						// 				if strings.Contains(flexAction, ":") {
						// 					flexActionElements := strings.Split(flexActionLabel, ":")
						// 					flexAction = flexActionElements[len(flexActionElements)-1]
						// 				} else {
						// 					flexAction = flexActionLabel
						// 				}
						// 				lineFlexAction := linebot.NewMessageTemplateAction(flexActionLabel, flexAction)
						// 				lineFlexBodyButton := &linebot.ButtonComponent{
						// 					Type:   linebot.FlexComponentTypeButton,
						// 					Action: lineFlexAction,
						// 				}
						// 				lineFlexBodyVerticalComponent = append(lineFlexBodyVerticalComponent, lineFlexBodyButton)
						// 			} else {
						// 				lineFlexText := &linebot.TextComponent{
						// 					Type: linebot.FlexComponentTypeText,
						// 					Text: flexBodyVerticalComponent,
						// 				}
						// 				lineFlexBodyVerticalComponent = append(lineFlexBodyVerticalComponent, lineFlexText)
						// 			}
						// 		}
						// 		lineFlexBodyVertical = &linebot.BoxComponent{
						// 			Type:     linebot.FlexComponentTypeBox,
						// 			Layout:   "vertical",
						// 			Contents: lineFlexBodyVerticalComponent,
						// 		}
						// 		flexBodyContent = append(flexBodyContent, lineFlexBodyVertical)
						// 	}
						// }
						// fmt.Println("flex body content", flexBodyContent)
						// fmt.Println("number of flex body content", flexBodyContent)
						// lineFlexBody = &linebot.BoxComponent{
						// 	Type:     linebot.FlexComponentTypeBox,
						// 	Layout:   "vertical",
						// 	Contents: flexBodyContent,
						// }
						// flexBubbleContainer.Body = lineFlexBody
					} else if strings.Contains(flexComponent, "Footer{") {
						//TODO
					}
				}

				// contents := &linebot.BubbleContainer{
				// 	Type: linebot.FlexContainerTypeBubble,
				// 	Body: &linebot.BoxComponent{
				// 		Type:   linebot.FlexComponentTypeBox,
				// 		Layout: linebot.FlexBoxLayoutTypeHorizontal,
				// 		Contents: []linebot.FlexComponent{
				// 			&linebot.TextComponent{
				// 				Type: linebot.FlexComponentTypeText,
				// 				Text: "Hello,",
				// 			},
				// 			&linebot.TextComponent{
				// 				Type: linebot.FlexComponentTypeText,
				// 				Text: "World!",
				// 			},
				// 		},
				// 	},
				// }
				flexBubbleContainer = &linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Hero: lineFlexHero,
				}
				if _, err = bot.ReplyMessage(
					event.ReplyToken,
					linebot.NewFlexMessage("Flex alt text", flexBubbleContainer),
				).Do(); err != nil {
					log.Print(err)
				}

			}
		}
	}
}
