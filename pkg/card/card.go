package card

import (
	"fmt"
	"log"
	"sync"
)

type Holder struct {
	Id   int
	Name string
}

type Card struct {
	Id       int
	Number   string
	Issuer   string
	HolderId int
	Type     string
}

type Service struct {
	mu      sync.RWMutex
	cards   []*Card
	holders []*Holder
}

func NewService() *Service { return &Service{} }

func (s *Service) AddHolder(name string) {
	var holder Holder

	holder.Name = name
	holder.Id = len(s.holders)

	s.holders = append(s.holders, &holder)
}

func (s *Service) AddCard(issuer string, holder int, image string) {
	var card Card

	card.Id = len(s.cards)
	card.Number = fmt.Sprintf("000%d", len(s.cards))
	card.Issuer = issuer
	card.HolderId = holder
	card.Type = image

	s.cards = append(s.cards, &card)
}

func (s *Service) AllCards() []*Card {
	s.mu.RLock()
	defer s.mu.RLock()
	return s.cards
}

func (s *Service) HolderCards(holderid int) []*Card {
	s.mu.RLock()
	defer s.mu.RLock()

	cards := make([]*Card, 0)
	for _, c := range s.cards {
		if c.HolderId == holderid {
			cards = append(cards, c)
		}
	}
	return cards
}

func (s *Service) AddHolderCard(issuer string, holder int, image string) {

	if (issuer == "visa") || (issuer == "master") {
		if (image == "plastic") || (image == "virtual") {
			for _, c := range s.cards {
				if holder == c.HolderId {
					s.AddCard(issuer, holder, image)
					break
				}
			}
		} else {
			log.Println("wrong card type")
			return
		}

	} else {
		log.Println("wrong card type")
		return
	}

	log.Println("card is added")
	return
}
