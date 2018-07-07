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
			case *linebot.FlexMessage:

			case *linebot.TextMessage:
				// form

				curText := "FlexForm{form ~ order\n______\nNama (sesuai line):Tolongin\nTempat Pemesanan atau Pengambilan Barang (patokan):KFC pasirkaliki\nAlamat Yang Dituju (patokan):Jl Dago Asri, Bandung\nNo Telepon (wajib diisi):081122334455\nCatatan:cepetan yah`Confirm{Yes|No}}"
				curText = strings.Replace(strings.TrimSuffix(curText, "}"), "FlexForm{", "", -1)
				formText := strings.Split(curText, "`")[0]
				form := strings.Split(formText, "\n")
				formFooter := strings.Split(curText, "`")[1]
				formFooter = strings.Replace(strings.TrimSuffix(formFooter, "}"), "Confirm{", "", -1)
				formFooterComponent := strings.Split(formFooter, "|")
				var flexBubbleContainer *linebot.BubbleContainer
				var flexFormHeader *linebot.BoxComponent
				var flexFormBody *linebot.BoxComponent
				var flexFormFooter *linebot.BoxComponent
				var bodyComponent []linebot.FlexComponent
				for _, row := range form {
					if strings.Contains(row, "form ~ ") {
						var headerComponent []linebot.FlexComponent
						var header *linebot.TextComponent
						titleElements := strings.Split(row, " ~ ")
						formTitle := strings.Title(titleElements[len(titleElements)-1])
						header = &linebot.TextComponent{
							Type:   linebot.FlexComponentTypeText,
							Text:   formTitle,
							Size:   linebot.FlexTextSizeTypeXl,
							Align:  linebot.FlexComponentAlignTypeCenter,
							Weight: linebot.FlexTextWeightTypeBold,
						}
						headerComponent = append(headerComponent, header)
						flexFormHeader = &linebot.BoxComponent{
							Type:     linebot.FlexComponentTypeBox,
							Layout:   linebot.FlexBoxLayoutTypeVertical,
							Contents: headerComponent,
							Margin:   linebot.FlexComponentMarginTypeSm,
						}
					} else if strings.Contains(row, ":") {
						var bodyContent *linebot.BoxComponent
						var bodyContentComponent []linebot.FlexComponent

						formBodyContent := strings.Split(row, ":")
						for index, text := range formBodyContent {
							var bodyLabelValue *linebot.TextComponent
							if index == 0 {
								bodyLabelValue = &linebot.TextComponent{
									Type:   linebot.FlexComponentTypeText,
									Text:   text,
									Weight: linebot.FlexTextWeightTypeBold,
									Wrap:   true,
									Size:   linebot.FlexTextSizeTypeXs,
								}
								bodyContentComponent = append(bodyContentComponent, bodyLabelValue)
							} else {
								bodyLabelValue = &linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  text,
									Wrap:  true,
									Align: linebot.FlexComponentAlignTypeEnd,
									Size:  linebot.FlexTextSizeTypeXs,
								}
								bodyContentComponent = append(bodyContentComponent, bodyLabelValue)
							}
						}
						bodyContent = &linebot.BoxComponent{
							Type:     linebot.FlexComponentTypeBox,
							Layout:   linebot.FlexBoxLayoutTypeHorizontal,
							Contents: bodyContentComponent,
							Spacing:  linebot.FlexComponentSpacingTypeMd,
						}
						bodyComponent = append(bodyComponent, bodyContent)
						separator := &linebot.SeparatorComponent{
							Type: linebot.FlexComponentTypeSeparator,
						}
						bodyComponent = append(bodyComponent, separator)
					}
				}

				var footerComponent []linebot.FlexComponent
				for _, button := range formFooterComponent {
					var color string
					if button == "Yes" {
						color = "#2E874A"
					} else {
						color = "#BA2424"
					}
					buttonFlex := &linebot.ButtonComponent{
						Type:   linebot.FlexComponentTypeButton,
						Action: linebot.NewMessageTemplateAction(button, button+" ~ "+formText),
						Style:  linebot.FlexButtonStyleTypePrimary,
						Color:  color,
						Height: linebot.FlexButtonHeightTypeSm,
					}
					footerComponent = append(footerComponent, buttonFlex)
				}

				flexFormFooter = &linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeHorizontal,
					Spacing:  linebot.FlexComponentSpacingTypeSm,
					Contents: footerComponent,
				}
				flexFormBody = &linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Contents: bodyComponent,
					Spacing:  linebot.FlexComponentSpacingTypeSm,
				}

				//Carousel

				// newFlex := "Flex{Header{}|Hero{https://firebasestorage.googleapis.com/v0/b/talkabot-a9388.appspot.com/o/getlargeimage.png?alt=media&token=add458d4-fc1e-459b-b92d-04355d2392db}|Body{Horizontal{FlexAction{1:1:1:1:1:1:1:Rincian Toko}};Vertical{FlexAction{1:1:1:1:1:1:2:Menu}~FlexAction{1:1:1:1:1:1:3:Tolongin Beliin}}}|Footer{}}"
				// newFlex = strings.Replace(strings.TrimSuffix(newFlex, "}"), "Flex{", "", -1)
				// flex := strings.Split(newFlex, "|")
				// fmt.Println("flex component", flex)
				// var flexBubbleContainer *linebot.BubbleContainer
				// var lineFlexHeader *linebot.BoxComponent
				// var lineFlexHero *linebot.ImageComponent
				// var lineFlexBody *linebot.BoxComponent
				// var lineFlexFooter *linebot.BoxComponent
				// for _, flexComponent := range flex {
				// 	// var lineFlexHero *linebot.ImageComponent
				// 	// var lineFlexBody *linebot.BoxComponent
				// 	if strings.Contains(flexComponent, "Header{") {
				// 		//TODO
				// 	} else if strings.Contains(flexComponent, "Hero{") {
				// 		//TODO
				// 		flexHero := strings.Replace(strings.TrimSuffix(flexComponent, "}"), "Hero{", "", -1)
				// 		lineFlexHero = &linebot.ImageComponent{
				// 			Type: linebot.FlexComponentTypeImage,
				// 			URL:  flexHero,
				// 			Size: linebot.FlexImageSizeTypeFull,
				// 		}
				// 		// flexBubbleContainer.Hero = lineFlexHero
				// 	} else if strings.Contains(flexComponent, "Body{") {
				// 		//TODO
				// 		var flexBodyContent []linebot.FlexComponent
				// 		flexBody := strings.Split(strings.Replace(strings.TrimSuffix(flexComponent, "}"), "Body{", "", -1), ";")
				// 		for _, flexBodyComponent := range flexBody {
				// 			// var flexBodyContentBox linebot.BoxComponent
				// 			if strings.Contains(flexBodyComponent, "Horizontal{") {
				// 				var lineFlexBodyHorizontal *linebot.BoxComponent
				// 				var lineFlexBodyHorizontalComponent []linebot.FlexComponent
				// 				flexBodyHorizontal := strings.Split(strings.Replace(strings.TrimSuffix(flexBodyComponent, "}"), "Horizontal{", "", -1), "~")
				// 				for _, flexBodyHorizontalComponent := range flexBodyHorizontal {
				// 					if strings.Contains(flexBodyHorizontalComponent, "FlexAction{") {
				// 						flexAction := strings.Replace(strings.TrimSuffix(flexBodyHorizontalComponent, "}"), "FlexAction{", "", -1)
				// 						var flexActionLabel string
				// 						if strings.Contains(flexAction, ":") {
				// 							flexActionElements := strings.Split(flexAction, ":")
				// 							flexActionLabel += flexActionElements[len(flexActionElements)-1]
				// 						} else {
				// 							flexActionLabel += flexAction
				// 						}
				// 						lineFlexAction := linebot.NewMessageTemplateAction(flexActionLabel, flexAction)
				// 						lineFlexBodyButton := &linebot.ButtonComponent{
				// 							Type:   linebot.FlexComponentTypeButton,
				// 							Action: lineFlexAction,
				// 							Style:  linebot.FlexButtonStyleTypePrimary,
				// 							Color:  "#000000",
				// 							Margin: linebot.FlexComponentMarginTypeSm,
				// 						}
				// 						lineFlexBodyHorizontalComponent = append(lineFlexBodyHorizontalComponent, lineFlexBodyButton)
				// 					} else {
				// 						lineFlexText := &linebot.TextComponent{
				// 							Type: linebot.FlexComponentTypeText,
				// 							Text: flexBodyHorizontalComponent,
				// 						}
				// 						lineFlexBodyHorizontalComponent = append(lineFlexBodyHorizontalComponent, lineFlexText)
				// 					}
				// 				}
				// 				lineFlexBodyHorizontal = &linebot.BoxComponent{
				// 					Type:     linebot.FlexComponentTypeBox,
				// 					Layout:   linebot.FlexBoxLayoutTypeHorizontal,
				// 					Contents: lineFlexBodyHorizontalComponent,
				// 				}
				// 				flexBodyContent = append(flexBodyContent, lineFlexBodyHorizontal)
				// 			} else if strings.Contains(flexBodyComponent, "Vertical{") {
				// 				var lineFlexBodyVertical *linebot.BoxComponent
				// 				var lineFlexBodyVerticalComponent []linebot.FlexComponent
				// 				flexBodyVertical := strings.Split(strings.Replace(strings.TrimSuffix(flexBodyComponent, "}"), "Vertical{", "", -1), "~")
				// 				for _, flexBodyVerticalComponent := range flexBodyVertical {
				// 					if strings.Contains(flexBodyVerticalComponent, "FlexAction{") {
				// 						flexAction := strings.Replace(strings.TrimSuffix(flexBodyVerticalComponent, "}"), "FlexAction{", "", -1)
				// 						var flexActionLabel string
				// 						if strings.Contains(flexAction, ":") {
				// 							flexActionElements := strings.Split(flexAction, ":")
				// 							flexActionLabel += flexActionElements[len(flexActionElements)-1]
				// 						} else {
				// 							flexActionLabel += flexAction
				// 						}
				// 						lineFlexAction := linebot.NewMessageTemplateAction(flexActionLabel, flexAction)
				// 						lineFlexBodyButton := &linebot.ButtonComponent{
				// 							Type:   linebot.FlexComponentTypeButton,
				// 							Action: lineFlexAction,
				// 							Style:  linebot.FlexButtonStyleTypeSecondary,
				// 							Color:  "#000000",
				// 							Margin: linebot.FlexComponentMarginTypeSm,
				// 						}
				// 						lineFlexBodyVerticalComponent = append(lineFlexBodyVerticalComponent, lineFlexBodyButton)
				// 					} else {
				// 						lineFlexText := &linebot.TextComponent{
				// 							Type: linebot.FlexComponentTypeText,
				// 							Text: flexBodyVerticalComponent,
				// 						}
				// 						lineFlexBodyVerticalComponent = append(lineFlexBodyVerticalComponent, lineFlexText)
				// 					}
				// 				}
				// 				lineFlexBodyVertical = &linebot.BoxComponent{
				// 					Type:     linebot.FlexComponentTypeBox,
				// 					Layout:   linebot.FlexBoxLayoutTypeVertical,
				// 					Contents: lineFlexBodyVerticalComponent,
				// 				}
				// 				flexBodyContent = append(flexBodyContent, lineFlexBodyVertical)
				// 			}
				// 		}
				// 		fmt.Println("flex body content", flexBodyContent)
				// 		fmt.Println("number of flex body content", flexBodyContent)
				// 		lineFlexBody = &linebot.BoxComponent{
				// 			Type:     linebot.FlexComponentTypeBox,
				// 			Layout:   linebot.FlexBoxLayoutTypeVertical,
				// 			Contents: flexBodyContent,
				// 		}
				// 		// flexBubbleContainer.Body = lineFlexBody
				// 	} else if strings.Contains(flexComponent, "Footer{") {
				// 		//TODO
				// 	}
				// }
				// flexBubbleContainer = &linebot.BubbleContainer{
				// 	Type:   linebot.FlexContainerTypeBubble,
				// 	Header: lineFlexHeader,
				// 	Hero:   lineFlexHero,
				// 	Body:   lineFlexBody,
				// 	Footer: lineFlexFooter,
				// }
				flexBubbleContainer = &linebot.BubbleContainer{
					Type:   linebot.FlexContainerTypeBubble,
					Header: flexFormHeader,
					Body:   flexFormBody,
					Footer: flexFormFooter,
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
