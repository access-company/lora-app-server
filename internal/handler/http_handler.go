package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"

	log "github.com/Sirupsen/logrus"
)

// HttpHandler implements a HTTP handler for sending and receiving data by
// an application.
type HttpHandler struct {
	dataDownChan chan DataDownPayload
	callbackURL  string
}

// NewHttpHandler creates a new HttpHandler.
func NewHttpHandler(url string) (Handler, error) {
	h := HttpHandler{
		dataDownChan: make(chan DataDownPayload),
		callbackURL:  url,
	}
	return &h, nil
}

// Close stops the handler.
func (h *HttpHandler) Close() error {
	log.Info("handler/http: closing handler")
	close(h.dataDownChan)
	return nil
}

// SendDataUp sends a DataUpPayload.
func (h *HttpHandler) SendDataUp(payload DataUpPayload) error {
	log.Info("handler/http: SendDataUp")
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/http: data-up payload marshal error: %s", err)
	}
	fmt.Printf( "payload %+v\n", b)

	resp, err := http.Post(h.CallbackURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("handler/http: fail to request: %s", err)
	}
	fmt.Printf( "response %+v\n", resp)
	return nil
}

// SendJoinNotification sends a JoinNotification.
func (h *HttpHandler) SendJoinNotification(payload JoinNotification) error {
	log.Info("handler/http: SendJoinNotification")
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/http: join notification marshal error: %s", err)
	}
	fmt.Printf( "payload %+v\n", b)
	return nil
}

// SendACKNotification sends an ACKNotification.
func (h *HttpHandler) SendACKNotification(payload ACKNotification) error {
	log.Info("handler/http: SendACKNotification")
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/http: ack notification marshal error: %s", err)
	}
	fmt.Printf( "payload %+v\n", b)
	return nil
}

// SendErrorNotification sends an ErrorNotification.
func (h *HttpHandler) SendErrorNotification(payload ErrorNotification) error {
	log.Info("handler/http: SendErrorNotification")
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/http: error notification marshal error: %s", err)
	}
	fmt.Printf( "payload %+v\n", b)
	return nil
}

// DataDownChan returns the channel containing the received DataDownPayload.
func (h *HttpHandler) DataDownChan() chan DataDownPayload {
	log.Info("handler/http: DataDownChan")
	return h.dataDownChan
}
