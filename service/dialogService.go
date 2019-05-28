package service

import (
	"diploma/domain"
)

type DialogService struct {
	Repository domain.DialogStore
}

func (dService DialogService) GetAllDialogs(userToken string) []domain.Dialog {
	var dialogs []domain.Dialog
	dialogs = dService.Repository.ReadAllDialogs(userToken)
	return dialogs
}

func (dService DialogService) GetDialogByID(id string) []domain.Message {
	var messages []domain.Message
	messages = dService.Repository.ReadDialog(id)
	return messages
}
