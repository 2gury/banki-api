package chatgpt_client

import (
	"banks-api/internal/pkg/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ayush6624/go-chatgpt"
)

const (
	format = `
		{
			"company_name": string,
			"company_logo": string,
			"period_from": int, // если нет этой инфомации, то оставляй 0
			"period_to": int, // если нет этой инфомации, то оставляй 0
			"amount_from": int, // если нет этой инфомации, то оставляй 0
			"amount_to": int, // если нет этой инфомации, то оставляй 0
			"rate_from": float, // если нет этой инфомации, то оставляй 0
			"rate_to": float, // если нет этой инфомации, то оставляй 0
			"registration": [
				"string" // если нет этой инфомации, то оставляй пустым
			],
			"credit_purpose": [
				"string" // если нет этой инфомации, то оставляй пустым
			],
			"documents": [
				"string" // если нет этой инфомации, то оставляй пустым
			],
			"review_time": int, // если нет этой инфомации, то оставляй 0
			"obtain_method": [string], // если нет этой инфомации, то оставляй пустым
			"description": string // Краткое описание продукта/компании без цифр
		}
`
)

type ChatgptClient struct {
	cli *chatgpt.Client
}

func NewChatgptClient(key string) (*ChatgptClient, error) {
	cli, err := chatgpt.NewClient(key)
	if err != nil {
		return nil, err
	}

	return &ChatgptClient{
		cli: cli,
	}, nil
}

func (gpt *ChatgptClient) GrabText(ctx context.Context, text string) (*model.BankChatGPTInfo, error) {
	res, err := gpt.cli.SimpleSend(ctx, fmt.Sprintf(`
		Конвертируй информацию по продукту в следующем формате. В ответ пришли только json, так как я его буду парсить автоматически
		
		Информация по продукту: "%s"

		Формат: "%s"`, text, format))
	if err != nil {
		return nil, err
	}

	if len(res.Choices) == 0 {
		return nil, fmt.Errorf("error zero choices from chatgpt")
	}

	var convertedInfo model.BankChatGPTInfo
	bytesContent := []byte(res.Choices[0].Message.Content)
	if err := json.Unmarshal(bytesContent, &convertedInfo); err != nil {
		return nil, err
	}

	return &convertedInfo, nil
}

func (gpt *ChatgptClient) IsValidReview(ctx context.Context, content string) (bool, error) {
	res, err := gpt.cli.SimpleSend(ctx, fmt.Sprintf(`
		Провалидируй следующее ревью по следующим правилам:
		1. Не должно быть оскорблений
		2. Не должно быть мата
		3. Не должно быть призыва к насилию
		3. Не должно быть рекламы
		4. Не должно быть непонятных символов

		Если по отзыву не понятно валиден ли они или нет, есть сомнения, то отклоняй такой отзыв. 
		В ответ пришли: "true", "false"
		Если отзыв вылиден, то true, если не прошел по правилам то false
		
		Текст отзыва: "%s"`, content))
	if err != nil {
		return false, err
	}

	if len(res.Choices) == 0 {
		return false, fmt.Errorf("error zero choices from chatgpt")
	}

	if res.Choices[0].Message.Content == "true" {
		return true, nil
	}

	return false, nil
}
