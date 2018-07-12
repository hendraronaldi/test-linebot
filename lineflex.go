package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func LineFlexButton(curText string) *linebot.FlexMessage {
	var element []string
	var buttonText []string
	var carouselButtonComponent []*linebot.BubbleContainer
	title := " "
	curText = curText[7 : len(curText)-1]
	if strings.Contains(curText, ";") {
		element = strings.Split(curText, ";")
		title = element[0]
		if strings.Contains(element[1], "|") {
			buttonText = strings.Split(element[1], "|")
		} else {
			buttonText = append(buttonText, element[1])
		}
	} else {
		if strings.Contains(curText, "|") {
			buttonText = strings.Split(curText, "|")
		} else {
			buttonText = append(buttonText, curText)
		}
	}
	var templateHeaderComponent []linebot.FlexComponent
	headerTextComponent := &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   title,
		Weight: linebot.FlexTextWeightTypeBold,
		Wrap:   true,
	}
	templateHeaderComponent = append(templateHeaderComponent, headerTextComponent)
	templateHeader := &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: templateHeaderComponent,
	}

	var buttonCarousel []linebot.FlexComponent
	for index := range buttonText {
		fmt.Println("index", index)
		element := buttonText[index]
		var text string
		if strings.Contains(element, ":") {
			content := strings.Split(element, ":")
			text = content[len(content)-1]
		} else {
			text = element
		}
		buttonColumn := &linebot.ButtonComponent{
			Type:   linebot.FlexComponentTypeButton,
			Action: linebot.NewMessageTemplateAction(text, element),
		}
		fmt.Println("element", element)
		buttonCarousel = append(buttonCarousel, buttonColumn)
		if (index > 0 && index%4 == 0) || index == len(buttonText)-1 {
			var buttonTemplate *linebot.BoxComponent
			var start int
			var end int
			if index > 0 && index%4 == 0 {
				start = index - 4
				end = index
			} else {
				start = index - index%4
				end = len(buttonText)
			}
			fmt.Println("start", start)
			fmt.Println("end", end)
			buttonTemplate = &linebot.BoxComponent{
				Type:     linebot.FlexComponentTypeBox,
				Layout:   linebot.FlexBoxLayoutTypeVertical,
				Contents: buttonCarousel[start:end],
			}
			blockStyle := &linebot.BubbleStyle{
				Header: &linebot.BlockStyle{
					BackgroundColor: "#e5e5e5",
				},
				Footer: &linebot.BlockStyle{
					Separator: true,
				},
			}
			buttonFlexTemplate := &linebot.BubbleContainer{
				Type:   linebot.FlexContainerTypeBubble,
				Header: templateHeader,
				Footer: buttonTemplate,
				Styles: blockStyle,
			}
			carouselButtonComponent = append(carouselButtonComponent, buttonFlexTemplate)
			// buttonCarousel = buttonCarousel[:0]
		}
	}
	// var buttonTemplate *linebot.BoxComponent
	// if len(buttonCarousel) <= 4 {
	// 	buttonTemplate := &linebot.BoxComponent{
	// 		Type:     linebot.FlexComponentTypeBox,
	// 		Layout:   linebot.FlexBoxLayoutTypeVertical,
	// 		Contents: buttonCarousel,
	// 	}
	// } else {
	// 	for index := range buttonCarousel {
	// 		if (index > 0 && index %4 == 0) || index == len(buttonCarousel) - 1 {
	// 			if
	// 		}
	// 	}
	// }

	carouselButtonFlexTemplate := &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: carouselButtonComponent,
	}

	return linebot.NewFlexMessage("buttons", carouselButtonFlexTemplate)
}

