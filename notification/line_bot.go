package notification

type LineBotClient struct {
}

func NewLineBotClient() *LineBotClient {
	lbc := new(LineBotClient)
	return lbc
}

func (lbc LineBotClient) Notify() {

}
