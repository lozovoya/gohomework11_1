package app

import (
	"encoding/json"
	"github.com/lozovoya/gohomework11_1/cmd/bank/app/dto"
	"github.com/lozovoya/gohomework11_1/pkg/card"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	cardSvc *card.Service
	mux     *http.ServeMux
}

func NewServer(cardSvc *card.Service, mux *http.ServeMux) *Server {
	return &Server{cardSvc: cardSvc, mux: mux}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/addCard", s.addCard)
	s.mux.HandleFunc("/getAllCards", s.getAllCards)
	s.mux.HandleFunc("/getHolderCards", s.getHolderCards)
	s.mux.HandleFunc("/addHolderCard", s.addHolderCard)
}

func (s *Server) addCard(w http.ResponseWriter, r *http.Request) {
	log.Println("implement me")
}

func (s *Server) getAllCards(w http.ResponseWriter, r *http.Request) {
	cards := s.cardSvc.AllCards()
	if len(cards) == 0 {
		log.Println("no cards available")
		return
	}
	dtos := make([]*dto.CardDTO, len(cards))
	for i, c := range cards {
		dtos[i] = &dto.CardDTO{
			Id:       c.Id,
			Number:   c.Number,
			Issuer:   c.Issuer,
			HolderId: c.HolderId,
			Type:     c.Type,
		}
	}

	respBody, err := json.Marshal(dtos)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
		return
	}

}

func (s *Server) getHolderCards(w http.ResponseWriter, r *http.Request) {

	holderid, err := strconv.Atoi(r.FormValue("holderid"))
	if err != nil {
		log.Println(err)
		return
	}
	cards := s.cardSvc.HolderCards(holderid)
	if len(cards) == 0 {
		log.Println("no cards available")
		return
	}

	dtos := make([]*dto.CardDTO, len(cards))
	for i, c := range cards {
		dtos[i] = &dto.CardDTO{
			Id:       c.Id,
			Number:   c.Number,
			Issuer:   c.Issuer,
			HolderId: c.HolderId,
			Type:     c.Type,
		}
	}

	respBody, err := json.Marshal(dtos)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		log.Println(err)
		return
	}

}

func (s *Server) addHolderCard(w http.ResponseWriter, r *http.Request) {

	holderid, err := strconv.Atoi(r.FormValue("holderid"))
	if err != nil {
		log.Println(err)
		return
	}
	issuer := r.FormValue("issuer")
	if err != nil {
		log.Println(err)
		return
	}
	image := r.FormValue("image")
	if err != nil {
		log.Println(err)
		return
	}

	s.cardSvc.AddHolderCard(issuer, holderid, image)
	return
}
