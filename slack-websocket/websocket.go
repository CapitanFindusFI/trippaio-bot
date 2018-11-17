package slack_websocket

import (
	"fmt"
	"github.com/CapitanFindusFI/trippaio-bot/slack-websocket/users"
	"github.com/nlopes/slack"
	"math/rand"
	"time"
)

func shuffle(vals []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

func erudisci() []string {
	listaErudizione := []string{
		"crostini di poppa",
		"guancia di vitello giovane",
		"insalatina di trippa e nervetti con battuto di prezzemolo e aglio",
		"lampredotto classico",
		"lampredotto rifatto con le patate",
		"lampredotto rifatto all'uccelletto",
		"lampredotto alla cacciatora con le olive",
		"polpettine di lampredotto",
		"sushi di lampredotto",
		"caciucco di lampredotto",
		"stracotto di vitellino giovane",
	}

	return shuffle(listaErudizione)
}

func complimenta() string {
	listaComplimenti := []string{
		"vi ringrazio ragazzi",
		"gentilissimi",
		"sì eh? grazie ragazzi",
	}

	randIndex := rand.Int31n(int32(len(listaComplimenti)))

	return listaComplimenti[randIndex]
}

func offendiFiglio() string {
	listaOffese := []string{
		"lapo sei sempre nì mezzo dio bono!",
		"lapo che ti sposti un po'!",
		"lapo che hai belle fatto con loro o no?",
		"allora lapo, t'ho detto 5 e 50!",
	}

	randIndex := rand.Int31n(int32(len(listaOffese)))

	return listaOffese[randIndex]
}

func wave() string {
	listaSaluti := []string{
		"ciao lillo",
		"ciao ragazzi",
		"buongiorno signori!",
		"eccoli!",
		"carissimi buongiorno!",
	}

	randIndex := rand.Int31n(int32(len(listaSaluti)))

	return listaSaluti[randIndex]
}

func pagare() []string {
	listaPagare := []string{
		"allora...",
		"hai preso?",
		"stracot...",
		"vaschetta o panino?",
		"nana bianca?",
		"aah allora...",
		"polpett... 5 e 50...",
		"nana bianca 5 e basta",
		"vaschett... 6 e 50",
		"allora 7 e 50",
	}

	return listaPagare
}

func senza() string {
	listaSenza := []string{
		"come senza? o che un tu sta bene?",
		"ebreo",
		"vaia senza, o che sei grullo?",
	}

	randIndex := rand.Int31n(int32(len(listaSenza)))

	return listaSenza[randIndex]
}

type UsersMap map[string]UserHandler

func getUserInfo(rtm *slack.RTM, event *slack.MessageEvent) {
	userInfo, _ := rtm.GetUserInfo(event.User)

	userName := userInfo.Name

	um := UsersMap{
		"andrea.martini": users.AndreaHandler{},
		"giusepping":     users.GiuseppeHandler{},
		"alessio.melani": users.MeloHandler{},
		"lomi":           users.LomiHandler{},
		"alberto":        users.BiniHandler{},
	}

	if method, present := um[userName]; present {
		messages := method.Handle(rtm, event)
		for _, msg := range messages {
			rtm.SendMessage(rtm.NewOutgoingMessage(msg, event.Channel))
			time.Sleep(time.Second * time.Duration(rand.Int31n(3)))
		}
	} else {
		fmt.Printf("User: %s\n Name: %s", userName, userInfo.Profile.DisplayName)
	}
}

func WebsocketHandler(rtm *slack.RTM) {
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)

		case *slack.MessageEvent:
			getUserInfo(rtm, ev)

		default:

			// Ignore other slack-events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
