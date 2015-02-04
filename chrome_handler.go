package chrome

type BaseClientHandler struct {
	clientHandlerT ClientHandlerT
	lifeSpan       LifeSpanHandlerT
	request        RequestHandlerT
	display        DisplayHandlerT
	download       DownloadHandlerT
}

func (ch *BaseClientHandler) GetContextMenuHandler() ContextMenuHandlerT {
	return ContextMenuHandlerT{nil}
}
func (ch *BaseClientHandler) GetDialogHandler() DialogHandlerT {
	return DialogHandlerT{nil}
}
func (ch *BaseClientHandler) GetDisplayHandler() DisplayHandlerT {
	return ch.display
}
func (ch *BaseClientHandler) SetDisplayHandler(dsp DisplayHandlerT) {
	ch.display = dsp
	return
}

func (ch *BaseClientHandler) GetDownloadHandler() DownloadHandlerT {
	return ch.download
}

func (ch *BaseClientHandler) SetDownloadHandler(download DownloadHandlerT) {
	ch.download = download
	return
}

func (ch *BaseClientHandler) GetDragHandler() DragHandlerT {
	return DragHandlerT{nil}
}
func (ch *BaseClientHandler) GetFocusHandler() FocusHandlerT {
	return FocusHandlerT{nil}
}
func (ch *BaseClientHandler) GetGeoLocationHandler() GeolocationHandlerT {
	return GeolocationHandlerT{nil}
}
func (ch *BaseClientHandler) GetJsDialogHandler() JsdialogHandlerT {
	return JsdialogHandlerT{nil}
}
func (ch *BaseClientHandler) GetKeyboardHandler() KeyboardHandlerT {
	return KeyboardHandlerT{nil}
}

func (ch *BaseClientHandler) SetLifeSpanHandler(lsh LifeSpanHandlerT) {
	ch.lifeSpan = lsh
	return
}

func (ch *BaseClientHandler) GetLifeSpanHandler() LifeSpanHandlerT {
	return ch.lifeSpan
}

func (ch *BaseClientHandler) GetLoadHandler() LoadHandlerT {
	return LoadHandlerT{nil}
}
func (ch *BaseClientHandler) GetRenderHandler() RenderHandlerT {
	return RenderHandlerT{nil}
}

func (ch *BaseClientHandler) SetRequestHandler(rqh RequestHandlerT) {
	ch.request = rqh
	return
}

func (ch *BaseClientHandler) GetRequestHandler() RequestHandlerT {
	return ch.request
}

func (ch *BaseClientHandler) GetClientHandlerT() ClientHandlerT {
	return ch.clientHandlerT
}
func (ch *BaseClientHandler) SetClientHandlerT(cht ClientHandlerT) {
	ch.clientHandlerT = cht
	return
}
