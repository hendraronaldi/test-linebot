package main

import (
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func LineConfirmCarousel(curText string) *linebot.FlexMessage {
	curText = strings.TrimRight(strings.TrimLeft(curText, "Confirm{"), "}")
	var confirmationText string
	// var onSave string
	var index int
	// if strings.Contains(curText, ";") {
	element := strings.Split(curText, ";")
	confirmationText = element[0]
	if len(element) > 2 {
		// onSave = element[1]
		index = 2
	} else {
		index = 1
	}
	// }
	confirmText := strings.Split(element[index], "|")

	var confirmFlexTemplate *linebot.BubbleContainer
	var templateBody *linebot.BoxComponent
	var templateBodyComponent []linebot.FlexComponent
	var confirmTextComponent *linebot.TextComponent
	var separator *linebot.SeparatorComponent
	var confirmButtonTemplate *linebot.BoxComponent
	var confirmButtonTemplateComponent []linebot.FlexComponent
	var confirmButtonYes *linebot.ButtonComponent
	var confirmButtonNo *linebot.ButtonComponent

	confirmTextComponent = &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   confirmationText,
		Weight: linebot.FlexTextWeightTypeBold,
		Wrap:   true,
	}
	separator = &linebot.SeparatorComponent{
		Type: linebot.FlexComponentTypeSeparator,
	}
	templateBodyComponent = append(templateBodyComponent, confirmTextComponent, separator)

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

	confirmFlexTemplate = &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: templateBody,
	}

	return linebot.NewFlexMessage("confirmation", confirmFlexTemplate)

	// var template *linebot.ConfirmTemplate
	// if strings.Contains(strings.Split(strings.Split(logics[i-1], "\n")[0], " ~ ")[1], "reminder") {
	// 	template = linebot.NewConfirmTemplate(
	// 		confirmationText,
	// 		linebot.NewPostbackTemplateAction(confirmText[0], onSave, confirmText[0]+" ~ "+logics[i-1], ""),
	// 		linebot.NewMessageTemplateAction(confirmText[1], confirmText[1]+" ~ "+logics[i-1]),
	// 	)
	// } else {
	// 	template = linebot.NewConfirmTemplate(
	// 		confirmationText,
	// 		linebot.NewMessageTemplateAction(confirmText[0], confirmText[0]+" ~ "+logics[i-1]),
	// 		linebot.NewMessageTemplateAction(confirmText[1], confirmText[1]+" ~ "+logics[i-1]),
	// 	)
	// }
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
		// var action []linebot.TemplateAction
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
				//TODO hero
				image = element[0]
				lineFlexHero = &linebot.ImageComponent{
					Type: linebot.FlexComponentTypeImage,
					URL:  image,
					Size: linebot.FlexImageSizeTypeFull,
				}

				//TODO body
				title = element[1]
				titleComponent := &linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   title,
					Size:   linebot.FlexTextSizeTypeXl,
					Weight: linebot.FlexTextWeightTypeBold,
					Wrap:   true,
				}
				description = element[2]
				descriptionComponent := &linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: description,
					Wrap: true,
				}
				lineFlexBodyComponent = append(lineFlexBodyComponent, titleComponent, descriptionComponent)
				lineFlexBody = &linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Contents: lineFlexBodyComponent,
				}
			} else {
				//TODO footer
				var footerButton *linebot.ButtonComponent
				var buttonAction linebot.TemplateAction
				if element[2] == "url" {
					// action = append(action, linebot.NewURITemplateAction(element[0], element[1]))
					buttonAction = linebot.NewURITemplateAction(element[0], element[1])
				} else if element[2] == "button" {
					// action = append(action, linebot.NewMessageTemplateAction(element[0], element[1]))
					buttonAction = linebot.NewMessageTemplateAction(element[0], element[1])
				}
				footerButton = &linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: buttonAction,
					Style:  linebot.FlexButtonStyleTypeLink,
					Margin: linebot.FlexComponentMarginTypeSm,
				}
				lineFooterComponent = append(lineFooterComponent, footerButton)
			}
		}
		lineFlexFooter = &linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: lineFooterComponent,
		}

		flexBubbleContainer = &linebot.BubbleContainer{
			Type:   linebot.FlexContainerTypeBubble,
			Header: lineFlexHeader,
			Hero:   lineFlexHero,
			Body:   lineFlexBody,
			Footer: lineFlexFooter,
		}
		carousel = append(carousel, flexBubbleContainer)

		// carouselColumn := linebot.NewCarouselColumn(image,
		// 	title,
		// 	description,
		// 	action...,
		// )
		// carousel = append(carousel, carouselColumn)
	}
	// template := linebot.NewCarouselTemplate(
	// 	carousel...,
	// )

	carouselContainer = &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: carousel,
	}
	return linebot.NewFlexMessage("flex", carouselContainer)
}

func LineFlexForm(curText string) *linebot.FlexMessage {
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
			Action: linebot.NewMessageTemplateAction(button, button),
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

	flexBubbleContainer = &linebot.BubbleContainer{
		Type:   linebot.FlexContainerTypeBubble,
		Header: flexFormHeader,
		Body:   flexFormBody,
		Footer: flexFormFooter,
	}
	return linebot.NewFlexMessage("flex", flexBubbleContainer)
}