func LineFlexConfirm(curText string) *linebot.FlexMessage {
	curText = strings.TrimRight(strings.TrimLeft(curText, "Confirm{"), "}")
	var confirmationText string
	var index int
	element := strings.Split(curText, ";")
	confirmationText = element[0]
	if len(element) > 2 {
		index = 2
	} else {
		index = 1
	}
	confirmText := strings.Split(element[index], "|")

	var confirmFlexTemplate *linebot.BubbleContainer
	var templateHeader *linebot.BoxComponent
	var templateBody *linebot.BoxComponent
	var templateHeaderComponent []linebot.FlexComponent
	var templateBodyComponent []linebot.FlexComponent
	var confirmTextComponent *linebot.TextComponent
	var confirmButtonTemplate *linebot.BoxComponent
	var confirmButtonTemplateComponent []linebot.FlexComponent
	var confirmButtonYes *linebot.ButtonComponent
	var confirmButtonNo *linebot.ButtonComponent
	var blockStyle *linebot.BubbleStyle

	confirmTextComponent = &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   confirmationText,
		Weight: linebot.FlexTextWeightTypeBold,
		Wrap:   true,
	}
	templateHeaderComponent = append(templateHeaderComponent, confirmTextComponent)
	templateHeader = &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: templateHeaderComponent,
	}

	confirmButtonYes = &linebot.ButtonComponent{
		Type:   linebot.FlexComponentTypeButton,
		Action: linebot.NewMessageTemplateAction(confirmText[0], confirmText[0]),
		Style:  linebot.FlexButtonStyleTypeLink,
	}
	confirmButtonNo = &linebot.ButtonComponent{
		Type:   linebot.FlexComponentTypeButton,
		Action: linebot.NewMessageTemplateAction(confirmText[1], confirmText[1]),
		Style:  linebot.FlexButtonStyleTypeLink,
	}
	confirmButtonTemplateComponent = append(confirmButtonTemplateComponent, confirmButtonYes, confirmButtonNo)
	confirmButtonTemplate = &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeHorizontal,
		Contents: confirmButtonTemplateComponent,
	}

	templateBodyComponent = append(templateBodyComponent, confirmButtonTemplate)
	templateBody = &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: templateBodyComponent,
	}

	blockStyle = &linebot.BubbleStyle{
		Header: &linebot.BlockStyle{
			BackgroundColor: "#e5e5e5",
		},
		Footer: &linebot.BlockStyle{
			Separator: true,
		},
	}

	confirmFlexTemplate = &linebot.BubbleContainer{
		Type:   linebot.FlexContainerTypeBubble,
		Header: templateHeader,
		Footer: templateBody,
		Styles: blockStyle,
	}

	return linebot.NewFlexMessage("confirmation", confirmFlexTemplate)
}

func LineFlexCarousel(curText string) *linebot.FlexMessage {
	curText = strings.TrimRight(strings.TrimLeft(curText, "carousel{"), "}")
	carouselContent := strings.Split(curText, "|")
	var carousel []*linebot.BubbleContainer
	var carouselContainer *linebot.CarouselContainer
	for index := 0; index < len(carouselContent); index++ {
		var image string
		var title string
		var description string
		var flexBubbleContainer *linebot.BubbleContainer
		var lineFlexHeader *linebot.BoxComponent
		var lineFlexHero *linebot.ImageComponent
		var lineFlexBody *linebot.BoxComponent
		var lineFlexBodyComponent []linebot.FlexComponent
		var lineFlexFooter *linebot.BoxComponent
		var lineFooterComponent []linebot.FlexComponent
		content := strings.Split(carouselContent[index], ";")
		for i := 0; i < len(content); i++ {
			element := strings.Split(content[i], "~")
			if i == 0 {
				//hero
				image = element[0]
				lineFlexHero = &linebot.ImageComponent{
					Type:       linebot.FlexComponentTypeImage,
					URL:        image,
					Size:       linebot.FlexImageSizeTypeFull,
					AspectMode: linebot.FlexImageAspectModeTypeCover,
				}

				//body
				title = element[1]
				titleComponent := &linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   title,
					Size:   linebot.FlexTextSizeTypeLg,
					Weight: linebot.FlexTextWeightTypeBold,
					Wrap:   true,
				}
				description = element[2]
				descriptionComponent := &linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  description,
					Wrap:  true,
					Size:  linebot.FlexTextSizeTypeSm,
					Color: "#b5b5b5",
				}
				lineFlexBodyComponent = append(lineFlexBodyComponent, titleComponent, descriptionComponent)
				lineFlexBody = &linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Contents: lineFlexBodyComponent,
				}
			} else {
				//footer
				var footerButton *linebot.ButtonComponent
				var buttonAction linebot.TemplateAction
				if element[2] == "url" {
					buttonAction = linebot.NewURITemplateAction(element[0], element[1])
				} else if element[2] == "button" {
					buttonAction = linebot.NewMessageTemplateAction(element[0], element[1])
				}
				footerButton = &linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: buttonAction,
					Style:  linebot.FlexButtonStyleTypeLink,
					Margin: linebot.FlexComponentMarginTypeXs,
					Height: linebot.FlexButtonHeightTypeSm,
				}
				lineFooterComponent = append(lineFooterComponent, footerButton)
			}
		}
		lineFlexFooter = &linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: lineFooterComponent,
		}

		blockStyle := &linebot.BubbleStyle{
			Footer: &linebot.BlockStyle{
				Separator: true,
			},
		}

		flexBubbleContainer = &linebot.BubbleContainer{
			Type:   linebot.FlexContainerTypeBubble,
			Header: lineFlexHeader,
			Hero:   lineFlexHero,
			Body:   lineFlexBody,
			Footer: lineFlexFooter,
			Styles: blockStyle,
		}
		carousel = append(carousel, flexBubbleContainer)
	}

	carouselContainer = &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: carousel,
	}
	return linebot.NewFlexMessage("flex", carouselContainer)
}

