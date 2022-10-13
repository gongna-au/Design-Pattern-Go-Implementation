package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorNotification_Notify(t *testing.T) {
	sender1 := NewEmailMsgSender([]string{"test@test.com"})
	n1 := NewErrorNotification(sender1)
	err1 := n1.Notify("test msg")
	assert.Nil(t, err1)
	sender2 := NewEmailMsgSender([]string{"15002597110"})
	n2 := NewErrorNotification(sender2)
	err2 := n2.Notify("test msg")
	assert.Nil(t, err2)

}
