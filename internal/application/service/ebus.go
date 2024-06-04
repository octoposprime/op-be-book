package application

import (
	"context"

	me "github.com/octoposprime/op-be-book/internal/domain/model/entity"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// This is the event listener handler of the application layer.
func (a *Service) EventListen() *Service {
	go a.Listen(context.Background(), smodel.ChannelCreatePage, a.EventListenerCallBack)
	go a.Listen(context.Background(), smodel.ChannelDeletePage, a.EventListenerCallBack)
	return a
}

// This is a call-back function of the event listener handler of the application layer.
func (a *Service) EventListenerCallBack(channelName string, page me.Page) {
	if channelName == smodel.ChannelCreatePage {
		a.CreatePage(context.Background(), page)
	} else if channelName == smodel.ChannelDeletePage {
		a.DeletePage(context.Background(), page)
	} else {
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "EventListenerCallBack", channelName, smodel.ErrorChannelNameNotValid.Error()))
	}
}
