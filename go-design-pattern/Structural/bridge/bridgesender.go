package bridge

type IMsgSender interface {
	Send(msg string) error
}

type EmailMsgSender struct {
	emails []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{
		emails: emails,
	}
}
func (e *EmailMsgSender) Send(msg string) error {
	return nil
}

type PhoneMsgSender struct {
	numbers []string
}

func NewPhoneMsgSender(numbers []string) *PhoneMsgSender {
	return &PhoneMsgSender{
		numbers: numbers,
	}

}
func (e *PhoneMsgSender) Send(msg string) error {
	return nil
}

type WechatMsgSender struct {
	number []string
}

func (e *WechatMsgSender) Send(msg string) error {
	return nil
}
func NewWechatMsgSender() *WechatMsgSender {
	return &WechatMsgSender{}
}

type INotification interface {
	Notify(msg string) error
}

// ErrorNotification 错误通知
// 后面可能还有 warning 各种级别
type ErrorNotification struct {
	sender IMsgSender
}

// NewErrorNotification NewErrorNotification
func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

// Notify 发送通知
func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}

// WarnNotification 警告通知
type WarnNotification struct {
	sender IMsgSender
}

// NewWarnNotification
func NewWarnNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

// Notify 发送通知
func (n *WarnNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