func LineFlexForm(curText string) *linebot.FlexMessage {
	curText = strings.Replace(strings.TrimSuffix(curText, "}"), "FlexForm{", "", -1)
	formText := strings.Split(curText, "`")[0]
	NumberUnderScore := strings.Count(formText, "_")
	separator := "\n" + strings.Repeat("_", NumberUnderScore)
	formTypeText := strings.Split(formText, separator)[0]
	bodyText := strings.Split(formText, separator)[1]

	//Header
	var flexFormHeader *linebot.BoxComponent

	//Body
	re := regexp.MustCompile(`\n.*:`)
	bodyLabel := re.FindAllString(bodyText, -1)
	bodyContent := re.Split(bodyText, -1)
	bodyContent = bodyContent[1:len(bodyContent)]
	fmt.Println(bodyContent)

	var bodyContentBox *linebot.BoxComponent
	var bodyComponent []linebot.FlexComponent
	var bodyLabelComponent *linebot.TextComponent
	var bodyValueComponent *linebot.TextComponent

	var formTypeFlexComponent []linebot.FlexComponent
	formType := strings.Title(strings.Split(formTypeText, " ~ ")[1])
	formTypeComponent := &linebot.TextComponent{
		Type:    linebot.FlexComponentTypeText,
		Text:    formType,
		Size:    linebot.FlexTextSizeTypeXl,
		Weight:  linebot.FlexTextWeightTypeBold,
		Gravity: linebot.FlexComponentGravityTypeCenter,
	}
	spacer := &linebot.SpacerComponent{
		Type: linebot.FlexComponentTypeSpacer,
		Size: linebot.FlexSpacerSizeTypeXl,
	}
	formTypeFlexComponent = append(formTypeFlexComponent, formTypeComponent, spacer)
	formTypeBox := &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Margin:   linebot.FlexComponentMarginTypeXl,
		Contents: formTypeFlexComponent,
	}
	bodyComponent = append(bodyComponent, formTypeBox)

	for index := range bodyLabel {
		var bodyContentComponent []linebot.FlexComponent
		label := strings.TrimSuffix(strings.TrimSpace(bodyLabel[index]), ":")
		fmt.Println("label", label)
		value := strings.TrimSpace(bodyContent[index])
		fmt.Println("value", value)
		if strings.Contains(strings.ToLower(label), "invoice") {
			//Header Component
			var headerFieldComponent []linebot.FlexComponent
			var headerComponent []linebot.FlexComponent
			var headerLabel *linebot.TextComponent
			var headerValue *linebot.TextComponent
			headerLabel = &linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   label,
				Wrap:   true,
				Size:   linebot.FlexTextSizeTypeXxs,
				Weight: linebot.FlexTextWeightTypeBold,
			}
			headerValue = &linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   value,
				Wrap:   true,
				Size:   linebot.FlexTextSizeTypeXxs,
				Weight: linebot.FlexTextWeightTypeBold,
				Align:  linebot.FlexComponentAlignTypeEnd,
			}
			headerFieldComponent = append(headerFieldComponent, headerLabel, headerValue)
			formHeaderBox := &linebot.BoxComponent{
				Type:     linebot.FlexComponentTypeBox,
				Layout:   linebot.FlexBoxLayoutTypeHorizontal,
				Contents: headerFieldComponent,
			}
			headerComponent = append(headerComponent, formHeaderBox)
			flexFormHeader = &linebot.BoxComponent{
				Type:     linebot.FlexComponentTypeBox,
				Layout:   linebot.FlexBoxLayoutTypeVertical,
				Contents: headerComponent,
			}
		} else {
			//Body Component
			bodyLabelComponent = &linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   label,
				Weight: linebot.FlexTextWeightTypeBold,
				Wrap:   true,
				Size:   linebot.FlexTextSizeTypeXs,
			}
			bodyValueComponent = &linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Text:  value,
				Wrap:  true,
				Align: linebot.FlexComponentAlignTypeEnd,
				Size:  linebot.FlexTextSizeTypeXs,
			}
			bodyContentComponent = append(bodyContentComponent, bodyLabelComponent, bodyValueComponent)
			bodyContentBox = &linebot.BoxComponent{
				Type:     linebot.FlexComponentTypeBox,
				Layout:   linebot.FlexBoxLayoutTypeHorizontal,
				Contents: bodyContentComponent,
				Spacing:  linebot.FlexComponentSpacingTypeMd,
			}
			bodyComponent = append(bodyComponent, bodyContentBox)
			if index < len(bodyLabel)-1 {
				separator := &linebot.SeparatorComponent{
					Type: linebot.FlexComponentTypeSeparator,
				}
				bodyComponent = append(bodyComponent, separator)
			}
		}
	}
	flexFormBody := &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: bodyComponent,
		Spacing:  linebot.FlexComponentSpacingTypeSm,
	}

	//Footer
	formFooter := strings.Split(curText, "`")[1]
	formFooter = strings.Replace(strings.TrimSuffix(formFooter, "}"), "Confirm{", "", -1)
	formFooterComponent := strings.Split(formFooter, "|")
	var footerComponent []linebot.FlexComponent
	for _, button := range formFooterComponent {
		var color string
		var style linebot.FlexButtonStyleType
		if button == "Yes" {
			color = "#6d6d6d"
			style = linebot.FlexButtonStyleTypePrimary
		} else {
			color = "#e2e2e2"
			style = linebot.FlexButtonStyleTypeSecondary
		}
		buttonFlex := &linebot.ButtonComponent{
			Type:   linebot.FlexComponentTypeButton,
			Action: linebot.NewMessageTemplateAction(button, button),
			Style:  style,
			Color:  color,
			Height: linebot.FlexButtonHeightTypeSm,
		}
		footerComponent = append(footerComponent, buttonFlex)
	}
	flexFormFooter := &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeHorizontal,
		Spacing:  linebot.FlexComponentSpacingTypeSm,
		Contents: footerComponent,
	}

	//Style
	blockStyle := &linebot.BubbleStyle{
		Header: &linebot.BlockStyle{
			BackgroundColor: "#e5e5e5",
		},
		Body: &linebot.BlockStyle{
			Separator: true,
		},
		Footer: &linebot.BlockStyle{
			Separator: true,
		},
	}

	//Container
	flexBubbleContainer := &linebot.BubbleContainer{
		Type:   linebot.FlexContainerTypeBubble,
		Header: flexFormHeader,
		Body:   flexFormBody,
		Footer: flexFormFooter,
		Styles: blockStyle,
	}
	return linebot.NewFlexMessage("flex", flexBubbleContainer)
}
